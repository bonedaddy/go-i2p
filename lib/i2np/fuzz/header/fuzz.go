package exportable

import "github.com/bonedaddy/go-i2p/lib/i2np"

func Fuzz(data []byte) int {
	i2np.ReadI2NPNTCPHeader(data)
	return 0
}
