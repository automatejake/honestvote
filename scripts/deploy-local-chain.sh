go run main.go --tcp 7002 --http 7003 --role producer --collection-prefix a_ --registry true --institution-name "Honestvote" --private-key "790bc20440d6dc17bff7a7b17547d4f5ea7782a4d30d4707878d5bb95cbd9676"--public-key "02e9e343da7bb7fbe2caf18315fa5c3f907f71c23caf9532accb225cb3aed1f6a3"

sleep 5

go run main.go --tcp 7004 --http 7005 --role producer --collection-prefix b_ --registry-host 127.0.0.1 --registry-port 7002 --private-key "03c49530245e86f122fbdeb33929d8f1280f5ef6024f122ba87caa90818f637747" --public-key "f7d08c91d295b461f4906e71b90644bdca831e891eaaec98f4d7bc2a65b4c760" & \

sleep 10

go run main.go --tcp 7006 --http 7007 --role producer --collection-prefix c_ --registry-host 127.0.0.1 --registry-port 7002 --private-key "0394ff99af0abe7de260046bf330ca60d341d23d0890ceb8d3ab9ac839e40ee10b" --public-key "285f65cfcbba3e3716c54cd5d006f4913f4cfc9c67232e44429d17a4cc8d0f4b" & \