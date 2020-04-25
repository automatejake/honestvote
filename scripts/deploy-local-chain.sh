go run main.go --tcp 7002 --http 7003 --role producer --collection-prefix a_ --registry true --institution-name "Honestvote"

sleep 5

go run main.go --tcp 7004 --http 7005 --role producer --collection-prefix b_ --registry-host 127.0.0.1 --registry-port 7002 --private-key "02113abfd3e15ef47bc8a3ccf3d2fc73bbc0a2cc1b8c8fbf22aca574510840b441" --public-key "4e5f04e07d45830d7e5732d4639636421d02b5ed06b7fbf69df569b248a9060a" & \

sleep 10

go run main.go --tcp 7006 --http 7007 --role producer --collection-prefix c_ --registry-host 127.0.0.1 --registry-port 7002 --private-key "020fc659d422d41878b39e2f5f5ed6a1fad98e5069ad63de4416b567a7d0c42b63" --public-key "80bf5a80b414d58688716917d700cee767b7e273d1a021458899fdb777622d68" & \