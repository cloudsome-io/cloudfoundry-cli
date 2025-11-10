//go:build linux && arm64

package cfcli

import (
	_ "embed"
)

//go:embed embed/cf7-cli_7.8.0_linux_arm64
var cfLinuxArm64 []byte

func init() {
	platformBinary = embeddedBinary{
		data: cfLinuxArm64,
		name: "cf",
	}
}
