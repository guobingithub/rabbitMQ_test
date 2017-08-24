package main

import (
	"bytes"
	"fmt"
	"github.com/streadway/amqp"
	"logger"
)

var conn *amqp.Connection

//发送方channel
var channel_SkillGroup *amqp.Channel
var channel_GroupAgentConfig *amqp.Channel
var channel_VccConfig *amqp.Channel
var channel_EvaluateOptions *amqp.Channel
var channel_SessionRoutes *amqp.Channel
var channel_WorkingHours *amqp.Channel
var channel_AgentConfig *amqp.Channel
var channel_WebChannel *amqp.Channel
var channel_WechatChannel *amqp.Channel

//接收方channel
var channel_connSuccess *amqp.Channel
var channel_agentBreak *amqp.Channel
var channel_agentBreakPassive *amqp.Channel
var channel_userBreak *amqp.Channel
var channel_connRedirect *amqp.Channel
var channel_session *amqp.Channel
var channel_evaluate *amqp.Channel
var channel_message *amqp.Channel
var channel_agentStatus *amqp.Channel

var count uint64 = 0

//接收的待同步消息
var connSuccessMsg = make(<-chan amqp.Delivery)
var agentBreakMsg = make(<-chan amqp.Delivery)
var agentBreakPassiveMsg = make(<-chan amqp.Delivery)
var userBreakMsg = make(<-chan amqp.Delivery)
var connRedirectMsg = make(<-chan amqp.Delivery)
var sessionMsg = make(<-chan amqp.Delivery)
var evaluateMsg = make(<-chan amqp.Delivery)
var messageMsg = make(<-chan amqp.Delivery)
var agentStatusMsg = make(<-chan amqp.Delivery)

const (
	RECEIVE  = "receive"
	EMIT     = "emit"
	exchange = "test.msg.ex"
	mqurl    = "amqp://guest:guest@localhost:5672/"

	//接收方
	queue_ConnSuccess       = "queue_ConnSuccess"
	queue_AgentBreak = "queue_AgentBreak"
	queue_AgentBreakPassive = "queue_AgentBreakPassive"
	queue_UserBreak = "queue_UserBreak"
	queue_ConnRedirect = "queue_ConnRedirect"
	queue_Session = "queue_Session"
	queue_Evaluate = "queue_Evaluate"
	queue_Message = "queue_Message"
	queue_AgentStatus = "queue_AgentStatus"


	//路由定义，需要与后台那边接收方定义一致
	routeKey_SkillGroup       = "routeKey_SkillGroup"
	routeKey_GroupAgentConfig = "routeKey_GroupAgentConfig"
	routeKey_VccConfig        = "routeKey_VccConfig"
	routeKey_EvaluateOptions  = "routeKey_EvaluateOptions"
	routeKey_SessionRoutes    = "routeKey_SessionRoutes"
	routeKey_WorkingHours     = "routeKey_WorkingHours"
	routeKey_AgentConfig      = "routeKey_AgentConfig"
	routeKey_WebChannel       = "routeKey_WebChannel"
	routeKey_WechatChannel    = "routeKey_WechatChannel"


	//路由定义，需要与后台那边发送方定义一致
	routeKey_connSuccess       = "routeKey_connSuccess"
	routeKey_agentBreak        = "routeKey_agentBreak"
	routeKey_agentBreakPassive = "routeKey_agentBreakPassive"
	routeKey_userBreak         = "routeKey_userBreak"
	routeKey_connRedirect      = "routeKey_connRedirect"
	routeKey_session           = "routeKey_session"
	routeKey_evaluate          = "routeKey_evaluate"
	routeKey_message           = "routeKey_message"
	routeKey_agentStatus       = "routeKey_agentStatus"
)

func failOnErr(err error, msg string) {
	if err != nil {
		logger.Fatal(fmt.Sprintf("%s:%s", msg, err))
		panic(fmt.Sprintf("%s:%s", msg, err))
	}
}

func BytesToString(b *[]byte) *string {
	s := bytes.NewBuffer(*b)
	r := s.String()
	return &r
}

//连接rabbitmq server
func mqConnect(kind string) {
	if conn == nil {
		var err error
		conn, err = amqp.Dial(mqurl)
		failOnErr(err, "failed to connect tp rabbitmq")
	}

	if kind == RECEIVE {
		openChannelForReceive(conn)
	} else if kind == EMIT {
		openChannelForEmit(conn)
	}
}

func openChannelForReceive(conn *amqp.Connection) {
	var err error
	channel_connSuccess, err = conn.Channel()
	failOnErr(err, "failed to open a channel")
	channel_agentBreak, err = conn.Channel()
	failOnErr(err, "failed to open a channel")
	channel_agentBreakPassive, err = conn.Channel()
	failOnErr(err, "failed to open a channel")
	channel_userBreak, err = conn.Channel()
	failOnErr(err, "failed to open a channel")
	channel_connRedirect, err = conn.Channel()
	failOnErr(err, "failed to open a channel")
	channel_session, err = conn.Channel()
	failOnErr(err, "failed to open a channel")
	channel_evaluate, err = conn.Channel()
	failOnErr(err, "failed to open a channel")
	channel_message, err = conn.Channel()
	failOnErr(err, "failed to open a channel")
	channel_agentStatus, err = conn.Channel()
	failOnErr(err, "failed to open a channel")
}

func openChannelForEmit(conn *amqp.Connection) {
	var err error
	channel_SkillGroup, err = conn.Channel()
	failOnErr(err, "failed to open a channel")
	channel_GroupAgentConfig, err = conn.Channel()
	failOnErr(err, "failed to open a channel")
	channel_VccConfig, err = conn.Channel()
	failOnErr(err, "failed to open a channel")
	channel_EvaluateOptions, err = conn.Channel()
	failOnErr(err, "failed to open a channel")
	channel_SessionRoutes, err = conn.Channel()
	failOnErr(err, "failed to open a channel")
	channel_WorkingHours, err = conn.Channel()
	failOnErr(err, "failed to open a channel")
	channel_AgentConfig, err = conn.Channel()
	failOnErr(err, "failed to open a channel")
	channel_WebChannel, err = conn.Channel()
	failOnErr(err, "failed to open a channel")
	channel_WechatChannel, err = conn.Channel()
	failOnErr(err, "failed to open a channel")
}

func close() {
	channel_SkillGroup.Close()
	channel_GroupAgentConfig.Close()
	channel_VccConfig.Close()
	channel_EvaluateOptions.Close()
	channel_SessionRoutes.Close()
	channel_WorkingHours.Close()
	channel_AgentConfig.Close()
	channel_WebChannel.Close()
	channel_WechatChannel.Close()

	channel_connSuccess.Close()
	channel_agentBreak.Close()
	channel_agentBreakPassive.Close()
	channel_userBreak.Close()
	channel_connRedirect.Close()
	channel_session.Close()
	channel_evaluate.Close()
	channel_message.Close()
	channel_agentStatus.Close()

	conn.Close()
}

func checkChannelNil() bool {
	if channel_SkillGroup == nil ||
		channel_GroupAgentConfig == nil ||
		channel_VccConfig == nil ||
		channel_EvaluateOptions == nil ||
		channel_SessionRoutes == nil ||
		channel_WorkingHours == nil ||
		channel_AgentConfig == nil ||
		channel_WebChannel == nil ||
		channel_WechatChannel == nil ||
		channel_connSuccess == nil ||
		channel_agentBreak == nil ||
		channel_agentBreakPassive == nil ||
		channel_userBreak == nil ||
		channel_connRedirect == nil ||
		channel_session == nil ||
		channel_evaluate == nil ||
		channel_message == nil ||
		channel_agentStatus == nil {
		return true
	} else {
		return false
	}
}
