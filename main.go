package main

import (
	"fmt"
	"log"
	"strings"

	"fyne.io/systray"
	probing "github.com/prometheus-community/pro-bing"
)

var pinger *probing.Pinger

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetTitle("Initializing...")

	mLatencyLabel := systray.AddMenuItem("Latency", "Current Latency")
	mLatencyLabel.Disable()

	systray.AddSeparator()
	mQuit := systray.AddMenuItem("Quit", "Quit the app")

	var err error
	pinger, err = probing.NewPinger("8.8.8.8")
	if err != nil {
		panic(err)
	}
	pinger.SetLogger(nil)

	// Update tray title
	pinger.OnRecv = func(pkt *probing.Packet) {
		latency := pkt.Rtt.Milliseconds()

		switch {
		case latency < 50:
			systray.SetTitle(fmt.Sprintf("🟢 %d ms", latency))
		case latency < 75:
			systray.SetTitle(fmt.Sprintf("🟠 %d ms", latency))
		default:
			systray.SetTitle(fmt.Sprintf("🔴 %d ms", latency))
		}

		mLatencyLabel.SetTitle(fmt.Sprintf("%s: %d ms", pinger.IPAddr().String(), latency))
	}

	pinger.OnSendError = func(_ *probing.Packet, err error) {
		systray.SetTitle("🔴 Network unavailable")

		if strings.Contains(err.Error(), "sendto") {
			parts := strings.Split(err.Error(), ": ")
			if len(parts) >= 2 {
				mLatencyLabel.SetTitle(fmt.Sprintf("8.8.8.8: %s", parts[len(parts)-1]))
				return
			}
		}

		mLatencyLabel.SetTitle(fmt.Sprintf("8.8.8.8: %s", err.Error()))
	}

	// Running the ping
	go func() {
		if err := pinger.Run(); err != nil {
			log.Println("Pinger has crashed")
			systray.Quit()
		}
	}()

	// mQuit click
	go func() {
		<-mQuit.ClickedCh
		systray.Quit()
	}()
}

func onExit() {
	pinger.Stop()
}
