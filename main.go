package main

import (
	"fmt"
	"log"
	"runtime"
	"strings"

	"fyne.io/systray"
	probing "github.com/prometheus-community/pro-bing"

	"ipingtray/icons"
)

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetTitle("Initializing...")
	systray.SetIcon(icons.White)

	mLatencyLabel := systray.AddMenuItem("Latency", "Current Latency")
	mLatencyLabel.Disable()

	systray.AddSeparator()
	mQuit := systray.AddMenuItem("Quit", "Quit the app")

	// Update tray title
	pinger.OnRecv = func(pkt *probing.Packet) {
		latency := pkt.Rtt.Milliseconds()
		icon := icons.Red

		switch {
		case latency < 50:
			icon = icons.Green
		case latency < 75:
			icon = icons.Orange
		}

		updateTray(fmt.Sprintf("%d ms", latency), icon)
		mLatencyLabel.SetTitle(fmt.Sprintf("%s: %d ms", pinger.IPAddr().String(), latency))
	}

	pinger.OnSendError = func(_ *probing.Packet, err error) {
		updateTray("Network Unavailable", icons.Red)

		if strings.Contains(err.Error(), "sendto") {
			parts := strings.Split(err.Error(), ": ")
			if len(parts) >= 2 {
				mLatencyLabel.SetTitle(fmt.Sprintf("%s: %s", pinger.Addr(), parts[len(parts)-1]))
				return
			}
		}

		mLatencyLabel.SetTitle(fmt.Sprintf("%s: %s", pinger.Addr(), err.Error()))
	}

	// Running the ping
	go func() {
		if err := pinger.Run(); err != nil {
			log.Printf("pinger has crashed: %v\n", err)
			systray.Quit()
		}
	}()

	// mQuit click
	go func() {
		<-mQuit.ClickedCh
		systray.Quit()
	}()
}

func updateTray(title string, icon []byte) {
	systray.SetTitle(title)
	systray.SetIcon(icon)

	if runtime.GOOS == "windows" {
		systray.SetTooltip(title)
	}
}

func onExit() {
	pinger.Stop()
}
