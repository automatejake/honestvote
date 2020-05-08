go run main.go --tcp 7002 --http 7003 --role producer --collection-prefix a_ --registry true --institution-name "Honestvote" --private-key "790bc20440d6dc17bff7a7b17547d4f5ea7782a4d30d4707878d5bb95cbd9676" --public-key "02e9e343da7bb7fbe2caf18315fa5c3f907f71c23caf9532accb225cb3aed1f6a3"  

sleep 5

go run main.go --tcp 7004 --http 7005 --role producer --collection-prefix b_ --registry-host 127.0.0.1 --registry-port 7002 --private-key "02d534be559432dc14bd59923a7cd30ac0365676272c4e24a481cc01d2a521e9ba" --public-key "e6effdcc6bb47ed6bdf462afa757433970a32794a5717bae1931b86f87558a20" & \

sleep 10

go run main.go --tcp 7006 --http 7007 --role producer --collection-prefix c_ --registry-host 127.0.0.1 --registry-port 7002 --private-key "03f546690e5c2cdf1b4cb3168d558564461ed59ead3259bd01cc251f37e70a2725" --public-key "c5b0aff1adcb84b54e3892471c505eef2fab95de5469054f499224e9884440d4" & \