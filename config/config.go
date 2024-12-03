package config

import (
	"encoding/base64"
	"fmt"
	"os"
)

func LoadEnvVariables() {
	os.Setenv("SERVER_PORT", "6345")
	os.Setenv("MOCK_API_URL", "https://challenge.trio.dev/api/v1/")
	os.Setenv("MOCK_API_CONTACTS_ENDPOINT", "contacts")
	os.Setenv("MAILCHIMP_API_KEY", "M2I2ZThjYmM5NDQ2MjNiNTBlODQ0NWFkNjdmMDE4NWItdXM5")
	os.Setenv("MAILCHIMP_API_URL", "https://us9.api.mailchimp.com/3.0")
	os.Setenv("MAILCHIMP_LIST_NAME", "Leonardo Alves de Paula e Silva")
	os.Setenv("MAILCHIMP_LIST_ID", "5c2f756353")
	os.Setenv("MAILCHIMP_MAX_SIMULTANEOS_CONNECTIONS", "10")
}

func GetEnvVariable(key string) string {
	return os.Getenv(key)
}

func EncodeBase64(str string) string {
	return base64.RawStdEncoding.EncodeToString([]byte(str))
}

func DecodeBase64(s string) (string, error) {
	ret, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		ret, err = base64.RawStdEncoding.DecodeString(s)
		if err != nil {
			return "", fmt.Errorf("erro [%s] ao decodificar os dados de base64", err.Error())
		}
	}
	return fmt.Sprintf("%s", ret), nil
}
