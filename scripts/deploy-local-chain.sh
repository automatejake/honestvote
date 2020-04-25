go run main.go --tcp 7002 --http 7003 --role producer --collection-prefix a_ --registry true --institution-name "Honestvote"

sleep 5

go run main.go --tcp 7004 --http 7005 --role producer --collection-prefix b_ --registry-host 127.0.0.1 --registry-port 7002 --private-key "03d66f44c59e365af0bdf1a7f132eeb7c9b969678baa6f9a8f3cb308f3e3c8d0e4" --public-key "60b3a8e99644e950dfba521ee6dfa59095ac43841f3f19d34ce0e62851a2b180" & \

sleep 10

go run main.go --tcp 7006 --http 7007 --role producer --collection-prefix c_ --registry-host 127.0.0.1 --registry-port 7002 --private-key "02d36f7b78ddb926eae011fb30bdc4729a4235c84b88a66de125bbfba20bf30bbe" --public-key "866228ba46518ac3aa52ccf0a12e1f89f310f876c136b0141362a6d90d614a13" & \