#include <algorithm>
#include <cstdio>
#include <cmath>

using namespace std;

struct Student {
    int id;
    int score[4];
};

int now = 0;

bool cmp(Student a, Student b) { return a.score[now] > b.score[now]; }

int main() {
    Student stu[2000];
    char rank[1000000][4] = {0};
    // Student *stu = new Student[2000];
    // int (*rank)[4] = new int[1000000][4]();
    char Course[] = "ACME";
    int N, M;

    scanf("%d%d", &N, &M);

    for (int i = 0; i < N; i++) {
        scanf("%d%d%d%d", &stu[i].id, &stu[i].score[1], &stu[i].score[2],
              &stu[i].score[3]);
        stu[i].score[0] =
            round((stu[i].score[1] + stu[i].score[2] + stu[i].score[3]) / 3.0) + 0.5;
    }

    for (now = 0; now < 4; now++) {
        sort(stu, stu + N, cmp);
        rank[stu[0].id][now] = 1;
        for (int i = 1; i < N; i++) {
            if (stu[i].score[now] == stu[i - 1].score[now])
                rank[stu[i].id][now] = rank[stu[i - 1].id][now];
            else
                rank[stu[i].id][now] = i + 1;
        }
    }
    for (int i = 0; i < M; i++) {
        int id;
        scanf("%d", &id);
        if (rank[id][0] == 0)
            printf("N/A\n");
        else {
            int c = 0;
            int r = 9999;
            for (int j = 0; j < 4; j++){
                if (rank[id][j] < r) {
                    c = j;
                    r = rank[id][j];
                }
            }

            printf("%d %c\n", r, Course[c]);
        }
    }

    return 0;
}