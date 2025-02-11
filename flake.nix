{
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };
  outputs =
    {
      nixpkgs,
      flake-utils,
      ...
    }:
    flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
        mainPkg = pkgs.buildGo123Module rec {
          pname = "osquery-nix-extension";
          version = "unversioned";
          vendorHash = "sha256-B3XfbMPMKvSz4rRfzYvLLEg5NObanpq69t/5Vv/JMjE=";

          src = ./.;
          ldflags = [
            "-s"
            "-w"
            "-X cmd.version=${version}"
            "-X cmd.builtBy=flake"
          ];
          doCheck = false;
          modRoot = "./.";
          subPackages = [ "cmd/osquery-nix-extension" ];

          fixupPhase = ''
            cp -r $out/bin/osquery-nix-extension $out/bin/osquery-nix-extension.ext
          '';
        };
      in
      {
        packages = rec {
          osquery-nix-extension = mainPkg;
          default = osquery-nix-extension;
        };

        apps = rec {
          osquery-nix-extension = {
            type = "app";
            program = "${mainPkg}/bin/osquery-nix-extension";
          };
          default = osquery-nix-extension;
        };

        devShells.default = pkgs.mkShellNoCC {
          packages = [
            pkgs.go_1_23
            pkgs.gofumpt
            pkgs.golangci-lint
            pkgs.gopls
            pkgs.goreleaser
            pkgs.gotools
            pkgs.just
            pkgs.nix-prefetch
          ];
          shellHook = "go mod tidy";
        };
      }
    );
}
