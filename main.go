package main

import (
	"fmt"
	"os"

	config "github.com/hashicorp/terraform-config-inspect/tfconfig"
)

func printProviderVersion(file string) error {
	module, diag := config.LoadModule(file)
	if diag != nil && diag.HasErrors() {
		return diag
	}

	fmt.Println("Required Providers:")
	for name, provider := range module.RequiredProviders {
		fmt.Printf("- Name:     %s\n", name)
		fmt.Printf("  Directry: %s\n", file)
		fmt.Printf("  Source:   %s\n", provider.Source)
		fmt.Printf("  Version:  %s\n", provider.VersionConstraints)
	}
	return nil
}

func run() error {
	if len(os.Args) != 2 {
		return fmt.Errorf("usage: %s <file>", os.Args[0])
	}

	return printProviderVersion(os.Args[1])
}

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}
