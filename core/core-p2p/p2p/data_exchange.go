package p2p

import (
	"context"
	"encoding/json"
	"net"
	"strconv"

	"github.com/jneubaum/honestvote/core/core-crypto/crypto"
	"github.com/jneubaum/honestvote/core/core-database/database"
	"github.com/jneubaum/honestvote/tests/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func SendIndex(index int, conn net.Conn) {
	write := new(Message)
	write.Message = "grab data"
	write.Data = []byte(strconv.Itoa(index))

	jWrite, err := json.Marshal(write)

	if err == nil {
		conn.Write(jWrite)
	}
}

//GrabDocuments is here due to circular import error in database_exchange
func GrabDocuments(client *mongo.Client, conn net.Conn, old_index string) {
	current_index := GrabBlocks(client, conn, old_index)

	GrabTransactions(client, conn, current_index, "elections")

	GrabTransactions(client, conn, current_index, "registrations")

	GrabTransactions(client, conn, current_index, "votes")
}

func GrabBlocks(client *mongo.Client, conn net.Conn, old_index string) int {
	var block database.Block

	collection := client.Database("honestvote").Collection(database.CollectionPrefix + "blockchain")

	index, _ := strconv.Atoi(old_index)

	result, err := collection.Find(context.TODO(), bson.M{"index": bson.M{"$gt": index}})

	if err != nil {
		logger.Println("database_exchange.go", "GrabDocuments()", err.Error())
	}

	for result.Next(context.TODO()) {
		err = result.Decode(&block)

		MoveBlock(conn, block)
	}

	return index
}

func GrabTransactions(client *mongo.Client, conn net.Conn, index int, tranType string) {

	collection := client.Database("honestvote").Collection(database.CollectionPrefix + tranType)

	result, err := collection.Find(context.TODO(), bson.M{"blockIndex": bson.M{"$gt": index}})

	if err != nil {
		logger.Println("database_exchange.go", "GrabDocuments()", err.Error())
	}

	switch tranType {
	case "elections":
		for result.Next(context.TODO()) {
			var election database.Election
			err = result.Decode(&election)

			//Move the elections as necessary
			MoveTransaction(conn, election, "elections")
		}
	case "registrations":
		for result.Next(context.TODO()) {
			var registration database.Registration
			err = result.Decode(&registration)

			//Move the registrations as necessary
			MoveTransaction(conn, registration, "registrations")
		}
	case "votes":
		for result.Next(context.TODO()) {
			var vote database.Vote
			err = result.Decode(&vote)

			//Move the votes as necessary
			MoveTransaction(conn, vote, "votes")
		}
	}
}

//Send the data to the full/peer node
func MoveBlock(conn net.Conn, block database.Block) {

	write := new(Message)
	write.Message = "receive block"
	write.Data, _ = json.Marshal(block)

	jWrite, err := json.Marshal(write)

	if err == nil {
		logger.Println("sync_database.go", "MoveDocuments", "Moving Documents")
		conn.Write(jWrite)
	} else {
		logger.Println("sync_database.go", "MoveDocuments", err.Error())
	}
}

//Send the data to the full/peer node
func MoveTransaction(conn net.Conn, transaction interface{}, tranType string) {

	write := new(Message)
	write.Message = "receive transaction"
	write.Data, _ = json.Marshal(transaction)
	write.Type = tranType

	jWrite, err := json.Marshal(write)

	if err == nil {
		logger.Println("sync_database.go", "MoveDocuments", "Moving Documents")
		conn.Write(jWrite)
	} else {
		logger.Println("sync_database.go", "MoveDocuments", err.Error())
	}
}

//Send a block out to be verified by other peers
func ProposeBlock(block database.Block) {
	j, err := json.Marshal(block)

	write := new(Message)
	write.Message = "verify block"
	write.Data = j

	jWrite, err := json.Marshal(write)

	if err == nil {
		for _, node := range Nodes {
			node.Write(jWrite)
		}
	}

	ProposedBlock = database.Block{}

}

func ProposeTransaction(transaction interface{}, tranType string) {
	j, err := json.Marshal(transaction)

	write := new(Message)
	write.Message = "send transaction"
	write.Type = tranType
	write.Data = j

	jWrite, err := json.Marshal(write)

	if err == nil {
		for _, node := range Nodes {
			node.Write(jWrite)
		}
	}
}

//gets latest block, sends it to GrabDocuments which
func LatestHashAndIndex(client *mongo.Client) database.Block {
	var block database.Block
	//collection := client.Database("honestvote").Collection("a_blockchain")
	collection := client.Database("honestvote").Collection(database.CollectionPrefix + "blockchain")

	cur, _ := collection.CountDocuments(context.TODO(), bson.M{})
	filter := bson.M{"index": cur}
	documentReturned := collection.FindOne(context.TODO(), filter)

	documentReturned.Decode(&block)

	return block

}

func SendRegistrationTransaction(registrant database.AwaitingRegistration) error {
	registration := database.Registration{
		Type:        "Registration",
		Election:    registrant.ElectionName,
		Receiver:    registrant.Sender,
		RecieverSig: registrant.SenderSig,
		Sender:      Self.PublicKey,
	}
	encoded, err := registration.Encode()
	if err != nil {
		logger.Println("data_exchange.go", "SendRegistrationTransaction()", err)
		return err
	}

	hash := crypto.CalculateHash(encoded)

	signature, err := crypto.Sign([]byte(hash), PrivateKey)
	if err != nil {
		logger.Println("data_exchange.go", "SendRegistrationTransaction()", err)
		return err
	}

	registration.Signature = signature

	// Add transaction to queue
	Enqueue(registration)

	return nil
}
