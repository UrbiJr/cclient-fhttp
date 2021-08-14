package cclient

import (
	"github.com/useflyent/fhttp/cookiejar"

	http "github.com/useflyent/fhttp"
	"golang.org/x/net/proxy"

	utls "github.com/refraction-networking/utls"
)

func NewClient(clientHello utls.ClientHelloID, allowRedirects bool, proxyUrl ...string) (http.Client, error) {
	jar, _ := cookiejar.New(nil)
	if len(proxyUrl) > 0 && len(proxyUrl[0]) > 0 {
		dialer, err := newConnectDialer(proxyUrl[0])
		if err != nil {
			if !allowRedirects {
				return http.Client{Jar: jar, CheckRedirect: func(req *http.Request, via []*http.Request) error {
					return http.ErrUseLastResponse
				}}, err
			} else {
				return http.Client{Jar: jar}, err
			}
		}
		if !allowRedirects {
			return http.Client{
				Transport: newRoundTripper(clientHello, dialer), Jar: jar, CheckRedirect: func(req *http.Request, via []*http.Request) error {
					return http.ErrUseLastResponse
				},
			}, nil
		} else {
			return http.Client{
				Transport: newRoundTripper(clientHello, dialer), Jar: jar,
			}, nil
		}
	} else {
		if !allowRedirects {
			return http.Client{
				Transport: newRoundTripper(clientHello, proxy.Direct), Jar: jar, CheckRedirect: func(req *http.Request, via []*http.Request) error {
					return http.ErrUseLastResponse
				},
			}, nil
		} else {
			return http.Client{
				Transport: newRoundTripper(clientHello, proxy.Direct), Jar: jar,
			}, nil
		}
	}
}
