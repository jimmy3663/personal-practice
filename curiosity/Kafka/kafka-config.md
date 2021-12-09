# Kafka Config 

- Kafka 설정에 대해 써보려고 한다. 

```
public static Map<Object, Object> getConsumerProperties(String bootstrapServers) {
        Map<Object, Object> props = new HashMap<>();
        props.put(ConsumerConfig.BOOTSTRAP_SERVERS_CONFIG, bootstrapServers);
        props.put(ConsumerConfig.ENABLE_AUTO_COMMIT_CONFIG, false); 
        props.put(ConsumerConfig.AUTO_OFFSET_RESET_CONFIG, "earliest"); 
        props.put(ConsumerConfig.KEY_DESERIALIZER_CLASS_CONFIG, StringDeserializer.class);
        props.put(ConsumerConfig.VALUE_DESERIALIZER_CLASS_CONFIG, JsonDeserializer.class);
        props.put(ConsumerConfig.ISOLATION_LEVEL_CONFIG, "read_committed");
        props.put(JsonDeserializer.TRUSTED_PACKAGES,"*");

        return props;

    }

    public static Map<Object, Object> getProducerProperties(String bootstrapServers){
        Map<Object, Object> props = new HashMap<>();

        props.put(ProducerConfig.BOOTSTRAP_SERVERS_CONFIG, bootstrapServers);
        props.put(ProducerConfig.ACKS_CONFIG, "all");
        props.put(ProducerConfig.KEY_SERIALIZER_CLASS_CONFIG, StringSerializer.class);
        props.put(ProducerConfig.VALUE_SERIALIZER_CLASS_CONFIG, JsonSerializer.class);
        props.put(ProducerConfig.TRANSACTIONAL_ID_CONFIG, UUID.randomUUID().toString());

        return props;
    }
```

## 상세 내용 

- 카프카 설정에는 크게 Consumer, Producer로 나뉘어 설정해준다. 

### Producer Config 

- ProducerConfig.ACKS_CONFIG : 브로커는 파티션에 대한 복제본을 만들 수 있다. 이를 replica라고 부르게 되는데 데이터베이스와 동일하게 master 와 slave 로 나뉘어 master에게는 파티션에 r/w 할 권한이 있고 slave들에게는 read의 권한 밖에 존재하지 않느다. 여기서 ACKS_CONFIG에는 0, 1, all 이라는 옵션이 있게 되는데 0 일때는 producer가 보낸 메세지에 대해 master 든 slave 든 잘 받았는지 확인을 안한다. 1 일때는 보낸 메세지에 대해 master가 제대로 받았는지 ack를 확인 한다. all 일때는 master는 물론이고 master가 slave에게 까지 write를 제대로 했는지에 대한 ack를 받는다. 성능면에선 0 > 1 > all 순으로 좋지만 강건한 순으로 보자면 all > 1 > 0 순으로 좋아지게 된다. 

- ProducerConfig.KEY_SERIALIZER_CLASS_CONFIG : Producer가 message를 보낼 때 key와 data 형태로 보내게 된다. 이때 key를 어떻게 직렬화를 하여 보낼 껀지에 대한 config이다. Key Serializer에는 String, Integer, Json, Confluent의 Avro serializer들이 존재하는데 각 장단점 보다는 어떤 타입의 key를 쓰고 싶느냐에 따라 설정해주면 된다. 

- ProducerConfig.VALUE_SERIALIZER_CLASS_CONFIG : Producer가 message를 보낼 때 data에 해당 하는 message를 직렬화 하는 형태이다. Value Serializer 또한 key Serializer와 같은 serializer들이 존재한다. 

(깊이 안파보면 모르는 사실이 직렬화를 무엇을 쓰냐에 따라 message header에 들어가는 정보들이 달라진다. 이 내용은 나중에 trouble shooting에 관련된 내용을 쓸떄 기술 할 것이다.)

- ProducerConfig.TRANSACTIONAL_ID_CONFIG : Producer의 transaction ID를 명시해주는 prefix이다. transaction id는 굉장히 중요한 설정인데 이는 kafka producer의 transaction의 관리에 대한 유무를 결정하기 때문이다. 명시해주지 않는다면 kafka의 transaction은 spring platform transaction manager에 의해 관리가 안될 뿐더러 rollback에 관한 기능이 없어진다. 

(이 또한 trouble shooting을 기술 할때 자세히 기술 할 것이다. )

### Consumer Config 







