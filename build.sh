python3 -m grpc_tools.protoc -I./proto/ --python_out=. --pyi_out=. --grpc_python_out=./solvers/2015 proto/advent.proto