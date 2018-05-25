package models

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
)

func Strtomd5(str string) string  {
	h := md5.New()
	h.Write([]byte(str))
	re := hex.EncodeToString(h.Sum(nil))
	return  re
}

func Pwdhash(str string) string {
	return Strtomd5(str)
}

func StringsToJson(str string) string {
	rs := []rune(str)
	jsons := ""
	for _, r := range rs {
		rint := int(r)
		if rint < 128 {
			jsons += string(r)
		} else {
			jsons += "\\u" + strconv.FormatInt(int64(rint), 16) // json
		}
	}
	
	return jsons
}