go run main.go --tcp 7002 --http 7003 --role producer --collection-prefix a_ --registry true --institution-name "West Chester University" \
--private-key "3077020101042090239de4607d93045ee60d077b310d5e4b3cc203a658b4f0b48a9a8bb26a8d6fa00a06082a8648ce3d030107a14403420004600e1045fdcd45ec2d59eeb396bad9549db0077c9b53627d066b6952d93b3c5e85d1fe4f9b2c712aff3d8234bb4e07cd5ef65b795ec7582eec4d097f4b62bb81" \
--public-key "3059301306072a8648ce3d020106082a8648ce3d03010703420004600e1045fdcd45ec2d59eeb396bad9549db0077c9b53627d066b6952d93b3c5e85d1fe4f9b2c712aff3d8234bb4e07cd5ef65b795ec7582eec4d097f4b62bb81" & \

sleep 5

go run main.go --tcp 7004 --http 7005 --role producer --collection-prefix b_ --registry-host 127.0.0.1 --registry-port 7002  & \

sleep 10

go run main.go --tcp 7006 --http 7007 --role producer --collection-prefix c_ --registry-host 127.0.0.1 --registry-port 7002  & \