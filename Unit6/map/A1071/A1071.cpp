#include <iostream>
#include <string>
#include <map>

using namespace std;

bool isCharacter(char c) {
	if(c >= '0' && c <= '9')
		return true;
	else if (c >= 'a' && c <= 'z') 
		return true;
	else if (c >= 'A' && c <= 'Z')
		return true;
	else
		return false;
}

int main(){
	
	map<string, int> count;

	string str;

	getline(cin, str);

	for(int i = 0; i < str.length();) {
		string word;

		while(isCharacter(str[i]) && i < str.length()) {
			if(str[i] >= 'A' && str[i] <= 'Z')
			str[i] += 32;
			word += str[i];
			i++;
		}

		while(!isCharacter(str[i]) && i < str.length()) {
			i++;
		}
		
		if(word != "") {
			if(count.find(word) != count.end())
				count[word]++;
			else
				count[word] = 1;
		}
		
	}

	int max = 0;
	string word;

	for(map<string, int>::iterator it = count.begin(); it != count.end(); it++) {
		if (it->second > max) {
			word = it->first;
			max = it->second;
		}
	}

	cout << word << ' ' << max << endl;

	return 0;
}