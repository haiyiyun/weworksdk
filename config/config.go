package config

type Config struct {
	CorpId              string `json:"corp_id"`
	ProviderSecret      string `json:"provider_secret"`
	SuiteId             string `json:"suite_id"`
	SuiteSecret         string `json:"suite_secret"`
	SuiteToken          string `json:"suite_token"`
	SuiteEncodingAesKey string `json:"suite_encoding_aes_key"`
}
