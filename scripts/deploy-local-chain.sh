cd ../build

go run main.go --tcp 7002 --http 7003 --role producer --collection-prefix a_ --registry true --institution-name "West Chester University" --private-key "30770201010420d67cdc95f32fd2623704e24c10e2de29afbb5941a6f98d8d519e448dcaed1754a00a06082a8648ce3d030107a144034200048c2be6467d4a477ac8b5cbbded6528af7b6c44291853467448e585a4e57e3c7cdcb52646d192959a54c770f2c79cb6e7a0c3b716275588b4e7433aeb0128eac2" --public-key "3059301306072a8648ce3d020106082a8648ce3d030107034200048c2be6467d4a477ac8b5cbbded6528af7b6c44291853467448e585a4e57e3c7cdcb52646d192959a54c770f2c79cb6e7a0c3b716275588b4e7433aeb0128eac2" & \

sleep 5

go run main.go --tcp 7004 --http 7005 --role producer --collection-prefix b_ --registry-host 127.0.0.1 --registry-port 7002  & \

sleep 10

go run main.go --tcp 7006 --http 7007 --role producer --collection-prefix c_ --registry-host 127.0.0.1 --registry-port 7002  & \