#include <algorithm>
#include <cstdio>
#include <cstring>

#define loop(nn) for (int i = 0; i < nn; i++)

using namespace std;

const int maxn = 50;

struct Student {
    char name[11];
    char id[11];
    int score;
} stu[maxn];

bool cmp(Student a, Student b) { return a.score > b.score; }

int main() {
    int n, low, high;

    scanf("%d", &n);

    for (int i = 0; i < n; i++) {
        scanf("%s %s %d", stu[i].name, stu[i].id, &stu[i].score);
    }
    scanf("%d%d", &low, &high);

    sort(stu, stu + n, cmp);

    bool flag = false;
    for (int i = 0; i < n; i++) {
        if (stu[i].score >= low && stu[i].score <= high) {
            printf("%s %s\n", stu[i].name, stu[i].id);
            flag = true;
        }
    }

    if (!flag) {
        printf("NONE\n");
    }

    return 0;
}