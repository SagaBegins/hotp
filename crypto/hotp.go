package hotp

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/binary"
	"hash"
	"math"
	"time"
)

// Calculates the HMAC based one time password using the given 'encryption' for a 'secretKey', initial time 't0'
// and the 'interval' between each successive password and returns a password with a length of 'passLength'.
func CalcHotp(encryption string, secretKey []byte, t0 int64, interval int64, passLength int) (res uint64) {
	var offset int
	var hmac_result []byte
	var hasher hash.Hash

	counter := uint64(time.Now().Unix() / interval)

	bytearr := make([]byte, 8)
	binary.BigEndian.PutUint64(bytearr, counter)

	switch encryption {
	case "sha1":
		hasher = hmac.New(sha1.New, secretKey)
	case "sha256":
		hasher = hmac.New(sha256.New, secretKey)
	case "sha512":
		hasher = hmac.New(sha512.New, secretKey)
	default:
		hasher = hmac.New(sha512.New, secretKey)
	}

	hasher.Write(bytearr)
	hmac_result = hasher.Sum(nil)
	offset = int(hmac_result[len(hmac_result)-1] & 0xf)

	bin_code := []byte{byte(0x0), byte(0x0), byte(0x0), byte(0x0),
		byte((hmac_result[offset] & 0x7f)),
		byte((hmac_result[offset+1]) & 0xff),
		byte((hmac_result[offset+2]) & 0xff),
		byte((hmac_result[offset+3]) & 0xff)}

	res = binary.BigEndian.Uint64(bin_code)
	return res % uint64(math.Pow10(passLength))
}
