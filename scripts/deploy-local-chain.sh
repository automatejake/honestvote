cd ../build

go run main.go --tcp 7002 --http 7003 --role producer --collection-prefix a_ --registry true --institution-name "Honestvote" --private-key "2edfd293fd384a0036f519083c0636db111eb4d7949dd22927201e30d84755c8" --public-key "03effee3be38c8f30b1d29c4f7c4ec16b4e03d27599bb6c5efe0e26d42458358dc"  

sleep 5

go run main.go --tcp 7004 --http 7005 --role producer --collection-prefix b_ --registry-host 127.0.0.1 --registry-port 7002 --private-key "030168787ff44ecf8d8676db18da2ccacae85bb284607372a90987edc1c8d3d64a" --public-key "d41ff4f99b0737697e57ad7931926a978c76dcb06fc5727006212a8d2e9308fe" --registry false & \

sleep 10

go run main.go --tcp 7006 --http 7007 --role producer --collection-prefix c_ --registry-host 127.0.0.1 --registry-port 7002 --private-key "03484041b1c8bc03fe134af7b643d310c3e2dd70d6fc0940d11aa17c0cd42eb97c" --public-key "301b5d78f02f565070b9d1b7aef978f7e8fd2c3473e1bdae2070eabd325ea6be" --registry false & \