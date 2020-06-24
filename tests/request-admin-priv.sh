echo "Election Transaction:"

curl --header "Content-Type: application/json" --request POST --data '{"message":"cmVxdWVzdGluZyBhZG1pbmlzdHJhdG9yIHByaXZpbGVnZXM=","publickey":"029c7ba15cb5c452bd3aa39b891b7adb841493898b589916ee3306fcddc826ac10","signature":"304402201e83351f68a19340f4ae4f5137f43bb50373bfae90d305f4e5571a3a4fad98590220420419fad9122b1fd3c48a599897a537d619413028edc080c0fee8d4cca3a1dd"}' http://localhost:7003/administrator/request-privileges

