

build-nix-flake: 
    nix build .

fmt:
    gofumpt -w ./cmd ./nixpkg

lint:
    golangci-lint run --config .golangci.yml ./cmd ./nixpkg

install-dependencies:
    go mod tidy
   
setup-git-hooks:
    cp -f scripts/pre-commit.sh .git/hooks/pre-commit

update-dependencies: update-nix-dependencies update-go-dependencies update-go-vendor-hash

update-go-dependencies:
    go get -u ./...
    go mod tidy

update-go-vendor-hash:
    ./scripts/set-nix-flake-vendor-hash.sh

update-nix-dependencies:
    nix flake update

