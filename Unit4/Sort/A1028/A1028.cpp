#include <algorithm>
#include <cstdio>
#include <cstring>

using namespace std;

struct Student
{
    char id[8];
    char name[12];
    int score;
};

bool cmp1(Student a, Student b) {
    return strcmp(a.id, b.id) < 0;
}

bool cmp2(Student a, Student b) {
    int s = strcmp(a.name, b.name);
    if (s != 0) return s < 0;
    else return strcmp(a.id, b.id) < 0;
}

bool cmp3(Student a, Student b) {
    if (a.score != b.score) return a.score < b.score;
    else return strcmp(a.id, b.id) < 0;
}

int main() {
    int N, C;
    Student stu[100010];
    scanf("%d%d", &N, &C);

    for (int i = 0; i < N; i++) {
        scanf("%s%s%d", stu[i].id, stu[i].name, &stu[i].score);
    }

    switch (C) {
        case 1:
            sort(stu, stu + N, cmp1);
            break;
        case 2:
            sort(stu, stu + N, cmp2);
            break;
        default:
            sort(stu, stu + N, cmp3);
            break;
    }

    for (int i = 0; i < N; i++)
        printf("%s %s %d\n", stu[i].id, stu[i].name, stu[i].score);

    return 0;
}