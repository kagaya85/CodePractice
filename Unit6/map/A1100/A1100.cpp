#include <string>
#include <iostream>
#include <map>

using namespace std;

const string lowDigit[] = { "tret", "jan", "feb", "mar", "apr", "may", "jun", "jly", "aug", "sep", "oct", "nov", "dec"};
const string highDigit[] = {"tret", "tam", "hel", "maa", "huh", "tou", "kes", "hei", "elo", "syy", "lok", "mer", "jou"};

string numToStr[170];
map<string, int> strToNum;

void tb() {

	for(int i = 0; i < 13; i++) {
		numToStr[i] = lowDigit[i];
		strToNum[lowDigit[i]] = i;

		numToStr[13 * i] = highDigit[i];
		strToNum[highDigit[i]] = 13 * i;
	}

	for(int i = 1; i < 13; i++) {
		for(int j = 1; j < 13; j++) {
			string tmp = highDigit[i] + " " + lowDigit[j];
			numToStr[i * 13 + j] = tmp;
			strToNum[tmp] = i * 13 + j;
		}
	}
	

}

int main() {

	tb();

	int n;

	scanf("%d%*c", &n); // %*c 表示读入一个char但不赋值，用于读取换行符
	
	while(n--) {
		string s;
		getline(cin, s);

		if(s[0] >= '0' && s[0] <= '9') {
			printf("%s\n", numToStr[atoi(s.c_str())].c_str());
		}
		else
			printf("%d\n", strToNum[s]);
	
	}

	return 0;
}