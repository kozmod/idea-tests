package main

import (
	"fmt"
	. "github.com/kozmod/idea-tests/http-client-server/http2-client/cmd"
	_ "github.com/kozmod/idea-tests/http-client-server/http2-client/pkg"
	"github.com/kozmod/idea-tests/http-client-server/http2-client/pkg/client"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "http2 client",
	Short: "test http2 client",
	Run: func(cmd *cobra.Command, args []string) {
		start(DefaultServerAddr)
	},
}

var defaultValsCmd = &cobra.Command{
	Use:   "dval",
	Short: "print default values",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("DefaultServerAddr: " + DefaultServerAddr)
	},
}

var serverAddrCmd = &cobra.Command{
	Use:   "s [server address]",
	Short: "define server address and start start http2 client",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		start(args[0])
	},
}

func main() {
	rootCmd.AddCommand(defaultValsCmd)
	rootCmd.AddCommand(serverAddrCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func start(serverAddr string) {
	h2c := client.New(serverAddr)
	h2c.LogGet(DefaultServerAddr + "/t")
	h2c.LogPostJson(DefaultServerAddr+"/t", `{"foo":"bar"}`)
}
