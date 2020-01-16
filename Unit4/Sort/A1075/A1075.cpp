#include <cstdio>
#include <algorithm>

using namespace std;

struct Student {
    int id;
    int score[5];
    int fullMarkNum = 0;
};

bool cmp (Student a, Student b) {
    
}

int main() {
    int N, K, M;
    int p[5];

    Student stu[10010];

    scanf("%d%d%d", &N, &K, &M);

    for (int i = 0; i < K; i++)
        scanf("%d", &p[i]);

    for (int i = 0; i < M; i++) {
        int id, n, score;
        scanf("%d%d%d", &id, &n, &score);
        stu[id].id = id;
        stu[id].score[n] = score;
        if (score == p[n])
            stu[id].fullMarkNum++;
    }


    return 0;
}
