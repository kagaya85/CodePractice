#include <cstdio>
#include <cstring>

using namespace std;

int main() {
    bool hashTable[128] = {false};
    char str1[100], str2[100];

    scanf("%s", str1);
    scanf("%s", str2);

    int len1 = strlen(str1);
    int len2 = strlen(str2);

    char c1, c2;

    for (int i = 0, j = 0; i < len1; i++) {
        c1 = str1[i];
        if (c1 >= 'a' && c1 <= 'z') c1 -= 32;

        if (j < len2) {
            c2 = str2[j];
            if (c2 >= 'a' && c2 <= 'z') c2 -= 32;
            if (c1 == c2) {
                j++;
                continue;
            }
        }

        if (!hashTable[c1]) {
            hashTable[c1] = true;
            putchar(c1);
        }
    }

    return 0;
}