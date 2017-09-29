package main

import (
	"fmt"

	"xojoc.pw/useragent"
)

type uaPackage struct {
	path string
	ok   int
	max  int
}

func (pkg *uaPackage) Result() string {
	percent := float64(pkg.ok) / float64(pkg.max)
	percent = percent * 100
	r := fmt.Sprintf("%q: %d / %d = %.2f%%",
		pkg.path,
		pkg.ok,
		pkg.max,
		percent)
	return r
}

func (pkg *uaPackage) OK() {
	pkg.ok++
}

type xojocPackage struct {
	uaPackage
}

const xojocPath = "xojoc.pw/useragent"

func newXojocPackage(max int) xojocPackage {
	return xojocPackage{
		uaPackage{
			path: xojocPath,
			max:  max}}
}

func (pkg *xojocPackage) Parse(in string) string {
	ua := useragent.Parse(in)
	if ua == nil {
		return "<nil>"
	}
	pkg.ok++
	return fmt.Sprintf("%q\n%+v\n", pkg.path, ua)
}
