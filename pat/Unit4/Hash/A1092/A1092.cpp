#include <cstdio>
#include <cstring>

using namespace std;

int main() {
    int hashTable[100] = {};
    char sell[1010];
    char need[1010];
    int miss = 0;

    scanf("%s", sell);
    scanf("%s", need);

    int len1 = strlen(sell);
    int len2 = strlen(need);

    for (int i = 0; i < len1; i++)
        hashTable[sell[i] - '0']++;

    for (int i = 0; i < len2; i++) {
        int index = need[i] - '0';
        if (--hashTable[index] < 0) 
            miss++;
    }

    if (miss)
        printf("No %d", miss);
    else
        printf("Yes %d", len1 - len2);

    return 0;
}