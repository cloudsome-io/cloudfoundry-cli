//go:build darwin && arm64

package cfcli

import (
	_ "embed"
)

//go:embed embed/cf7-cli_7.8.0_macosarm
var cfDarwinArm64 []byte

func init() {
	platformBinary = embeddedBinary{
		data: cfDarwinArm64,
		name: "cf",
	}
}
