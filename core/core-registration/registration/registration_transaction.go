package registration

import (
	"encoding/json"

	"github.com/jneubaum/honestvote/core/core-crypto/crypto"
	"github.com/jneubaum/honestvote/core/core-database/database"
	"github.com/jneubaum/honestvote/core/core-p2p/p2p"
)

func SendRegistrationTransaction(registrant database.AwaitingRegistration) error {
	registration := database.Registration{
		Type:        "Registration",
		Election:    registrant.ElectionName,
		Receiver:    registrant.Sender,
		RecieverSig: registrant.SenderSig,
		Sender:      p2p.Self.PublicKey,
	}
	encoded, err := registration.Encode()
	if err != nil {
		return err
	}

	hash := crypto.CalculateHash(encoded)

	signature, err := crypto.Sign([]byte(hash), p2p.PrivateKey)
	if err != nil {
		return err
	}

	registration.Signature = signature

	data, err := json.Marshal(registration)
	if err != nil {
		return err
	}

	p2p.ReceiveTransaction("Registration", data)
	return nil
}
