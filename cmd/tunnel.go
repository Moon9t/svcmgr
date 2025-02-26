package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/moon9t/svcmgr/internal/config"
	"github.com/moon9t/svcmgr/internal/services"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	localPort  int
	remoteHost string
	remotePort int
)

var tunnelCmd = &cobra.Command{
	Use:   "tunnel [service]",
	Short: "Create cosmic tunnel connection",
	Long:  `Establish secure tunnel through celestial gateways`,
	Args:  cobra.ExactArgs(1),
	Run:   runTunnel,
}

func init() {
	tunnelCmd.Flags().IntVarP(&localPort, "local-port", "l", 8080,
		"Local port for cosmic access")
	tunnelCmd.Flags().StringVarP(&remoteHost, "remote-host", "r", "localhost",
		"Remote host in the target constellation")
	tunnelCmd.Flags().IntVarP(&remotePort, "remote-port", "p", 80,
		"Remote port in the target constellation")
	rootCmd.AddCommand(tunnelCmd)
}

func runTunnel(cmd *cobra.Command, args []string) {
	serviceName := args[0]
	service, err := config.LoadService(serviceName)
	if err != nil {
		log.Fatal().Err(err).Msg("Service not found in cosmic registry")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Handle cosmic signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChan
		log.Info().Msg("Closing cosmic tunnel...")
		cancel()
	}()

	log.Info().Msgf("🌌 Tunnel active: localhost:%d -> %s:%d via %s",
		localPort, remoteHost, remotePort, service.Name)

	tunnelConfig := services.TunnelConfig{
		Host: remoteHost,
		Port: remotePort,
	}

	if err := services.CreateTunnel(ctx, localPort, tunnelConfig); err != nil {
		log.Fatal().Err(err).Msg("Cosmic tunnel collapsed")
	}
}
