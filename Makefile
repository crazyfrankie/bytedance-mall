export ROOT_MOD=github.com/crazyfrankie/bytedance-mall

.PHONY:gen-frontend
gen-frontend:
	@cd ./app/frontend && cwgo server --I ../../idl/ --type HTTP --service frontend --module ${ROOT_MOD}/app/frontend --idl ../../idl/frontend/home.proto

.PHONY: gen-user
gen-user:
	@cd ./rpc_gen && cwgo client --type RPC --service user --module ${ROOT_MOD}/rpc_gen -I ../idl --idl ../idl/user.proto 
	@cd ./app/user && cwgo server --type RPC --service user --module ${ROOT_MOD}/app/user --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/user.proto

.PHONY: gen-product
gen-product:
	@cd ./rpc_gen && cwgo client --type RPC --service product --module ${ROOT_MOD}/rpc_gen -I ../idl --idl ../idl/product.proto 
	@cd ./app/product && cwgo server --type RPC --service product --module ${ROOT_MOD}/app/product --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/product.proto
