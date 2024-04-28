package main

import (
	"math/rand/v2"
)

type Color struct {
	r uint8
	g uint8
	b uint8
	a uint8
}

var Colors = map[string][]uint8{
	"Black":  {0, 0, 0},
	"White":  {255, 255, 255},
	"Red":    {255, 0, 0},
	"Green":  {0, 128, 0},
	"Lime":   {0, 255, 0},
	"Blue":   {0, 0, 255},
	"Yellow": {255, 255, 0},
	"Cyan":   {0, 255, 255},
	"Brown":  {150, 75, 0},
}

func RandomColor() []uint8 {
	return []uint8{
		uint8(rand.IntN(256)),
		uint8(rand.IntN(256)),
		uint8(rand.IntN(256)),
	}
}

func InitColor() *Color {
	color := new(Color)

	color.r = 0
	color.g = 0
	color.b = 0
	color.a = 0

	return color
}

func NewColor(red, green, blue, alpha uint8) *Color {
	return &Color{
		r: red,
		g: green,
		b: blue,
		a: alpha,
	}
}

func (c *Color) SetColor(red, green, blue uint8) {
	c.r = red
	c.g = green
	c.b = blue
}

func (c *Color) SetColorWithAlpha(red, green, blue, alpha uint8) {
	c.SetColor(red, green, blue)
	c.a = alpha
}

func (c *Color) GetR() uint8 {
	return c.r
}

func (c *Color) GetG() uint8 {
	return c.g
}

func (c *Color) GetB() uint8 {
	return c.b
}

func (c *Color) GetAlpha() uint8 {
	return c.a
}

func (c *Color) SetR(red uint8) {
	c.r = red
}

func (c *Color) SetG(green uint8) {
	c.g = green
}

func (c *Color) SetB(blue uint8) {
	c.b = blue
}

func (c *Color) SetAlpha(alpha uint8) {
	c.a = alpha
}
