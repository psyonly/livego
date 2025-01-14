package heart

import (
	"fmt"
	"github.com/gwuhaolin/livego/configure"
	"net/http"
	"time"
)

// 定期向注册中心发送心跳的函数
func SendHeart() {
	target := configure.Config.Get("register")
	url := fmt.Sprintf("http://%s/lg/heart?port=%s", target, configure.Config.Get("api_addr"))
	for {
		time.Sleep(time.Second*5)
		_, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
		}
	}
}

// SendStatus 是推流信号发送
func SendStatus(signal string) {
	target := configure.Config.Get("register")
	url := fmt.Sprintf("http://%s/lg/status?port=%s&status=%s", target, configure.Config.Get("api_addr"), signal)

	_, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
}
