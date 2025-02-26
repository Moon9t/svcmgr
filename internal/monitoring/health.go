package monitoring

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/moon9t/svcmgr/internal/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/rs/zerolog/log"
)

var (
	serviceStatus = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "cosmic_service_status",
			Help: "Current status of monitored services",
		},
		[]string{"service"},
	)
)

func init() {
	prometheus.MustRegister(serviceStatus)
}

func StartHealthChecks(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			services, _ := config.LoadServices()
			for _, s := range services {
				status := checkServiceHealth(config.Service{
					Name: s.Name,
					Type: s.Type,
					Host: s.Host,
					Port: s.Port,
				})
				serviceStatus.WithLabelValues(s.Name).Set(status)
				log.Debug().
					Str("service", s.Name).
					Float64("status", status).
					Msg("Cosmic health update")
			}
		}
	}
}

func checkServiceHealth(s config.Service) float64 {
	// Implementation varies by service type
	switch s.Type {
	case "http":
		return checkHTTPHealth(s)
	case "ssh":
		return checkSSHHealth(s)
	default:
		return 0
	}
}

func checkHTTPHealth(s config.Service) float64 {
	url := fmt.Sprintf("http://%s:%d/health", s.Host, s.Port)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		return 0
	}
	return 1
}

func checkSSHHealth(_ config.Service) float64 {
	// Placeholder implementation for SSH health check
	// You should replace this with actual SSH health check logic
	return 1
}
