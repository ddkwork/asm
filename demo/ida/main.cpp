#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <ctype.h>
#include "defs.h"

void hexdump(const char *title, const unsigned char *data, size_t size) {
    printf("%20s\n", title);
    for (size_t i = 0; i < size; i += 16) {
        printf("%08lx  ", (unsigned long) i);
        for (size_t j = 0; j < 8; j++) {
            if (i + j < size) {
                printf("%02x ", data[i + j]);
            } else {
                printf("   ");
            }
        }
        printf("  ");
        for (size_t j = 8; j < 16; j++) {
            if (i + j < size) {
                printf("%02x ", data[i + j]);
            } else {
                printf("   ");
            }
        }
        printf(" | ");
        for (size_t j = 0; j < 16; j++) {
            if (i + j < size) {
                printf("%c", isprint(data[i + j]) ? data[i + j] : '.');
            }
        }
        printf("\n");
    }
}

// void   asm1(int a1@<ebx>, int a2@<edi>, int a3@<esi>)
void asm1() {
    int v3; // eax
    int v4; // edx
    int v5; // ecx
    unsigned int v6; // kr00_4
    _DWORD *v7; // eax
    unsigned __int64 v8; // rcx
    unsigned int *v9; // esi
    unsigned __int64 v10; // kr30_8
    __int64 v11; // rax
    __int64 v12; // rax
    unsigned __int64 v13; // kr48_8
    int v14; // edi
    unsigned int v15; // eax
    unsigned int v16; // ebx
    __int64 v17; // rax
    unsigned int v18; // kr68_4
    bool v19; // sf
    __int64 v20; // [esp+0h] [ebp-274h] BYREF
    __int64 v21; // [esp+10h] [ebp-264h]
    __int64 v22; // [esp+20h] [ebp-254h]
    __int64 v23; // [esp+30h] [ebp-244h] BYREF
    __int64 v24; // [esp+40h] [ebp-234h]
    __int64 v25; // [esp+50h] [ebp-224h]
    __int64 v26; // [esp+60h] [ebp-214h]
    __int64 v27; // [esp+70h] [ebp-204h]
    __int64 v28; // [esp+80h] [ebp-1F4h]
    __int64 v29; // [esp+90h] [ebp-1E4h]
    unsigned __int64 v30; // [esp+A0h] [ebp-1D4h]
    int v31; // [esp+A8h] [ebp-1CCh]
    __int64 v33; // [esp+B0h] [ebp-1C4h]
    __int64 v34; // [esp+C0h] [ebp-1B4h]
    __int64 v36; // [esp+E0h] [ebp-194h]
    __int64 v37; // [esp+F0h] [ebp-184h]
    __int64 v38; // [esp+100h] [ebp-174h]
    //    unsigned char HWID[8] = {0x9, 0x99, 0x8a, 0x7b, 0xfe, 0x46, 0xc2, 0xf0};

    int data[8]; //={0x9, 0x99, 0x8a, 0x7b, 0xfe, 0x46, 0xc2, 0xf0};; // [esp+110h] [ebp-164h] BYREF
    __int64 v40; // [esp+130h] [ebp-144h]
    unsigned int v41; // [esp+138h] [ebp-13Ch]
    int v42; // [esp+148h] [ebp-12Ch]
    unsigned __int8 *v43; // [esp+15Ch] [ebp-118h]
    unsigned int v44; // [esp+210h] [ebp-64h]
    int v45; // [esp+214h] [ebp-60h]
    int v46; // [esp+218h] [ebp-5Ch]
    int v47; // [esp+21Ch] [ebp-58h]
    int v48; // [esp+220h] [ebp-54h]
    int *v49; // [esp+224h] [ebp-50h]
    int *v50; // [esp+228h] [ebp-4Ch]
    int v51; // [esp+22Ch] [ebp-48h]
    int v52; // [esp+230h] [ebp-44h]
    int v53; // [esp+234h] [ebp-40h] BYREF
    unsigned __int8 anzhuang[8]; // [esp+244h] [ebp-30h] BYREF
    unsigned __int8 code3Buf[8]; // [esp+254h] [ebp-20h] BYREF
    unsigned __int8 HWID[8] = {0x9, 0x99, 0x8a, 0x7b, 0xfe, 0x46, 0xc2, 0xf0};; // [esp+264h] [ebp-10h] BYREF
    int savedregs; // [esp+274h] [ebp+0h] BYREF

    // HWID[0] = 9;
    // HWID[1] = 0x99;
    // HWID[2] = 0x8A;
    // HWID[3] = 0x7B;
    // HWID[4] = 0xFE;
    // HWID[5] = 0x46;
    // HWID[6] = 0xC2;
    // HWID[7] = 0xF0;
    memset(code3Buf, 0, sizeof(code3Buf));
    hexdump("title", HWID, 8u);
    // v52 = a3;
    // v51 = a2;
    v50 = &savedregs;
    v49 = &v53;
    // v48 = a1;
    v47 = v5;
    v46 = v4;
    v45 = v3;
    // v6 = __readeflags();
    v44 = v6;
    qmemcpy(data, HWID, 8u);
    v43 = anzhuang;
    data[6] = (unsigned __int64) (0x1302i64 * (BYTE1(data[0]) + 0xF366)) >> 0x20;
    *(_QWORD *) anzhuang = 0x71B793 * (BYTE1(data[0]) + 0x1302i64 * (BYTE1(data[0]) + 0xF366)) + 0x7CFF86;
    v40 = *(_QWORD *) anzhuang;
    data[2] = *(_QWORD *) anzhuang >> 0x12;
    v42 = *(_DWORD *) &anzhuang[4] >> 0x12;
    (v36) = (unsigned __int64) (0x1634i64 * (BYTE2(v37) + 0xF366)) >> 0x20;
    v7 = (_DWORD *) HIDWORD(v37);
    HIDWORD(v8) = ((BYTE2(v37) + 0x1634i64 * (BYTE2(v37) + 0xF366) + 1) * (_QWORD) v38
                   + 2 * (*(_QWORD *) anzhuang / 0x6381BE9Ai64 + __PAIR64__(v41, (v38)))
                   + 0x2D1F65) >> 0x20;
    *(_DWORD *) HIDWORD(v37) = (BYTE2(v37) + 0x1634 * (BYTE2(v37) + 0xF366) + 1) * v38
                               + 2 * (*(_QWORD *) anzhuang / 0x6381BE9Ai64 + (v38))
                               + 0x2D1F65;
    LODWORD(v8) = *v7;
    v7[1] = HIDWORD(v8);
    *(_QWORD *) &v37 = v8;
    v31 = (unsigned __int64) (0x1968i64 * (BYTE3(v33) + 0xF366)) >> 0x20;
    v9 = (unsigned int *) HIDWORD(v33);
    v10 = (BYTE3(v33) + 0x1968i64 * (BYTE3(v33) + 0xF366) + 1) * v34
          + 3 * ((__int64) v8 / 0x6381BE9A + (v8 >> 0x12) + 0x21D78D);
    *((_QWORD *) &v26 + 1) = 0x6381BE9Ai64;
    (v26) = HIDWORD(v10);
    *(_DWORD *) (HIDWORD(v33) + 4) = HIDWORD(v10);
    LODWORD(v26) = v10;
    *v9 = v10;
    HIDWORD(v29) = HIDWORD(v10) >> 0x12;
    v11 = (__int64) v26 / *((_QWORD *) &v26 + 1);
    v30 = v11 + __PAIR64__(HIDWORD(v28), v10 >> 0x12);
    (v27) = ((unsigned __int64) BYTE4(v26) + 0x1C9Ei64 * (BYTE4(v26) + 0xF366) + 1) >> 0x20;
    HIDWORD(v24) = v9[1];
    (v24) = v10;
    v12 = (BYTE4(v26) + 0x1C9Ei64 * (BYTE4(v26) + 0xF366) + 1) * *((_QWORD *) &v24 + 1);
    v13 = 4 * v28 + v12 + 0xB47D9D;
    *(_QWORD *) v9 = v13;
    v9[1] = (((unsigned __int16) v13 & 0xFFF0) + (v13 >> 0x13) + HIDWORD(v13) + (unsigned __int64) HIWORD(HIDWORD(v13)))
            >> 0x20;
    v14 = 7;
    *v9 = (v13 & 0xFFF0) + (v13 >> 0x13) + HIDWORD(v13) + HIWORD(HIDWORD(v13));
    HIDWORD(v24) = 7;
    (v28) = 0x18F - ((_DWORD) &v23 + 7);
    do {
        v15 = *v9;
        v16 = v9[1];
        *(_QWORD *) &v28 = __PAIR64__(v16 ^ (v16 >> 7), *v9 ^ (unsigned int)(*(_QWORD *)v9 >> 7)) >> 0x19;
        HIDWORD(v26) = *((unsigned __int8 *) &v22 + v14);
        *(_QWORD *) &v26 = __SPAIR64__(v16, v15) / 0x6A + v27;
        (v23) = (unsigned __int64) (v14 * v14 * (_QWORD) v26) >> 0x20;
        *((_QWORD *) &v20 + 1) = __PAIR64__(v16, *v9);
        v17 = (SHIDWORD(v25) + 1i64) * *((_QWORD *) &v20 + 1);
        v18 = v17 + v14 * v14 * v26;
        LODWORD(v17) = (v17 + __PAIR64__((v22), v14 * v14 * (int)v26)) >> 0x20;
        (v22) = v17;
        HIDWORD(v17) = HIDWORD(v21);
        *(_QWORD *) v9 = HIDWORD(v24) * HIDWORD(v24) * HIDWORD(v24) + __PAIR64__(v17, v18);
        v14 = HIDWORD(v20) - 1;
        v19 = HIDWORD(v20) - 1 < 0;
        *(_QWORD *) v9 += (__int64) v24 / 0x14C9
                + HIDWORD(v17)
                * *((unsigned __int8 *) &v20 + HIDWORD(v17))
                * *((unsigned __int8 *) &v20 + HIDWORD(v17))
                * ((int) &v20 + HIDWORD(v17) + (v25));
        HIDWORD(v20) = v14;
    } while (!v19);
    qmemcpy(code3Buf, v9, sizeof(code3Buf));
    hexdump("aAsm1ForCode3", (const unsigned __int8 *) ((v38) - 0x20), 8u);
}

int main() {
    asm1();
    return 0;
}
