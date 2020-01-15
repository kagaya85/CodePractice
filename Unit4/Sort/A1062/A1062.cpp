#include <algorithm>
#include <cstdio>
#include <cstring>
#include <fstream>

using namespace std;

struct Student {
    char id[12];
    int dscore;
    int cscore;
    int tot_score;
    int flag;
};

bool cmp(Student a, Student b) {
    if (a.flag != b.flag)
        return a.flag < b.flag;
    else if (a.tot_score != b.tot_score)
        return a.tot_score > b.tot_score;
    else if (a.dscore != b.dscore)
        return a.dscore > b.dscore;
    else
        return strcmp(a.id, b.id) < 0;
}

int main() {
    int N, L, H;

    scanf("%d%d%d", &N, &L, &H);

    Student stu[100010];
    int passNum = N;

    for (int i = 0; i < N; i++) {
        scanf("%s%d%d", stu[i].id, &stu[i].dscore, &stu[i].cscore);
        stu[i].tot_score = stu[i].dscore + stu[i].cscore;
        if (stu[i].dscore < L || stu[i].cscore < L) {
            stu[i].flag = 5;
            passNum--;
        }
        else if (stu[i].dscore >= H && stu[i].cscore >= H) 
            stu[i].flag = 1;
        else if (stu[i].dscore >= H && stu[i].cscore < H) 
            stu[i].flag = 2;
        else if (stu[i].dscore >= stu[i].cscore)
            stu[i].flag = 3;
        else 
            stu[i].flag = 4; 
    }

    sort(stu, stu + N, cmp);

    printf("%d\n", passNum);
    for (int i = 0; i < passNum; i++) {
        printf("%s %d %d\n", stu[i].id, stu[i].dscore, stu[i].cscore);
    }

    return 0;
}