package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kozmod/idea-tests/http-client-server/http2-server/pkg/utils"

	. "github.com/kozmod/idea-tests/http-client-server/http2-server/cmd"
	_ "github.com/kozmod/idea-tests/http-client-server/http2-server/pkg"
	"github.com/kozmod/idea-tests/http-client-server/http2-server/pkg/server"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "http2 client",
	Short: "start http2 server",
	Run: func(cmd *cobra.Command, args []string) {
		http2server := server.ConfigureAndServe(DefaultServerPort)
		log.Fatal(http2server.ListenAndServe())
	},
}

var defaultValsCmd = &cobra.Command{
	Use:   "dval",
	Short: "print default values",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println(fmt.Sprintf("DefaultServerPort=%s;", DefaultServerPort))
	},
}

var handleFuncMapCmd = &cobra.Command{
	Use:   "hmap",
	Short: "handle function map",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println(fmt.Sprintf("Handle function map:\n%s", utils.AsString(server.HandleFunctionMap)))
	},
}

var portCmd = &cobra.Command{
	Use:   "p [server port]",
	Short: "define port and start start http2 server",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		server.ConfigureAndServe(":" + args[0])
	},
}

func main() {
	rootCmd.AddCommand(defaultValsCmd)
	rootCmd.AddCommand(portCmd)
	rootCmd.AddCommand(handleFuncMapCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
