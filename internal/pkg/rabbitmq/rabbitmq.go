package rabbitmq

import (
	"fmt"
	"log"
	"xuetu-project/config"

	"github.com/streadway/amqp"
)

// RabbitMQ 全局RabbitMQ连接和通道
var (
	Connection *amqp.Connection
	Channel    *amqp.Channel
)

// InitRabbitMQ 初始化RabbitMQ连接
func InitRabbitMQ() {
	cfg := config.AppConfig.RabbitMQ

	// 构建连接字符串
	connStr := fmt.Sprintf("amqp://%s:%s@%s:%d/", cfg.Username, cfg.Password, cfg.Host, cfg.Port)

	// 建立连接
	conn, err := amqp.Dial(connStr)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	// 创建通道
	channel, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}

	// 设置全局变量
	Connection = conn
	Channel = channel

	log.Println("Successfully connected to RabbitMQ")
}

// Publish 发布消息到队列
func Publish(queueName, message string) error {
	q, err := Channel.QueueDeclare(
		queueName, // 队列名称
		true,      // durable - 持久化
		false,     // autoDelete - 自动删除
		false,     // exclusive - 排他性
		false,     // noWait - 非阻塞
		nil,       // 参数
	)
	if err != nil {
		return fmt.Errorf("failed to declare a queue: %v", err)
	}

	err = Channel.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	if err != nil {
		return fmt.Errorf("failed to publish a message: %v", err)
	}

	return nil
}

// Consume 从队列消费消息
func Consume(queueName string, handler func(msgs <-chan amqp.Delivery)) error {
	q, err := Channel.QueueDeclare(
		queueName, // 队列名称
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare a queue: %v", err)
	}

	msgs, err := Channel.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		return fmt.Errorf("failed to register a consumer: %v", err)
	}

	// 启动消费者处理
	handler(msgs)

	return nil
}

// Close 关闭RabbitMQ连接
func Close() error {
	if Channel != nil {
		if err := Channel.Close(); err != nil {
			return err
		}
	}
	if Connection != nil {
		if err := Connection.Close(); err != nil {
			return err
		}
	}
	return nil
}

// 使用例子:
//err := rabbitmq.Publish("notifications", notificationMsg)
