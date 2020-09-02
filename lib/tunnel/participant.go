package tunnel

import (
	"github.com/bonedaddy/go-i2p/lib/crypto"
)

type Participant struct {
	decryption *crypto.Tunnel
}
