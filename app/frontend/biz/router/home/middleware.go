// Code generated by hertz generator.

package home

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/crazyfrankie/bytedance-mall/app/frontend/middleware"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _homeMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _aboutMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		middleware.Auth(),
	}
}
