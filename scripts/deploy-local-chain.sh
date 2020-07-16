go run main.go --tcp 7002 --http 7003 --role producer --collection-prefix a_ --registry true --institution-name "Honestvote" --private-key "af92b4297523f83eb1ca057445cdd0b3582c5717e0b48c68603220f398cc9fce" --public-key "03e4973aa880e3dbd88210a063a3878ec80d5a688083f60f63577c3d6f0db784c0"  

sleep 5

go run main.go --tcp 7004 --http 7005 --role producer --collection-prefix b_ --registry-host 127.0.0.1 --registry-port 7002 --private-key "03c99d6df78e165e8a251913e648de4ee54a19fc4c78e9e87cd4288b42929b5791" --public-key "1732f64aa4acf3a44ea8da58be0c02c097c56c0dc680e5042f61bd8d3dca1fe6" --registry false & \

sleep 10

go run main.go --tcp 7006 --http 7007 --role producer --collection-prefix c_ --registry-host 127.0.0.1 --registry-port 7002 --private-key "029d37c9faec61b5686c688719610007f742fd825d92e2df3638c8b69ee2e65465" --public-key "c18d8df78a5689c2702ba1c19af873729835ac8872956068e8f2e64cc5972121" --registry false & \