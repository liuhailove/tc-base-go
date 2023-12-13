protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --proto_path=/ /*.proto



protoc --twirp_out=. --proto_path=proto/ proto/*.proto


protoc --twirp_out=. --proto_path=. *.proto



pbjs --ts --keep-case --out tc_models.ts tc_models.proto


