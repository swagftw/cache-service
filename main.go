package main

import (
	"fmt"
	"os/exec"
)

func main() {
	out, err := exec.Command("").Output()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(out))
}
