# Kafka에 대한 전반적인 내용 
내가 현재까지 알고 있는것들에 대한 (21.11.28 기준)
생각해 보니 카프카에 대한 지식 보단 카프카를 어떻게 사용하는지에 대한 내용을 기술 할 예정이다. 
(카프카 고유에 대한 기술적 지식은 아직 더 공부가 필요한듯..)

## Kafka Cluster 

- Kafka broker가 1개 이상 있는 broker의 묶음을 kafka cluster 라고 하는데 해당 클러스터에는 zookeeper가 필수적으로 존재해야한다. 
- 정확히 Zookeeper가 어떤 원리로 동작하는진 잘 모르나 zookeeper는 브로커들의 관리와 메타데이터들의 관리를 해준다고 한다. 

## Kafka Broker

- Kafka broker는 흔히 카프카라고도 불린다. 카프카 브로커가 실질적으로 produer로 부터 message를 받고 consumer가 가져 갈 수 있는 환경을 만들기 때문이다. 
- 하나의 브로커는 복제를 할 수 있다. 이때 leader, slave로 나뉘게 되는데 leader는 r/w 둘다 가능하고 slave는 r만 수행한다. Leader와 slave는 주키퍼가 지정해 주며 leader 브로커가 죽었을 때 자동으로 주키퍼가 slave 중 하나를 선택하여 leader를 선정해준다.

## Parition 

- 파티션은 메세지가 쌓이는 실질적인 message queue이다.
- 파티션 또한 여러개가 생성될 수 있는데 초기에 파티션의 갯수를 정해놓고 가는 전략과 listener가 늘어남에 따라 파티션을 늘리는 전략이 존재한다. 
- 파티션은 append only로 한번 늘어나면 줄어들지 않는다. (이렇게 파티션이 늘어나는 과정을 rebalancing이라 한다. rebalancing이 된 후에는 새롭게 만들어진 파티션과 listener들을 binding 시켜준다. rebalancing이 일어날때의 경우도 자알 대비해야한다.ㅎㅎ..)
- 하나의 파티션은 listener 그룹 당 하나의 서비스와 binding 될 수 있다. 이는 pod가 늘어 났을 때 파티션의 갯수가 고정적이라면 늘어난 pod는 파티션과 binding 되지 않는아 메세지를 가져올 수 없게 된다.  

#### 예시 

1) partition 수와 consumer의 수가 동일 

![parti-rule-1](https://user-images.githubusercontent.com/43136526/143767201-18a19d58-39fd-46e5-8f57-4d08ca77da99.png)

- 모든 consumer가 파티션으로 부터 메세지를 받을 수 있다. 

2) parition의 수가 consumer의 수 보다 많을 경우 

![parti-rule-2](https://user-images.githubusercontent.com/43136526/143767207-32a7d9a7-ff8c-4560-89d5-d2f69ab76741.png)

- 모든 consumer가 파티션으로 부터 메세지를 받을 수 있다. Consumer가 어느 파티션으로부터 메세지를 poll 해올지는 정해져 있지 않다. (매번 다르다. 그래서 순서성 보장에 따른 처리 로직이 필요)

3) parition의 수가 consumer의 수 보다 적을 경우 

![parti-rule-3](https://user-images.githubusercontent.com/43136526/143767212-c63c224b-e1a8-4e33-b9bc-4e4c069b7291.png)

- 위 그림에서 order-4가 다른 1,2,3 보다 늦게 생성되었다는 가정을 해보자. 그런 경우에는 order-4 consumer는 parition에 binding 되지 않아 메세지를 poll 해 올 수 없다. 

## Group Id 

- Consumer 입장에서 group id는 괴장히 중요하다. 
- 한 consumer가 replica를 만들었을 경우 같은 일을 하는 서비스가 추가 된 것이다. 이때 한 메세지에 대해서 두 컨슈머가 같이 sub 하여 로직을 처리한다면 해당 로직에 중복 처리 현상이 생기게 될 것이다. 이때 group id가 중요한 역할을 하게 되는데 한 컨슈머가 메세지를 sub 하면 같은 group id를 갖고 있는 다른 컨슈머 들은 해당 메세지에 대해 sub을 하지 않게 된다. 




