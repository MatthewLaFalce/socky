# 🧩 Socky

**Socky** is a simple Go-based WebSocket client utility — like `wscat`, but written in Go for portability and ease of use.

It allows you to connect to any WebSocket server (including secure `wss://` endpoints), optionally authenticate with a token, and interactively send and receive messages from the terminal.

---

## 🚀 Features

- ✅ Connect to any WebSocket (`ws://` or `wss://`)
- 🔐 Optional Bearer token authentication
- 🎨 Colorized terminal output for clarity
- 🖥️ Interactive mode: send messages directly from stdin
- ⚡ Graceful shutdown on Ctrl+C

---

## 🧰 Installation

### Prerequisites
- Go 1.18 or newer installed on your system.

### Clone & Build
```bash
git clone https://github.com/MatthewLaFalce/socky.git
cd socky
./0_build.sh
```

This will produce binaries in the `./dist/` directory:
- `socky-linux`
- `socky-mac`

Alternatively, you can build manually with:
```bash
go build -o socky main.go
```

---

## 💡 Usage

```bash
./socky --addr example.com:8080 --path /ws --token "Bearer YOUR_TOKEN"
```

### Options

| Flag | Description | Default |
|------|--------------|----------|
| `--addr` | WebSocket server address (host:port) | `localhost:8080` |
| `--path` | WebSocket path | `/` |
| `--token` | Optional authentication token | *(none)* |

Example interactive session:
```bash
$ ./socky --addr echo.websocket.events --path /
Connecting to wss://echo.websocket.events/
>> hello
recv: hello
Interrupt detected, closing connection...
```

---

## 🧪 Development

To run directly without building:
```bash
go run main.go --addr echo.websocket.events --path /
```

---

## ⚠️ Notes

- `wss://` requires a valid SSL certificate.
- If you encounter connection issues, verify your server supports WebSocket upgrade (`HTTP 101 Switching Protocols`).
