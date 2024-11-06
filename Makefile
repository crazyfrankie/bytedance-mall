.PHONY:gen-frontend
gen-frontend:
	@cd ./app/frontend && cwgo server --I ../../idl/ --type HTTP --service frontend --module frontend --idl ../../idl/frontend/auth.proto