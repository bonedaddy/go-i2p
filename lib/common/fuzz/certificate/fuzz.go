package exportable

import "github.com/bonedaddy/go-i2p/lib/common"

func Fuzz(data []byte) int {
	cert := common.Certificate(data)
	cert.Data()
	cert.Length()
	cert.Type()
	return 0
}
