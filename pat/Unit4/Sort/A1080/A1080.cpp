#include <algorithm>
#include <cstdio>

#define loop(i, n) for (int i = 0; i < n; i++)

using namespace std;

struct Student {
    int id;
    int ge;
    int gi;
    int tot_score;
    int rank;
    int choice[6];
};

struct School {
    int quota;
    int stu_num = 0;
    int stu_id[40010];
    int last_id = -1;
};

School sch[101];
Student stu[40010];

bool cmpStu(Student a, Student b) {
    if (a.tot_score != b.tot_score)
        return a.tot_score > b.tot_score;
    else
        return a.ge > b.ge;
}

bool cmpId(int a, int b) {
    return stu[a].id < stu[b].id;
}

int main() {
    int N, M, K;

    scanf("%d%d%d", &N, &M, &K);

    loop(i, M) { scanf("%d", &sch[i].quota); }

    loop(i, N) {
        scanf("%d %d", &stu[i].ge, &stu[i].gi);
        loop(j, K) { scanf("%d", &stu[i].choice[j]); }
        stu[i].id = i;
        stu[i].tot_score = stu[i].ge + stu[i].gi;
    }

    sort(stu, stu + N, cmpStu);

    loop(i, N) {
        if (i > 0 && stu[i].tot_score == stu[i - 1].tot_score &&
            stu[i].ge == stu[i - 1].ge)
            stu[i].rank = stu[i - 1].rank;
        else
            stu[i].rank = i;
    }

    loop(i, N) {
        loop(j, K) {
            int choice = stu[i].choice[j];
            int last_id = sch[choice].last_id;
            int num = sch[choice].stu_num;

            if (sch[choice].quota > num ||
                (last_id != -1 && stu[last_id].rank == stu[i].rank)) {
                sch[choice].stu_id[num] = i;
                sch[choice].last_id = i;
                sch[choice].stu_num++;
                break;
            }
        }
    }

    loop(i, M) {
        if (sch[i].stu_num > 0) {
            sort(sch[i].stu_id, sch[i].stu_id + sch[i].stu_num, cmpId);
            loop(j, sch[i].stu_num) {
                printf("%d", stu[sch[i].stu_id[j]].id);
                if (j < sch[i].stu_num - 1)
                    printf(" ");
            }
        }
        printf("\n");
    }

    return 0;
}