
#include<stdio.h>
#include<string.h>

void demo_strings() {
    char* name = "Bill Gates";
    printf("Initial: %c %c \n", name[0], name[5]);
    return;
}

int main() {

    demo_strings();

    char* cmd;
    printf("> ");
    fgets(cmd, 10, stdin);
    printf("%lu\n", strlen(cmd));

    
    return 0;
} 