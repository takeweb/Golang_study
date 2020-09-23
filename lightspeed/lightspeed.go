package main

import "fmt"

const lightsSpeed = 299792 // km/秒

func main() {
	var distance = 56000000 // km
	fmt.Println(calcLightSpeed(distance), "秒")

	distance = 401000000 // km
	fmt.Println(calcLightSpeed(distance), "秒")
}

func calcLightSpeed(distance int) int {
	return distance / lightsSpeed
}
