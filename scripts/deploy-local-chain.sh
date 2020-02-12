go run main.go --tcp 7002 --http 7003 --role producer --collection-prefix a_ --registry true --institution-name "West Chester University" \
--private-key "30770201010420624e00a5fe1c2aad3c7916df88e0292c576a085a677278cea12167f7adb30093a00a06082a8648ce3d030107a144034200043269053de5c0ee04f06a2aa54087fbe0603a92c1ac4ffd8556ac1687e9f65d48e5c2a0b1b220199a700c1c695e5cf2c2be38f1d042aa4beb770b27e8350cae33" \
--public-key "3059301306072a8648ce3d020106082a8648ce3d030107034200043269053de5c0ee04f06a2aa54087fbe0603a92c1ac4ffd8556ac1687e9f65d48e5c2a0b1b220199a700c1c695e5cf2c2be38f1d042aa4beb770b27e8350cae33" & \

sleep 5

go run main.go --tcp 7004 --http 7005 --role producer --collection-prefix b_ --registry-host 127.0.0.1 --registry-port 7002  & \

sleep 10

go run main.go --tcp 7006 --http 7007 --role producer --collection-prefix c_ --registry-host 127.0.0.1 --registry-port 7002  & \