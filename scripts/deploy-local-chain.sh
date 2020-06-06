go run main.go --tcp 7002 --http 7003 --role producer --collection-prefix a_ --registry true --institution-name "Honestvote" --private-key "790bc20440d6dc17bff7a7b17547d4f5ea7782a4d30d4707878d5bb95cbd9676" --public-key "02e9e343da7bb7fbe2caf18315fa5c3f907f71c23caf9532accb225cb3aed1f6a3"  

sleep 5

go run main.go --tcp 7004 --http 7005 --role producer --collection-prefix b_ --registry-host 127.0.0.1 --registry-port 7002 --private-key "0321450893cb2800d3333916483d9161eb4e3081f5e628559d070fbe17e9a5b8c5" --public-key "4cea8d29922d9a55adf282205eefe8f87988dc64b435797a8e4d8e63b6857c2c" --registry false & \

sleep 10

go run main.go --tcp 7006 --http 7007 --role producer --collection-prefix c_ --registry-host 127.0.0.1 --registry-port 7002 --private-key "031b9c237e5eb9663ebabdfec106c52aa475a4dc0be0866d9f84e44eefde2b39d4" --public-key "78d8d5e3f861e960c1c4d195dcfa0a7c8a20b9f5ced6ab05c48c3479fcb49af2" --registry false & \