cd ../build

go run main.go --tcp 7002 --http 7003 --role producer --collection-prefix a_ --registry true --institution-name "West Chester University" \
--private-key "307702010104201c1799de2939b9acbb32051199e960e57a2039405d01dc1aee440f1eb95630e0a00a06082a8648ce3d030107a1440342000417ef2c220ac8c4a69fef5850cc599ef39a9cf2e5ee9298c6bae8a1fc0a8faff93eb2522968faf352e8b96da4f6ea758d61e0c231da525a50dd5fe9f2c63cdec1" \
--public-key "3059301306072a8648ce3d020106082a8648ce3d0301070342000417ef2c220ac8c4a69fef5850cc599ef39a9cf2e5ee9298c6bae8a1fc0a8faff93eb2522968faf352e8b96da4f6ea758d61e0c231da525a50dd5fe9f2c63cdec1" & \

sleep 5

go run main.go --tcp 7004 --http 7005 --role producer --collection-prefix b_ --registry-host 127.0.0.1 --registry-port 7002  & \

sleep 10

go run main.go --tcp 7006 --http 7007 --role producer --collection-prefix c_ --registry-host 127.0.0.1 --registry-port 7002  & \