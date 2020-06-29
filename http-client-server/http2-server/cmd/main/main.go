package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/kozmod/idea-tests/http-client-server/http2-server/version"

	. "github.com/kozmod/idea-tests/http-client-server/http2-server/cmd"
	_ "github.com/kozmod/idea-tests/http-client-server/http2-server/pkg"
	"github.com/kozmod/idea-tests/http-client-server/http2-server/pkg/server"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "http2 client",
		Short: "start http2 server",
		Run: func(cmd *cobra.Command, args []string) {
			http2server := server.ConfigureAndServe(tryReadPort())
			log.Fatal(http2server.ListenAndServe())
		},
	}

	defaultValsCmd = &cobra.Command{
		Use:   "dval",
		Short: "print default values",
		Run: func(cmd *cobra.Command, args []string) {
			log.Println(fmt.Sprintf("DefaultServerPort=%s;", DefaultServerPort))
		},
	}

	portCmd = &cobra.Command{
		Use:   "p [server port]",
		Short: "define port and start start http2 server",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			server.ConfigureAndServe(":" + args[0])
		},
	}
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "version of app (+ golang version and build time)",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(fmt.Sprintf("Go Version:         %s", runtime.Version()))
			fmt.Println(fmt.Sprintf("Build Version:      %s", version.Version))
			fmt.Println(fmt.Sprintf("Build Version Time: %s", version.Time))
		},
	}
)

func main() {
	rootCmd.AddCommand(defaultValsCmd)
	rootCmd.AddCommand(portCmd)
	rootCmd.AddCommand(versionCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func tryReadPort() string {
	port := ":"
	if conf, err := server.ReadConfig("config", "yml", "./etc/config"); err == nil {
		port = port + conf.ServerAddr.Port
	} else {
		log.Println(err)
		port = port + DefaultServerPort
	}
	return port
}
