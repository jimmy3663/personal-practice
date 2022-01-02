# Kafka Configuration 

- Kafka property들을 주입하여 Producer와 Consumer에 대한 구체적인 설정을 하는 부분이다. 
- 해당 Config에는 publish를 담당하는 producer를 관리하는 producer factory와 sub을 담당하는 consumer를 관리하는 consumer factory 그리고 해당 factory를 container로 관리하는 container factory에 대해 설정한다. 

## Producer Factory 

```java
@Bean
public ProducerFactory<Object, Object> producerFactory() {
    return new DefaultKafkaProducerFactory<>(KafkaProperties.getProducerProperties(bootstrapServers));
}

```

- 