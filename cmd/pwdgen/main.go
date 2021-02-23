package main

import (
	"fmt"
	"os"

	"github.com/function61/gokit/app/dynversion"
	"github.com/function61/gokit/crypto/randompassword"
	"github.com/function61/gokit/os/osutil"
	"github.com/spf13/cobra"
)

func main() {
	length := 16

	cmd := &cobra.Command{
		Use:     os.Args[0],
		Short:   "Password generator",
		Version: dynversion.Version,
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			osutil.ExitIfError(generate(length))
		},
	}

	cmd.Flags().IntVarP(&length, "length", "l", length, "Length of password")

	osutil.ExitIfError(cmd.Execute())
}

func generate(length int) error {
	_, err := fmt.Fprintln(os.Stdout, randompassword.Build(randompassword.DefaultAlphabet, length))
	return err
}
