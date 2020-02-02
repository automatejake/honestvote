go run main.go --tcp 7002 --http 7003 --role producer --collection-prefix a_ --registry true --institution-name "West Chester University" \
--private-key "307702010104207a5c8c2138a4cf0b07baa5123ee8e9bdf961b198ff9cda0f7fb71d66105240d9a00a06082a8648ce3d030107a144034200045b1c14302eb26a7e7b7158a998b74f444e9838bc645672291447b239f9cd7a15e6a9e676d66c340b69a37134682f527d293da7f708d900530cb8ad2b9a749e69" \
--public-key "3059301306072a8648ce3d020106082a8648ce3d030107034200045b1c14302eb26a7e7b7158a998b74f444e9838bc645672291447b239f9cd7a15e6a9e676d66c340b69a37134682f527d293da7f708d900530cb8ad2b9a749e69" & \

sleep 5

go run main.go --tcp 7004 --http 7005 --role producer --collection-prefix b_ --registry-host 127.0.0.1 --registry-port 7002  & \

sleep 10

go run main.go --tcp 7006 --http 7007 --role producer --collection-prefix c_ --registry-host 127.0.0.1 --registry-port 7002  & \