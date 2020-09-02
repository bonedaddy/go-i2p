package exportable

import "github.com/bonedaddy/go-i2p/lib/common"

func Fuzz(data []byte) int {
	destination := common.Destination(data)
	destination.Base32Address()
	destination.Base64()
	return 0
}
