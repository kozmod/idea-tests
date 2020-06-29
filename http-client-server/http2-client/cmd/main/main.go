package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/kozmod/idea-tests/http-client-server/http2-client/pkg/config"
	"github.com/kozmod/idea-tests/http-client-server/http2-client/version"

	. "github.com/kozmod/idea-tests/http-client-server/http2-client/cmd"
	_ "github.com/kozmod/idea-tests/http-client-server/http2-client/pkg"
	"github.com/kozmod/idea-tests/http-client-server/http2-client/pkg/client"
	"github.com/kozmod/idea-tests/http-client-server/http2-client/pkg/utils"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "http2 client",
		Short: "run with env args",
		Run: func(cmd *cobra.Command, args []string) {
			delay := 3 * time.Second
			log.Printf("Run with env var. Delay before start %v\n", delay)
			defaultValsCmd.Run(cmd, args)
			<-time.After(3 * time.Second)
			startEnvCmd.Run(cmd, args)
		},
	}

	defaultValsCmd = &cobra.Command{
		Use:   "dval",
		Short: "print default values",
		Run: func(cmd *cobra.Command, args []string) {
			c := config.FromEnv()
			log.Println(
				"\n" + "ENV:" +
					"\n" + fmt.Sprintf("ServerAddrEnv=%s;", config.ServerAddrEnv) +
					"\n" + fmt.Sprintf("RequestQuantityEnv=%s;", config.RequestQuantityEnv) +
					"\n" + fmt.Sprintf("RequestFrequencySec=%s;", config.RequestFrequencySecEnv) +
					"\n" + fmt.Sprintf("PostWithPayloadRtl=%s;", config.PostWithPayloadUrlEnv) +
					"\n" + "Conf val:" +
					"\n" + fmt.Sprintf("ServerAddrEnv=%s;", c.ServerAddr()) +
					"\n" + fmt.Sprintf("RequestQuantityEnv=%d;", c.RequestQuantity()) +
					"\n" + fmt.Sprintf("RequestFrequencySec=%s;", c.RequestFrequency()) +
					"\n" + fmt.Sprintf("PostWithPayloadRtl=%s;", c.PostWithPayloadUrl()) +
					"\n" + "Vars:" +
					"\n" + fmt.Sprintf("DefaultLogFilePath=%s;", DefaultLogFilePath),
			)
		},
	}

	startCmd = &cobra.Command{
		Use:     "start [addr postAddress, request quantity, frequency (seconds)]",
		Short:   "define server address and start start http2 client",
		Example: "./app start http://localhost:8080/tp 1 10",
		Args:    cobra.MinimumNArgs(3),
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

	startEnvCmd = &cobra.Command{
		Use:   "startEnv",
		Short: "start with data from env",
		Run: func(cmd *cobra.Command, args []string) {
			c := config.FromEnv()
			addr := fmt.Sprintf("%s%s",
				os.Getenv(c.ServerAddr()),
				os.Getenv(c.PostWithPayloadUrl()))
			start(addr, c.RequestQuantity(), c.RequestFrequency())
		},
	}

	postEnvCmd = &cobra.Command{
		Use:   "postEnv",
		Short: "single post to server use env",
		Run: func(cmd *cobra.Command, args []string) {
			addr := os.Getenv(config.ServerAddrEnv)
			postJson(client.New(), addr, fmt.Sprintf(`{"single":"to %s"}`, addr))
		},
	}

	postCmd = &cobra.Command{
		Use:     "post [addr postAddress,payload payload]",
		Short:   "single post with payload",
		Example: "./app post `{\"payload\":\"xxx\"}`",
		Args:    cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			addr, payload := args[0], args[1]
			if strings.TrimSpace(addr) == "" || strings.TrimSpace(payload) == "" {
				log.Fatal(errors.New(
					fmt.Sprintf("address or payload is empty:[addr=%s, payload=%s]", addr, payload)))
			}
			postJson(client.New(), addr, payload)
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
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(startEnvCmd)
	rootCmd.AddCommand(postCmd)
	rootCmd.AddCommand(postEnvCmd)
	rootCmd.AddCommand(versionCmd)
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

func postJson(client *client.H2client, postPayloadAddr string, json string) {
	client.LogPostJsonRs(postPayloadAddr, json)
}
