package main

import (
	"fmt"
	"github.com/8Monkeys/lib/id"
	"github.com/spf13/cobra"
)

var (
	hashCommand = &cobra.Command{
		Use:   "monkid",
		Short: "monkid creates hashes for various object types",
		Run: func(cmd *cobra.Command, args []string) {
			for file := range args {
				fmt.Printf("%s\n", id.FromFile(file))
			}
		},
	}
)

func main() {
	hashCommand.Execute()
}
