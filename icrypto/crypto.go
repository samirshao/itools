package icrypto

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

type Encrypt struct {
	data []byte
	err  error
}

func Sha256(data []byte) *Encrypt {
	h := sha256.New()
	_, err := h.Write(data)
	if err != nil {
		return &Encrypt{data: nil, err: err}
	}
	return &Encrypt{data: h.Sum(nil), err: nil}
}

func Md5(data []byte) *Encrypt {
	h := md5.New()
	_, err := h.Write(data)
	if err != nil {
		return &Encrypt{data: nil, err: err}
	}
	return &Encrypt{data: h.Sum(nil), err: nil}
}

func (t *Encrypt) ToHex() (string, error) {
	return hex.EncodeToString(t.data), t.err
}

func (t *Encrypt) ToByte() ([]byte, error) {
	return t.data, t.err
}
