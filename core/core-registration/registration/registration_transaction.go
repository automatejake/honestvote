package registration

import (
	"encoding/json"

	"github.com/jneubaum/honestvote/core/core-crypto/crypto"
	"github.com/jneubaum/honestvote/core/core-database/database"
	"github.com/jneubaum/honestvote/core/core-p2p/p2p"
	"github.com/jneubaum/honestvote/core/core-validation/validation"
)

func SendRegistrationTransaction(registrant database.AwaitingRegistration) error {
	registration := database.Registration{
		Type:        "Registration",
		Election:    registrant.ElectionName,
		Receiver:    registrant.Sender,
		RecieverSig: registrant.SenderSig,
		Sender:      p2p.Self.PublicKey,
	}

	headers := validation.GenerateRegistrationHeaders(registration)
	registration.Signature, _ = crypto.Sign([]byte(headers), p2p.PrivateKey)

	data, err := json.Marshal(registration)
	if err != nil {

	}

	p2p.ReceiveTransaction("Registration", data)
	return nil
}
