#include <algorithm>
#include <cstdio>
#include <cstring>

using namespace std;

struct People {
    char name[10];
    int age;
    int worth;
};

bool cmp(People a, People b) {
    if (a.worth != b.worth)
        return a.worth > b.worth;
    else if (a.age != b.age)
        return a.age < b.age;
    else
        return strcmp(a.name, b.name) < 0;
}

int main() {
    int N, K;
    People peo[100010];

    scanf("%d%d", &N, &K);

    for (int i = 0; i < N; i++)
        scanf("%s%d%d", peo[i].name, &peo[i].age, &peo[i].worth);

    sort(peo, peo + N, cmp);

    int M, Amin, Amax;
    for (int i = 0; i < K; i++) {
        scanf("%d%d%d", &M, &Amin, &Amax);
        printf("Case #%d:\n", i + 1);
        int flag = 0;
        int count = 0;
        for (int j = 0; j < N; j++) {
            if (count < M && peo[j].age >= Amin && peo[j].age <= Amax) {
                flag = 1;
                count++;
                printf("%s %d %d\n", peo[j].name, peo[j].age, peo[j].worth);
            }
        }
        if (!flag)
            printf("None\n");
    }

    return 0;
}