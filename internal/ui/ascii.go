package ui

import (
	"fmt"
	"os"
	"time"
)

var (
	version = "1.0.0-stellar"
	build   = "20240115-lunar"
)

func Banner() string {
	return fmt.Sprintf(`
    ███████╗████████╗███████╗███╗   ██╗███████╗
    ██╔════╝╚══██╔══╝██╔════╝████╗  ██║██╔════╝
    ███████╗   ██║   █████╗  ██╔██╗ ██║█████╗  
    ╚════██║   ██║   ██╔══╝  ██║╚██╗██║██╔══╝  
    ███████║   ██║   ███████╗██║ ╚████║███████╗
    ╚══════╝   ╚═╝   ╚══════╝╚═╝  ╚═══╝╚══════╝
    
    Stellar Service Manager %s
    Build: %s | %s
    
    `, version, build, time.Now().Format("2006-01-02 15:04:05 MST"))
}

func SuccessIcon() string {
	if os.Getenv("TERM") == "xterm-256color" {
		return "✨"
	}
	return "[OK]"
}

func ErrorArt() string {
	return `
    ( •_•)
    ( •_•)>⌐■-■
    (⌐■_■)
    ERROR!`
}
