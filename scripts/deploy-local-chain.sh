cd ../build


#----
go run main.go --tcp 7000 --udp 7001 --role registry  & \
sleep 5
go run main.go --tcp 7002 --http 7003 --role peer --collection-prefix a_ --registry-host 127.0.0.1 --registry-port 7001 & \
sleep 5
go run main.go --tcp 7004 --http 7005 --role peer --collection-prefix b_ --registry-host 127.0.0.1 --registry-port 7001  & \
sleep 5
go run main.go --tcp 7006 --http 7007 --role peer --collection-prefix c_ --registry-host 127.0.0.1 --registry-port 7001  & \
sleep 5
go run main.go --tcp 7008 --http 7009 --role full --collection-prefix d_ --registry-host 127.0.0.1 --registry-port 7001 & \

