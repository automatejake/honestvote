cd ../build
go run main.go --peer 7000 --http 7001 --role peer & \
go run main.go --peer 7002 --http 7003
