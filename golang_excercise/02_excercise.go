package main

import "fmt"

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
)

func main() {
	fmt.Printf("%-5s %-20d %-64b\n", "Unit", "Decimal Value", "Binary Representation")
	fmt.Println("-------------------------------------------------------------------")
	fmt.Printf("%-5s %-20d %-64b\n", "KB", KB, KB)
	fmt.Printf("%-5s %-20d %-64b\n", "MB", MB, MB)
	fmt.Printf("%-5s %-20d %-64b\n", "GB", GB, GB)
	fmt.Printf("%-5s %-20d %-64b\n", "TB", TB, TB)
	fmt.Printf("%-5s %-20d %-64b\n", "PB", PB, PB)
	fmt.Printf("%-5s %-20d %-64b\n", "EB", EB, EB)
}
