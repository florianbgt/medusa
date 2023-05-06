package gcode

import (
	"regexp"
	"strconv"

	"github.com/tarm/serial"
)

type GCodeSender struct {
	MinX   float64
	MaxX   float64
	MinY   float64
	MaxY   float64
	MinZ   float64
	MaxZ   float64
	Serial *serial.Port
}

type Position struct {
	X float64
	Y float64
	Z float64
}

func NewGCodeSender(name string, baud int) *GCodeSender {
	config := &serial.Config{Name: name, Baud: baud}
	s, err := serial.OpenPort(config)
	if err != nil {
		panic(err)
	}
	gcodeSender := &GCodeSender{
		MinX:   0,
		MaxX:   235,
		MinY:   0,
		MaxY:   235,
		MinZ:   0,
		MaxZ:   250,
		Serial: s,
	}
	return gcodeSender
}

func (g *GCodeSender) Close() {
	g.Serial.Close()
}

func (g *GCodeSender) Send(command string) string {
	_, err := g.Serial.Write([]byte(command + "\n"))
	if err != nil {
		panic(err)
	}

	buffer := make([]byte, 128)
	n, err := g.Serial.Read(buffer)
	if err != nil {
		panic(err)
	}

	return string(buffer[:n])
}

func (g *GCodeSender) GetCurrentPosition(retried bool) Position {
	out := g.Send("M114")
	regex := regexp.MustCompile(`X:(-?[0-9.]+) Y:(-?[0-9.]+) Z:(-?[0-9.]+) E:0.0`)

	matches := regex.FindStringSubmatch(out)

	// retry once
	// have not figure out why yet but M114 sometimes return weird output
	if len(matches) < 4 && !retried {
		return g.GetCurrentPosition(true)
	}

	x := matches[1]
	y := matches[2]
	z := matches[3]

	posX, _ := strconv.ParseFloat(x, 64)
	posY, _ := strconv.ParseFloat(y, 64)
	posZ, _ := strconv.ParseFloat(z, 64)

	position := Position{X: posX, Y: posY, Z: posZ}

	return position
}

func (g *GCodeSender) RelativeMove(x float64, y float64, z float64) {
	current_position := g.GetCurrentPosition(false)

	g.Send("G90") // set sbsolute mode

	if current_position.X+x > g.MaxX {
		x = g.MaxX
	} else if current_position.X+x < g.MinX {
		x = g.MinX
	} else {
		x = current_position.X + x
	}

	if current_position.Y+y > g.MaxY {
		y = g.MaxY
	} else if current_position.Y+y < g.MinY {
		y = g.MinY
	} else {
		y = current_position.Y + y
	}

	if current_position.Z+z > g.MaxZ {
		z = g.MaxZ
	} else if current_position.Z+z < g.MinZ {
		z = g.MinZ
	} else {
		z = current_position.Z + z
	}

	g.Send(`G0 X` + strconv.FormatFloat(x, 'f', 2, 64) + ` Y` + strconv.FormatFloat(y, 'f', 2, 64) + ` Z` + strconv.FormatFloat(z, 'f', 2, 64))
}
