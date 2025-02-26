package services

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/moon9t/svcmgr/internal/config"
	"github.com/rs/zerolog/log"
)

type RedisHandler struct{}

func (h *RedisHandler) Connect(s config.Service, pwd string) error {
	args := []string{"-h", s.Host, "-p", fmt.Sprintf("%d", s.Port)}
	if pwd != "" {
		args = append(args, "-a", pwd)
	}

	cmd := exec.Command("redis-cli", args...)
	log.Debug().Str("command", cmd.String()).Msg("Launching Redis CLI")

	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
