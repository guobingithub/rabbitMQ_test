package main

import (
	"encoding/json"
	"github.com/streadway/amqp"
)

/***
提供给main方法的接口:::连接rabbitmq server，向mq中发送消息，初始化
*/
func Push() {
	if checkChannelNil() {
		mqConnect(EMIT)
	}

	initEmit()
}

/***
内部接口:::mq发送端初始化
*/
func initEmit() {
	registerEmit(channel_SkillGroup, exchange)
	registerEmit(channel_GroupAgentConfig, exchange)
	registerEmit(channel_VccConfig, exchange)
	registerEmit(channel_EvaluateOptions, exchange)
	registerEmit(channel_SessionRoutes, exchange)
	registerEmit(channel_WorkingHours, exchange)
	registerEmit(channel_AgentConfig, exchange)
	registerEmit(channel_WebChannel, exchange)
	registerEmit(channel_WechatChannel, exchange)
}

/***
内部接口:::给相应的channel声明exchange
*/
func registerEmit(channel *amqp.Channel, exchangeName string) {
	var err error
	err = channel.ExchangeDeclare(exchangeName, "topic", true, false, false, false, nil)
	failOnErr(err, "registerEmit, Failed to declare an exchange")
}

/***
内部接口:::技能组增/删/改，变化
*/
func skillGroupMsgEmit(msg interface{}) {
	msgContent, err := json.Marshal(msg)
	failOnErr(err, "skillGroupMsgEmit, Failed to json format")

	err = channel_SkillGroup.Publish(exchange, routeKey_SkillGroup, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(msgContent),
	})
	failOnErr(err, "skillGroupMsgEmit, Failed to publish a message")
	//if err == nil {
	//	logger.Info("skillGroupMsgEmit, Succed to publish a msg...")
	//}
}

func Test_skillGroupMsgEmit(count uint64)  {
	msgMap := map[string] interface{} {
		"Data":"I am skillGroupMsgEmit FFFFFFFF",
		"Count":count,
	}
	skillGroupMsgEmit(msgMap)
}

/***
内部接口:::技能组-坐席配置变化，维护相应表，通知排队
*/
func groupAgentConfigMsgEmit(msg interface{}) {
	msgContent, err := json.Marshal(msg)
	failOnErr(err, "groupAgentConfigMsgEmit, Failed to json format")

	err = channel_GroupAgentConfig.Publish(exchange, routeKey_GroupAgentConfig, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(msgContent),
	})
	failOnErr(err, "groupAgentConfigMsgEmit, Failed to publish a message")
	//if err == nil {
	//	logger.Info("groupAgentConfigMsgEmit, Succed to publish a msg...")
	//}
}

func Test_groupAgentConfigMsgEmit(count uint64)  {
	msgMap := map[string] interface{} {
		"Data":"I am groupAgentConfigMsgEmit FFFFFFFF",
		"Count":count,
	}
	groupAgentConfigMsgEmit(msgMap)
}

/***
内部接口:::企业配置改变，维护相应表，涉及到排队关心项则需通知排队
*/
func vccConfigMsgEmit(msg interface{}) {
	msgContent, err := json.Marshal(msg)
	failOnErr(err, "vccConfigMsgEmit, Failed to json format")

	err = channel_VccConfig.Publish(exchange, routeKey_VccConfig, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(msgContent),
	})
	failOnErr(err, "vccConfigMsgEmit, Failed to publish a message")
	//if err == nil {
	//	logger.Info("vccConfigMsgEmit, Succed to publish a msg...")
	//}
}

func Test_vccConfigMsgEmit(count uint64)  {
	msgMap := map[string] interface{} {
		"Data":"I am vccConfigMsgEmit FFFFFFFF",
		"Count":count,
	}
	vccConfigMsgEmit(msgMap)
}

/***
内部接口:::满意度选项表改变，维护相应表，涉及到排队关心项则需通知排队
*/
func evaluateOptionsMsgEmit(msg interface{}) {
	msgContent, err := json.Marshal(msg)
	failOnErr(err, "evaluateOptionsMsgEmit, Failed to json format")

	err = channel_EvaluateOptions.Publish(exchange, routeKey_EvaluateOptions, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(msgContent),
	})
	failOnErr(err, "evaluateOptionsMsgEmit, Failed to publish a message")
	//if err == nil {
	//	logger.Info("evaluateOptionsMsgEmit, Succed to publish a msg...")
	//}
}

func Test_evaluateOptionsMsgEmit(count uint64)  {
	msgMap := map[string] interface{} {
		"Data":"I am evaluateOptionsMsgEmit FFFFFFFF",
		"Count":count,
	}
	evaluateOptionsMsgEmit(msgMap)
}

/***
内部接口:::会话路由表改变，维护相应表，涉及到排队关心项则需通知排队
*/
func sessionRoutesMsgEmit(msg interface{}) {
	msgContent, err := json.Marshal(msg)
	failOnErr(err, "sessionRoutesMsgEmit, Failed to json format")

	err = channel_SessionRoutes.Publish(exchange, routeKey_SessionRoutes, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(msgContent),
	})
	failOnErr(err, "sessionRoutesMsgEmit, Failed to publish a message")
	//if err == nil {
	//	logger.Info("sessionRoutesMsgEmit, Succed to publish a msg...")
	//}
}

func Test_sessionRoutesMsgEmit(count uint64)  {
	msgMap := map[string] interface{} {
		"Data":"I am sessionRoutesMsgEmit FFFFFFFF",
		"Count":count,
	}
	sessionRoutesMsgEmit(msgMap)
}

/***
内部接口:::工作时间表改变，维护相应表，涉及到排队关心项则需通知排队
*/
func workingHoursMsgEmit(msg interface{}) {
	msgContent, err := json.Marshal(msg)
	failOnErr(err, "workingHoursMsgEmit, Failed to json format")

	err = channel_WorkingHours.Publish(exchange, routeKey_WorkingHours, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(msgContent),
	})
	failOnErr(err, "workingHoursMsgEmit, Failed to publish a message")
	//if err == nil {
	//	logger.Info("workingHoursMsgEmit, Succed to publish a msg...")
	//}
}

func Test_workingHoursMsgEmit(count uint64)  {
	msgMap := map[string] interface{} {
		"Data":"I am workingHoursMsgEmit FFFFFFFF",
		"Count":count,
	}
	workingHoursMsgEmit(msgMap)
}

/***
内部接口:::坐席规则配置改变，维护相应表，涉及到排队关心项则需通知排队
*/
func agentConfigMsgEmit(msg interface{}) {
	msgContent, err := json.Marshal(msg)
	failOnErr(err, "agentConfigMsgEmit, Failed to json format")

	err = channel_AgentConfig.Publish(exchange, routeKey_AgentConfig, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(msgContent),
	})
	failOnErr(err, "agentConfigMsgEmit, Failed to publish a message")
	//if err == nil {
	//	logger.Info("agentConfigMsgEmit, Succed to publish a msg...")
	//}
}

func Test_agentConfigMsgEmit(count uint64)  {
	msgMap := map[string] interface{} {
		"Data":"I am agentConfigMsgEmit FFFFFFFF",
		"Count":count,
	}
	agentConfigMsgEmit(msgMap)
}

/***
内部接口:::网站渠道信息改变，维护相应表，涉及到排队关心项则需通知排队
*/
func webChannelMsgEmit(msg interface{}) {
	msgContent, err := json.Marshal(msg)
	failOnErr(err, "webChannelMsgEmit, Failed to json format")

	err = channel_WebChannel.Publish(exchange, routeKey_WebChannel, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(msgContent),
	})
	failOnErr(err, "webChannelMsgEmit, Failed to publish a message")
	//if err == nil {
	//	logger.Info("webChannelMsgEmit, Succed to publish a msg...")
	//}
}

func Test_webChannelMsgEmit(count uint64)  {
	msgMap := map[string] interface{} {
		"Data":"I am webChannelMsgEmit FFFFFFFF",
		"Count":count,
	}
	webChannelMsgEmit(msgMap)
}

/***
内部接口:::微信渠道信息改变，维护相应表，涉及到排队关心项则需通知排队
*/
func wechatChannelMsgEmit(msg interface{}) {
	msgContent, err := json.Marshal(msg)
	failOnErr(err, "wechatChannelMsgEmit, Failed to json format")

	err = channel_WechatChannel.Publish(exchange, routeKey_WechatChannel, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(msgContent),
	})
	failOnErr(err, "wechatChannelMsgEmit, Failed to publish a message")
	//if err == nil {
	//	logger.Info("wechatChannelMsgEmit, Succed to publish a msg...")
	//}
}

func Test_wechatChannelMsgEmit(count uint64)  {
	msgMap := map[string] interface{} {
		"Data":"I am wechatChannelMsgEmit FFFFFFFF",
		"Count":count,
	}
	wechatChannelMsgEmit(msgMap)
}