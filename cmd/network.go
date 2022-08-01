/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p/p2p/protocol/ping"
	"github.com/multiformats/go-multiaddr"
	"github.com/spf13/cobra"
)

// networkCmd represents the network command
var networkCmd = &cobra.Command{
	Use:   "network",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var node host.Host
		var p *peer.AddrInfo

		listen := cmd.Flag("listen").Value.String()
		connect := cmd.Flag("connect").Value.String()

		if listen != "" {
			port, err := strconv.ParseUint(listen, 10, 16)

			if err != nil {
				log.Fatalf("invalid listen port: %s", listen)
			}

			node, err = libp2p.New(
				libp2p.ListenAddrStrings(fmt.Sprintf("/ip4/127.0.0.1/tcp/%d", port)),
				libp2p.Ping(false),
			)
			if err != nil {
				log.Fatalf("failed to create libp2p node: %s", err)
			}

			pingService := &ping.PingService{Host: node}
			node.SetStreamHandler(ping.ID, pingService.PingHandler)

			peerInfo := peer.AddrInfo{
				ID:    node.ID(),
				Addrs: node.Addrs(),
			}

			if connect != "" {
				addr, err := multiaddr.NewMultiaddr(connect)

				if err != nil {
					log.Fatalf("invalid connect address: %s", addr)
				}

				p, err = peer.AddrInfoFromP2pAddr(addr)

				if err != nil {
					log.Fatalf("invalid connect address: %s", err)
				}

				if err := node.Connect(context.Background(), *p); err != nil {
					log.Fatalf("failed to connect to %s: %s", p.ID, err)
				}

				fmt.Println("Connecting to:", addr)
			}

			// print the node's listening addresses
			addrs, err := peer.AddrInfoToP2pAddrs(&peerInfo)
			if err != nil {
				log.Fatalf("failed to convert peer.AddrInfo: %s", err)
			}
			fmt.Println("libp2p node address:", addrs[0])

			ch := make(chan os.Signal, 1)

			signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
			<-ch
			fmt.Println("Received signal, shutting down...")

			// shut the node down
			if err := node.Close(); err != nil {
				log.Fatalf("failed to close libp2p node: %s", err)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(networkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	networkCmd.PersistentFlags().Uint16P("listen", "l", 3000, "listen port")
	networkCmd.PersistentFlags().StringP("connect", "c", "", "connect to address")
	networkCmd.PersistentFlags().StringP("data", "d", "", "data to connect")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// networkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
