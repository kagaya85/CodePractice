#define _CRT_SECURE_NO_WARNINGS

#include <cstdio>
#include <iostream>
#include <string>

using namespace std;


string preProcess(string s, int& e, int validNum) {
	
	int k = 0;
	e = 0;

	while (s.length() > 0 && s[0] == '0')
		s.erase(s.begin());

	if(s[0] == '.') {
		s.erase(s.begin());
		while(s.length() > 0 && s[0] =='0'){
			s.erase(s.begin());
			e--;
		}
	}
	else {
		while(s.length() > k && s[k] != '.') {
			k++;
			e++;
		}
		if(s.length() > k)
			s.erase(s.begin() + k);
	}
	
	if(s.length() == 0) {
		e = 0;
		return "0";
	}

	if(s.length() >= validNum) {
		s.erase(s.begin() + validNum, s.end());
	}
	else {
		int len = s.length();
		for(int i = 0; i < validNum - len; i++)
			s += '0';

	}

	return s;
}


int main() {
	int n;
	string s1, s2, s11, s22;

	cin >> n >> s1 >> s2;

	int e1, e2;

	s11 = preProcess(s1, e1, n);
	s22 = preProcess(s2, e2, n);

	if(s11 == s22 && e1 == e2) {
		
		printf("YES 0.%s*10^%d\n", s11.c_str(), e1);
	
	}
	else
		printf("NO 0.%s*10^%d 0.%s*10^%d\n", s11.c_str(), e1, s22.c_str(), e2);

}

