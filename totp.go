// Generates time-based one-time tokens
package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"math"
	"strconv"
	"time"
)

var period = 30
var digits = 6

// GetToken calculates time-based one-time password based on `secret` and time
func GetToken(secret string) (string, error) {
	t := time.Now().Unix()
	interval := uint64(math.Floor(float64(t) / float64(period)))
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, interval)

	// fmt.Printf("interval=%v; decode secret=%x; buf=%x\n", interval, secretBytes, buf)
	secretBytes, err := base32.StdEncoding.DecodeString(secret)
	if err != nil {
		return "", err
	}

	mac := hmac.New(sha1.New, secretBytes)
	mac.Write(buf)
	h := mac.Sum(nil)

	// fmt.Printf("hash=%x\n", h)

	o := h[19] & 15

	// fmt.Printf("o=%v (%[1]T)\n", o)

	unpackUint32 := binary.BigEndian.Uint32(h[o : o+4])

	// fmt.Printf("unpackUint32=%v\n", unpackUint32)

	token := (unpackUint32 & 0x7fffffff) % 1000000

	return strconv.FormatUint(uint64(token), 10), nil

	// fmt.Printf("token=%v", token)

}
