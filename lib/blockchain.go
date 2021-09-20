package lib

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"strings"
	"time"
)

const FirstHash = `0000000000000000000000000000000000000000000000000000000000000000`
const Difficulty = `0000`

type Block struct {
	Time     time.Time
	HashPrev string
	HashOwn  string
	Date     string
	Nonce    int64
}

func NewBlock(data string, hashPrev string) *Block {
	t := time.Now()
	hashOwn, nonce := func() (string, int64) {
		var cnt int64
		for cnt = 0; ; cnt++ {
			n := strconv.FormatInt(int64(cnt), 10)
			hash := sha256.Sum256(
				[]byte(
					data + strconv.FormatInt(
						t.Unix(), 10,
					) + n,
				),
			)
			str := hex.EncodeToString(hash[:])
			if strings.HasPrefix(str, Difficulty) {
				return str, cnt
			}
		}
	}()
	return &Block{
		Time:     t,
		HashPrev: hashPrev,
		HashOwn:  hashOwn,
		Date:     data,
		Nonce:    nonce,
	}
}
