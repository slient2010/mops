# mops
1)

set -e

rm -rf ./msg/pb/*.pb.go

cd ./msg/proto

protoc ./*.proto --go_out=../pb


2)

cd server/front

go-bindata -pkg=front -nocompress=true -debug=true html/...

go-bindata -pkg=front -nocompress=true html/...

go install mops

