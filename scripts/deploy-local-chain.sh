go run main.go --tcp 7002 --http 7003 --role producer --collection-prefix a_ --registry true --institution-name "West Chester University" \
--private-key "30770201010420685a8b0d114ea3fe8d0a158db6c6c918d6991a8b3a7a3819e06e5eeedf361b50a00a06082a8648ce3d030107a144034200041ca9ef0f92d2a1143cea968205160e3ebb24dad7b63feef36a231445201cd42d89a2303b24819266d53848db9b60bff507560f7df362c0fd82ca1fc19408656f" \
--public-key "3059301306072a8648ce3d020106082a8648ce3d030107034200041ca9ef0f92d2a1143cea968205160e3ebb24dad7b63feef36a231445201cd42d89a2303b24819266d53848db9b60bff507560f7df362c0fd82ca1fc19408656f" & \

sleep 5

go run main.go --tcp 7004 --http 7005 --role producer --collection-prefix b_ --registry-host 127.0.0.1 --registry-port 7002  & \

sleep 10

go run main.go --tcp 7006 --http 7007 --role producer --collection-prefix c_ --registry-host 127.0.0.1 --registry-port 7002  & \