#this script should install binaries with go install
cd ../core/core-consensus/consensus
go install

cd ../../core-crypto/crypto
go install

cd ../../core-database/database
go install

# cd ../../core-discovery/src
# go install

# cd ../../core-election/src
# go install

cd ../../core-http/http
go install

cd ../../core-p2p/p2p
go install

