# Clean Architecture

## Clean Architecture는 System Architecture의 발전 concept이라 생각.
- System Architecture란? 
- structure, behavior, view를 분리하여 시스템 목적을 달성하기 위해 각 컴포넌트들이 어떻게 상호작용하고 정보가 어떻게 공유되는지 설명해줌. 
- 여기서 중요시 여기는 부분은 바로 '관심사의 분리'

## Clean Architecture란? 
- System Architecture적인 concept을 깔끔히 정리한 구조이다. (내가 느끼기론)

- Clean Architecture에서는 크게 Domain, Infrastructure로 나뉜다. 
- Domain은 Business Logic을 들고 있어 기능적인 수행을 해주는 녀석이고 Infrastructure는 UI, DB, APIs, Framework들이 존재한다. 
- 여기서 관심사를 분리하는 규칙을 Dependency Rule이라 한다. 

### Dependency Rule 

- 양파 구조처럼 생긴 구조에서 의존성은 외부 껍질에서 내부 껍질 쪽으로 향해야 한다.
- Domain 은 안쪽 껍질이고 Infrasturcture은 바깥 껍질이다. 여기서 Domain은 infra를 몰라야하고 domain은 일종에 플러그인 마냥 마구 가져다 쓸 수 있어야한다는 것 같다. (아마?)

- 사실 Domain에서도 여러 층으로 나뉘게 되는데 BR(business rule)의 핵심 정보를 담고 있는 Entity, logic을 처리하는 UseCase, 바깥 껍질인 Infra와 소통을 할 수 있게 해주는 Adapter로 나뉜다. // 2021.05.28



## DTO (Data Transfer Object)

- 계층간 데이터 교환을 위한 객체 ex) java bean 
- DB에서 데이터를 얻어 Servie나 controller에 보낼 때 이용하는 객체 
- 로직을 갖지 않은 순수한 객체로 getter,setter만 존재
- but DB의 데이터를 수정 할 필요가 없기 때문에 setter는 없음 
- VO(Value Object)와 다른점은 VO는 read only 다. 또한 VO는 특정 비지니스 값을 담고 있는 객체이고 DTO는 layer간 통신 용도로 오고가는 객체이다. 

- https://gmlwjd9405.github.io/2018/12/25/difference-dao-dto-entity.html 

## CQRS (Command Query Responsibility Segregation)

- 시스템에서 명령을 처리하는 책임과 조회하는 책임을 분리하는 것이 s핵심. 
- CQRS의 명령 : 시스템의 상태를 변경하는 작업을 의미 
- CQRS의 조회 : 시스템의 상태를 반환하는 작업을 의미 

- https://justhackem.wordpress.com/2016/09/17/what-is-cqrs/

## Dependency Injection 

### Dependency 
- 한 객체가 다른 객체와 상호작용을 이루고 있다면 다른 객체에 의존성을 띈다. 
	- 이는 tight coupling이라고도 함. 
	- 하나의 모듈을 바꾸면 의존하고 있는 다른 모듈을 모두 변경 해야 한다. 

### DI 
- 객체 자체가 아니라 Framework에 의해 객체의 의존성이 주입되는 설계 패텬 

## MVC pattern 

- https://m.blog.naver.com/jhc9639/220967034588