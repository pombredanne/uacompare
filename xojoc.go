package main

import (
	xojocPkg "xojoc.pw/useragent"
)

const xojocName = "xojoc"

func xojocParse(pkg *uaPackage, in string) *uaResult {

	fixUnknown := func(in string) string {
		if in == "unknown" || in == "0.0.0" {
			return ""
		}
		return in
	}

	ua := xojocPkg.Parse(in)
	if ua == nil {
		return &uaResult{}
	}

	osName := fixUnknown(ua.OS)
	osVersion := fixUnknown(ua.OSVersion.String())
	browserName := fixUnknown(ua.Name)
	browserVersion := fixUnknown(ua.Version.String())

	r := uaResult{
		os:             osName,
		osVersion:      osVersion,
		browser:        browserName,
		browserVersion: browserVersion,
		mobile:         ua.Mobile,
		tablet:         ua.Tablet,
	}
	if r.Valid() {
		pkg.Inc()
	}
	return &r

}
