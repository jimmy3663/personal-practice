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

## 문자열 갖고 놀기 

- istringstream 

```c++ 
#include<iostream>
#include<sstream>

int main() {
        std::istringstream iss("test 123 123hah ahha");
        std::string str, str1, str2;
        int inta;

        std::cout<< iss.str() <<std::endl;
        iss >> str >> str1 >> str2 >> inta;
        std::cout<<"str  : "<< str  <<std::endl;
        std::cout<<"str1 : "<< str1 <<std::endl;
        std::cout<<"str2 : "<< str2 <<std::endl;
        std::cout<<"inta : "<< inta <<std::endl;

}
```

- ostringstream

```c++
#include<iostream>
#include<sstream>

int main() {
        std::ostringstream oss;
        int inta = 10;
        std::string str = " test_string";
        oss << "test " << inta << str <<std::endl;
        std::cout<< oss.str();
}
```

- sstringstream

## 탐색 

- lower_bound, upper_bound (이진 탐색)

