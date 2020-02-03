#include <cstdio>
#include <cstring>
#include <iostream>

using namespace std;

int main() {
    bool hashTable[100] = {};
    char str1[10010];
    char str2[10010];
    
    cin.getline(str1, 10010);
    cin.getline(str2, 10010);

    for (int i = 0; i < strlen(str2); i++) 
        hashTable[str2[i] - ' '] = true;

    for (int i = 0; i < strlen(str1); i++)
        if (!hashTable[str1[i] - ' '])
            putchar(str1[i]);

    return 0;
}