#include <algorithm>
#include <cstdio>

using namespace std;

const int maxn = 10010;
struct Student {
    int id;
    int score[5] = {-1, -1, -1, -1, -1};
    int fullMarkNum = 0;
    int totScore;
    bool outputFlag = false;
};

bool cmp(Student a, Student b) {
    if (a.totScore != b.totScore)
        return a.totScore > b.totScore;
    else if (a.fullMarkNum != b.fullMarkNum)
        return a.fullMarkNum > b.fullMarkNum;
    else 
        return a.id < b.id;
}

int main() {
    int N, K, M;
    int p[6];

    Student stu[maxn];

    scanf("%d%d%d", &N, &K, &M);

    for (int i = 1; i <= K; i++) scanf("%d", &p[i]);

    int s_id, p_id, score;
    for (int i = 0; i < M; i++) {
        scanf("%d%d%d", &s_id, &p_id, &score);
        stu[s_id].id = s_id;
        if (score != -1) 
            stu[s_id].outputFlag = true;
        if (score == -1 && stu[s_id].score[p_id] == -1)
            stu[s_id].score[p_id] = 0;
        if (score == p[p_id] && stu[s_id].score[p_id] < p[p_id])
            stu[s_id].fullMarkNum++;
        if (score > stu[s_id].score[p_id])
            stu[s_id].score[p_id] = score;
    }

    for (int i = 1; i <= N; i++) {
        for (int j = 1; j <= K; j++) {
            if (stu[i].score[j] != -1)
                stu[i].totScore += stu[i].score[j];
        }
    }

    sort(stu + 1, stu + N + 1, cmp);

    int rank = 1;
    for (int i = 1; i <= N && stu[i].outputFlag == true; i++) {
        // printf("ok\n");
        if (i > 1 && stu[i].totScore != stu[i - 1].totScore)
            rank = i;
        printf("%d %05d %d", rank, stu[i].id, stu[i].totScore);
        for (int j = 1; j <= K; j++) {
            if (stu[i].score[j] == -1)
                printf(" -");
            else
                printf(" %d", stu[i].score[j]);
        }
        printf("\n");
    }

    return 0;
}
