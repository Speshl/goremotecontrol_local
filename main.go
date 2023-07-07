package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/Speshl/goremotecontrol_local/client/controllers/ground"
	"github.com/Speshl/goremotecontrol_local/server"
	"golang.org/x/sync/errgroup"
)

func main() {
	listJoysticks := flag.Bool("listjoys", false, "List available joysticks")
	showJoyStats := flag.Bool("joystats", false, "Shows states of connected joysticks")

	listSerial := flag.Bool("listserial", false, "List available serial devices")
	serialPort := flag.String("serial", "COM3", "Serial Port") //windows: COM3 //linux: /dev/ttyUSB0
	baudRate := flag.Int("baudrate", 115200, "Serial baudrate")

	flag.Parse()

	if listJoysticks != nil && *listJoysticks {
		// _, err := client.GetJoysticks()
		// if err != nil {
		// 	log.Fatal(err)
		// }
	} else if showJoyStats != nil && *showJoyStats {
		// _, err := client.ShowJoyStats()
		// if err != nil {
		// 	log.Fatal(err)
		// }
	} else if listSerial != nil && *listSerial {
		err := server.GetSerialDevices()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		ctx, _ := context.WithCancel(context.Background())
		controller := ground.NewGroundControllerG27()
		controller.Start()

		errGroup, ctx := errgroup.WithContext(ctx)

		server := server.NewServer(serialPort, baudRate)
		errGroup.Go(func() error {
			return server.Start(ctx)
		})

		updateTicker := time.NewTicker(100 * time.Millisecond)
		logTicker := time.NewTicker(10 * time.Second)

		errGroup.Go(func() error {
			for {
				select {
				case <-ctx.Done():
					log.Println("Client context closed")
					return ctx.Err()
				case <-updateTicker.C:
					server.SendGroundCommand(controller.GetControlData())
				case <-logTicker.C:
					fmt.Printf("Control Status: %+v\n\n", controller.GetControlData())
				}
			}
		})

		/*time.Sleep(time.Minute * 5) //TODO: Remove
		cancelCtx()
		updateTicker.Stop()
		logTicker.Stop()*/

		err := errGroup.Wait()
		if err != nil {
			log.Fatal("server error: %w", err)
		}

		err = controller.Stop()
		if err != nil {
			log.Fatal(err)
		}
	}
}
