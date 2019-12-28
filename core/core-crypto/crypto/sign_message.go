package crypto

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/asn1"
	"encoding/base64"
	"encoding/json"
	"errors"
	"math/big"
)

// Represents the two mathematical components of an ECDSA signature once
// decomposed.
type ECDSASignature struct {
	R, S *big.Int
}

// Encapsulates the overall message we're trying to decode and validate.
type Envelope struct {
	RawMessage json.RawMessage `json:"message"`
	Message    interface{}     `json:"-"`
	Signature  string          `json:"signature"`
}

// The body of the message to be contained in the Message field of our Envelope
// structure.
type MessageBody struct {
	Type        string          `json:"type"`
	UserID      uint32          `json:"userId"`
	Transaction json.RawMessage `json:"transaction"`
}

// Helper function to compute the SHA256 hash of the given string of bytes.
func hash(b []byte) []byte {
	h := sha256.New()
	// hash the body bytes
	h.Write(b)
	// compute the SHA256 hash
	return h.Sum(nil)
}

// Attempts to create a new envelope structure from the given JSON string.
func NewEnvelopeFromJSON(s string) (*Envelope, error) {
	var e Envelope
	if err := json.Unmarshal([]byte(s), &e); err != nil {
		return nil, err
	}
	// now attempt to unmarshal the message body itself from the raw message
	var body MessageBody
	if err := json.Unmarshal(e.RawMessage, &body); err != nil {
		return nil, err
	}
	e.Message = body
	return &e, nil
}

// The central validation routine that validates this message against the given
// public key. On success, returns nil, on failure returns a relevant error.
func (e *Envelope) Validate(publicKey *ecdsa.PublicKey) error {
	// first decode the signature to extract the DER-encoded byte string
	der, err := base64.StdEncoding.DecodeString(e.Signature)
	if err != nil {
		return err
	}
	// unmarshal the R and S components of the ASN.1-encoded signature into our
	// signature data structure
	sig := &ECDSASignature{}
	_, err = asn1.Unmarshal(der, sig)
	if err != nil {
		return err
	}
	// compute the SHA256 hash of our message
	h := hash(e.RawMessage)
	// validate the signature!
	valid := ecdsa.Verify(
		publicKey,
		h,
		sig.R,
		sig.S,
	)
	if !valid {
		return errors.New("Signature validation failed")
	}
	// signature is valid
	return nil
}
