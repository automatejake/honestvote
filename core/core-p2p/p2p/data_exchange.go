package p2p

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/jneubaum/honestvote/core/core-database/database"
	"github.com/jneubaum/honestvote/tests/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func SendIndex(index int64, conn net.Conn) {
	write := new(Message)
	write.Message = "grab data"
	write.Data = []byte(string(index))

	jWrite, err := json.Marshal(write)

	if err == nil {
		conn.Write(jWrite)
	}
}

//GrabDocuments is here due to circular import error in database_exchange
func GrabDocuments(client *mongo.Client, conn net.Conn, old_index string) {

	var block database.Block

	collection := client.Database("honestvote").Collection(database.CollectionPrefix + "blockchain")

	index, _ := strconv.Atoi(old_index)
	current, err := collection.CountDocuments(context.TODO(), bson.M{})

	difference := current - int64(index)

	if difference > 0 && err == nil {
		fmt.Println("Here's the difference: ", difference)

		result, err := collection.Find(context.TODO(), bson.M{"index": bson.M{"$gt": index}})

		if err != nil {
			logger.Println("database_exchange.go", "GrabDocuments()", err.Error())
		}

		for result.Next(context.TODO()) {
			err = result.Decode(&block)
<<<<<<< HEAD
			if tran, ok := block.Transaction.(primitive.D); ok {
				tempMap := tran.Map()
				if pos, ok := tempMap["positions"].(primitive.A); ok {
					for _, position := range pos {
						if info, ok := position.(primitive.D); ok {
							tMap := info.Map()
							// if cand, ok := tMap["candidates"].(primitive.A); ok {
							// 	for _, candidate := range cand {
							// 		if candInfo, ok := candidate.(primitive.D); ok {
							// 			aMap := candInfo.Map()
							// 			tMap["candidates"] = aMap
							// 		}
							// 	}
							// }
							tempMap["positions"] = tMap
						}
					}
				}
=======
			if t, ok := block.Transaction.(primitive.D); ok {
				tempMap := t.Map()
>>>>>>> 4038928d19af35107bacd1023ed6e5a2b639cdbf
				block.Transaction = tempMap
			}
			MoveDocuments(conn, block)
		}
	} else {
		fmt.Println("Indexes are equal!")
	}
}

func FixBSON(block database.Block) {
	if tran, ok := block.Transaction.(primitive.D); ok {
		tempMap := tran.Map()
		if pos, ok := tempMap["positions"].(primitive.A); ok {
			for _, position := range pos {
				if info, ok := position.(primitive.D); ok {
					tempMap["positions"] = info.Map()
				}
			}
		}
		block.Transaction = tempMap
	}
}

//Send the data to the full/peer node
func MoveDocuments(conn net.Conn, block database.Block) {

	write := new(Message)
	write.Message = "receive data"
	write.Data, _ = json.Marshal(block)

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

	fmt.Println("proposed block")
	write := new(Message)
	write.Message = "verify transaction"
	write.Data = j
	write.Type = database.TransactionType(block.Transaction)

	jWrite, err := json.Marshal(write)

	if err == nil {
		for _, node := range Nodes {
			node.Write(jWrite)
		}
	}

	ProposedBlock = database.Block{}

}

func DecideType(data []byte, mType string, conn net.Conn) {
	var block database.Block

	if mType == "Vote" {
		vote := &database.Vote{}
		block = database.Block{Transaction: vote}
	} else if mType == "Election" {
		election := &database.Election{}
		block = database.Block{Transaction: election}
	} else if mType == "Registration" {
		registration := &database.Registration{}
		block = database.Block{Transaction: registration}
	}

	json.Unmarshal(data, &block)
	logger.Println("peer_routes.go", "HandleConn()", "Verifying")
	VerifyBlock(block, conn)
}

//Decide if the block sent is valid
func VerifyBlock(block database.Block, conn net.Conn) {

}
func LatestHashAndIndex(client *mongo.Client) {
	var block database.Block

	collection := client.Database("honestvote").Collection("b_blockchain")
	//collection := client.Database("honestvote").Collection(database.CollectionPrefix + "blockchain")

	ctx := context.Background()
	var bsonMap bson.M

	filter := bson.M{"xx": bsonMap}

	// Pass the filter to Find() to return a MongoDB cursor
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		log.Fatal("collection. Find ERROR:", err)
	}
	// iterate through all documents
	for cursor.Next(ctx) {
		// Decode the document
		if err := cursor.Decode(&block); err != nil {
			log.Fatal("cursor. Decode ERROR:", err)
			return
		}
	}
	PreviousBlock = block
	fmt.Println(PreviousBlock)

}
