go run main.go --tcp 7002 --http 7003 --role producer --collection-prefix a_ --registry true --institution-name "Honestvote" --private-key "790bc20440d6dc17bff7a7b17547d4f5ea7782a4d30d4707878d5bb95cbd9676" --public-key "02e9e343da7bb7fbe2caf18315fa5c3f907f71c23caf9532accb225cb3aed1f6a3"  

sleep 5

go run main.go --tcp 7004 --http 7005 --role producer --collection-prefix b_ --registry-host 127.0.0.1 --registry-port 7002 --private-key "03caf7342368a6582b24278aa1aaa91a721a77f9b5bb27ec772868b0c140a7fd15" --public-key "4a451c08789c0f889d63b0c15945a17bdef8dc82632b1cd20a35112405d99250" --registry false & \

sleep 10

go run main.go --tcp 7006 --http 7007 --role producer --collection-prefix c_ --registry-host 127.0.0.1 --registry-port 7002 --private-key "037c812f1c155e5e6f6cba6ff4e6d1a27f07241d45c12ad541fa8133079989c09c" --public-key "e5be6ffdfaac59a9ff7931a424325a00d12c3f378a6f48bc7a11d9f933afd119" --registry false & \