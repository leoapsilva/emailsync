package config

import "os"

func LoadEnvVariables() {
	os.Setenv("SERVER_PORT", "6345")
	os.Setenv("MOCK_API_URL", "https://challenge.trio.dev/api/v1/")
	os.Setenv("MOCK_API_CONTACTS_ENDPOINT", "contacts")
	os.Setenv("MAILCHIMP_API_KEY", "755380e5c4e7fe0534e0460c2fadfa9f-us9")
	os.Setenv("MAILCHIMP_API_URL", "https://https://us9.admin.mailchimp.com")
	os.Setenv("MAILCHIMP_LIST_NAME", "Leonardo Alves de Paula e Silva")
	os.Setenv("MAILCHIMP_LIST_ID", "5c2f756353")
}

func GetEnvVariable(key string) string {
	return os.Getenv(key)
}
