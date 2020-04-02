#define _CRT_SECURE_NO_WARNINGS
#include <iostream>
#include <cstdio>
#include <map>
#include <string>
#include <set>

using namespace std;

map<string, set<int>> titleMap, authorMap, keyMap, pubMap, yearMap;

void query(map<string, set<int>> &m, string &key) {
	
	if(m.find(key) != m.end()) {
		for(set<int>::iterator it = m[key].begin(); it != m[key].end(); it++) {
			printf("%07d\n", *it);
		}

	} else {
		printf("Not Found\n");
	}

}

int main() {

	int n, id;
	string title, author, key, pub, year;

	scanf("%d", &n);

	for(int i = 0; i < n; i++){
		scanf("%d%*c", &id);

		getline(cin, title);
		titleMap[title].insert(id);

		getline(cin, author);
		authorMap[author].insert(id);

		while(cin >> key) {
			keyMap[key].insert(id);
			if(getchar() == '\n')
				break;
		}

		getline(cin, pub);
		pubMap[pub].insert(id);

		getline(cin, year);
		yearMap[year].insert(id);
	
	}

	string str;
	int type, m;

	scanf("%d", &m);

	for(int i = 0; i < m; i++){
		
		scanf("%d: ", &type);
		getline(cin, str);

		printf("%d: %s\n", type, str.c_str());

		switch(type){
			case 1:	
				query(titleMap, str);
				break;
			case 2:	
				query(authorMap, str);
				break;
			case 3: 
				query(keyMap, str);
				break;
			case 4: 
				query(pubMap, str);
				break;
			default:
				query(yearMap, str);
				break;
	
		}
	} 

	
	return 0;
}