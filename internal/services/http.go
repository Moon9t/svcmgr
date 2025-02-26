package services

import (
    "fmt"
    "os/exec"
    "runtime"

    "github.com/moon9t/svcmgr/internal/config"
    "github.com/rs/zerolog/log"
)

type HTTPHandler struct{}

func (h *HTTPHandler) Connect(s config.Service, _ string) error {
    url := fmt.Sprintf("http://%s:%d", s.Host, s.Port)
    if s.Port == 443 {
        url = fmt.Sprintf("https://%s", s.Host)
    }

    var cmd *exec.Cmd
    switch runtime.GOOS {
    case "darwin":
        cmd = exec.Command("open", url)
    case "windows":
        cmd = exec.Command("cmd", "/c", "start", url)
    default:
        cmd = exec.Command("xdg-open", url)
    }

    log.Debug().Str("url", url).Msg("Opening cosmic gateway")
    return cmd.Start()
}