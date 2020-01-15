#include <algorithm>
#include <cstdio>
#include <cstring>

using namespace std;

struct Time {
    int MM, dd, hh, mm;
};

struct Record {
    char name[30];
    Time time;
    char state[10];
};

int toll_of_time[24];

int cal_mins(const Time start, const Time end) {
    int dayMins = 24 * 60;
    int startMins = start.dd * dayMins + start.hh * 60 + start.mm;
    int endMins = end.dd * dayMins + end.hh * 60 + end.mm;

    return endMins - startMins;
}

bool cmp(Record a, Record b) {
    int s = strcmp(a.name, b.name);

    if (s != 0)
        return s < 0;
    else if (a.time.MM != b.time.MM)
        return a.time.MM < b.time.MM;
    else if (a.time.dd != b.time.dd)
        return a.time.dd < b.time.dd;
    else if (a.time.hh != b.time.hh)
        return a.time.hh < b.time.hh;
    else
        return a.time.mm < b.time.mm;
}

int cal_toll(Time start, Time end) {
    int toll = 0;

    while (start.dd < end.dd || start.hh < end.hh) {
        toll += toll_of_time[start.hh] * (60 - start.mm);
        start.hh++;
        if (start.hh == 24) {
            start.hh = 0;
            start.dd++;
        }
        start.mm = 0;
    }

    toll += toll_of_time[end.hh] * end.mm;

    return toll;
}

int main() {
    int N;
    Record rec[1010];

    for (int i = 0; i < 24; i++) scanf("%d", &toll_of_time[i]);

    scanf("%d", &N);

    for (int i = 0; i < N; i++) {
        scanf("%s", rec[i].name);
        scanf("%d:%d:%d:%d", &rec[i].time.MM, &rec[i].time.dd, &rec[i].time.hh,
              &rec[i].time.mm);
        scanf("%s", rec[i].state);
    }

    sort(rec, rec + N, cmp);

    int flag = 0;
    int tot_toll = 0;
    Time start;
    char nameNow[30];
    strcpy(nameNow, rec[0].name);
    printf("%s %02d\n", nameNow, rec[0].time.MM);

    for (int i = 0; i < N; i++) {
        if (strcmp(nameNow, rec[i].name)) {
            printf("Total amount: $%.2f\n", tot_toll / 100.0);
            printf("%s %02d\n", rec[i].name, rec[i].time.MM);
            strcpy(nameNow, rec[i].name);
            tot_toll = 0;
            flag = 0;
        }

        if (!strcmp(rec[i].state, "on-line")) {
            flag = 1;
            start = rec[i].time;
        } else if (flag == 1) {
            Time end = rec[i].time;
            int toll = cal_toll(start, end);
            tot_toll += toll;
            printf("%02d:%02d:%02d %02d:%02d:%02d %d $%.2f\n", start.dd, start.hh, start.mm,
                   end.dd, end.hh, end.mm, cal_mins(start, end),
                   toll / 100.0);
            flag = 0;
        }
    }
    printf("Total amount: $%.2f\n", tot_toll / 100.0);

    return 0;
}