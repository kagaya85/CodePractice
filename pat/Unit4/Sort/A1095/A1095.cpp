#include <algorithm>
#include <cstdio>
#include <cstring>
#include <map>
#include <string>

#define IN 0
#define OUT 1

using namespace std;

const int maxn = 10010;

struct Car {
    char id[8];
    int time;
    bool status;  // 0 - in 1 - out
} cars[maxn], valid[maxn];

bool cmpByIdAndTime(Car a, Car b) {
    int r = strcmp(a.id, b.id);
    if (r != 0)
        return r < 0;
    else
        return a.time < b.time;
}

bool cmpByTime(Car a, Car b) { return a.time < b.time; }

int timeToSec(int h, int m, int s) { return h * 3600 + m * 60 + s; }

int main() {
    int N, K, hh, mm, ss;
    char status[4];

    scanf("%d%d", &N, &K);

    for (int i = 0; i < N; i++) {
        scanf("%s %d:%d:%d %s", cars[i].id, &hh, &mm, &ss, status);
        cars[i].time = timeToSec(hh, mm, ss);
        cars[i].status = strcmp(status, "in") ? OUT : IN;
    }

    sort(cars, cars + N, cmpByIdAndTime);

    int maxTime = -1;
    int validNum = 0;
    map<string, int> parkTime;

    for (int i = 0; i < N - 1; i++) {
        if (!strcmp(cars[i].id, cars[i + 1].id) && cars[i].status == IN &&
            cars[i + 1].status == OUT) {
            valid[validNum++] = cars[i];
            valid[validNum++] = cars[i + 1];
            int inTime = cars[i + 1].time - cars[i].time;
            if (parkTime.count(cars[i].id) == 0) {
                parkTime[cars[i].id] = 0;
            }
            parkTime[cars[i].id] += inTime;
            maxTime = max(maxTime, parkTime[cars[i].id]);
        }
    }

    sort(valid, valid + validNum, cmpByTime);

    int now = 0, carNum = 0;
    for (int i = 0; i < K; i++) {
        scanf("%d:%d:%d", &hh, &mm, &ss);
        int time = timeToSec(hh, mm, ss);
        while(now < validNum && valid[now].time <= time) {
            if (valid[now].status == IN)
                carNum++;
            else 
                carNum--;
            now++;
        }
        printf("%d\n", carNum);
    }

    map<string, int>::iterator it;
    for (it = parkTime.begin(); it != parkTime.end(); it++) {
        if (it->second == maxTime)
            printf("%s ", it->first.c_str());
    }

    printf("%02d:%02d:%02d\n", maxTime/3600, maxTime%3600/60, maxTime%60);
    return 0;
}