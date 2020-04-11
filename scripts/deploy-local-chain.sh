go run main.go --tcp 7002 --http 7003 --role producer --collection-prefix a_ --registry true --institution-name "Honestvote"

sleep 5

go run main.go --tcp 7004 --http 7005 --role producer --collection-prefix b_ --registry-host 127.0.0.1 --registry-port 7002 --private-key "024bb2ffffef540f7850761e100f9aacda2317ea759da0800ce5bcb76fa8e5d4f5" --public-key "ecc6643d208e7f549ed674a0fc05cae69153f6a8ad30e902e123bb6a3b946110" & \

sleep 10

go run main.go --tcp 7006 --http 7007 --role producer --collection-prefix c_ --registry-host 127.0.0.1 --registry-port 7002 --private-key "03bb7128975bd6797e450e7e59c365052aff0f3a5655d1b92f9802cb828ff73fcd" --public-key "11724f46611ddf1b834efc72223a40bf07a7e65399f29341b50593fdd372f208" & \