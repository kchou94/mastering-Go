package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	buffer := []byte("Hello, world\n")

	f1, err := os.Create("/tmp/file1")
	if err != nil {
		fmt.Println("Cannot create file", err)
		return
	}
	defer f1.Close()
	fmt.Fprintf(f1, string(buffer))

	f2, err := os.Create("/tmp/file2")
	if err != nil {
		fmt.Println("Cannot create file", err)
		return
	}
	defer f2.Close()
	n, err := f2.WriteString(string(buffer))
	fmt.Printf("Wrote %d bytes\n", n)

	f3, err := os.Create("/tmp/file3")
	if err != nil {
		fmt.Println("Cannot create file", err)
		return
	}
	defer f3.Close()
	w := bufio.NewWriter(f3)
	n, err = w.WriteString(string(buffer))
	fmt.Printf("Wrote %d bytes\n", n)
	w.Flush()

	f := "/tmp/file4"
	f4, err := os.Create(f)
	if err != nil {
		fmt.Println("Cannot create file", err)
		return
	}
	defer f4.Close()
	for i := 0; i < 5; i++ {
		n, err = io.WriteString(f4, string(buffer))
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Wrote %d bytes\n", n)
	}
	// Append to file
	f4, err = os.OpenFile(f, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f4.Close()
	// Write() needs a byte slice
	n, err = f4.Write([]byte("\nNew line"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("wrote %d bytes\n", n)
}
