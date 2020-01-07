#ln -s $(pwd)/pingpong-proto/ go-grpc-server
#ln -s $(pwd)/pingpong-proto/ go-grpc-client
cp $(pwd)/pingpong-proto/generated/grpc/*.pb go-grpc-server/api/
cp $(pwd)/pingpong-proto/generated/grpc/*.pb go-grpc-client/api/