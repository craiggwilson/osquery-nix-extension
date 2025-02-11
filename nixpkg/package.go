package nixpkg

import (
	"bufio"
	"bytes"
	"fmt"
	"iter"
	"os/exec"
	"regexp"
)

var storePathRegex = regexp.MustCompile(`^/nix/store/[a-z0-9]+-(.+?)(-([0-9].*?))?(\.drv)?$`)

func ListFromClosure(closure string) (iter.Seq[Package], error) {
	cmd := exec.Command("nix-store", "--query", "--requisites", closure)
	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("querying the nix-store: %w", err)
	}

	return parseOutput(out), nil
}

func parseOutput(out []byte) iter.Seq[Package] {
	scanner := bufio.NewScanner(bytes.NewReader(out))

	return func(yield func(Package) bool) {
		for scanner.Scan() {
			pkg := parseOutputLine(scanner.Text())
			if !yield(pkg) {
				return
			}
		}

		// TODO: handle scanner.Err()
	}
}

func parseOutputLine(line string) Package {
	matches := storePathRegex.FindStringSubmatch(line)

	pkg := Package{
		StorePath: StorePath(line),
	}

	if len(matches) > 0 {
		pkg.Name = matches[1]
	}
	if len(matches) > 2 {
		pkg.Version = Version(matches[3])
	}

	return pkg
}

type Package struct {
	Name      string
	Version   Version
	StorePath StorePath
}

type (
	StorePath string
	Version   string
)
