cd ../build

go run main.go --tcp 7002 --http 7003 --role producer --collection-prefix a_ --registry true --institution-name "Honestvote" --private-key "2edfd293fd384a0036f519083c0636db111eb4d7949dd22927201e30d84755c8" --public-key "03effee3be38c8f30b1d29c4f7c4ec16b4e03d27599bb6c5efe0e26d42458358dc"  

# sleep 5

# go run main.go --tcp 7004 --http 7005 --role producer --collection-prefix b_ --registry-host 127.0.0.1 --registry-port 7002 --private-key "021dc40041f99fd762b97205ac2c29dad52b34a6ed87313893d3fe39d101be3588" --public-key "04087dec4028eb48aa73d8c06a44eba2e183d6f3712ce8e8c535b4594a6e5bf1" --registry false & \

# sleep 10

# go run main.go --tcp 7006 --http 7007 --role producer --collection-prefix c_ --registry-host 127.0.0.1 --registry-port 7002 --private-key "02f4da1f5334c1c91a9569d3f914759410b31176411cf7686e8004417750ef2b83" --public-key "93e7b16c3fe6f6a9fe163c4f69f22676038bf3564d23cbd6a9de4a7e941b2965" --registry false & \
