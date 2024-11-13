
#define _STDINT

#include <vcruntime.h>

#if _VCRT_COMPILER_PREPROCESSOR

#pragma warning(push)
#pragma warning(disable: _VCRUNTIME_DISABLED_WARNINGS)

typedef signed char        int8_t;
typedef short              int16_t;
typedef int                int32_t;
typedef long long          int64_t;
typedef unsigned char      uint8_t;
typedef unsigned short     uint16_t;
typedef unsigned int       uint32_t;
typedef unsigned long long uint64_t;

typedef signed char        int_least8_t;
typedef short              int_least16_t;
typedef int                int_least32_t;
typedef long long          int_least64_t;
typedef unsigned char      uint_least8_t;
typedef unsigned short     uint_least16_t;
typedef unsigned int       uint_least32_t;
typedef unsigned long long uint_least64_t;

typedef signed char        int_fast8_t;
typedef int                int_fast16_t;
typedef int                int_fast32_t;
typedef long long          int_fast64_t;
typedef unsigned char      uint_fast8_t;
typedef unsigned int       uint_fast16_t;
typedef unsigned int       uint_fast32_t;
typedef unsigned long long uint_fast64_t;

typedef long long          intmax_t;
typedef unsigned long long uintmax_t;

// These macros must exactly match those in the Windows SDK's intsafe.h.
#define INT8_MIN         (-127i8 - 1)
#define INT16_MIN        (-32767i16 - 1)
#define INT32_MIN        (-2147483647i32 - 1)
#define INT64_MIN        (-9223372036854775807i64 - 1)
#define INT8_MAX         127i8
#define INT16_MAX        32767i16
#define INT32_MAX        2147483647i32
#define INT64_MAX        9223372036854775807i64
#define UINT8_MAX        0xffui8
#define UINT16_MAX       0xffffui16
#define UINT32_MAX       0xffffffffui32
#define UINT64_MAX       0xffffffffffffffffui64

#define INT_LEAST8_MIN   INT8_MIN
#define INT_LEAST16_MIN  INT16_MIN
#define INT_LEAST32_MIN  INT32_MIN
#define INT_LEAST64_MIN  INT64_MIN
#define INT_LEAST8_MAX   INT8_MAX
#define INT_LEAST16_MAX  INT16_MAX
#define INT_LEAST32_MAX  INT32_MAX
#define INT_LEAST64_MAX  INT64_MAX
#define UINT_LEAST8_MAX  UINT8_MAX
#define UINT_LEAST16_MAX UINT16_MAX
#define UINT_LEAST32_MAX UINT32_MAX
#define UINT_LEAST64_MAX UINT64_MAX

#define INT_FAST8_MIN    INT8_MIN
#define INT_FAST16_MIN   INT32_MIN
#define INT_FAST32_MIN   INT32_MIN
#define INT_FAST64_MIN   INT64_MIN
#define INT_FAST8_MAX    INT8_MAX
#define INT_FAST16_MAX   INT32_MAX
#define INT_FAST32_MAX   INT32_MAX
#define INT_FAST64_MAX   INT64_MAX
#define UINT_FAST8_MAX   UINT8_MAX
#define UINT_FAST16_MAX  UINT32_MAX
#define UINT_FAST32_MAX  UINT32_MAX
#define UINT_FAST64_MAX  UINT64_MAX

#ifdef _WIN64
    #define INTPTR_MIN   INT64_MIN
    #define INTPTR_MAX   INT64_MAX
    #define UINTPTR_MAX  UINT64_MAX
#else
    #define INTPTR_MIN   INT32_MIN
    #define INTPTR_MAX   INT32_MAX
    #define UINTPTR_MAX  UINT32_MAX
#endif

#define INTMAX_MIN       INT64_MIN
#define INTMAX_MAX       INT64_MAX
#define UINTMAX_MAX      UINT64_MAX

#define PTRDIFF_MIN      INTPTR_MIN
#define PTRDIFF_MAX      INTPTR_MAX

#ifndef SIZE_MAX
    // SIZE_MAX definition must match exactly with limits.h for modules support.
    #ifdef _WIN64
        #define SIZE_MAX 0xffffffffffffffffui64
    #else
        #define SIZE_MAX 0xffffffffui32
    #endif
#endif

#define SIG_ATOMIC_MIN   INT32_MIN
#define SIG_ATOMIC_MAX   INT32_MAX

#define WCHAR_MIN        0x0000
#define WCHAR_MAX        0xffff

#define WINT_MIN         0x0000
#define WINT_MAX         0xffff

#define INT8_C(x)    (x)
#define INT16_C(x)   (x)
#define INT32_C(x)   (x)
#define INT64_C(x)   (x ## LL)

#define UINT8_C(x)   (x)
#define UINT16_C(x)  (x)
#define UINT32_C(x)  (x ## U)
#define UINT64_C(x)  (x ## ULL)

#define INTMAX_C(x)  INT64_C(x)
#define UINTMAX_C(x) UINT64_C(x)

#pragma warning(pop) // _VCRUNTIME_DISABLED_WARNINGS

#endif // _VCRT_COMPILER_PREPROCESSOR



typedef unsigned char   undefined; //size = 1

typedef unsigned char    byte,uchar; //size = 1
typedef unsigned char    dwfenc; //size = 1
typedef unsigned short    word; //size = 2
typedef unsigned int    dword; //size = 4
typedef long long    longlong; //size = 8
typedef unsigned long long    qword; //size = 8
typedef unsigned short    ushort; //size = 2
typedef unsigned int    uint; //size = 4
typedef unsigned long    ulong; //size = 4
typedef unsigned long long    ulonglong; //size = 8

typedef long double float10;

// use these structs as return values instead of arrays
// i.e. you must get rid of "warning: address of stack memory associated with local variable 'abVar1' returned [-Wreturn-stack-address]"
// or you will get UB especially at -O1 and higher

typedef unsigned char           undefined1; //size = 1
typedef unsigned short          undefined2; //size = 2
typedef struct { byte arr[3]; } undefined3; //size = 3
typedef unsigned int            undefined4; //size = 4
typedef struct { byte arr[5]; } undefined5; //size = 5
typedef struct { byte arr[6]; } undefined6; //size = 6
typedef struct { byte arr[7]; } undefined7; //size = 7
typedef unsigned long long      undefined8; //size = 8

typedef struct { byte arr[16]; } arr16;
typedef struct { uint arr[16]; } arr64;


uint64_t mul(uint param_1, uint param_2, uint param_3, uint param_4) {
    uint64_t x = uint64_t(param_1) << 32 | param_2;
    uint64_t y = uint64_t(param_3) << 32 | param_4;
    return (x * y);
}


uint64_t div(uint param_1, uint param_2, uint param_3, uint param_4) {
    uint64_t x = uint64_t(param_1) << 32 | param_2;
    uint64_t y = uint64_t(param_3) << 32 | param_4;
    return (x / y);
}


uint64_t __cdecl asm1() {
    ulonglong uVar1;
    longlong lVar2;
    uint highXor;
    uint v;
    uint elem;
    uint uVar6;
    byte pbVar8;
    uint64_t puVar9;
    uint64_t uVar10;
    int64_t iVar11;
    uint64_t uVar12;
    byte data[8] = {0x9, 0x99, 0x8a, 0x7b, 0xfe, 0x46, 0xc2, 0xf0};
    int i;
    uint64_t out;
    uint uStack_30;

    out = 0;
    uVar10 = mul(data[1] + 0xf366, 0, 0x1302, 0);
    lVar2 = uVar10 + data[1];
    uVar10 = mul(lVar2, (lVar2 >> 0x20), 0x71b793, 0);
    out = uVar10 + 0x7cff86;
    highXor = data[2];
    uVar1 = out >> 0x12;
    uVar6 = out;
    elem = (out >> 0x20);
    iVar11 = div(uVar6, elem, 0x6381be9a, 0);
    lVar2 = uVar1 + iVar11;
    uVar10 = mul(lVar2, (lVar2 >> 0x20), 2, 0);
    uVar12 = mul(highXor + 0xf366, 0, 0x1634, 0);
    v = uVar12 + highXor;
    uVar12 = mul(v + 1, (uVar12 >> 0x20) + (uVar12 << 32 | highXor) +
                        (0xfffffffe < v), uVar6, elem);
    lVar2 = uVar12 + uVar10 + 0x2d1f65;
    highXor = (lVar2 >> 0x20);
    puVar9 = out;
    uVar6 = data[3];
    v = out >> 0x12;
    iVar11 = div(puVar9, highXor, 0x6381be9a, 0);
    lVar2 = iVar11 +
            ((highXor >> 0x12) << 32 | v) |
            (highXor << 0xe)
            + 0x21d78d;
    uVar10 = mul(lVar2, (lVar2 >> 0x20), 3, 0);
    uVar12 = mul(uVar6 + 0xf366, 0, 0x1968, 0);
    v = uVar12 + uVar6;
    uVar12 = mul(v + 1, (uVar12 >> 0x20) + (uVar12 << 32 | uVar6) + //TOARR
                        (0xfffffffe < v), puVar9, highXor);
    uVar1 = uVar12 + uVar10;
    puVar9 = uVar1;
    v = (uVar1 >> 0x20);
    uStack_30 = v;
    iVar11 = div(puVar9, v, 0x6381be9a, 0);
    lVar2 = (uVar1 >> 0x12) + iVar11;
    highXor = data[4];
    uVar10 = mul(highXor + 0xf366, 0, 0x1c9e, 0);
    v = uVar10 + highXor;
    uVar10 = mul(v + 1, (uVar10 >> 0x20) + (uVar10 << 32 | highXor) +
                        (0xfffffffe < v), puVar9, uStack_30);
    uVar12 = mul(lVar2, (lVar2 >> 0x20), 4, 0);
    uVar1 = uVar12 + uVar10 + 0xb47d9d;
    v = (uVar1 >> 0x20);
    out = (uVar1 >> 0x13) + ((v >> 0x10) << 32 | v) << 32 | ((v >> 0x10) + v) +
          (uVar1 & 0xfff0);
    i = 7;
    do {
        v = uStack_30;
        highXor = uStack_30 >> 7 ^ uStack_30;
        elem = uStack_30 << 0x19 ^ out;
        uVar6 = uStack_30 >> 0x19;
        iVar11 = div(out, uStack_30, 0x6a, 0);
        lVar2 = iVar11 + uVar6 << 32 | (elem >> 0x19) | (highXor << 7);
        uVar6 = lVar2;
        highXor = (lVar2 >> 0x20);
        elem = data[i];
        uVar10 = mul(i * i, i * i >> 0x1f, uVar6, highXor);
        uVar12 = mul(elem + 1, (0xfffffffe < elem), out, v);
        out = uVar10 + uVar12 + (elem * elem * elem);
        v = data[i];
        iVar11 = div(uVar6, highXor, 0x14c9, 0);
        out = iVar11 + ((data[i] + (0x18f - (data[7]))) * v * v * i) + out;
        i = i + -1;
    } while (-1 < i);
    return out;//13404165719009342968 [0xba0529b4000535f8]
}


int main(void) {
    asm1();
    return 0;
}
