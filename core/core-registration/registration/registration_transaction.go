package registration

import (
	"encoding/json"

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
		Signature:   "",
	}

	data, err := json.Marshal(registration)
	if err != nil {

	}

	p2p.ReceiveTransaction("Registration", data)
	return nil
}
