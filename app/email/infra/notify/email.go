package notify

import (
	"github.com/crazyfrankie/bytedance-mall/rpc_gen/kitex_gen/email"
	"github.com/kr/pretty"
)

type NoopEmail struct {
}

func (e *NoopEmail) Send(req *email.EmailReq) error {
	pretty.Printf("%v\n", req)
	return nil
}

func NewNoopEmail() NoopEmail {
	return NoopEmail{}
}
