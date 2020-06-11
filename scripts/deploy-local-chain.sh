cd ../build

go run main.go --tcp 7002 --http 7003 --role producer --collection-prefix a_ --registry true --institution-name "Honestvote" --private-key "2edfd293fd384a0036f519083c0636db111eb4d7949dd22927201e30d84755c8" --public-key "03effee3be38c8f30b1d29c4f7c4ec16b4e03d27599bb6c5efe0e26d42458358dc"  

sleep 5

go run main.go --tcp 7004 --http 7005 --role producer --collection-prefix b_ --registry-host 127.0.0.1 --registry-port 7002 --private-key "02b19aa5ee3db1701ef4a5e752a9731b7fb4195fee680b36efd63e8ddd08f10636" --public-key "df97c6b3fb39e1accdf60b88fe2fb045d8e5378c886e9225d9e0836cfe96df48" --registry false & \

sleep 10

go run main.go --tcp 7006 --http 7007 --role producer --collection-prefix c_ --registry-host 127.0.0.1 --registry-port 7002 --private-key "03de136311a3a2746baf0a41b428340b6b45f5065817baf6728927e7f20967107b" --public-key "faef2af5075770daa5edc30500bea44dc036c1467530d968942cf15faba4fd5a" --registry false & \
