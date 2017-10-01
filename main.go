package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	d := flag.String("data", "", "path to data file")
	v := flag.Bool("v", false, "verbose output")
	flag.Parse()
	os.Exit(run(*d, *v, os.Stdout, os.Stderr))
}

const (
	exitOK   = 0
	exitErr  = 1
	exitUser = 2
)

func run(path string, verbose bool, stdout, stderr io.Writer) int {
	lgr := newLogger(verbose, stdout, stderr)

	var data Data
	err := loadData(lgr, path, &data)
	if err != nil {
		lgr.Errorf("load raw data: %v", err)
		return exitErr
	}

	max := len(data)

	xojoc := newUAPackage(xojocPath, max)

	for _, ua := range data {
		lgr.Debugf("\n%s", ua)
		r := xojocParse(xojoc, ua)
		lgr.Debugf("%s", showResult(xojoc, r))
	}
	lgr.Infof("\n\n%v", xojoc.Total())

	return exitOK
}

func showResult(pkg *uaPackage, r *uaResult) string {
	valid := "-"
	if r.valid {
		valid = "+"
	}
	s := fmt.Sprintf("%s %s: %+v\n", valid, pkg.path, r)
	return s
}
