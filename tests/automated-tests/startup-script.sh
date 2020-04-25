go run main.go --tcp 7002 --http 7003 --role producer --collection-prefix a_ --registry true --institution-name "Honestvote" --private-key "790bc20440d6dc17bff7a7b17547d4f5ea7782a4d30d4707878d5bb95cbd9676" --public-key "02e9e343da7bb7fbe2caf18315fa5c3f907f71c23caf9532accb225cb3aed1f6a3"  

sleep 5

go run main.go --tcp 7004 --http 7005 --role producer --collection-prefix b_ --registry-host 127.0.0.1 --registry-port 7002 --private-key "0273f638382c1c53f00c4e9fa5bd5e2206ccc74c10f4dbbc0edd807757da4e8fff" --public-key "30e4caeec94efb11c7c1278d6caad518d1c0a50d7d5a0373c2decb501c65f50e" & \

sleep 10

go run main.go --tcp 7006 --http 7007 --role producer --collection-prefix c_ --registry-host 127.0.0.1 --registry-port 7002 --private-key "02795dbcbdf5ae53e5d27a5a987785120562a312b208b5f966abb0ddd0d9de5db1" --public-key "2ba163f6cc7bd8bce3149de944c756bad3a7d565d5385cf7343cfaafabf847ec" & \