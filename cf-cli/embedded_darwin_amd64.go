//go:build darwin && amd64

package cfcli

import (
	_ "embed"
)

//go:embed embed/cf7-cli_7.8.0_osx_x86_64
var cfDarwinAmd64 []byte

func init() {
	platformBinary = embeddedBinary{
		data: cfDarwinAmd64,
		name: "cf",
	}
}
