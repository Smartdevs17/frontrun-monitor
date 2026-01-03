# Frontrun Monitor

A high-performance frontrunning detection and monitoring system for Ethereum/DeFi. Monitor the mempool in real-time to detect and alert on suspicious transaction patterns that indicate frontrunning attempts.

## 🎯 Project Status

**Current Phase:** Foundation Complete - Ready for Core Implementation

✅ **Implemented:**
- HTTP server with graceful shutdown
- Configuration management with validation
- Health check endpoints
- Build automation (Makefile)
- Go best practices (error handling, timeouts, context)

🚧 **In Progress:**
- Mempool monitoring implementation
- Frontrunning detection algorithms
- Database integration

📋 **Planned:**
- WebSocket real-time updates
- Frontend dashboard
- Advanced pattern detection

## 📁 Project Structure

```
frontrun-monitor/
├── backend/              # Go backend service
│   ├── main.go          # Application entry point
│   ├── server.go        # HTTP server implementation
│   ├── config.go        # Configuration management
│   ├── go.mod           # Go module definition
│   └── go.sum           # Dependency checksums
├── Makefile             # Build automation
├── .gitignore           # Git ignore rules
├── LICENSE              # MIT License
└── README.md            # This file
```

## 🛠 Technology Stack

### Backend
- **Language:** Go 1.24.0
- **Framework:** Standard library (`net/http`)
- **Dependencies:**
  - `github.com/ethereum/go-ethereum` - Ethereum client library
  - `github.com/gorilla/websocket` - WebSocket support
  - `github.com/sirupsen/logrus` - Structured logging
  - `golang.org/x/time/rate` - Rate limiting

### Planned
- **Database:** PostgreSQL (configured, not connected)
- **Frontend:** Modern UI framework (not started)
- **Real-time:** WebSocket connections

## 🚀 Getting Started

### Prerequisites

- **Go 1.21+** (currently using Go 1.24.0)
- **PostgreSQL** (for future database integration)
- **Ethereum Node Access** (Infura, Alchemy, or local node)

### Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/Smartdevs17/frontrun-monitor.git
   cd frontrun-monitor
   ```

2. **Install dependencies:**
   ```bash
   make deps
   ```
   Or manually:
   ```bash
   cd backend
   go mod download
   ```

3. **Build the project:**
   ```bash
   make build
   ```

4. **Run the server:**
   ```bash
   make run
   ```
   Or use the binary:
   ```bash
   ./backend/frontrun-monitor
   ```

## ⚙️ Configuration

The application uses environment variables for configuration:

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT` | `8080` | HTTP server port (1-65535) |
| `ENV` | `development` | Environment: `development`, `staging`, or `production` |
| `DATABASE_URL` | `postgres://localhost/frontrun_monitor` | PostgreSQL connection string |
| `LOG_LEVEL` | `info` | Log level: `debug`, `info`, `warn`, or `error` |

### Example Configuration

```bash
export PORT=8080
export ENV=production
export DATABASE_URL=postgres://user:password@localhost:5432/frontrun_monitor
export LOG_LEVEL=info
```

## 📡 API Endpoints

### Current Endpoints

#### `GET /health`
Health check endpoint.

**Response:**
```json
{
  "status": "healthy"
}
```

#### `GET /`
API information endpoint.

**Response:**
```json
{
  "message": "Frontrun Monitor API",
  "version": "1.0.0"
}
```

### Planned Endpoints

- `GET /api/v1/detections` - List detected frontrunning attempts
- `GET /api/v1/detections/:id` - Get detection details
- `GET /api/v1/transactions` - List monitored transactions
- `GET /api/v1/stats` - Statistics and metrics
- `WS /ws` - WebSocket connection for real-time updates

## 🔧 Makefile Commands

The project includes a comprehensive Makefile for common tasks:

### Build Commands
```bash
make build          # Build the backend binary
make build-linux    # Build for Linux
make build-windows  # Build for Windows
make build-darwin   # Build for macOS
make build-all      # Build for all platforms
```

### Development Commands
```bash
make run            # Run the server
make dev            # Run with auto-reload (requires air)
make deps           # Install dependencies
make fmt            # Format code
make vet            # Run go vet
make lint           # Run golangci-lint (if installed)
```

### Testing Commands
```bash
make test           # Run tests
make test-coverage  # Run tests with coverage report
```

### Utility Commands
```bash
make clean          # Clean build artifacts
make install        # Build and install binary
make help           # Show all available commands
```

## 🏗 Architecture

### Current Implementation

```
┌─────────────────┐
│   main.go       │  Entry point, graceful shutdown
│                 │  Signal handling, context management
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│   server.go     │  HTTP server with timeouts
│                 │  Route registration, graceful shutdown
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│   config.go     │  Configuration loading & validation
│                 │  Environment variable handling
└─────────────────┘
```

### Planned Architecture

```
┌─────────────────┐
│   main.go       │
└────────┬────────┘
         │
         ├──► Server ──► API Endpoints
         │
         ├──► MempoolMonitor ──► Ethereum Node
         │              │
         │              ├──► Transaction Analysis
         │              │
         │              └──► Detection Engine
         │
         ├──► Database ──► PostgreSQL
         │
         └──► WebSocket ──► Real-time Updates
```

## ✨ Features

### Implemented Features

- ✅ **Graceful Shutdown** - Proper signal handling and resource cleanup
- ✅ **Configuration Validation** - Validates all configuration values
- ✅ **HTTP Server Timeouts** - Prevents resource exhaustion
- ✅ **Error Handling** - Proper error wrapping and propagation
- ✅ **Health Checks** - Monitoring endpoint for service health
- ✅ **Build Automation** - Comprehensive Makefile

### Planned Features

- 🔄 **Mempool Monitoring** - Real-time transaction monitoring
- 🔍 **Frontrunning Detection** - Pattern recognition algorithms
- 💾 **Database Storage** - Persistent storage for detections
- 🔌 **WebSocket Support** - Real-time updates to clients
- 📊 **Dashboard** - Frontend visualization interface


 
##  Development

### Running in Development Mode

```bash
# Install air for hot reload
go install github.com/cosmtrek/air@latest

# Run with auto-reload
make dev
```

### Code Quality

```bash
# Format code
make fmt

# Run static analysis
make vet

# Run linter (if installed)
make lint
```

### Testing

```bash
# Run tests
make test

# Generate coverage report
make test-coverage
```

## 📦 Dependencies

### Core Dependencies

- `github.com/ethereum/go-ethereum` - Ethereum blockchain interaction
- `github.com/gorilla/websocket` - WebSocket protocol support
- `github.com/sirupsen/logrus` - Structured logging
- `golang.org/x/time/rate` - Rate limiting

### Installation

Dependencies are automatically installed when you run:
```bash
make deps
```

Or manually:
```bash
cd backend
go get github.com/ethereum/go-ethereum
go get github.com/gorilla/websocket
go get github.com/sirupsen/logrus
go get golang.org/x/time/rate
```

## 🚦 Project Roadmap

### Phase 1: Foundation ✅ (Current)
- [x] HTTP server setup
- [x] Configuration management
- [x] Graceful shutdown
- [x] Build automation

### Phase 2: Core Monitoring (Next)
- [ ] Ethereum node connection
- [ ] Mempool subscription
- [ ] Transaction processing
- [ ] Basic detection algorithms

### Phase 3: Advanced Features
- [ ] Database integration
- [ ] WebSocket server
- [ ] Advanced pattern detection
- [ ] Analytics and metrics

### Phase 4: Frontend & Polish
- [ ] Dashboard UI
- [ ] Real-time visualization
- [ ] API documentation
- [ ] Production deployment

## 🤝 Contributing

Contributions are welcome! Please feel free to submit pull requests.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- Built with [Go](https://golang.org/)
- Uses [go-ethereum](https://github.com/ethereum/go-ethereum) for Ethereum integration
- Inspired by the need for better MEV/frontrunning detection tools

## 📞 Support

For issues, questions, or contributions, please open an issue on GitHub.

---

**Status:** 🚧 Active Development | **Version:** 1.0.0 | **License:** MIT
