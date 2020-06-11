go run main.go --tcp 7002 --http 7003 --role producer --collection-prefix a_ --registry true --institution-name "Honestvote" --private-key "790bc20440d6dc17bff7a7b17547d4f5ea7782a4d30d4707878d5bb95cbd9676" --public-key "02e9e343da7bb7fbe2caf18315fa5c3f907f71c23caf9532accb225cb3aed1f6a3"  

sleep 5

go run main.go --tcp 7004 --http 7005 --role producer --collection-prefix b_ --registry-host 127.0.0.1 --registry-port 7002 --private-key "021f5d2717cc09d33fbbdee82d4abd95fbe365a80c84bc99ce4aa469697a21a7ab" --public-key "6514388d720c07b694377dc0236637e50e4a0cbb5c88d017d9df8c4f62c4ef10" --registry false & \

sleep 10

go run main.go --tcp 7006 --http 7007 --role producer --collection-prefix c_ --registry-host 127.0.0.1 --registry-port 7002 --private-key "0215d5ed4faa64e879009ac16b8703f599f57d87ef47bc9b0c2b487c4013f25368" --public-key "697f14c7d4c13158a1ac6feb8550d599fbba8f85f0b09df0196d038259a05996" --registry false & \