// gpbkdf2 project main.go
package main

import (
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"
)

type options struct {
	Help            bool   `short:"h" long:"help" description:"Shows this help message"`
	Passphrase      string `short:"p" long:"passphrase" description:"Passphrase to encode into PBKDF2"`
	Salt            string `short:"s" long:"salt" description:"Salt"`
	DigestAlgorithm string `short:"d" long:"digest-algorithm" default:"sha512" description:"Digest algorithm, eg. sha512"`
	DigestRounds    int    `short:"r" long:"digest-rounds" default:"6000" description:"Number of rounds of a digest algorithm eg. 6000"`
}

func main() {
	var opts options
	p := flags.NewParser(&opts, flags.Default&^flags.HelpFlag)
	_, err := p.Parse()
	if err != nil {
		fmt.Printf("fail to parse args: %v", err)
		os.Exit(1)
	}
	if opts.Help {
		p.WriteHelp(os.Stdout)
		os.Exit(0)
	}

	println(EncodePasswordToPBKDF2(opts.Passphrase, opts.Salt, opts.DigestRounds, 16, opts.DigestAlgorithm))
}
