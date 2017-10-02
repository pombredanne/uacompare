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

	var osName, osVersion string
	if ua.OS != nil {
		osName = ua.OS.Name
		osVersion = ua.OS.Version
	}
	var browserName, browserVersion string
	if ua.Browser != nil {
		browserName = ua.Browser.Name
		browserVersion = ua.Browser.Version
	}

	var mobile, tablet bool
	if ua.DeviceType != nil {
		switch ua.DeviceType.Name {
		case "Phone":
			mobile = true
		case "Tablet":
			tablet = true
		}
	}
	r := uaResult{
		os:             osName,
		osVersion:      osVersion,
		browser:        browserName,
		browserVersion: browserVersion,
		mobile:         mobile,
		tablet:         tablet,
	}
	if r.Valid() {
		pkg.Inc()
	}
	return &r
}
