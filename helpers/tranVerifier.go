package helpers

import (
	secp "secp256k1-go"
)

func PubkeyFromSeckey(privateKey []byte) []byte {
	//

	return secp.PubkeyFromSeckey([]byte(privateKey))
}
