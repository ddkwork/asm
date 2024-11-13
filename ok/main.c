#include <stdio.h>
#include "defs.h"

// local variable allocation has failed, the output may be wrong!
__int64   asm() {
    __int64 v2; // rax
    int v3; // ecx
    unsigned int v4; // kr00_4
    __int64 v5; // kr20_8
    unsigned __int64 v6; // rcx
    __int64 v7; // kr40_8
    __int64 v8; // rax
    unsigned __int64 v9; // kr58_8
    int index; // edi
    bool stop; // sf
    __int64 v12; // rax
    unsigned __int8 data[8]; // [esp+0h] [ebp-164h] BYREF
    int i; // [esp+1Ch] [ebp-148h]
    int v16; // [esp+28h] [ebp-13Ch]
    int v17; // [esp+2Ch] [ebp-138h]
    __int64 v18; // [esp+40h] [ebp-124h]
    __int16 *elem; // [esp+4Ch] [ebp-118h]
    unsigned __int64 v20; // [esp+50h] [ebp-114h]
    int v21; // [esp+58h] [ebp-10Ch]
    unsigned int v22; // [esp+100h] [ebp-64h]
    __int64 v23; // [esp+104h] [ebp-60h]
    int v24; // [esp+10Ch] [ebp-58h]
    int v25; // [esp+110h] [ebp-54h]
    __int64 v28; // [esp+11Ch] [ebp-48h]
    int v29; // [esp+124h] [ebp-40h] BYREF
    unsigned __int64 out; // [esp+134h] [ebp-30h] OVERLAPPED BYREF
    int result[2]; // [esp+144h] [ebp-20h] BYREF
    int savedregs; // [esp+164h] [ebp+0h] BYREF

    data[0] = 9;
    data[1] = 0x99;
    data[2] = 0x8A;
    data[3] = 0x7B;
    data[4] = 0xFE;
    data[5] = 0x46;
    data[6] = 0xC2;
    data[7] = 0xF0;

    result[0] = 0;
    result[1] = 0;
    // elem = (__int16 *) &out;
    out = 0x71B793 * (data[1] + 0x1302i64 * (data[1] + 0xF366)) + 0x7CFF86;
    v5 = data[2] + 0x1634i64 * (data[2] + 0xF366) + 1;
    HIDWORD(v6) = (v5 * out + 2 * ((__int64) out / 0x6381BE9A + (out >> 0x12)) + 0x2D1F65) >> 0x20;
    LODWORD(out) = v5 * out + 2 * ((__int64) out / 0x6381BE9A + (out >> 0x12)) + 0x2D1F65;
    LODWORD(v6) = out;
    v7 = (data[3] + 0x1968i64 * (data[3] + 0xF366) + 1) * v6 + 3 * (
             (__int64) v6 / 0x6381BE9A + (v6 >> 0x12) + 0x21D78D);
    v17 = HIDWORD(v7) >> 0x12;
    v18 = v7 / 0x6381BE9A + ((unsigned __int64) v7 >> 0x12);
    v16 = ((unsigned __int64) data[4] + 0x1C9Ei64 * (data[4] + 0xF366) + 1) >> 0x20;
    v8 = (data[4] + 0x1C9Ei64 * (data[4] + 0xF366) + 1) * v7;
    v9 = 4 * v18 + v8 + 0xB47D9D;
    index = 7;
    out = ((4 * (_DWORD) v18 + (_DWORD) v8 + 0xB47D9D) & 0xFFF0)
          + (v9 >> 0x13)
          + HIDWORD(v9)
          + (unsigned __int64) HIWORD(HIDWORD(v9));
    i = 7;
    v21 = 0x18F - (_DWORD) & data[7];
    do {
        v20 = __PAIR64__(HIDWORD(out) ^ (HIDWORD(out) >> 7), (unsigned int) out ^ (unsigned int) (out >> 7)) >> 0x19;
        elem = (__int16 *) data[index];
        v18 = (__int64) out / 0x6A
              + (__PAIR64__(HIDWORD(out) ^ (HIDWORD(out) >> 7), (unsigned int) out ^ (unsigned int) (out >> 7)) >>
                 0x19);
        v16 = ((unsigned __int64) (__int16 *) ((char *) elem + 1) * out + index * index * v18) >> 0x20;
        out = (int) elem * (int) elem * (int) elem
              + (unsigned __int64) (__int16 *) ((char *) elem + 1) * out
              + index * index * v18;
        index = i - 1;
        stop = i - 1 < 0;
        out += v18 / 0x14C9 + i * data[i] * data[i] * (int) &data[i + v21];
        --i;
    } while (!stop);
    return out;
}

int main(void) {
 asm();
    return 0;
}
