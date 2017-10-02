package main

import (
	varstrPkg "github.com/varstr/uaparser"
)

const varstrName = "varstr"

func varstrParse(pkg *uaPackage, in string) *uaResult {
	ua := varstrPkg.Parse(in)
	if ua == nil {
		return &uaResult{}
	}
	var mobile, tablet bool
	switch ua.DeviceType.Name {
	case "Phone":
		mobile = true
	case "Tablet":
		tablet = true
	}
	r := uaResult{
		os:             ua.OS.Name,
		osVersion:      ua.OS.Version,
		browser:        ua.Browser.Name,
		browserVersion: ua.Browser.Version,
		mobile:         mobile,
		tablet:         tablet,
	}
	if r.Valid() {
		pkg.Inc()
	}
	return &r
}
