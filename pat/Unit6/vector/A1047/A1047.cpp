#define _CRT_SECURE_NO_WARNINGS
#include <cstdio>
#include <vector>
#include <algorithm>
#include <cstring>

using namespace std;

const int maxN = 40010;
const int maxK = 2510;
char name[maxN][5];
vector<int> course[maxK];

bool cmp(int a, int b) {
	return strcmp(name[a], name[b]) < 0;
}

int main() {
	int N, K;

	scanf("%d%d", &N, &K);
	
	for(int i = 0; i < N; i++){
		int n, courseId;
		scanf("%s%d", name[i], &n);
		for(int j = 0; j < n; j++) {
			scanf("%d", &courseId);
			course[courseId].push_back(i);
		}
	}

	for(int i = 1; i <= K; i++){
		printf("%d %d\n", i, course[i].size());
		sort(course[i].begin(), course[i].end(), cmp);
		for(int j = 0; j < course[i].size(); j++)
			printf("%s\n", name[course[i][j]]);
	}

	return 0;
}