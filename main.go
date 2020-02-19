package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	if err := run(os.Stdin, os.Stdout); err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(-1)
	}
}

const modePrefix = "mode: "
const modeAtomic = "mode: atomic"
const modeSet = "mode: set"

var lineRegex = regexp.MustCompile(`^(.+) (\d+)$`)

type coverage struct {
	mode  string
	lines map[string]int
}

func run(in io.Reader, out io.Writer) error {
	report := parse(in)
	write(out, report)
	return nil
}

func parse(in io.Reader) coverage {
	report := coverage{
		mode:  modeAtomic,
		lines: map[string]int{},
	}

	scanner := bufio.NewScanner(in)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// if this is a mode line – store it and continue
		if strings.HasPrefix(line, modePrefix) {
			report.mode = line
			continue
		}

		//

		// check that it is a valid-looking line and extract coverage metadata and hit count
		xs := lineRegex.FindAllStringSubmatch(line, 2)
		if xs == nil {
			continue
		}

		// and if it's a good-looking line – add its data to the coverage report
		meta := xs[0][1]
		hits, err := strconv.Atoi(xs[0][2])
		if err != nil {
			panic("Can not parse hit count " + xs[0][2])
		}
		report.lines[meta] = report.lines[meta] + hits
	}

	return report
}

// write rolled-up coverage report, sorting lines by metadata
func write(out io.Writer, report coverage) {
	fmt.Fprintln(out, report.mode)

	metas := make([]string, 0, len(report.lines))
	for k := range report.lines {
		metas = append(metas, k)
	}
	sort.Strings(metas)

	for _, meta := range metas {
		hits := report.lines[meta]
		if report.mode == modeSet && hits > 1 {
			hits = 1
		}
		fmt.Fprintf(out, "%s %d\n", meta, hits)
	}
}
