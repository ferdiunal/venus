package venus

import (
	"github.com/speps/go-hashids/v2"
)

type Venus struct {
	Len    int
	Salt   string
	hashid *hashids.HashID
}

func (v *Venus) init() *Venus {
	hd := hashids.NewData()
	hd.Salt = v.Salt
	hd.MinLength = v.Len
	h, _ := hashids.NewWithData(hd)

	v.hashid = h

	return v

}

func (v *Venus) DecodeInt64(data string) int64 {
	decoded, err := v.hashid.DecodeInt64WithError(data)
	if err != nil || err == nil && len(decoded) == 0 {
		return 0
	}

	return decoded[0]
}

func (v *Venus) Encode(data int64) string {
	hash, _ := v.hashid.EncodeInt64([]int64{data})

	return hash
}

func New(salt string, len int) *Venus {
	venus := &Venus{
		Salt: salt,
		Len:  len,
	}

	return venus.init()
}
