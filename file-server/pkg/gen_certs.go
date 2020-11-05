package pkg

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"log"
	"math/big"
	"net"
	"os"
	"strings"
	"time"

	"github.com/pkg/errors"
)

type param struct {
	host       string
	ecdsaCurve string
	validFrom  time.Time
	validFor   time.Duration
	isCA       bool
	ed25519Key bool
	rsaBits    int
	certPath   string
	keyPath    string
}

func newParam(opts ...Option) *param {
	param := &param{
		host:       "",
		ecdsaCurve: "P256",
		validFrom:  time.Now(),
		validFor:   365 * 24 * time.Hour,
		isCA:       false,
		ed25519Key: false,
		rsaBits:    2048,
	}
	for i, _ := range opts {
		opts[i](param)
	}
	return param
}

type Option func(p *param)

func WithHost(host string) Option {
	return func(p *param) {
		p.host = host
	}
}

func WithCertAndKeyPaths(certPath, keyPath string) Option {
	return func(p *param) {
		p.certPath = certPath
		p.keyPath = keyPath
	}
}

func GenCerts(opts ...Option) {
	param := newParam(opts...)

	if len(param.host) == 0 {
		log.Fatalf("Missing required --host parameter")
	}
	if len(param.certPath) == 0 {
		log.Fatalf("Failed to create certificate. Cert path is empty: %v", param.certPath)
	}

	if len(param.keyPath) == 0 {
		log.Fatalf("Failed to create certificate. Key path is empty: %v", param.keyPath)
	}
	createDir(param.certPath)
	createDir(param.keyPath)

	var priv interface{}
	var err error
	switch param.ecdsaCurve {
	case "":
		if param.ed25519Key {
			_, priv, err = ed25519.GenerateKey(rand.Reader)
		} else {
			priv, err = rsa.GenerateKey(rand.Reader, param.rsaBits)
		}
	case "P224":
		priv, err = ecdsa.GenerateKey(elliptic.P224(), rand.Reader)
	case "P256":
		priv, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	case "P384":
		priv, err = ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	case "P521":
		priv, err = ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	default:
		log.Fatalf("Unrecognized elliptic curve: %q", param.ecdsaCurve)
	}
	if err != nil {
		log.Fatalf("Failed to generate private key: %v", err)
	}

	// ECDSA, ED25519 and RSA subject keys should have the DigitalSignature
	// KeyUsage bits set in the x509.Certificate template
	keyUsage := x509.KeyUsageDigitalSignature
	// Only RSA subject keys should have the KeyEncipherment KeyUsage bits set. In
	// the context of TLS this KeyUsage is particular to RSA key exchange and
	// authentication.
	if _, isRSA := priv.(*rsa.PrivateKey); isRSA {
		keyUsage |= x509.KeyUsageKeyEncipherment
	}

	notBefore := param.validFrom
	notAfter := notBefore.Add(param.validFor)

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		log.Fatalf("Failed to generate serial number: %v", err)
	}

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"Acme Co"},
		},
		NotBefore: notBefore,
		NotAfter:  notAfter,

		KeyUsage:              keyUsage,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	hosts := strings.Split(param.host, ",")
	for _, h := range hosts {
		if ip := net.ParseIP(h); ip != nil {
			template.IPAddresses = append(template.IPAddresses, ip)
		} else {
			template.DNSNames = append(template.DNSNames, h)
		}
	}

	if param.isCA {
		template.IsCA = true
		template.KeyUsage |= x509.KeyUsageCertSign
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, publicKey(priv), priv)
	if err != nil {
		log.Fatalf("Failed to create certificate: %v", err)
	}

	certOut, err := os.Create(param.certPath)
	if err != nil {
		log.Fatalf("Failed to open %s for writing: %v", param.certPath, err)
	}
	if err := pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes}); err != nil {
		log.Fatalf("Failed to write data to %s: %v", param.certPath, err)
	}
	if err := certOut.Close(); err != nil {
		log.Fatalf("Error closing %s: %v", param.certPath, err)
	}
	log.Printf("wrote %s\n", param.certPath)

	keyOut, err := os.OpenFile(param.keyPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Failed to open %s for writing: %v", param.keyPath, err)
		return
	}
	privBytes, err := x509.MarshalPKCS8PrivateKey(priv)
	if err != nil {
		log.Fatalf("Unable to marshal private key: %v", err)
	}
	if err := pem.Encode(keyOut, &pem.Block{Type: "PRIVATE KEY", Bytes: privBytes}); err != nil {
		log.Fatalf("Failed to write data to %s: %v", param.keyPath, err)
	}
	if err := keyOut.Close(); err != nil {
		log.Fatalf("Error closing %s: %v", param.keyPath, err)
	}
	log.Printf("wrote %s\n", param.keyPath)
}

func publicKey(priv interface{}) interface{} {
	switch k := priv.(type) {
	case *rsa.PrivateKey:
		return &k.PublicKey
	case *ecdsa.PrivateKey:
		return &k.PublicKey
	case *ed25519.PrivateKey:
		return k.Public().(ed25519.PublicKey)
	default:
		return nil
	}
}

func createDir(filePath string) {
	paths := strings.Split(filePath, "/")
	fmt.Println(paths)
	paths = paths[:len(paths)-1]
	fmt.Println(paths)
	if len(paths) > 0 && paths[0] == "." {
		paths = paths[1:]
	}
	fmt.Println(paths)
	if len(paths) > 0 {
		tmpDir := ""
		for i, dir := range paths {
			if tmpDir != "" {
				tmpDir = tmpDir + "/" + dir
			} else {
				tmpDir = dir
			}
			_, err := os.Stat(tmpDir)
			if os.IsNotExist(err) {
				if err := os.Mkdir(tmpDir, 0777); err != nil {
					log.Fatal(errors.WithMessage(err, fmt.Sprintf("iteration: %d, path: %s", i, tmpDir)))
				}
				log.Printf("dir created %s", tmpDir)
			}
		}
	}
}
