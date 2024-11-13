#include <string.h>
#include "defs.h"

unsigned __int64 __stdcall mul(unsigned __int64 a1, unsigned __int64 a2) {
    return a1 * a2;
}

unsigned __int64 __stdcall div(unsigned __int64 a1, unsigned __int64 a2) { return a1 / a2; }

__int64 asm() {
    __int64 v2=0; // rax
    int v3=0; // ecx
    unsigned int v4=0; // kr00_4
    int v5=0; // edi
    unsigned __int64 v6; // rax
    int v7=0; // edi
    unsigned __int64 v8=0; // rax
    unsigned __int64 v9=0; // kr18_8
    unsigned __int64 v10=0; // rax
    unsigned __int64 v11=0; // rax
    unsigned int v12=0; // kr04_4
    unsigned __int64 v13=0; // rcx
    int v14=0; // esi
    unsigned int v15=0; // edi
    __int64 v16=0; // rax
    unsigned __int64 v17=0; // kr28_8
    unsigned __int64 v18=0; // rax
    unsigned __int64 v19=0; // rax
    _DWORD *v20; // esi
    int v21=0; // edi
    unsigned int v22=0; // eax
    bool v23; // cf
    int v24=0; // edx
    __int64 v25=0; // kr38_8
    unsigned __int64 v26=0; // rax
    unsigned __int64 v27=0; // kr40_8
    int index=0; // edi
    unsigned int v29=0; // ebx
    __int64 v30=0; // rax
    unsigned __int64 v31=0; // rax
    unsigned __int64 v32=0; // kr58_8
    int i=0; // edx
    __int64 v34=0; // rax
    unsigned int v35=0; // ebx
    unsigned int v36=0; // edi
    __int64 v37=0; // rax
    unsigned __int64 v38=0; // kr60_8
    bool stop; // sf
    unsigned __int64 v41=0; // [esp-10h] [ebp-174h]
    unsigned __int64 v42=0; // [esp-10h] [ebp-174h]
    unsigned __int64 v43=0; // [esp-10h] [ebp-174h]
    __int64 v44=0; // [esp-8h] [ebp-16Ch]
    __int64 v45=0; // [esp-8h] [ebp-16Ch]
    unsigned char data[8]; // [esp+0h] [ebp-164h] BYREF
    int num=0; // [esp+1Ch] [ebp-148h]
    unsigned int v48=0; // [esp+28h] [ebp-13Ch]
    int v49; // [esp+2Ch] [ebp-138h]
    unsigned __int64 v50=0; // [esp+40h] [ebp-124h]
    unsigned __int64 *v51; // [esp+4Ch] [ebp-118h]
    unsigned __int64 v52=0; // [esp+50h] [ebp-114h]
    unsigned int v53=0; // [esp+58h] [ebp-10Ch]
    __int64 v60; // [esp+11Ch] [ebp-48h]
    unsigned __int64 v62=0; // [esp+134h] [ebp-30h] BYREF
    unsigned __int64 out=0; // [esp+144h] [ebp-20h] BYREF

    data[0] = 9;
    data[1] = 0x99;
    data[2] = 0x8A;
    data[3] = 0x7B;
    data[4] = 0xFE;
    data[5] = 0x46;
    data[6] = 0xC2;
    data[7] = 0xF0;

    v62 = 0x71B793i64;
    v5 = data[1];
    v51 = &v62;


    v6 = mul(data[1] + 0xF366, 0x1302i64);
    v48 = HIDWORD(v6);
    v62 = mul(v5 + v6, 0x71B793i64) + 0x7CFF86;
    v7 = data[2];
    v50 = v62;
    v48 = v62 >> 0x12;
    v53 = HIDWORD(v62) >> 0x12;
    LODWORD(v8) = div(v62, 0x6381BE9Ai64);
    // v9 = mul(v8 + __PAIR64__(v53, v48), 2i64); //todo bug need 4
    int xp=v8+v62>>18;
    xp=xp;
    v9 = mul(v8+(v62>>18), 2i64); //todo bug need 4
    v10 = mul(v7 + 0xF366, 0x1634i64);
    v48 = HIDWORD(v10);
    v11 = mul(v7 + v10 + 1, v50);


    v12 = v11;
    LODWORD(v11) = v51;
    HIDWORD(v13) = (__PAIR64__(HIDWORD(v11), v12) + v9 + 0x2D1F65) >> 0x20;
    *(_DWORD *) v51 = v12 + v9 + 0x2D1F65;
    LODWORD(v13) = v11;
    // *(_DWORD *) (v11 + 4) = HIDWORD(v13);
    HIDWORD(v11) = HIDWORD(v13);

    v14 = data[3];
    v41 = v13;
    v50 = v13;
    v15 = HIDWORD(v13) >> 0x12;
    HIDWORD(v13) = v13 >> 0x12;
    LODWORD(v16) = div(v41, 0x6381BE9Ai64); //todo bug


    v17 = mul(v16 + __PAIR64__(v15, HIDWORD(v13)) + 0x21D78D, 3i64); //48F7E51C 00008421


    v18 = mul(v14 + 0xF366, 0x1968i64); //18340C68
    v48 = HIDWORD(v18);
    v19 = mul(v14 + v18 + 1, *(_DWORD *) v51);

    v20 = v51;
    HIDWORD(v13) = (v19 + v17) >> 0x20;
    v21 = v19 + v17;
    HIDWORD(v19) = HIDWORD(v13) >> 0x12;
    *((_DWORD *) v51 + 1) = HIDWORD(v13);
    v42 = __PAIR64__(HIDWORD(v13), (int)v19 + (int)v17);
    *v20 = v19 + v17;
    HIDWORD(v13) = __PAIR64__(HIDWORD(v13), (int)v19 + (int)v17) >> 0x12;
    v49 = HIDWORD(v19);
    v22 = div(*v20, 0x6381BE9Ai64);//todo bug
    // v23 = __CFADD__(v22, HIDWORD(v13));
    v23 = v22 + HIDWORD(v13);

    LODWORD(v50) = v22 + HIDWORD(v13);
    HIDWORD(v13) = data[4];
    // HIDWORD(v50) = v24 + v23 + v49;
    (v50) = HIDWORD(v22) + v49;
    v25 = SHIDWORD(v13) + mul(data[4] + 0xF366, 0x1C9Ei64) + 1;
    v48 = HIDWORD(v25);
    HIDWORD(v44) = v20[1];
    LODWORD(v44) = v21;
    v26 = mul(v25, v44);
    v27 = mul(v50, 4i64) + v26 + 0xB47D9D;
    *(_QWORD *) v20 = v27;
    v20[1] = (((unsigned __int16) v27 & 0xFFF0) + (v27 >> 0x13) + HIDWORD(v27) + (unsigned __int64)
              HIWORD(HIDWORD(v27))) >> 0x20;
    index = 7;
    *v20 = (v27 & 0xFFF0) + (v27 >> 0x13) + HIDWORD(v27) + HIWORD(HIDWORD(v27));
    num = 7;
    v53 = 0x18F - (_DWORD) &data[7];
    do {
        v29 = v20[1];
        v43 = *(_QWORD *) v20;
        v52 = __PAIR64__(v29 ^ (v29 >> 7), *v20 ^ (unsigned int)(*(_QWORD *)v20 >> 7)) >> 0x19;
        LODWORD(v30) = div(v43, 0x6Ai64);
        v51 = (unsigned __int64 *) data[index];
        v50 = v30 + v52;
        v31 = mul(index * index, v30 + v52);
        v48 = HIDWORD(v31);
        HIDWORD(v45) = v29;
        LODWORD(v45) = *v20;
        v32 = mul((__int64) v51 + 1, v45) + v31;
        v48 = HIDWORD(v32);
        i = num;
        *(_QWORD *) v20 = (int) v51 * (int) v51 * (int) v51 + v32;
        v34 = i * data[i] * data[i] * (int) &data[i + v53];
        v35 = HIDWORD(v34);
        v36 = v34;
        LODWORD(v37) = div(v50, 0x14C9i64);
        v38 = v37 + __PAIR64__(v35, v36) + *(_QWORD *) v20;
        index = num - 1;
        stop = num - 1 < 0;
        *(_QWORD *) v20 = v38;
        num = index;
    } while (!stop);
    memcpy(out, v20, sizeof(out));
    return out;
}


int main(void) {
    asm();
    return 0;
}
