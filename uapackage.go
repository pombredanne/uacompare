package main

import "fmt"

type uaPackage struct {
	path string
	ok   int
	max  int
}

func newUAPackage(path string, max int) *uaPackage {
	pkg := uaPackage{
		path: path,
		max:  max,
	}
	return &pkg
}

func (pkg *uaPackage) Total() string {
	percent := float64(pkg.ok) / float64(pkg.max)
	percent = percent * 100
	r := fmt.Sprintf("%q: %d / %d = %3.2f%%",
		pkg.path,
		pkg.ok,
		pkg.max,
		percent)
	return r
}

func (pkg *uaPackage) Inc() {
	pkg.ok++
}
