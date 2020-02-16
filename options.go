package https

import (
	"time"
)

// GenerateOptions is used to configure the keys.
type GenerateOptions struct {

	// Comma-separated hostnames and IPs to generate a certificate for.
	Host string

	// Creation date formatted as "Jan 1 15:04:05 2011".
	// Default is time.Now().
	ValidFrom string

	// Duration that certificate is valid for.
	// Default is 365*24*time.Hour.
	ValidFor time.Duration

	// Whether this cert should be its own Certificate Authority
	// Default is false.
	IsCA bool

	// Size of RSA key to generate. Ignored if ECDSACurve is set
	// Default is 2048.
	RSABits int

	// ECDSA curve to use to generate a key. Valid values are P224, P256 (recommended), P384, P521
	// Default is "".
	ECDSACurve string

	// Generate an Ed25519 key
	// Default is false.
	ED25519Key bool
}
