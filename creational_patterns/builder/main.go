package main

import (
	"builder/structs"
	"fmt"
)

func main() {
	director := structs.Director{}

	gamingBuilder := structs.GamingPCBuilder{}
	director.MakePC(&gamingBuilder, "AMD Ryzen 7 9800X3D", "NVIDIA GeForce RTX 5080", 32, 1024)
	gamingPC := gamingBuilder.GetPC()

	fmt.Println(gamingPC.CPU)
}
