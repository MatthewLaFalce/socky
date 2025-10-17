package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/fatih/color"
	"github.com/gorilla/websocket"
)

func main() {
	// Command line flags
	addr := flag.String("addr", "localhost:8080", "WebSocket server address")
	path := flag.String("path", "/", "WebSocket path")
	token := flag.String("token", "", "Authentication token (e.g. Bearer token)")
	scheme := flag.String("scheme", "wss", "WebSocket scheme (ws or wss)")
	flag.Parse()

	// Re-parse to include DB flags even if they appear after the binary name
	flag.CommandLine.Parse(os.Args[1:])

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: *scheme, Host: *addr, Path: *path}
	color.New(color.FgCyan).Printf("Connecting to %s\n", u.String())

	// Set token header
	headers := http.Header{}
	if *token != "" {
		headers.Add("Authorization", *token)
	}

	// Connect to WebSocket
	c, _, err := websocket.DefaultDialer.Dial(u.String(), headers)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	// Read messages from the server
	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				color.New(color.FgRed).Printf("read: %v\n", err)
				return
			}
			color.New(color.FgGreen).Printf("recv: %s\n", message)
		}
	}()

	// Send messages typed into stdin
	go func() {
		for {
			var msg string
			fmt.Print(">> ")
			fmt.Scanln(&msg)
			err := c.WriteMessage(websocket.TextMessage, []byte(msg))
			if err != nil {
				color.New(color.FgRed).Printf("write: %v\n", err)
				return
			}
		}
	}()

	// Wait for Ctrl+C to close
	for {
		select {
		case <-done:
			return
		case <-interrupt:
			color.New(color.FgYellow).Printf("Interrupt detected, closing connection...\n")
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				color.New(color.FgRed).Printf("close: %v\n", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
