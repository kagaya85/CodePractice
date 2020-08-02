#define _CRT_SECURE_NO_WARNINGS
#include <cstdio>
#include <set>

using namespace std;

const int maxN = 51;
set<int> st[maxN];

int main(){
	int n, k, x;

	scanf("%d", &n);
	for(int i = 1; i <= n; i++){
		scanf("%d", &k);
		for(int j = 0; j < k; j++){
			scanf("%d", &x);
			st[i].insert(x);
		}
	}
	
	int q, st1, st2;
	scanf("%d", &q);
	
	for(int i = 0; i < q; i++) {
		scanf("%d %d", &st1, &st2);
		int totNum = st[st2].size();
		int sameNum = 0;
		for(set<int>::iterator it = st[st1].begin(); it != st[st1].end(); it++) {
			if(st[st2].find(*it) != st[st2].end())
				sameNum++;
			else
				totNum++;
		}
		printf("%.1f%%\n", sameNum * 100.0 / totNum);
	}


	return 0;
}