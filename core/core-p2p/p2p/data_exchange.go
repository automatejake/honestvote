package p2p

import (
	"context"
	"encoding/json"
	"fmt"
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

			if tran, ok := block.Transaction.(primitive.D); ok {
				tranMap := tran.Map()

				if tranMap["type"].(string) == "Election" {
					if pos, ok := tranMap["positions"].(primitive.A); ok {
						var posElements primitive.A
						tranMap["positions"] = nil

						for _, position := range pos {
							if posInfo, ok := position.(primitive.D); ok {
								posMap := posInfo.Map()

								if cand, ok := posMap["candidates"].(primitive.A); ok {
									var candElements primitive.A
									posMap["candidates"] = nil

									for _, candidate := range cand {
										if candInfo, ok := candidate.(primitive.D); ok {
											candMap := candInfo.Map()
											candElements = append(candElements, candMap)
										}
									}

									posMap["candidates"] = candElements
								}

								posElements = append(posElements, posMap)
							}
						}

						tranMap["positions"] = posElements
					}
				}

				block.Transaction = tranMap
			}

			MoveDocuments(conn, block)
		}
	} else {
		fmt.Println("Indexes are equal!")
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

//gets latest block, sends it to GrabDocuments which
func LatestHashAndIndex(client *mongo.Client, conn net.Conn) {
	var block database.Block
	//collection := client.Database("honestvote").Collection("a_blockchain")
	collection := client.Database("honestvote").Collection(database.CollectionPrefix + "blockchain")

	cur, _ := collection.CountDocuments(context.TODO(), bson.M{})
	filter := bson.M{"index": cur}
	documentReturned := collection.FindOne(context.TODO(), filter)

	documentReturned.Decode(&block)

	//fmt.Println(block)

	s := strconv.Itoa(block.Index)

	PreviousBlock = block

	GrabDocuments(database.MongoDB, conn, s)

}
