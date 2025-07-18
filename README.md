# p2ptest

A minimal Go demo for experimenting with [libp2p](https://libp2p.io/) peer-to-peer networking.

This project creates a simple peer node using [go-libp2p](https://github.com/libp2p/go-libp2p). You can use it to start two or more nodes locally or on separate machines, connect them, and exchange simple messages.

## Features

- Start a libp2p node with a random peer ID.
- Listen for incoming connections on a fixed port (TCP/4001).
- Optionally connect to a peer by providing their multiaddress.
- Print incoming messages to the console.

## Getting Started

### Prerequisites

- Go 1.18 or later installed ([Download Go](https://go.dev/dl/))

### Installation

1. **Clone the repository:**

    ```bash
    git clone https://github.com/lrogana/p2ptest.git
    cd p2ptest
    ```

2. **Initialize Go modules and install dependencies:**

    ```bash
    go mod tidy
    ```

---

## Usage

You can run multiple instances to simulate a small peer-to-peer network.

### **Start the First Node**

Run the following in a terminal:
```go run main.go```

You should see output similar to:
Your Peer ID: 12D3KooW...
Listening on:
/ip4/0.0.0.0/tcp/4001/p2p/12D3KooW...

### Start a Second Node and Connect to the First
Copy the multiaddress shown by the first node (e.g. /ip4/127.0.0.1/tcp/4001/p2p/12D3KooW...)
Then, on another terminal or computer, run:
```go run main.go /ip4/127.0.0.1/tcp/4001/p2p/12D3KooW...``` 