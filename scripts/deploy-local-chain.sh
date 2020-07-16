cd ../build

go run main.go --tcp 7002 --http 7003 --role producer --collection-prefix a_ --registry true --institution-name "Honestvote" --private-key "2edfd293fd384a0036f519083c0636db111eb4d7949dd22927201e30d84755c8" --public-key "03effee3be38c8f30b1d29c4f7c4ec16b4e03d27599bb6c5efe0e26d42458358dc" & \

sleep 5

go run main.go --tcp 7004 --http 7005 --role producer --collection-prefix b_ --registry-host 127.0.0.1 --registry-port 7002 --private-key "03cef0af01459608364ad5fa61283ba1448d20154d2112b21aa69b5bbe550cd190" --public-key "eb903f8c6e3e492229edb6d2579ebf8f11114aa84c366035d06ac2122411715a" --registry false & \

sleep 10

go run main.go --tcp 7006 --http 7007 --role producer --collection-prefix c_ --registry-host 127.0.0.1 --registry-port 7002 --private-key "03cd84b0a959e42729a152e637abad8d8d00be6d1cd47b5435d180e57f9daede3f" --public-key "3eb6823fb91703e22f3cfd24325ae9ef49ee694e8bc136cd11204c391c2ae835" --registry false & \