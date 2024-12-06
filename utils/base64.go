package utils

import (
	"encoding/base64"
	"fmt"
)

func EncodeBase64(str string) string {
	return base64.RawStdEncoding.EncodeToString([]byte(str))
}

func DecodeBase64(s string) (string, error) {
	ret, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		ret, err = base64.RawStdEncoding.DecodeString(s)
		if err != nil {
			return "", fmt.Errorf("erro [%s] on decode data from base64", err.Error())
		}
	}
	return fmt.Sprintf("%s", ret), nil
}
