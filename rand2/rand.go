package main

import (
	crand "crypto/rand"
	"encoding/binary"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var s int64
	if err := binary.Read(crand.Reader, binary.LittleEndian, &s); err != nil {
		s = time.Now().UnixNano()
	}
	rand.Seed(s)

	var num = rand.Intn(11)
	fmt.Println(num)

	num = rand.Intn(11)
	fmt.Println(num)
}
