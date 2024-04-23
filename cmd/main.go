package main

import (
	"github.com/Skisocks/mapper/pkg/cmd"
)

func main() {
	rootCmd := cmd.Main()
	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}
