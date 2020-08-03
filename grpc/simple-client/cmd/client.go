package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/kozmod/idea-tests/grpc/proto/generated/api"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var (
	addr               = "127.0.0.1:8080"
	defaultQuantity    = 10
	randomDurationFunc = func() time.Duration {
		return time.Duration(rand.Intn(5)) * time.Second
	}
)

func main() {
	rootCmd.AddCommand(startCmd)
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

var (
	rootCmd = &cobra.Command{
		Use:   "grpc client",
		Short: "run with 1 send",
		Run: func(cmd *cobra.Command, args []string) {
			start(defaultQuantity)
		},
	}
	startCmd = &cobra.Command{
		Use:     "start [request quantity]",
		Short:   "define server address and start start grpc client",
		Example: "./app start http://localhost:8080/tp 1 10",
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			q, err := strconv.Atoi(args[0])
			if err != nil {
				log.Fatal(err)
			}
			start(q)

		},
	}
)

func start(quantity int) {
	cctx, cancel := context.WithCancel(context.TODO())
	var wg sync.WaitGroup
	wg.Add(quantity)
	for i := 0; i < quantity; i++ {
		go func(i int) {
			ctx, _ := context.WithTimeout(cctx, 1*time.Second)
			//time.Sleep(randomDurationFunc())
			time.Sleep(5 * time.Second)
			client, connection := NewBidiServiceClient(addr, grpc.WithInsecure())
			defer connection.Close()
			rs, err := client.Execute(ctx, &api.Rq{Uid: strconv.Itoa(i), Val: fmt.Sprintf("val-%d", i)})
			if err != nil {
				log.Println(fmt.Sprintf("execution error: %v", err))
			} else {
				log.Println(fmt.Sprintf("resp: %v", rs))
			}
			wg.Done()
		}(i)
	}
	go func() {
		time.Sleep(20 * time.Second)
		cancel()
	}()
	wg.Wait()
}

func NewBidiServiceClient(addr string, opts ...grpc.DialOption) (api.SimpleServiceClient, *grpc.ClientConn) {
	connection, err := grpc.Dial(addr, opts...)
	if err != nil {
		log.Fatal("Fail to connect - \n", err)
	}
	return api.NewSimpleServiceClient(connection), connection
}
