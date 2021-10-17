# 의미 

- C++ STL 중 자꾸 까먹는 library를 메모해 놓자.
- 자세 설명은 해당 알고리즘 폴더 안에 적어 둘 예정.

# 자주 쓰는 C++ 기법

## string split

```c++
#include <sstream>
#include <string>

vector<string> split(string input){ // 띄어쓰기로 자르고 싶은 string
    vector<string> answer; // return vector 
    stringstream ss(input); // input string을 stringstream 객체로 변환
    string temp; // 쪼개질 string 
 
    while (getline(ss, temp, ' ')) {
        answer.push_back(temp);
    }
 
    return answer;
}
```

# C++ STL 

## 탐색 

- lower_bound, upper_bound (이진 탐색)

