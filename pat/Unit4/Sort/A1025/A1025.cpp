#include <algorithm>
#include <cstring>
#include <cstdio>

using namespace std;

struct Student {
    char id[16];
    int score;
    int final_rank;
    int location_number;
    int local_rank;
} student[10240];

bool cmp(Student a, Student b) {
    if (a.score != b.score)
        return a.score > b.score;
    else
        return strcmp(a.id, b.id) < 0;
}

int main() {
    int locationNum, totStuNum = 0;

    scanf("%d",&locationNum);

    int stuNum, index = 0;
    for (int i = 0; i < locationNum; i++) {
        scanf("%d",&stuNum);
        totStuNum += stuNum;

        for (int j = 0; j < stuNum; j++) {
            scanf("%s%d",student[index].id, &student[index].score);
            student[index].location_number = i + 1;
            index++;
        }

        sort(student + totStuNum - stuNum, student + totStuNum, cmp);

        int base = index - stuNum;
        student[base].local_rank = 1;
        for (int tmpIdx = 1; tmpIdx < stuNum; tmpIdx++) {
            if (student[base + tmpIdx].score !=
                student[base + tmpIdx - 1].score) {
                student[base + tmpIdx].local_rank = tmpIdx + 1;
            } else {
                student[base + tmpIdx].local_rank =
                    student[base + tmpIdx - 1].local_rank;
            }
        }
    }

    sort(student, student + totStuNum, cmp);
    student[0].final_rank = 1;
    for (int idx = 1; idx < totStuNum; idx++) {
        if (student[idx].score != student[idx - 1].score) {
            student[idx].final_rank = idx + 1;
        } else {
            student[idx].final_rank = student[idx - 1].final_rank;
        }
    }

    printf("%d\n", totStuNum);
    for (int i = 0; i < totStuNum; i++) {
        printf("%s %d %d %d\n",student[i].id, student[i].final_rank, student[i].location_number, student[i].local_rank);
    }
    
}