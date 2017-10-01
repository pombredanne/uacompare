package main

import (
	xojocPkg "xojoc.pw/useragent"
)

const xojocPath = "xojoc.pw/useragent"

func xojocParse(pkg *uaPackage, in string) *uaResult {
	ua := xojocPkg.Parse(in)
	if ua == nil {
		return &uaResult{}
	}
	r := uaResult{
		valid:          true,
		os:             ua.OS,
		osVersion:      ua.OSVersion.String(),
		browser:        ua.Name,
		browserVersion: ua.Version.String(),
		mobile:         ua.Mobile,
		tablet:         ua.Tablet,
	}
	pkg.Inc()
	return &r
}
