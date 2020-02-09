go run main.go --tcp 7002 --http 7003 --role producer --collection-prefix a_ --registry true --institution-name "West Chester University" \
--private-key "30770201010420014f2048786ee5425751d4e3c65faecf5647bde07c042e35174d7b7b017688a7a00a06082a8648ce3d030107a14403420004cb023af175587633569b16d2d88a65678e60d69901048405cb35b016681433d16767abfffaecfa4de90b6e8403b01bfb8164d648b6d937f48da874a4f4c10a5c" \
--public-key "3059301306072a8648ce3d020106082a8648ce3d03010703420004cb023af175587633569b16d2d88a65678e60d69901048405cb35b016681433d16767abfffaecfa4de90b6e8403b01bfb8164d648b6d937f48da874a4f4c10a5c" & \

sleep 5

go run main.go --tcp 7004 --http 7005 --role producer --collection-prefix b_ --registry-host 127.0.0.1 --registry-port 7002  & \

sleep 10

go run main.go --tcp 7006 --http 7007 --role producer --collection-prefix c_ --registry-host 127.0.0.1 --registry-port 7002  & \