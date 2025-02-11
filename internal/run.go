package internal

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/craiggwilson/osquery-nix-extension/nixpkg"
	"github.com/osquery/osquery-go"
	"github.com/osquery/osquery-go/plugin/table"
)

type Args struct {
	Socket   string
	Timeout  int
	Interval int
	Closure  string
}

func Run(args Args) error {
	if args.Socket == "" {
		log.Fatalln("Missing required --socket argument")
	}
	serverTimeout := osquery.ServerTimeout(
		time.Second * time.Duration(args.Timeout),
	)
	serverPingInterval := osquery.ServerPingInterval(
		time.Second * time.Duration(args.Interval),
	)

	server, err := osquery.NewExtensionManagerServer(
		"osquery_nix_extension",
		args.Socket,
		serverTimeout,
		serverPingInterval,
	)
	if err != nil {
		return fmt.Errorf("creating extension: %w", err)
	}

	server.RegisterPlugin(table.NewPlugin("nix_packages", schema(), generateData(args.Closure)))
	if err := server.Run(); err != nil {
		return fmt.Errorf("running server: %w", err)
	}

	return nil
}

func schema() []table.ColumnDefinition {
	return []table.ColumnDefinition{
		table.TextColumn("name"),
		table.TextColumn("version"),
		table.TextColumn("store_path"),
	}
}

func generateData(closure string) func(context.Context, table.QueryContext) ([]map[string]string, error) {
	return func(context.Context, table.QueryContext) ([]map[string]string, error) {
		pkgs, err := nixpkg.ListFromClosure(closure)
		if err != nil {
			return nil, fmt.Errorf("listing nix packages from closure %q: %w", closure, err)
		}

		var result []map[string]string

		for pkg := range pkgs {
			result = append(result, map[string]string{
				"name":       pkg.Name,
				"version":    string(pkg.Version),
				"store_path": string(pkg.StorePath),
			})
		}

		return result, nil
	}
}
