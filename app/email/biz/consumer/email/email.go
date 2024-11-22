package email

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"github.com/crazyfrankie/bytedance-mall/app/email/infra/mq"
	"github.com/crazyfrankie/bytedance-mall/app/email/infra/notify"
	"github.com/crazyfrankie/bytedance-mall/rpc_gen/kitex_gen/email"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

func ConsumerInit() {
	sub, err := mq.Nc.Subscribe("email", func(msg *nats.Msg) {
		var req email.EmailReq
		err := proto.Unmarshal(msg.Data, &req)
		if err != nil {
			klog.Error(err)
			return
		}

		noopEmail := notify.NewNoopEmail()
		_ = noopEmail.Send(&req)
	})

	if err != nil {
		panic(err)
	}

	server.RegisterShutdownHook(func() {
		sub.Unsubscribe()
		mq.Nc.Close()
	})
}
