go run main.go --tcp 7002 --http 7003 --role producer --collection-prefix a_ --registry true --institution-name "Honestvote" --private-key "790bc20440d6dc17bff7a7b17547d4f5ea7782a4d30d4707878d5bb95cbd9676" --public-key "02e9e343da7bb7fbe2caf18315fa5c3f907f71c23caf9532accb225cb3aed1f6a3"  

sleep 5

go run main.go --tcp 7004 --http 7005 --role producer --collection-prefix b_ --registry-host 127.0.0.1 --registry-port 7002 --private-key "038f4b53b94a36597949f5d9910749c97d4f5c951ed04e3c65691bac0c2738117b" --public-key "cef19189de1f3c135edc448c03424fbdb1f59efbe47cf829c706d48fcc5bd2aa" --registry false & \

sleep 10

go run main.go --tcp 7006 --http 7007 --role producer --collection-prefix c_ --registry-host 127.0.0.1 --registry-port 7002 --private-key "034d0264bb8c09c0eed070ee3b85fc35cd11d2e22eb3b15c4bddea33fa27058b42" --public-key "5a3610bfc758fb8d5653023cce6a2b06a15d89d95806b993c4caaa112947a5d8" --registry false & \