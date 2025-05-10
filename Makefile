start_build_proto:
	mkdir coffeeshop_proto
	protoc --go_out=./coffeeshop_proto --go_opt=paths=source_relative --go-grpc_out=./coffeeshop_proto --go-grpc_opt=paths=source_relative coffee_shop.proto

build_proto:
	protoc --go_out=./coffeeshop_proto --go_opt=paths=source_relative --go-grpc_out=./coffeeshop_proto --go-grpc_opt=paths=source_relative coffee_shop.proto