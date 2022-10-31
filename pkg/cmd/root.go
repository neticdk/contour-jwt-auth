package cmd

import (
	"crypto/tls"
	"fmt"
	"net"
	"os"

	"github.com/neticdk/contour-jwt-auth/pkg/server"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	listenAddr string
	logLevel   int

	rootCmd = &cobra.Command{
		Use:   "server",
		Short: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.SetLevel(log.Level(logLevel))

			lis, err := net.Listen("tcp", listenAddr)
			if err != nil {
				log.Fatalf("failed to listen: %v", err)
			}

			opts := []grpc.ServerOption{
				grpc.MaxConcurrentStreams(1 << 20),
				grpc.Creds(credentials.NewTLS(&tls.Config{
					MinVersion:   tls.VersionTLS13,
					Certificates: []tls.Certificate{},
					RootCAs:      nil,
				})),
			}
			s := server.NewServer(opts...)
			s.Serve(lis)

			return nil
		},
	}
)

// Execute root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	viper.AutomaticEnv()
	viper.SetDefault("LISTEN_ADDRESS", ":8080")

	rootCmd.Flags().StringVarP(&listenAddr, "address", "a", viper.GetString("LISTEN_ADDRESS"), "Listen address")
}
