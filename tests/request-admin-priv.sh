echo "Election Transaction:"

curl --header "Content-Type: application/json" --request POST --data '{"message":"cmVxdWVzdGluZyBhZG1pbmlzdHJhdG9yIHByaXZpbGVnZXM=","publickey":"025575ed73799ed416f4dc9e85294d4ba34f75fc34cb3606e62e98c58cc3373669","signature":"304502203e7796c9a2e5c0a34ed5eb3140476c7f46d91a788746db5f0ab8384b50208b7202210087fd963512b4ee5ad316fcc7383466e83bb019e898af5398d0e981fc279f55cb"}' http://localhost:7003//administrator/request-privileges



