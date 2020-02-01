cd ../build


#----
# go run main.go --tcp 7000 --udp 7001 --role registry  & \
# sleep 5
go run main.go --tcp 7002 --http 7003 --role producer --collection-prefix a_ --registry true --institution-name "West Chester University" --private-key "307702010104207cec04b9ca5d29678b431a97cb1d642f4e2ada9e01c960e048ea484beab2417ea00a06082a8648ce3d030107a144034200047c1d5cd365f84677317b8aa975421ad08d1242b66d5163534e9e20f814e40a14e464f300658aaf0ff5e1551bf805290ba6c7ca4ee3b0b0c8abf920fe51861ca6" --public-key "3059301306072a8648ce3d020106082a8648ce3d030107034200047c1d5cd365f84677317b8aa975421ad08d1242b66d5163534e9e20f814e40a14e464f300658aaf0ff5e1551bf805290ba6c7ca4ee3b0b0c8abf920fe51861ca6" & \
sleep 5
go run main.go --tcp 7004 --http 7005 --role producer --collection-prefix b_ --registry-host 127.0.0.1 --registry-port 7002  & \
sleep 10
go run main.go --tcp 7006 --http 7007 --role producer --collection-prefix c_ --registry-host 127.0.0.1 --registry-port 7002  & \
#sleep 5
#go run main.go --tcp 7008 --http 7009 --role full --collection-prefix d_ --registry-host 127.0.0.1 --registry-port 7001 & \

