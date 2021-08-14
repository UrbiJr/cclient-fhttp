package cclient

import (
	"github.com/useflyent/fhttp/cookiejar"

	http "github.com/useflyent/fhttp"
	"golang.org/x/net/proxy"

	utls "github.com/refraction-networking/utls"
)

func NewClient(clientHello utls.ClientHelloID, proxyUrl ...string) (http.Client, error) {
	jar, _ := cookiejar.New(nil)
	if len(proxyUrl) > 0 && len(proxyUrl[0]) > 0 {
		dialer, err := newConnectDialer(proxyUrl[0])
		if err != nil {
			return http.Client{Jar: jar}, err
		}
		return http.Client{
			Transport: newRoundTripper(clientHello, dialer), Jar: jar,
		}, nil
	} else {
		return http.Client{
			Transport: newRoundTripper(clientHello, proxy.Direct), Jar: jar,
		}, nil
	}
}
