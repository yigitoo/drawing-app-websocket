package main

import "fmt"

func GenerateColor(RGB []uint8) string {
	R := RGB[0]
	G := RGB[1]
	B := RGB[2]
	return fmt.Sprintf("#%02x%02x%02x", R, G, B)
}
