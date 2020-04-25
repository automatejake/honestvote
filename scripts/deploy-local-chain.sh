go run main.go --tcp 7002 --http 7003 --role producer --collection-prefix a_ --registry true --institution-name "Honestvote" --private-key "790bc20440d6dc17bff7a7b17547d4f5ea7782a4d30d4707878d5bb95cbd9676" --public-key "02e9e343da7bb7fbe2caf18315fa5c3f907f71c23caf9532accb225cb3aed1f6a3"  

sleep 5

go run main.go --tcp 7004 --http 7005 --role producer --collection-prefix b_ --registry-host 127.0.0.1 --registry-port 7002 --private-key "022c52e4f90a35d44d9df7d363fd2a6160da753a2e1b44d361108cf6cb9a20562d" --public-key "b4099068941f7871f98dbd0eaee334d6dea464f7490cc9226a68f278be56ff64" & \

sleep 10

go run main.go --tcp 7006 --http 7007 --role producer --collection-prefix c_ --registry-host 127.0.0.1 --registry-port 7002 --private-key "03be32d337962a88f8847bceac4c4af51b98f05124ad8200b2e8bfe8d3e9097e31" --public-key "6494ea971123a1767b638d02369302cb053c787bdf6ac8b6c2ccaeb71cb068c9" & \