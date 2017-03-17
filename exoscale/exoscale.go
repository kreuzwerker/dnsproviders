// Package exoscale adapts the lego Exoscale DNS provider
// for Caddy. Importing this package plugs it in.
package exoscale

import (
	"errors"

	"github.com/mholt/caddy/caddytls"
	"github.com/xenolf/lego/acme"
	"github.com/xenolf/lego/providers/dns/exoscale"
)

func init() {
	caddytls.RegisterDNSProvider("exoscale", NewDNSProvider)
}

// NewDNSProvider returns a new Exoscale DNS challenge provider.
// The credentials are interpreted as follows:
//
// len(0): use credentials from environment
// len(2): credentials[0] = API Key
//         credentials[1] = API Secret
func NewDNSProvider(credentials ...string) (acme.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return exoscale.NewDNSProvider()
	case 2:
		return exoscale.NewDNSProviderClient(credentials[0], credentials[1], "")
	default:
		return nil, errors.New("invalid credentials length")
	}
}
