package main

import (
	"flag"
	"log"

	"github.com/craiggwilson/osquery-nix-extension/internal"
)

func main() {
	var args internal.Args

	flag.StringVar(&args.Socket, "socket", "", "Path to the extensions UNIX domain socket")
	flag.IntVar(&args.Timeout, "timeout", 3, "Seconds to wait for autoloaded extensions")
	flag.IntVar(&args.Interval, "interval", 3, "Seconds delay between connectivity checks")
	flag.StringVar(&args.Closure, "closure", "/run/current-system", "Nix closure to list packages from")
	flag.Parse()

	err := internal.Run(args)
	if err != nil {
		log.Fatal(err)
	}
}
