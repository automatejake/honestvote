#!/bin/bash

go get github.com/joho/godotenv
go get go.mongodb.org/mongo-driver/bson
go get go.mongodb.org/mongo-driver/mongo
go get github.com/gorilla/mux
go get github.com/mitchellh/mapstructure



#this script should install binaries with go install
cd ../core/core-administrator/administrator
go install

cd ../../core-consensus/consensus
go install

cd ../../core-crypto/crypto
go install

cd ../../core-database/database
go install

cd ../../core-discovery/discovery
go install

cd ../../core-http/http
go install

cd ../../core-p2p/p2p
go install

cd ../../core-registration/registration
go install

cd ../../core-registry/registry
go install

cd ../../../tests/logger
go install