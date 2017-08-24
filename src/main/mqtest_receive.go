package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"logger"
)

/***
提供给main方法的接口:::连接rabbitmq server，从mq中取出数据
*/
func Receive() {
	if checkChannelNil() {
		mqConnect(RECEIVE)
	}

	initReceive()

	forever := make(chan bool)
	go func() {
			fmt.Println("before connSuccessHandle=============")
			go connSuccessHandle(connSuccessMsg)
			go agentBreakHandle(agentBreakMsg)
			go agentBreakPassiveHandle(agentBreakPassiveMsg)
			go userBreakHandle(userBreakMsg)
			go connRedirectHandle(connRedirectMsg)
			go sessionHandle(sessionMsg)
			go evaluateHandle(evaluateMsg)
			go messageHandle(messageMsg)
			go agentStatusHandle(agentStatusMsg)
	}()

	logger.Info(" [*] Waiting for messages...")
	<-forever
}

/***
内部接口:::mq接收端初始化
*/
func initReceive() {
	registerConsume(channel_connSuccess, exchange, queue_ConnSuccess, routeKey_connSuccess, &connSuccessMsg)
	registerConsume(channel_agentBreak, exchange, queue_AgentBreak, routeKey_agentBreak, &agentBreakMsg)
	registerConsume(channel_agentBreakPassive, exchange, queue_AgentBreakPassive, routeKey_agentBreakPassive, &agentBreakPassiveMsg)
	registerConsume(channel_userBreak, exchange, queue_UserBreak, routeKey_userBreak, &userBreakMsg)
	registerConsume(channel_connRedirect, exchange, queue_ConnRedirect, routeKey_connRedirect, &connRedirectMsg)
	registerConsume(channel_session, exchange, queue_Session, routeKey_session, &sessionMsg)
	registerConsume(channel_evaluate, exchange, queue_Evaluate, routeKey_evaluate, &evaluateMsg)
	registerConsume(channel_message, exchange, queue_Message, routeKey_message, &messageMsg)
	registerConsume(channel_agentStatus, exchange, queue_AgentStatus, routeKey_agentStatus, &agentStatusMsg)
}

/***
内部接口:::注册消费者
*/
func registerConsume(channel *amqp.Channel, exchangeName, queueName, routeKey string, msgReceive *<-chan amqp.Delivery) {
	var err error
	var queue amqp.Queue

	err = channel.ExchangeDeclare(exchangeName, "topic", true, false, false, false, nil)
	failOnErr(err, "Failed to declare an exchange")

	queue, err = channel.QueueDeclare(queueName, false, false, true, false, nil)
	failOnErr(err, "Failed to declare a queue")

	err = channel.QueueBind(queue.Name, routeKey, exchangeName, false, nil)
	failOnErr(err, "Failed to bind a queue")

	*msgReceive, err = channel.Consume(queueName, "", true, false, false, false, nil)
	failOnErr(err, "")
}

/***
内部接口:::坐席和用户连接成功的消息处理
*/
func connSuccessHandle(msgs <-chan amqp.Delivery) {
	//fmt.Println("connSuccessHandle enter...msgs:", msgs)
	for d := range msgs {
		//fmt.Println("connSuccessHandle msg parse...")
		s := BytesToString(&(d.Body))
		count++
		fmt.Printf("receve msg is :%s -- %d\n", *s, count)
		//TODO::维护相应的数据表
	}
}

/***
内部接口:::坐席和用户断开连接，坐席方主动断开的消息处理
*/
func agentBreakHandle(msgs <-chan amqp.Delivery) {
	//fmt.Println("agentBreakHandle enter...msgs:", msgs)
	for d := range msgs {
		//fmt.Println("agentBreakHandle msg parse...")
		s := BytesToString(&(d.Body))
		count++
		fmt.Printf("receve msg is :%s -- %d\n", *s, count)
		//TODO::维护相应的数据表
	}
}

/***
内部接口:::坐席方被动断开(异常断开)的消息处理
*/
func agentBreakPassiveHandle(msgs <-chan amqp.Delivery) {
	//fmt.Println("agentBreakPassiveHandle enter...msgs:", msgs)
	for d := range msgs {
		//fmt.Println("agentBreakPassiveHandle msg parse...")
		s := BytesToString(&(d.Body))
		count++
		fmt.Printf("receve msg is :%s -- %d\n", *s, count)
		//TODO::维护相应的数据表
	}
}

/***
内部接口:::坐席和用户断开连接，用户方主动断开的消息处理
*/
func userBreakHandle(msgs <-chan amqp.Delivery) {
	//fmt.Println("userBreakHandle enter...msgs:", msgs)
	for d := range msgs {
		//fmt.Println("userBreakHandle msg parse...")
		s := BytesToString(&(d.Body))
		count++
		fmt.Printf("receve msg is :%s -- %d\n", *s, count)
		//TODO::维护相应的数据表
	}
}

/***
内部接口:::坐席触发转接的消息处理
*/
func connRedirectHandle(msgs <-chan amqp.Delivery) {
	//fmt.Println("connRedirectHandle enter...msgs:", msgs)
	for d := range msgs {
		//fmt.Println("connRedirectHandle msg parse...")
		s := BytesToString(&(d.Body))
		count++
		fmt.Printf("receve msg is :%s -- %d\n", *s, count)
		//TODO::维护相应的数据表
	}
}

/***
内部接口:::会话结束的消息处理
*/
func sessionHandle(msgs <-chan amqp.Delivery) {
	//fmt.Println("sessionHandle enter...msgs:", msgs)
	for d := range msgs {
		//fmt.Println("sessionHandle msg parse...")
		s := BytesToString(&(d.Body))
		count++
		fmt.Printf("receve msg is :%s -- %d\n", *s, count)
		//TODO::维护相应的数据表
	}
}

/***
内部接口:::满意度评价的消息处理
*/
func evaluateHandle(msgs <-chan amqp.Delivery) {
	//fmt.Println("evaluateHandle enter...msgs:", msgs)
	for d := range msgs {
		//fmt.Println("evaluateHandle msg parse...")
		s := BytesToString(&(d.Body))
		count++
		fmt.Printf("receve msg is :%s -- %d\n", *s, count)
		//TODO::维护相应的数据表
	}
}

/***
内部接口:::留言的消息处理
*/
func messageHandle(msgs <-chan amqp.Delivery) {
	//fmt.Println("messageHandle enter...msgs:", msgs)
	for d := range msgs {
		//fmt.Println("messageHandle msg parse...")
		s := BytesToString(&(d.Body))
		count++
		fmt.Printf("receve msg is :%s -- %d\n", *s, count)
		//TODO::维护相应的数据表
	}
}

/***
内部接口:::坐席状态变更的消息处理
*/
func agentStatusHandle(msgs <-chan amqp.Delivery) {
	//fmt.Println("agentStatusHandle enter...msgs:", msgs)
	for d := range msgs {
		//fmt.Println("agentStatusHandle msg parse...")
		s := BytesToString(&(d.Body))
		count++
		fmt.Printf("receve msg is :%s -- %d\n", *s, count)
		//TODO::维护相应的数据表
	}
}