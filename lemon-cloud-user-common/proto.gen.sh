# generate dto
rm dto/**.pb.go
protoc -I . --go_out=plugins=grpc:dto dto/*.proto
cp $(find dto/ -type f -name "*.pb.go") dto/
rm -rf dto/github.com

# generate service
rm service/**.pb.go
protoc -I . --go_out=plugins=grpc:service service/*.proto
cp $(find service/ -type f -name "*.pb.go") service/
rm -rf service/github.com
