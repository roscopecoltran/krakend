package utils

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"

	// "github.com/roscopecoltran/krakend/logging"
	"github.com/speps/go-hashids"
)

// HashString TODO
func HashString(s string, length int) string {
	hasher := md5.New()
	if _, err := hasher.Write([]byte(s)); err != nil {
		return ""
	}

	rs := hex.EncodeToString(hasher.Sum(nil))

	if length == -1 {
		return rs
	}

	return rs[:length]
}

// GenerateString a string with n character
func GenerateString(n int) string {
	const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	symbols := big.NewInt(int64(len(alphanum)))
	states := big.NewInt(0)
	states.Exp(symbols, big.NewInt(int64(n)), nil)

	r, err := rand.Int(rand.Reader, states)
	if err != nil {
		panic(err)
	}

	var bytes = make([]byte, n)
	r2 := big.NewInt(0)
	symbol := big.NewInt(0)

	for i := range bytes {
		r2.DivMod(r, symbols, symbol)
		r, r2 = r2, r
		bytes[i] = alphanum[symbol.Int64()]
	}

	return string(bytes)
}

// GenerateOrderCode is the fun that help
// to generate a code for order with 6 characters
func GenerateCode() string {
	hd := hashids.NewData()
	hd.Alphabet = "0123456789abcdefghijklmnopqrstuvwxyz"
	h := hashids.NewWithData(hd)
	timeStamp := time.Now().UnixNano() / int64(time.Second)
	rs, _ := h.Encode([]int{int(timeStamp)})
	return rs
}
