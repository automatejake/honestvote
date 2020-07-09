go run main.go --tcp 7002 --http 7003 --role producer --collection-prefix a_ --registry true --institution-name "Honestvote" --private-key "af92b4297523f83eb1ca057445cdd0b3582c5717e0b48c68603220f398cc9fce" --public-key "03e4973aa880e3dbd88210a063a3878ec80d5a688083f60f63577c3d6f0db784c0"  

sleep 5

go run main.go --tcp 7004 --http 7005 --role producer --collection-prefix b_ --registry-host 127.0.0.1 --registry-port 7002 --private-key "02cc8ea7ccd568ae2d2994a8bd4d0d678e26303a622fcb54d340495ef0a6ea1140" --public-key "9dc836c00a33684c159259270f56d6b7064bbd0415771a720c5c8ff75b1a5a05" --registry false & \

sleep 10

go run main.go --tcp 7006 --http 7007 --role producer --collection-prefix c_ --registry-host 127.0.0.1 --registry-port 7002 --private-key "035fde7b91496e0e3d9d0b1fca6ff0123753230107a86632cfe8e746b7c5c72cf6" --public-key "061abdec0f6041f08e62b3d9fbae55700b4f5daf2931312cebca04c34e6910cb" --registry false & \