 enum Color {
     RED,
     GREEN,
     YELLOW =99
 };

typedef enum
{
    SEG_DEFAULT,
    SEG_ES,
    SEG_DS,
    SEG_FS,
    SEG_GS,
    SEG_CS,
    SEG_SS
} SEGMENTREG;


 struct Car {
     char make[50];
     int year;
 };

int add (int a, int b){
    return a+b;
};