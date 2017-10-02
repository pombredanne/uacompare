package main

import (
	xojocPkg "xojoc.pw/useragent"
)

const xojocName = "xojoc"

func xojocParse(pkg *uaPackage, in string) *uaResult {
	ua := xojocPkg.Parse(in)
	if ua == nil {
		return &uaResult{}
	}
	r := uaResult{
		os:             ua.OS,
		osVersion:      ua.OSVersion.String(),
		browser:        ua.Name,
		browserVersion: ua.Version.String(),
		mobile:         ua.Mobile,
		tablet:         ua.Tablet,
	}
	if r.Valid() {
		pkg.Inc()
	}
	return &r
}
