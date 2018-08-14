package main

import (
	"bytes"
	"log"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/ajstarks/openvg"
)

func vgfloat(num int) openvg.VGfloat {
	return openvg.VGfloat(num)
}

func timeFormat(i int) string {
	if i < 10 {
		return "0" + strconv.Itoa(i)
	}
	return strconv.Itoa(i)
}

func genClock(hour, min, sec int) {
	timeStr := timeFormat(hour) + " : " + timeFormat(min) + " : " + timeFormat(sec)

	openvg.FillColor("blue")
	openvg.Text(100, 100, timeStr, "serif", 56)
}

func main() {
	var (
		buf    bytes.Buffer
		logger = log.New(&buf, "logger: ", log.Lshortfile)
	)

	logger.Print("Log Started")

	width, height := openvg.Init()

	ticker := time.NewTicker(1 * time.Second)
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)

	openvg.Start(width, height)

	for {
		select {
		case <-ticker.C:
			hr, min, sec := time.Now().Clock()

			openvg.Background(255, 255, 255)
			genClock(hr, min, sec)

			openvg.End()
		case <-sigint:
			openvg.Finish()
			return
		}

	}
}
