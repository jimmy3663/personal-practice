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

## 순열 

- 4개 중 4개 중복 없는 순열 

```c++
#include <iostream>
#include <algorithm>
#include <vector>

using namespace std;

int main(){
	vecotr<int> v = {1,2,3,4};

	do{
		for(int i=0; i<v.size(); i++){
			cout<<v[i]<<endl;
		}
	}
	while(next_permutation(v.begin(), v.end()));
}
```

- 4개 중 n개 중복 없는 순열 

```c++
#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

int main(){
	vector<int> v = {1,2,3,4};

	vector<int> temp = {1,1,0,0}; // 4개 중 2개 뽑는 형태 

	do{
		for(int i=0; i<v.size(); i++){
			if(temp[i]==1){
				cout<<v[i]<<endl;
			}
		}
	}while(next_permutation(temp.begin(), temp.end()));
}
```

- n개 중복 있는 순열 

```c++ 
#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

vector<string> vocab;
string word = "AEIOU"; // 예를 들어 
void recurr(string s, int length){
	if(s.size() == length){
		vocab.push_back(s);
		return;
	}
	for(int i=0; i< n; i++){
		recurr(s+word[i] , length);
	}
}

int main(){



	for(int i=1; i<=n; i++){
		string s = "";
		recurr(s, i);
	}

	for(int i=0; i< vocab.size(); i++){
		cout<<vocab[i]<<endl;
	}

}
```
