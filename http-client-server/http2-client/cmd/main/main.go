package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	. "github.com/kozmod/idea-tests/http-client-server/http2-client/cmd"
	_ "github.com/kozmod/idea-tests/http-client-server/http2-client/pkg"
	"github.com/kozmod/idea-tests/http-client-server/http2-client/pkg/client"
	"github.com/kozmod/idea-tests/http-client-server/http2-client/pkg/utils"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "http2 client",
	Short: "run with env args",
	Run: func(cmd *cobra.Command, args []string) {
		var delay time.Duration = 3 * time.Second
		log.Printf("Run with env var. Delay before start %v\n", delay)
		defaultValsCmd.Run(cmd, args)
		<-time.After(3 * time.Second)
		startEnvCmd.Run(cmd, args)
	},
}

var defaultValsCmd = &cobra.Command{
	Use:   "dval",
	Short: "print default values",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println(
			"\n" + fmt.Sprintf("ServerAddrEnv=%s; val=%s", ServerAddrEnv, os.Getenv(ServerAddrEnv)) +
				"\n" + fmt.Sprintf("RequestQuantityEnv=%s; val=%s", RequestQuantityEnv, os.Getenv(RequestQuantityEnv)) +
				"\n" + fmt.Sprintf("RequestFrequencySec=%s; val=%s", RequestFrequencySec, os.Getenv(RequestFrequencySec)) +
				"\n" + fmt.Sprintf("DefaultLogFilePath=%s;", DefaultLogFilePath))
	},
}

var startCmd = &cobra.Command{
	Use:   "start [server address, request quantity, frequency (seconds)]",
	Short: "define server address and start start http2 client",
	Args:  cobra.MinimumNArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		q, err := strconv.Atoi(args[1])
		if err != nil {
			log.Fatal(err)
		}
		f, err := utils.AsSeconds(args[2])
		if err != nil {
			log.Fatal(err)
		}
		start(args[0], q, f)
	},
}

var startEnvCmd = &cobra.Command{
	Use:   "startEnv",
	Short: "start with data from env",
	Run: func(cmd *cobra.Command, args []string) {
		q, err := strconv.Atoi(os.Getenv(RequestQuantityEnv))
		if err != nil {
			log.Fatal(err)
		}
		f, err := utils.AsSeconds(os.Getenv(RequestFrequencySec))
		if err != nil {
			log.Fatal(err)
		}
		start(os.Getenv(ServerAddrEnv), q, f)
	},
}

var postEnvCmd = &cobra.Command{
	Use:   "postEnv",
	Short: "single post to server use env",
	Run: func(cmd *cobra.Command, args []string) {
		addr := os.Getenv(ServerAddrEnv)
		postJson(client.New(), addr, fmt.Sprintf(`{"single":"to %s"}`, addr))
	},
}

func main() {
	rootCmd.AddCommand(defaultValsCmd)
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(startEnvCmd)
	rootCmd.AddCommand(postEnvCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func start(serverAddr string, quantity int, frequency time.Duration) {
	h2c := client.New()
	for {
		for i := 1; i <= quantity; i++ {
			postJson(h2c, serverAddr, fmt.Sprintf(`{"iteration":"%d"}`, i))
		}
		<-time.After(frequency)
	}
}

func postJson(client *client.H2client, serverAddr string, json string) {
	client.LogPostJsonRs(serverAddr+"/tp", json)
}
