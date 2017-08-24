package main

import (
	"fmt"
	"time"
)

func main()  {
	fmt.Println("Test rabbitMQ start......")
	Push()			//send init
	go func() {
		for {
			count++
			Test_skillGroupMsgEmit(count)
			Test_groupAgentConfigMsgEmit(count)
			Test_vccConfigMsgEmit(count)
			Test_evaluateOptionsMsgEmit(count)
			Test_sessionRoutesMsgEmit(count)
			Test_workingHoursMsgEmit(count)
			Test_agentConfigMsgEmit(count)
			Test_webChannelMsgEmit(count)
			Test_wechatChannelMsgEmit(count)
			fmt.Println("==========================分割线====================================")
			fmt.Println("==========================分割线====================================")
			time.Sleep(15 * time.Second)
		}
	}()
	Receive()		//receive forever loop...
	fmt.Println("Test rabbitMQ end......")
	close()
}