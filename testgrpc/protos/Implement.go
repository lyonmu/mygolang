package protos

// 远程调用接口的实现文件
import (
	"context"
	"log"
)

var Serve = &server{}

type server struct {
	UnimplementedGreeterServer
}

// 真正实现远程调用所需的结果的方法
func (s *server) ConToPro(ctx context.Context, myConsumer *MyConsumer) (*MyProduct, error) {
	log.Printf("发送的数字为: %v", myConsumer.GetMmber())
	return &MyProduct{Nnber: 18}, nil
}
