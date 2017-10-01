package main

import (
	"fmt"

	xojocPkg "xojoc.pw/useragent"
)

const xojocPath = "xojoc.pw/useragent"

func xojocParse(pkg *uaPackage, in string) string {
	ua := xojocPkg.Parse(in)
	if ua == nil {
		return "<nil>"
	}
	pkg.Inc()
	return fmt.Sprintf("%q\n%+v\n", pkg.path, ua)
}
