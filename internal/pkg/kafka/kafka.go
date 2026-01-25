package kafka

//// KafkaProducer 全局Kafka生产者
//var KafkaProducer sarama.SyncProducer
//
//// KafkaConsumer 全局Kafka消费者
//var KafkaConsumer sarama.Consumer
//
//// 初始化互斥锁，确保线程安全
//var initOnce sync.Once
//
//// InitKafka 初始化Kafka连接
//func InitKafka() {
//	initOnce.Do(func() {
//		cfg := config.AppConfig.Kafka
//
//		// 配置Kafka生产者
//		producerConfig := sarama.NewConfig()
//		producerConfig.Producer.RequiredAcks = sarama.WaitForAll
//		producerConfig.Producer.Retry.Max = 5
//		producerConfig.Producer.Return.Successes = true
//
//		// 配置TLS（如果需要）
//		if cfg.TLS.Enable {
//			producerConfig.Net.TLS.Enable = true
//			producerConfig.Net.TLS.Config = &tls.Config{}
//		}
//
//		// 创建生产者
//		producer, err := sarama.NewSyncProducer(cfg.Brokers, producerConfig)
//		if err != nil {
//			log.Fatalf("Failed to create Kafka producer: %v", err)
//		}
//
//		// 配置Kafka消费者
//		consumerConfig := sarama.NewConfig()
//		consumerConfig.Consumer.Return.Errors = true
//
//		// 配置TLS（如果需要）
//		if cfg.TLS.Enable {
//			consumerConfig.Net.TLS.Enable = true
//			consumerConfig.Net.TLS.Config = &tls.Config{}
//		}
//
//		// 创建消费者
//		consumer, err := sarama.NewConsumer(cfg.Brokers, consumerConfig)
//		if err != nil {
//			log.Fatalf("Failed to create Kafka consumer: %v", err)
//		}
//
//		// 设置全局变量
//		KafkaProducer = producer
//		KafkaConsumer = consumer
//
//		log.Println("Successfully connected to Kafka")
//	})
//}
//
//// Publish 发布消息到Kafka主题
//func Publish(topic string, message []byte) error {
//	msg := &sarama.ProducerMessage{
//		Topic: topic,
//		Value: sarama.ByteEncoder(message),
//	}
//
//	partition, offset, err := KafkaProducer.SendMessage(msg)
//	if err != nil {
//		return fmt.Errorf("failed to send message to Kafka: %v", err)
//	}
//
//	log.Printf("Message sent to partition %d at offset %d", partition, offset)
//	return nil
//}
//
//// Consume 从Kafka主题消费消息
//func Consume(topics []string, groupId string, handler func(*sarama.ConsumerMessage)) error {
//	// 创建消费者组
//	config := sarama.NewConfig()
//	config.Consumer.Return.Errors = true
//
//	master, err := sarama.NewConsumer(config.AppConfig.Kafka.Brokers, config)
//	if err != nil {
//		return fmt.Errorf("failed to create consumer: %v", err)
//	}
//	defer master.Close()
//
//	for _, topic := range topics {
//		partitionList, err := master.Partitions(topic)
//		if err != nil {
//			return fmt.Errorf("failed to get partitions for topic %s: %v", topic, err)
//		}
//
//		var consumers []sarama.PartitionConsumer
//		for _, partition := range partitionList {
//			pc, err := master.ConsumePartition(topic, partition, sarama.OffsetNewest)
//			if err != nil {
//				return fmt.Errorf("failed to consume partition %d for topic %s: %v", partition, topic, err)
//			}
//			consumers = append(consumers, pc)
//		}
//
//		// 启动goroutine处理消息
//		go func(pc sarama.PartitionConsumer) {
//			defer pc.AsyncClose()
//			for {
//				select {
//				case msg := <-pc.Messages():
//					handler(msg)
//				case err := <-pc.Errors():
//					log.Printf("Consumer error: %v", err)
//				}
//			}
//		}(consumers[0]) // 简单示例，只处理第一个分区
//	}
//
//	return nil
//}
//
//// Close 关闭Kafka连接
//func Close() error {
//	var err error
//
//	if KafkaProducer != nil {
//		if closeErr := KafkaProducer.Close(); closeErr != nil {
//			err = closeErr
//		}
//	}
//
//	if KafkaConsumer != nil {
//		if closeErr := KafkaConsumer.Close(); closeErr != nil {
//			if err != nil {
//				err = fmt.Errorf("errors closing both producer and consumer: %v, %v", err, closeErr)
//			} else {
//				err = closeErr
//			}
//		}
//	}
//
//	return err
//}

//kafka:
//brokers:
//- localhost:9092
//tls:
//enable: false
//cert_file: ""
//key_file: ""
//ca_file: ""
//
//// KafkaTLSConfig Kafka TLS配置
//type KafkaTLSConfig struct {
//	Enable   bool   `mapstructure:"enable"`
//	CertFile string `mapstructure:"cert_file"`
//	KeyFile  string `mapstructure:"key_file"`
//	CAFile   string `mapstructure:"ca_file"`
//}
//
//// KafkaConfig Kafka相关配置
//type KafkaConfig struct {
//	Brokers []string       `mapstructure:"brokers"`
//	TLS     KafkaTLSConfig `mapstructure:"tls"`
//}
//kafka.InitKafka()
//defer kafka.Close()  // 确保程序退出时关闭连接

//err := kafka.Publish("notifications", []byte(notificationMsg))
