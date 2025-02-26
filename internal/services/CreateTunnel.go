package services

import (
	"context"
	"io"

	"fmt"

	"net"
)

func CreateTunnel(ctx context.Context, localPort int, config TunnelConfig) error {

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", localPort))

	if err != nil {

		return err

	}

	defer listener.Close()

	for {

		select {

		case <-ctx.Done():

			return ctx.Err()

		default:

			conn, err := listener.Accept()

			if err != nil {

				return err

			}

			go handleConnection(conn, config)

		}

	}

}

func handleConnection(conn net.Conn, config TunnelConfig) {

	defer conn.Close()

	remoteConn, err := net.Dial("tcp", net.JoinHostPort(config.Host, fmt.Sprintf("%d", config.Port)))

	if err != nil {

		return

	}

	defer remoteConn.Close()

	go func() {

		_, _ = io.Copy(remoteConn, conn)

	}()

	_, _ = io.Copy(conn, remoteConn)

}
