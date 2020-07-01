echo "Election Transaction:"

curl --header "Content-Type: application/json" --request POST --data '{"message":"cmVxdWVzdGluZyBhZG1pbmlzdHJhdG9yIHByaXZpbGVnZXM=","institution":"BizyLife","domain":"bizylife.com","publickey":"02ed4abcf66d24138db6428819047b004c653a404e3d36ac6cf87f297753986809","signature":"3044022044d095a1c3b7c5191d38560fcae3e12dd7d36bb4c4ea9b8b1e0bda1713de55660220780b82ad28d69c74bb162caaec2d741ed31ae73bd1fca0911312fa5a8fcd5bb1"}' http://localhost:7003/administrator/request-privileges



