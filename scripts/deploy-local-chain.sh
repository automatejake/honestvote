go run main.go --tcp 7002 --http 7003 --role producer --collection-prefix a_ --registry true --institution-name "Honestvote" --private-key "790bc20440d6dc17bff7a7b17547d4f5ea7782a4d30d4707878d5bb95cbd9676" --public-key "02e9e343da7bb7fbe2caf18315fa5c3f907f71c23caf9532accb225cb3aed1f6a3"  

sleep 5

go run main.go --tcp 7004 --http 7005 --role producer --collection-prefix b_ --registry-host 127.0.0.1 --registry-port 7002 --private-key "03fd71c32c0dec9d8eb9c8b2a97d796f2640259035e991c3d3336fbfd9f707da4c" --public-key "65397db76da6d419376a97540970e430c4e47e51170a09a5dfdac332c4b5afed" --registry false & \

sleep 10

go run main.go --tcp 7006 --http 7007 --role producer --collection-prefix c_ --registry-host 127.0.0.1 --registry-port 7002 --private-key "025ef2b1fe3808d3860584bbcc6c223a4af6aa61c50df97a0cccd675bebddb4406" --public-key "3991c6bf62f6eec84bf0c13ee864bdcfbcd25a559b9f88733f21c6fb1dc4af17" --registry false & \