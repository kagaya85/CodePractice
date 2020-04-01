#include <map>
#include <cstdio>

using namespace std;

int main() {

	int m, n, x;
	map<int, int> count;

	scanf("%d %d", &m, &n);

	for(int i = 0; i < m; i++) {
		for(int j = 0; j < n; j++){
			scanf("%d", &x);
			if(count.find(x) != count.end())
				count[x]++;
			else
				count[x] = 1;
		}
	}

	int max = 0;
	for(map<int, int>::iterator it = count.begin(); it != count.end(); it++) {
		if(it->second > max) {
			max = it->second;
			x = it->first;
		}
	}

	printf("%d\n", x);

	return 0;
}

