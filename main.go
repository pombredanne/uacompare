package main

import (
	"flag"
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

	avocet := newUAPackage(avocetName, max)
	mileusna := newUAPackage(mileusnaName, max)
	varstr := newUAPackage(varstrName, max)
	xojoc := newUAPackage(xojocName, max)

	for _, ua := range data {
		lgr.Debugf("\n%s", ua)
		r := xojocParse(xojoc, ua)
		lgr.Debugf("%s", xojoc.Result(r))
		r = mileusnaParse(mileusna, ua)
		lgr.Debugf("%s", mileusna.Result(r))
		r = varstrParse(varstr, ua)
		lgr.Debugf("%s", varstr.Result(r))
		r = avocetParse(avocet, ua)
		lgr.Debugf("%s", avocet.Result(r))
	}
	showTotals(lgr, avocet, mileusna, varstr, xojoc)

	return exitOK
}

func showTotals(lgr *Logger, pkgs ...*uaPackage) {
	lgr.Info("\n\nTotal:")
	for _, pkg := range pkgs {
		lgr.Infof("%v", pkg.Total())
	}
}
