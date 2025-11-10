//go:build linux && amd64

package cfcli

import (
	_ "embed"
)

//go:embed embed/cf7-cli_7.8.0_linux_x86-64
var cfLinuxAmd64 []byte

func init() {
	platformBinary = embeddedBinary{
		data: cfLinuxAmd64,
		name: "cf",
	}
}
