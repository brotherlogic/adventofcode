#python3 -m grpc_tools.protoc -I./proto/ --python_out=./solvers/2015 --pyi_out=. --grpc_python_out=./solvers/2015 proto/advent.proto
protoc --proto_path ../../../ -I=./proto --go_out ./proto --go_opt=paths=source_relative --go-grpc_out ./proto --go-grpc_opt paths=source_relative --go-grpc_opt=require_unimplemented_servers=false proto/advent.proto
#protoc --proto_path ./ --csharp_out ./proto --plugin=protoc-gen-go-grpc --grpc_out=./proto proto/advent.proto
#grpc_tools_ruby_protoc -I ./proto --ruby_out ./proto --grpc_out ./proto proto/advent.proto
