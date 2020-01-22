cd "C://Program Files (x86)/Nmap"

echo '{"message":"transaction", "type":"Election"}' | ./ncat localhost 7004 & \
sleep 1
echo '{"message":"transaction", "type":"Vote"}' | ./ncat localhost 7004 & \