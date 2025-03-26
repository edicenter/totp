// Generates time-based one-time tokens
package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"math"
)

var passwordPeriodSeconds = 30

// GetTOTP calculates time-based one-time password based on `secret` and time.
// `secret` is a shared secret between client and server.
func GetTOTP(secret string, unixTime int64) (string, error) {

	// 8-byte counter value, the moving factor.  This counter
	// MUST be synchronized between the HOTP generator (client)
	// and the HOTP validator (server).
	counter := uint64(math.Floor(float64(unixTime) / float64(passwordPeriodSeconds)))

	// HOTP: An HMAC-Based One-Time Password Algorithm
	// https://datatracker.ietf.org/doc/html/rfc4226

	counterBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(counterBytes, counter)
	secretBytes, err := base32.StdEncoding.DecodeString(secret)
	if err != nil {
		return "", err
	}
	// Step 1: Generate an HMAC-SHA-1 value
	// Let HS = HMAC-SHA-1(SECRET,COUNTER)
	// HS is a 20-byte string
	mac := hmac.New(sha1.New, secretBytes)
	mac.Write(counterBytes)
	HS := mac.Sum(nil)

	offset := HS[len(HS)-1] & 15

	// https://datatracker.ietf.org/doc/html/rfc4226#section-5.4
	// We treat the dynamic binary code as a 31-bit, unsigned, big-endian integer;
	binCode := binary.BigEndian.Uint32(HS[offset : offset+4])

	// the first byte is masked with a 0x7f.
	// We then take this number modulo 1,000,000 (10^6) to generate the 6-digit HOTP value.
	hotp := (binCode & 0x7fffffff) % 1000000

	password := fmt.Sprintf("%06d", hotp)

	return password, nil
}
