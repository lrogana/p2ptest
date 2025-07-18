package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	libp2p "github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	ma "github.com/multiformats/go-multiaddr"
)

func main() {
	ctx := context.Background()

	// 👇 Build host options conditionally
	var opts []libp2p.Option

	// If no address is provided, we're Peer A and will bind to fixed port 4001
	if len(os.Args) == 1 {
		listenAddr, _ := ma.NewMultiaddr("/ip4/0.0.0.0/tcp/4001")
		opts = append(opts, libp2p.ListenAddrs(listenAddr))
	}

	// ✅ Create the libp2p host
	h, err := libp2p.New(opts...)
	if err != nil {
		panic(err)
	}
	defer h.Close()

	fmt.Println("Your Peer ID:", h.ID())
	fmt.Println("Listening on:")
	for _, addr := range h.Addrs() {
		fmt.Printf("  %s/p2p/%s\n", addr, h.ID())
	}

	// 📥 Stream handler for incoming connections
	h.SetStreamHandler("/p2p/1.0.0", func(s network.Stream) {
		fmt.Println("New stream from:", s.Conn().RemotePeer())

		rw := bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))

		// Read loop
		go func() {
			for {
				str, err := rw.ReadString('\n')
				if err != nil {
					return
				}
				fmt.Printf("Received from %s: %s", s.Conn().RemotePeer(), str)
			}
		}()

		// Write loop
		go func() {
			for {
				fmt.Print(">> ")
				stdReader := bufio.NewReader(os.Stdin)
				text, _ := stdReader.ReadString('\n')
				rw.WriteString(text)
				rw.Flush()
			}
		}()
	})

	// 📤 Outgoing connection logic (if address provided)
	if len(os.Args) > 1 {
		addrStr := os.Args[1]

		// Parse multiaddr
		maddr, err := ma.NewMultiaddr(addrStr)
		if err != nil {
			panic(err)
		}

		// Extract peer info
		info, err := peer.AddrInfoFromP2pAddr(maddr)
		if err != nil {
			panic(err)
		}

		// Connect to peer
		if err := h.Connect(ctx, *info); err != nil {
			panic(err)
		}

		fmt.Println("✅ Connected to trusted peer", info.ID)

		// Open a new stream to the connected peer
		s, err := h.NewStream(ctx, info.ID, "/p2p/1.0.0")
		if err != nil {
			panic(err)
		}

		rw := bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))

		// Read loop
		go func() {
			for {
				str, err := rw.ReadString('\n')
				if err != nil {
					return
				}
				fmt.Printf("Received from %s: %s", s.Conn().RemotePeer(), str)
			}
		}()

		// Write loop
		go func() {
			for {
				fmt.Print(">> ")
				stdReader := bufio.NewReader(os.Stdin)
				text, _ := stdReader.ReadString('\n')
				rw.WriteString(text)
				rw.Flush()
			}
		}()
	}

	// Block forever to keep the app running
	select {}
}
