package main

import (
	"bufio"
	"github.com/ajstarks/openvg"
	"os"
)


func main(){
	width, height := openvg.Init()
	w2 := openvg.VGfloat(width/2)
	h2 := openvg.VGfloat(height/2)
	w := openvg.VGfloat(width)

	stops := []openvg.Offcolor{
		{0.0, openvg.RGB{44, 100, 232}, 1.0},
		{0.5, openvg.RGB{22, 50, 151}, 1.0},
		{1.0, openvg.RGB{88, 200, 255}, 1.0},
	}

	openvg.Start(width, height)
	openvg.BackgroundColor("black")
	openvg.FillRadialGradient(w2,0,w2,w2,w*.5,stops)
	openvg.Circle(w2,0,w)
	openvg.FillColor("white")
	openvg.TextMid(w2,h2,"DID IT!!!!!", "serif", width/10)
	openvg.SaveEnd("hvg.raw")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	openvg.Finish()
}

