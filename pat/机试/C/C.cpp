#include <cstdio>
#include <algorithm>

using namespace std;

struct Student {
    int no;
    int startTime;
    int comsumeTime;
    int limitTime;
    int actTime;    // 实际打饭时间
};

bool cmpTime(Student a, Student b) {
    return a.startTime < b.startTime;
}

bool cmpNo(Student a, Student b) {
    return a.no < b.no;
}

int main() {

    int n;
    Student stu[100001];

    scanf("%d", &n);

    for (int i = 0; i < n; i++) {
        scanf("%d%d%d", &stu[i].startTime, &stu[i].comsumeTime, &stu[i].limitTime);
        stu[i].no = i;
    }

    sort(stu, stu + n, cmpTime);

    int T = stu[0].startTime + stu[0].comsumeTime;
    stu[0].actTime = stu[0].startTime;

    for (int i = 1; i < n; i++) {
        if (stu[i].startTime + stu[i].limitTime < T) {
            stu[i].actTime = -1;
        }
        else {
            if (stu[i].startTime > T) {
                stu[i].actTime = stu[i].startTime;
            }
            else {
                stu[i].actTime = T;
            }
            T = stu[i].actTime + stu[i].comsumeTime;
        }
    }

    sort(stu, stu + n, cmpNo);

    for (int i = 0; i < n; i++) {
        printf("%d ", stu[i].actTime);
    }

    return 0;
}