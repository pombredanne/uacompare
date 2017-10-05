package main

import (
	"fmt"
	"strings"

	avocetPkg "github.com/avct/uasurfer"
)

const avocetName = "avocet"

func avocetParse(pkg *uaPackage, in string) *uaResult {
	fixUnknown := func(in string) string {
		if in == "Unknown" || in == "0.0.0" {
			return ""
		}
		return in
	}

	ua := avocetPkg.Parse(in)
	if ua == nil {
		return &uaResult{}
	}
	osName := ua.OS.Name.String()
	osName = strings.TrimPrefix(osName, "OS")
	osv := ua.OS.Version
	osVersion := fmt.Sprintf("%d.%d.%d", osv.Major, osv.Minor, osv.Patch)
	browserName := ua.Browser.Name.String()
	browserName = strings.TrimPrefix(browserName, "Browser")
	bv := ua.Browser.Version
	browserVersion := fmt.Sprintf("%d.%d.%d", bv.Major, bv.Minor, bv.Patch)
	var mobile, tablet bool
	if ua.DeviceType == avocetPkg.DevicePhone {
		mobile = true
	}
	if ua.DeviceType == avocetPkg.DeviceTablet {
		tablet = true
	}

	osName = fixUnknown(osName)
	osVersion = fixUnknown(osVersion)
	browserName = fixUnknown(browserName)
	browserVersion = fixUnknown(browserVersion)

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
