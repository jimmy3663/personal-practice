## 너낀점 

### 구현한 Code에 대한 
- 먼저 Go (Golang)을 이용한 MSA를 위한 clean architecture 코드를 구현한 느낌 부터 쓰고자 한다. 
- Clean Architecture이란 아직도 정확히 뭔지는 잘 모르겠다. 서버를 여러 layer들로 나눠서 하위 layer는 상위 layer에 대한 정보를 모르는 dependency를 최소한으로 하는 하나의 구조인거 같다는 느낌이다. OSI layer 처럼 5계층이라 하면 5계층이고 7계층이라 하면 7계층인거 처럼 layer는 나누기 따라 다른거 같다. 
- 지금 구현한 MSA를 위한 clean architecture는 Framework에서 처리 하는 Adapter layer와 비지니스 로직이 구현되어 지는 Domain layer로 나뉘어져 있다. (회사에서 model 2 code를 위한 template를 만들기 전에 감을 잡기 위해 간단한 MS 한두개를 구현하였다.)

### MSA에서의 Clean Architecture에 대한 
- Micro Service Architecture(MSA)에 대해서 개인적으로 찾아보고 회사에서 일하면서 느낀 생각은 대단한것이 아닌 하나의 큰 서버가 할일을 여러 서버로 나눠서 서로 네트워킹을 하며 서비스를 해주는 것 같다. 
- 회사에서는 MSA를 구축할 때 비지니스 로직이 들어가야 하는 부분이 4,5개 정도로 생각 하고 있다. 최소 서버 단위인 Bounded Context, Entity가 들어가 있는 Aggregate, Event를 발행하는 Event, 동작을 수행하는 command 그리고 여러 MS에 있는 정보를 볼 수 있는 view가 있다. 
- 위 5가지 fragment들이 Clean Architecture에서 말하는 Domain layer다.(아마?) 그리고 Rest API에 대한 handling이라던지 DB connection이라던지 이러한 것들은 Adapter layer라 한다.(라고 들었다. 사실 내가 찾아봤을 땐 adpater layer는 Domain layer -> Infra layer 혹은 Infra Layer -> Domain layer로 이동 시켜 주는 놈으로 나와 있었다.)
- 여튼 위 다섯가지 부분 들에는 오직 비즈니스 로직만이 들어가야 하고 아래 단에 대한 로직은 들어가 있어도, 볼 수 있어도 안된다고 한다. 이는 사실 agile 기법에서 나온 것이라 유추 되는데 빠른 프로젝트 진행을 위해 비개발자들(소위 클라이언트)과 개발자들이 같이 많은 소통을 하며 구현을 할 때 비개발자들도 대략적인 코드를 보고 이해할 수 있어야 해서 라고 한다. (물론 아닐 수도 있지만 난 그렇게 생각하고 있다.)

## MSA Modeling 

[image](https://user-images.githubusercontent.com/43136526/125438507-ee1b4b3d-e3f8-456d-aff0-e90ef35a7375.png)