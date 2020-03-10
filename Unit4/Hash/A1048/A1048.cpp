#include <cstdio>

using namespace std;

int main() {
    int n, m;
    int coins[1010] = {};
    scanf("%d %d", &n, &m);

    for (int i = 0; i < n; i++) {
        int value;
        scanf("%d", &value);
        coins[value]++;
    }
    
    for (int i = 1; i <= m / 2; i++) {
        if (coins[i] > 0) {
            coins[i]--;
            if (coins[m - i] > 0) {
                printf("%d %d", i, m - i);
                return 0;
            }
        }
    }

    printf("No Solution");
    return 0;
}