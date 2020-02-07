cd ../build

go run main.go --tcp 7002 --http 7003 --role producer --collection-prefix a_ --registry true --institution-name "West Chester University" \
--private-key "30770201010420919e2356211bc962d7c664dd353760c4b9f1c1c9cc82b5d0d5ca2fa86994b29ca00a06082a8648ce3d030107a1440342000414bce0bafd6895f36df5069a1f5f6f752b46e92792bbf2a5baf0192f3fb541f638348264ccd161b706bdad1f3706086945038376be40af3c355a76094a72166d" \
--public-key "3059301306072a8648ce3d020106082a8648ce3d0301070342000414bce0bafd6895f36df5069a1f5f6f752b46e92792bbf2a5baf0192f3fb541f638348264ccd161b706bdad1f3706086945038376be40af3c355a76094a72166d" & \

sleep 5

go run main.go --tcp 7004 --http 7005 --role producer --collection-prefix b_ --registry-host 127.0.0.1 --registry-port 7002  & \

sleep 10

go run main.go --tcp 7006 --http 7007 --role producer --collection-prefix c_ --registry-host 127.0.0.1 --registry-port 7002  & \