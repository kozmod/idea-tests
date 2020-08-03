package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	pb "github.com/kozmod/idea-tests/grpc/proto/generated/api"
	"github.com/spf13/cobra"
	"golang.org/x/exp/rand"
	"google.golang.org/grpc"
)

var (
	port               = ":8080"
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
			startCmd.Run(cmd, args)
		},
	}
	startCmd = &cobra.Command{
		Use:     "start [request quantity]",
		Short:   "define server address and start start grpc client",
		Example: "./app start http://localhost:8080/tp 1 10",
		Run: func(cmd *cobra.Command, args []string) {
			start()
		},
	}
)

func start() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterSimpleServiceServer(grpcServer, &timeoutServer{})

	// graceful shutdown
	sch := make(chan os.Signal, 1)
	signal.Notify(sch, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		for range sch {
			// sig is a ^Conf, handle it
			log.Println("shutting down gRPC service...")
			grpcServer.GracefulStop()
		}
	}()

	log.Printf("server start on port: %v\n", port)
	log.Fatal(grpcServer.Serve(listener))
}

type timeoutServer struct {
}

func (s *timeoutServer) Execute(ctx context.Context, rq *pb.Rq) (*pb.Rs, error) {
	log.Printf("get rq: %v\n", rq)
	time.Sleep(randomDurationFunc())
	rs := &pb.Rs{Uid: rq.Uid, Val: rq.Val}
	return rs, nil
}
