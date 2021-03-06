package translatableerror

type InvalidSSLCertError struct {
	URL              string
	SuggestedCommand string
}

func (InvalidSSLCertError) Error() string {
	return "Invalid SSL Cert for {{.API}}\nTIP: Use 'cf {{.SuggestedCommand}} --skip-ssl-validation' to continue with an insecure API endpoint"
}

func (e InvalidSSLCertError) Translate(translate func(string, ...interface{}) string) string {
	return translate(e.Error(), map[string]interface{}{
		"API":              e.URL,
		"SuggestedCommand": e.SuggestedCommand,
	})
}
