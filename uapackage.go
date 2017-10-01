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

type uaResult struct {
	valid          bool
	os             string
	osVersion      string
	browser        string
	browserVersion string
	mobile         bool
	tablet         bool
}

func (r *uaResult) String() string {
	s := fmt.Sprintf("os: %s(%s) br: %s(%s) m:%t t:%t",
		r.os, r.osVersion, r.browser, r.browserVersion, r.mobile, r.tablet)
	return s
}
