// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	hello "frontend/biz/router/hello"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	hello.Register(r)
}
