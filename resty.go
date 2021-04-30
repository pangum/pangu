package pangu

import (
	`crypto/tls`

	`github.com/go-resty/resty/v2`
	`github.com/storezhang/gox`
)

// NewResty 创建Http客户端
func NewResty(config gox.HttpClientConfig) (req *resty.Request) {
	client := resty.New()
	if "" != config.Proxy.Host {
		client.SetProxy(config.Proxy.Addr())
	}
	if 0 != config.Timeout {
		client.SetTimeout(config.Timeout)
	}
	if config.AllowGetPayload {
		client.SetAllowGetMethodPayload(true)
	}
	if config.Certificate.Skip {
		// nolint:gosec
		client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	} else {
		if "" != config.Certificate.Root {
			client.SetRootCertificate(config.Certificate.Root)
		}
		if 0 != len(config.Certificate.Clients) {
			certificates := make([]tls.Certificate, 0, len(config.Certificate.Clients))
			for _, c := range config.Certificate.Clients {
				certificate, err := tls.LoadX509KeyPair(c.Public, c.Private)
				if nil != err {
					continue
				}
				certificates = append(certificates, certificate)
			}
			client.SetCertificates(certificates...)
		}
	}
	if 0 != len(config.Headers) {
		client.SetHeaders(config.Headers)
	}
	if 0 != len(config.QueryParams) {
		client.SetQueryParams(config.QueryParams)
	}
	if 0 != len(config.FormData) {
		client.SetFormData(config.FormData)
	}
	if 0 != len(config.Cookies) {
		client.SetCookies(config.Cookies)
	}
	if "" != config.Auth.Type {
		switch config.Auth.Type {
		case gox.AuthTypeBasic:
			client.SetBasicAuth(config.Auth.Username, config.Auth.Password)
		case gox.AuthTypeToken:
			client.SetAuthToken(config.Auth.Token)
			if "" != config.Auth.Scheme {
				client.SetAuthScheme(config.Auth.Scheme)
			}
		}
	}

	// nolint:gosec
	req = client.R()

	return
}
