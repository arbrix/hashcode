package main

import (
	"bufio"
	"os"
)

func main() {
	fi, _ := os.Open("")
	fo, _ := os.Create("output.txt")
	scanner := bufio.NewScanner(fi)
	writer := bufio.NewWriter(fo)
	defer fi.Close()
	defer fo.Close()
	defer writer.Flush()
	scanner.Scan()
	writer.Write(scanner.Bytes())
}
