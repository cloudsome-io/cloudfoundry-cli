//go:build windows && amd64

package cfcli

import (
	_ "embed"
)

//go:embed embed/cf7-cli_7.8.0_winx64.exe
var cfWindowsAmd64 []byte

func init() {
	platformBinary = embeddedBinary{
		data: cfWindowsAmd64,
		name: "cf.exe",
	}
}
