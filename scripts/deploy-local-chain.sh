go run main.go --tcp 7002 --http 7003 --role producer --collection-prefix a_ --registry true --institution-name "Honestvote" --private-key "af92b4297523f83eb1ca057445cdd0b3582c5717e0b48c68603220f398cc9fce" --public-key "03e4973aa880e3dbd88210a063a3878ec80d5a688083f60f63577c3d6f0db784c0"  

sleep 5

go run main.go --tcp 7004 --http 7005 --role producer --collection-prefix b_ --registry-host 127.0.0.1 --registry-port 7002 --private-key "0214d0d2c497f2123353671d8b3a7b906e5a7b07c297fcce1b70abea8e76e28688" --public-key "869e09e1204034a9599c650dbfc4196058f3c09209d80438b4d1cff337c51a02" --registry false & \

sleep 10

go run main.go --tcp 7006 --http 7007 --role producer --collection-prefix c_ --registry-host 127.0.0.1 --registry-port 7002 --private-key "0266f622ff71b23203df33125e921e00c6a28aedfd6493af8609a0ed84a25c5489" --public-key "121cb22c1afe3722ab6dc30743f84016b9d4377d67ae339c295ca13ee69b3a2f" --registry false & \