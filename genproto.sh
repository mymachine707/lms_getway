protoc --go_out=./protogen \
    --go-grpc_out=./protogen \
    ./lms_proto/book_service/*.proto