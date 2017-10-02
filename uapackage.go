package main

import (
	"fmt"
)

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

// Total returns results total for package.
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

// Inc increments valid user agent counter.
func (pkg *uaPackage) Inc() {
	pkg.ok++
}

// Result returns formated result output.
func (pkg *uaPackage) Result(r *uaResult) string {
	valid := "-"
	if r.valid {
		valid = "+"
	}
	s := fmt.Sprintf("%s %s: %s\n", valid, pkg.path, r)
	return s
}

type uaResult struct {
	valid          bool
	os             string
	osVersion      string
	browser        string
	browserVersion string
	device         string
	mobile         bool
	tablet         bool
}

// Valid validates result's fields.
func (r *uaResult) Valid() bool {
	fields := []string{r.os, r.osVersion, r.browser, r.browserVersion}
	for _, f := range fields {
		if f == "" {
			return false
		}
	}
	if r.mobile && r.tablet {
		return false
	}
	return true
}

func (r *uaResult) String() string {
	var device, mobile, tablet string
	if r.device != "" {
		device = "d: " + r.device
	}
	if r.mobile {
		mobile = "mobile"
	}
	if r.tablet {
		tablet = "tablet"
	}
	s := fmt.Sprintf("os: %s(%s) br: %s(%s) %s %s %s",
		r.os, r.osVersion, r.browser, r.browserVersion, mobile, tablet, device)
	return s
}
