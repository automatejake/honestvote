# This is the section of code responsible for generating mock data that will be used to populate the database
go run automated_testing_data.go


# This is the section of code responsible for generating unit tests

cd ../../

cd core
# cd core-administrator/administrator-test
# go test -v

cd core-consensus/consensus-test
go test -v

cd ../../core-crypto/crypto-test
go test -v

cd ../../core-database/database-test
go test -v

# cd ../../core-discovery/discovery-test
# go test -v

# cd ../../core-election/election-test
# go test -v

cd ../../core-http/http-test
go test -v

cd ../../core-p2p/p2p-test
go test -v

# cd ../../core-registration/registration-test
# go test -v

# cd ../../core-registry/registry-test
# go test -v

cd ../../core-validation/validation-test
go test -v

# cd ../../core-websocket/websocket-test
# go test -v
