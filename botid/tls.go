package botid

import (
	"github.com/bogdanfinn/fhttp/http2"
	"github.com/bogdanfinn/tls-client/profiles"
	tls "github.com/bogdanfinn/utls"
)

var Brave_144 = profiles.NewClientProfile(
	tls.ClientHelloID{
		Client:  "Brave_144",
		Version: "1",
		Seed:    nil,
		SpecFactory: func() (tls.ClientHelloSpec, error) {
			return tls.ClientHelloSpec{
				CipherSuites: []uint16{
					tls.GREASE_PLACEHOLDER,
					tls.TLS_AES_128_GCM_SHA256,
					tls.TLS_AES_256_GCM_SHA384,
					tls.TLS_CHACHA20_POLY1305_SHA256,
					tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
					tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
					tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
					tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
					tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256,
					tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256,
					tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
					tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
					tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
					tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
					tls.TLS_RSA_WITH_AES_128_CBC_SHA,
					tls.TLS_RSA_WITH_AES_256_CBC_SHA,
				},
				// CompressionMethods is not implemented by tls.peet.ws, check manually
				CompressionMethods: []uint8{
					tls.CompressionNone,
				},
				Extensions: []tls.TLSExtension{
					&tls.UtlsGREASEExtension{},
					&tls.ExtendedMasterSecretExtension{},
					&tls.SupportedCurvesExtension{[]tls.CurveID{
						tls.CurveID(tls.GREASE_PLACEHOLDER),
						4588 /* X25519MLKEM768 (4588) */,
						tls.X25519,
						tls.CurveP256,
						tls.CurveP384,
					}},
					&tls.SignatureAlgorithmsExtension{SupportedSignatureAlgorithms: []tls.SignatureScheme{
						tls.ECDSAWithP256AndSHA256,
						tls.PSSWithSHA256,
						tls.PKCS1WithSHA256,
						tls.ECDSAWithP384AndSHA384,
						tls.PSSWithSHA384,
						tls.PKCS1WithSHA384,
						tls.PSSWithSHA512,
						tls.PKCS1WithSHA512,
					}},
					&tls.PSKKeyExchangeModesExtension{[]uint8{
						tls.PskModeDHE,
					}},
					&tls.SupportedVersionsExtension{[]uint16{
						tls.GREASE_PLACEHOLDER,
						tls.VersionTLS13,
						tls.VersionTLS12,
					}},
					tls.BoringGREASEECH(),
					&tls.ApplicationSettingsExtension{
						//codePoint:          tls.ExtensionALPSOld,
						SupportedProtocols: []string{"h2"},
					},
					&tls.SupportedPointsExtension{SupportedPoints: []byte{0x00}},
					&tls.UtlsCompressCertExtension{[]tls.CertCompressionAlgo{
						tls.CertCompressionBrotli,
					}},
					&tls.ALPNExtension{AlpnProtocols: []string{"h2", "http/1.1"}},
					&tls.RenegotiationInfoExtension{Renegotiation: tls.RenegotiateNever},
					&tls.SCTExtension{},
					&tls.SessionTicketExtension{},
					&tls.SNIExtension{},
					&tls.StatusRequestExtension{},
					&tls.KeyShareExtension{[]tls.KeyShare{
						{Group: tls.CurveID(tls.GREASE_PLACEHOLDER), Data: []byte{0} /* TLS_GREASE (0xaaaa) */},
						{Group: 4588 /* X25519MLKEM768 (4588) */},
						{Group: tls.X25519},
					}},
					&tls.UtlsGREASEExtension{},
					&tls.UtlsPreSharedKeyExtension{OmitEmptyPsk: true},
				},
			}, nil
		},
	},
	map[http2.SettingID]uint32{
		http2.SettingHeaderTableSize:   65536,
		http2.SettingEnablePush:        0,
		http2.SettingInitialWindowSize: 6291456,
		http2.SettingMaxHeaderListSize: 262144,
	},
	[]http2.SettingID{
		http2.SettingHeaderTableSize,
		http2.SettingEnablePush,
		http2.SettingInitialWindowSize,
		http2.SettingMaxHeaderListSize,
	},
	[]string{
		":method",
		":authority",
		":scheme",
		":path",
	},
	uint32(15663105),
	[]http2.Priority{},
	&http2.PriorityParam{
		StreamDep: 0,
		Exclusive: true,
		Weight:    255,
	},
)
