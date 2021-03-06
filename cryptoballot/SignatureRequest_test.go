package cryptoballot

import (
	"crypto/rand"
	"crypto/rsa"
	"testing"
)

var (
	goodRequest = []byte(`12345

475995cee84cf85cdcc9018bafb09a21d70dc1e5413ec7ad3a4ea4f2d07cae35c878332b71e1bcaa7103ed46d2d4f46864e7b454b0e44dec66c84c283f6df521

MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEA31GRu9r2QRA9PtIzMKyV3vloQlrmxRLYIgiUsNg6bNOmTOJ1og+HNpTY8XOujf3KpPS38F1XM3AAJQi3pUjcJEdeiqroFf8b7t2pas1V+Bg2XAWWbfKctpnMuxeIYuJE52KhUK4y+qGaLXI+53oT09w3V4CdeQNZllVL2a6q+6gjpdZ+/YOPQ+dncHtYCxNHu1Idub0EP/ZMkdcHLwpi/gmuw7qvdpQTeiw54krV3MoiZq50ZTxTFRCjFJ+C+pmrYaPygrkCkv3sj3v1Be8k0EBYsMH8yZoigbyE0/SlCH+RGLSiS1yAV+MHcoVMzPFbXnFv9usI3UNVSXrDSzsxYgiDaeX7KVrraKhJrM/LIypZbJDiKLpLzKFEx+SkSQ/3e8eSsedp7N5RSvcz9GU6K4sUYtvNdiwHZTTakoo7m8pBF7dE9Guxjtcc42vwBSArsYrfstFcMaVwwth1Ohh/vO1W5EmMzzsqqm7DYPCVFapwV7wlveYFyD5e9ZVb/im8s+2NHg6PY5L1ke+JN+zx75M54nGezk+1pJcy05r66a56Wyh85RgMUok1XMPbiVmhA8TVwlCZGnfXetsSsFKgFjAGD+DdLCdkj9TH2tG7pewlEDNjVM+iWJA8Tmt/H+n4tL1LedzGs1KkwEZKEcxZtxDdBxPWFQDK3UloOwaP6y0CAwEAAQ==

5e517adedc743d29e9c7099e8f02e8c0c2f286e8e6e0a5f0c1fbf22b01fb70fb67518c206659e38f41922459eae98bd81712220583d754b9cae8487692b4a9c9

T494uksBnCb6Ld45SmISlQNIYU9iB6iZG8GyeS4jPZvUm1xFN+vUesmzKS1kIrqGx1fz11Gii8mGp32CcA78lvtbxRKF1E/BXuVLx/B5rmL3ClZIJ6WT2SLuJ69CSA1h0k0I9DT0FSD9D9m2cF/oJbs2h4AEJdHU2/rLdCcZtLl/FZGhDBa1BjqGH1FZcJiD1geBJsQIHjA+eNKucxJKt+3u0TmWrFQ535NhccvDhpNaDjUEaCdFpGJS0jzohmxI6kFDzahG2rSwXdYvYnSsZExgs/EmWAQu7ZNYncZAHY/NHWWXrvGdmLiXBBNL1so4kCD+SJLID0/mU7qgZfyOTleXKHFfLayzhayKTKx4W1zuxu84DB258JdURoOnB/EJVpNFTud2q0Daqmyq9C8dffyLa6SjUqqfQVCPMuM/yjnOX2vZr4QhlJXyij71ZOFtTh68+i+dTkjrp74NQRlzFsLROf1KKpnfTKjm9WRzHn0k0Ckyb+YDprYpClZCoG65aozxiPRI71ev0Z99p61KDw2vqj94cASakgDCIGwHDInBsaz+oBUjsjXccCHnsPAX64szzhReO2xBL4TgdWPoo28y+FveybeWr002MM+Q8LxVgCQpxJRp5e5wKyeCnb5Msct/Q8nwV2is/fHwFPYQQn3A2m6pQ/lVQY6hgnxki5g=`)
)

func TestGoodSignatureRequest(t *testing.T) {
	// Set number of bits to 2048
	MinPublicKeySize = 2048

	req, err := NewSignatureRequest(goodRequest)
	if err != nil {
		t.Error(err)
		return
	}

	if string(goodRequest) != req.String() {
		t.Errorf("SignatureRequest round-trip from string and back again failed.")
	}

	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Errorf("failed to generate private key")
	}

	if _, err := req.SignBallot(key); err != nil {
		t.Errorf("failed to sign ballot")
	}
}
