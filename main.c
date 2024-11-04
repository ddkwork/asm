#include <stdio.h>
#include <ctype.h>

void hexdump(const char *title, const unsigned char *data, size_t size) {
    // 打印标题，右对齐 20 个字符
    printf("%20s\n", title);

    for (size_t i = 0; i < size; i += 16) {
        // 打印地址
        printf("%08lx  ", (unsigned long) i);

        // 打印十六进制字节
        for (size_t j = 0; j < 8; j++) {
            if (i + j < size) {
                printf("%02x ", data[i + j]);
            } else {
                printf("   "); // 填充空格
            }
        }

        // 在8个字节后添加额外空格
        printf("  ");

        for (size_t j = 8; j < 16; j++) {
            if (i + j < size) {
                printf("%02x ", data[i + j]);
            } else {
                printf("   "); // 填充空格
            }
        }

        // 打印 ASCII 字符
        printf(" | ");
        for (size_t j = 0; j < 16; j++) {
            if (i + j < size) {
                printf("%c", isprint(data[i + j]) ? data[i + j] : '.');
            }
        }
        printf("\n");
    }
}


void asm1() {
    unsigned char HWID[8] = {0x9, 0x99, 0x8a, 0x7b, 0xfe, 0x46, 0xc2, 0xf0};
    unsigned char code3Buf[8] = {0};
    unsigned char anzhuang[8] = {0};

    hexdump("xor", HWID, 8);
    _asm{
            pushad
            pushfd
            sub esp,0xff
            xor ebx,ebx
            lea esi,HWID
            lea edi,DWORD PTR SS:[ESP]//0x18A4
            mov ecx,0x8
            rep movsb
            lea ESI,DWORD PTR DS:[anzhuang] ; int
            MOV DWORD PTR DS:[ESI],0x71B793 ; 93B7710000000000
            mov DWORD PTR DS:[ESI+0x4],EBX
            MOVZX EDI,BYTE PTR SS:[ESP+0x1] ; 0x99 20 46 45 20 []byte{0x9, 0x99, 0x8a, 0x7b, 0xfe, 0x46, 0xc2, 0xf0}
            PUSH EBX
            LEA EAX,DWORD PTR DS:[EDI+0xF366] ; data[i] 0x00000047 ??
            CDQ
            PUSH 0x1302
            PUSH EDX //0
            PUSH EAX //0x00000047 CDQ = 0x0000f3ff
            MOV DWORD PTR SS:[ESP+0x5C],ESI
            CALL __allmul // 0x0000f3ff *1302 = 0x121dd4fe
            MOV ECX,EAX
            MOV EAX,EDX
            MOV DWORD PTR SS:[ESP+0x28],EAX
            MOV EAX,EDI //99
            CDQ
            ADD ECX,EAX
            MOV EAX,DWORD PTR SS:[ESP+0x28]
            PUSH EBX
            PUSH 0x71B793
            ADC EAX,EDX
            PUSH EAX
            PUSH ECX
            CALL __allmul
            ADD EAX,0x7CFF86
            ADC EDX,EBX
            MOV DWORD PTR DS:[ESI+0x4],EDX
            MOV DWORD PTR DS:[ESI],EAX
            MOVZX EDI,BYTE PTR SS:[ESP+0x2]
            MOV ESI,EDX
            MOV EDX,EAX
            PUSH EBX
            MOV ECX,ESI
            SHRD EDX,ECX,0x12
            PUSH 0x6381BE9A
            PUSH ESI
            SHR ECX,0x12
            PUSH EAX
            MOV DWORD PTR SS:[ESP+0x50],EAX
            MOV DWORD PTR SS:[ESP+0x54],ESI
            MOV DWORD PTR SS:[ESP+0x38],EDX
            MOV DWORD PTR SS:[ESP+0x68],ECX
            CALL __alldiv
            MOV ECX,DWORD PTR SS:[ESP+0x28]
            ADD ECX,EAX
            MOV EAX,DWORD PTR SS:[ESP+0x58]
            PUSH 0x0
            PUSH 0x2
            ADC EAX,EDX
            PUSH EAX
            PUSH ECX
            CALL __allmul
            MOV ESI,EAX
            PUSH 0x0
            MOV EBX,EDX
            LEA EAX,DWORD PTR DS:[EDI+0xF366]
            CDQ
            PUSH 0x1634
            PUSH EDX
            PUSH EAX
            CALL __allmul
            MOV ECX,EAX
            MOV EAX,EDX
            MOV DWORD PTR SS:[ESP+0x28],EAX
            MOV EAX,EDI
            CDQ
            ADD ECX,EAX
            MOV EAX,DWORD PTR SS:[ESP+0x28]
            ADC EAX,EDX
            MOV EDX,DWORD PTR SS:[ESP+0x44]
            ADD ECX,0x1
            PUSH EDX
            MOV EDX,DWORD PTR SS:[ESP+0x44]
            ADC EAX,0x0
            PUSH EDX
            PUSH EAX
            PUSH ECX
            CALL __allmul
            ADD ESI,EAX
            MOV EAX,DWORD PTR SS:[ESP+0x4C]
            ADC EBX,EDX
            ADD ESI,0x2D1F65
            ADC EBX,0x0
            MOV DWORD PTR DS:[EAX],ESI
            MOV ECX,DWORD PTR DS:[EAX]
            MOV DWORD PTR DS:[EAX+0x4],EBX
            MOVZX ESI,BYTE PTR SS:[ESP+0x3]
            MOV EAX,EBX
            PUSH 0x0
            MOV EDX,ECX
            MOV EDI,EAX
            SHRD EDX,EDI,0x12
            PUSH 0x6381BE9A
            PUSH EAX
            PUSH ECX
            MOV DWORD PTR SS:[ESP+0x50],ECX
            MOV DWORD PTR SS:[ESP+0x54],EAX
            SHR EDI,0x12
            MOV EBX,EDX
            CALL __alldiv
            ADD EBX,EAX
            ADC EDI,EDX
            PUSH 0x0
            ADD EBX,0x21D78D
            PUSH 0x3
            ADC EDI,0x0
            PUSH EDI
            PUSH EBX
            CALL __allmul
            MOV EDI,EAX
            PUSH 0x0
            MOV EBX,EDX
            LEA EAX,DWORD PTR DS:[ESI+0xF366]
            CDQ
            PUSH 0x1968
            PUSH EDX
            PUSH EAX
            CALL __allmul
            MOV ECX,EAX
            MOV EAX,EDX
            MOV DWORD PTR SS:[ESP+0x28],EAX
            MOV EAX,ESI
            CDQ
            ADD ECX,EAX
            MOV EAX,DWORD PTR SS:[ESP+0x28]
            ADC EAX,EDX
            MOV EDX,DWORD PTR SS:[ESP+0x44]
            PUSH EDX
            MOV EDX,DWORD PTR SS:[ESP+0x44]
            ADD ECX,0x1
            PUSH EDX
            ADC EAX,0x0
            PUSH EAX
            PUSH ECX
            CALL __allmul
            MOV ESI,DWORD PTR SS:[ESP+0x4C]
            ADD EDI,EAX
            ADC EBX,EDX
            MOV EAX,EBX
            MOV EDX,EBX
            PUSH 0x0
            MOV ECX,EDI
            SHRD ECX,EDX,0x12
            PUSH 0x6381BE9A
            PUSH EAX
            SHR EDX,0x12
            MOV DWORD PTR DS:[ESI+0x4],EBX
            PUSH EDI
            MOV DWORD PTR DS:[ESI],EDI
            MOV EBX,ECX
            MOV DWORD PTR SS:[ESP+0x3C],EDX
            CALL __alldiv
            ADD EBX,EAX
            MOV EAX,DWORD PTR SS:[ESP+0x2C]
            MOV DWORD PTR SS:[ESP+0x40],EBX
            MOVZX EBX,BYTE PTR SS:[ESP+0x4]
            ADC EAX,EDX
            MOV DWORD PTR SS:[ESP+0x44],EAX
            PUSH 0x0
            LEA EAX,DWORD PTR DS:[EBX+0xF366]
            CDQ
            PUSH 0x1C9E
            PUSH EDX
            PUSH EAX
            CALL __allmul
            MOV ECX,EAX
            MOV EAX,EDX
            MOV DWORD PTR SS:[ESP+0x28],EAX
            MOV EAX,EBX
            CDQ
            ADD ECX,EAX
            MOV EAX,DWORD PTR SS:[ESP+0x28]
            ADC EAX,EDX
            ADD ECX,0x1
            ADC EAX,0x0
            MOV DWORD PTR SS:[ESP+0x28],EAX
            MOV EAX,DWORD PTR DS:[ESI+0x4]
            PUSH EAX
            MOV EAX,DWORD PTR SS:[ESP+0x2C]
            PUSH EDI
            PUSH EAX
            PUSH ECX
            CALL __allmul
            PUSH 0x0
            MOV EBX,EDX
            MOV EDX,DWORD PTR SS:[ESP+0x48]
            PUSH 0x4
            MOV EDI,EAX
            MOV EAX,DWORD PTR SS:[ESP+0x48]
            PUSH EDX
            PUSH EAX
            CALL __allmul
            ADD EDI,EAX
            ADC EBX,EDX
            ADD EDI,0xB47D9D
            ADC EBX,0x0
            MOV EAX,EBX
            MOV ECX,EDI
            MOV EDX,EAX
            MOV DWORD PTR DS:[ESI],EDI
            XOR EDI,EDI
            MOV DWORD PTR DS:[ESI+0x4],EBX
            XOR EBX,EBX
            SHR EDX,0x10
            ADD EDX,EAX
            ADC EDI,EBX
            MOV EBX,ECX
            SHRD EBX,EAX,0x13
            SHR EAX,0x13
            ADD EDX,EBX
            ADC EDI,EAX
            AND ECX,0xFFF0
            XOR EAX,EAX
            ADD EDX,ECX
            ADC EDI,EAX
            MOV EAX,0x18F
            LEA ECX,DWORD PTR SS:[ESP+0x7]
            MOV DWORD PTR DS:[ESI+0x4],EDI
            MOV EDI,0x7
            SUB EAX,ECX
            MOV DWORD PTR DS:[ESI],EDX
            MOV DWORD PTR SS:[ESP+0x1C],EDI
            MOV DWORD PTR SS:[ESP+0x58],EAX
            JMP L78

            L78:
            MOV EAX,DWORD PTR DS:[ESI] ; asm
            MOV EBX,DWORD PTR DS:[ESI+0x4]
            MOV EDX,EAX
            MOV ECX,EBX
            SHRD EDX,ECX,0x7
            SHR ECX,0x7
            XOR ECX,EBX
            XOR EDX,EAX
            PUSH 0x0
            SHRD EDX,ECX,0x19
            PUSH 0x6A
            PUSH EBX
            SHR ECX,0x19
            PUSH EAX
            MOV DWORD PTR SS:[ESP+0x60],EDX
            MOV DWORD PTR SS:[ESP+0x64],ECX
            CALL __alldiv
            MOV ECX,DWORD PTR SS:[ESP+0x50]
            ADD ECX,EAX
            MOV EAX,DWORD PTR SS:[ESP+0x54]
            ADC EAX,EDX
            MOVZX EDX,BYTE PTR SS:[ESP+EDI]
            MOV DWORD PTR SS:[ESP+0x44],EAX
            MOV EAX,EDI
            IMUL EAX,EDI
            MOV EDI,DWORD PTR SS:[ESP+0x44]
            PUSH EDI
            MOV DWORD PTR SS:[ESP+0x50],EDX
            CDQ
            PUSH ECX
            PUSH EDX
            PUSH EAX
            MOV DWORD PTR SS:[ESP+0x50],ECX
            CALL __allmul
            MOV EDI,EAX
            MOV EAX,EDX
            MOV DWORD PTR SS:[ESP+0x28],EAX
            MOV EAX,DWORD PTR SS:[ESP+0x4C]
            CDQ
            MOV ECX,EAX
            MOV EAX,DWORD PTR DS:[ESI]
            PUSH EBX
            ADD ECX,0x1
            PUSH EAX
            ADC EDX,0x0
            PUSH EDX
            PUSH ECX
            CALL __allmul
            MOV ECX,DWORD PTR SS:[ESP+0x4C]
            ADD EDI,EAX
            MOV EAX,DWORD PTR SS:[ESP+0x28]
            ADC EAX,EDX
            MOV DWORD PTR SS:[ESP+0x28],EAX
            MOV EAX,ECX
            IMUL EAX,ECX
            IMUL EAX,ECX
            CDQ
            ADD EDI,EAX
            MOV EAX,DWORD PTR SS:[ESP+0x28]
            ADC EAX,EDX
            MOV EDX,DWORD PTR SS:[ESP+0x1C]
            MOV DWORD PTR DS:[ESI+0x4],EAX
            LEA EAX,DWORD PTR SS:[ESP+EDX]
            MOV DWORD PTR DS:[ESI],EDI
            MOVZX ECX,BYTE PTR DS:[EAX]
            MOV EDI,DWORD PTR SS:[ESP+0x58]
            ADD EAX,EDI
            IMUL EAX,ECX
            IMUL EAX,ECX
            IMUL EAX,EDX
            CDQ
            PUSH 0x0
            MOV EBX,EDX
            MOV EDX,DWORD PTR SS:[ESP+0x48]
            PUSH 0x14C9
            MOV EDI,EAX
            MOV EAX,DWORD PTR SS:[ESP+0x48]
            PUSH EDX
            PUSH EAX
            CALL __alldiv
            ADD EDI,EAX
            MOV ECX,DWORD PTR DS:[ESI+0x4]
            ADC EBX,EDX
            MOV EDX,DWORD PTR DS:[ESI]
            ADD EDX,EDI
            MOV EDI,DWORD PTR SS:[ESP+0x1C]
            ADC ECX,EBX
            DEC EDI
            MOV DWORD PTR DS:[ESI],EDX
            MOV DWORD PTR DS:[ESI+0x4],ECX
            MOV DWORD PTR SS:[ESP+0x1C],EDI
            JNS L78 ; ebx 1A8B^
            JMP LOK

            ////	<__allmul>  Src: []byte{0x9, 0x99, 0x8a, 0x7b, 0xfe, 0x46, 0xc2, 0xf0} //09998A7BFE46C2F0,
            __allmul:
            MOV EAX,DWORD PTR SS:[ESP+0x8]
            MOV ECX,DWORD PTR SS:[ESP+0x10]
            OR ECX,EAX
            MOV ECX,DWORD PTR SS:[ESP+0xC]
            JNZ L7
            MOV EAX,DWORD PTR SS:[ESP+0x4]
            MUL ECX
            RETN 0x10

            L7:
            PUSH EBX
            MUL ECX
            MOV EBX,EAX
            MOV EAX,DWORD PTR SS:[ESP+0x8]
            MUL DWORD PTR SS:[ESP+0x14]
            ADD EBX,EAX
            MOV EAX,DWORD PTR SS:[ESP+0x8]
            MUL ECX
            ADD EDX,EBX
            POP EBX
            RETN 0x10

            ////////////////////////////////////////////////
            //<__alldiv>
            __alldiv:
            PUSH EBX
            PUSH ESI
            MOV EAX,DWORD PTR SS:[ESP+0x18]
            OR EAX,EAX
            JNZ L8
            MOV ECX,DWORD PTR SS:[ESP+0x14]
            MOV EAX,DWORD PTR SS:[ESP+0x10]
            XOR EDX,EDX
            DIV ECX
            MOV EBX,EAX
            MOV EAX,DWORD PTR SS:[ESP+0xC]
            DIV ECX
            MOV EDX,EBX
            JMP BUGDEDE0005

            L8:
            MOV ECX,EAX
            MOV EBX,DWORD PTR SS:[ESP+0x14]
            MOV EDX,DWORD PTR SS:[ESP+0x10]
            MOV EAX,DWORD PTR SS:[ESP+0xC]

            BUGDEDE0002:
            SHR ECX,1
            RCR EBX,1
            SHR EDX,1
            RCR EAX,1
            OR ECX,ECX
            JNZ BUGDEDE0002
            DIV EBX
            MOV ESI,EAX
            MUL DWORD PTR SS:[ESP+0x18]
            MOV ECX,EAX
            MOV EAX,DWORD PTR SS:[ESP+0x14]
            MUL ESI
            ADD EDX,ECX
            JB BUGDEDE0003
            CMP EDX,DWORD PTR SS:[ESP+0x10]
            JA BUGDEDE0003
            JB BUGDEDE0004
            CMP EAX,DWORD PTR SS:[ESP+0xC]
            JBE BUGDEDE0004

            BUGDEDE0003:
            DEC ESI

            BUGDEDE0004:
            XOR EDX,EDX
            MOV EAX,ESI

            BUGDEDE0005:
            POP ESI
            POP EBX
            RETN 0x10
            //////////////////////////////////////////////////////////
            LOK:
            //  LEA AA,[ESI]
            //    LEA AA,DWORD PTR DS:[ESI]
            lea edi,[code3Buf]
            mov ECX,0x8
            rep movsb

            add esp,0xff
            popfd
            popad
            }
    hexdump("asm1 for code3", code3Buf, 8);
}

/*
/////////////////////// ghidra test /////////////////////////////////////////////////
typedef union _LARGE_INTEGER {
    struct {
        unsigned long LowPart;
        unsigned long HighPart;
    } u;

    unsigned long long QuadPart;
} LARGE_INTEGER;

LARGE_INTEGER make_large_integer( unsigned long high,  unsigned long low) {
    LARGE_INTEGER v;
    v.u.HighPart = high;
    v.u.LowPart = low;
    return v;
}


typedef unsigned char undefined; //size = 1

typedef unsigned char uchar; //size = 1
typedef unsigned char byte; //size = 1
typedef unsigned char dwfenc; //size = 1
typedef unsigned short word; //size = 2
typedef unsigned int dword; //size = 4
typedef long long longlong; //size = 8
typedef unsigned long long qword; //size = 8
typedef unsigned short ushort; //size = 2
typedef unsigned int uint; //size = 4
typedef unsigned long ulong; //size = 4
typedef unsigned long long ulonglong; //size = 8

typedef long double float10;


typedef unsigned char undefined1; //size = 1
typedef unsigned short undefined2; //size = 2
typedef struct {
    byte arr[3];
} undefined3; //size = 3
typedef unsigned int undefined4; //size = 4
typedef struct {
    byte arr[5];
} undefined5; //size = 5
typedef struct {
    byte arr[6];
} undefined6; //size = 6
typedef struct {
    byte arr[7];
} undefined7; //size = 7
typedef LARGE_INTEGER undefined8; //size = 8

typedef struct {
    byte arr[16];
} arr16;

typedef struct {
    uint arr[16];
} arr64;


LARGE_INTEGER mul(uint param_1, uint param_2, uint param_3, uint param_4) {
    LARGE_INTEGER v;
    if ((param_4 | param_2) == 0) {
        (ulonglong) param_1 * (ulonglong) param_3;
        v.QuadPart = param_1 * (ulonglong) param_3;
        return v;
    }
    v.u.HighPart = ((ulonglong) param_1 * (ulonglong) param_3 >> 0x20) +param_2 * param_3 + param_1 * param_4 << 32;
    v.u.HighPart = ((ulonglong) param_1 * (ulonglong) param_3);
    return v;
}

LARGE_INTEGER div(uint param_1, uint param_2, uint param_3, uint param_4) {
    ulonglong uVar1;
    longlong lVar2;
    uint uVar3;
    int iVar4;
    uint uVar5;
    uint uVar6;
    uint uVar7;
    uint uVar8;
    uint uVar9;

    uVar3 = param_1;
    uVar8 = param_4;
    uVar6 = param_2;
    uVar9 = param_3;
    if (param_4 == 0) {
        uVar3 = param_2 / param_3;
        iVar4 = (int) (((ulonglong) param_2 % (ulonglong) param_3 << 0x20 | (ulonglong) param_1) /
                       (ulonglong) param_3);
    } else {
        do {
            uVar5 = uVar8 >> 1;
            uVar9 = uVar9 >> 1 | (uint) ((uVar8 & 1) != 0) << 0x1f;
            uVar7 = uVar6 >> 1;
            uVar3 = uVar3 >> 1 | (uint) ((uVar6 & 1) != 0) << 0x1f;
            uVar8 = uVar5;
            uVar6 = uVar7;
        } while (uVar5 != 0);
        uVar1 = (uVar7 << 32 | uVar3) / (ulonglong) uVar9;
        iVar4 = (int) uVar1;
        lVar2 = (ulonglong) param_3 * (uVar1 & 0xffffffff);
        uVar3 = (uint) ((ulonglong) lVar2 >> 0x20);
        uVar8 = uVar3 + iVar4 * param_4;
        if (uVar3<(iVar4 * param_4) ||
            param_2 < uVar8 ||
            param_2 <= uVar8 && param_1 < (uint) lVar2){
            iVar4 += -1;
        }
        uVar3 = 0;
    }
    LARGE_INTEGER v;
    v.u.HighPart = uVar3 << 32;
    v.u.LowPart = iVar4;
    return v;
}


void __cdecl ghidraDecode() {
    LARGE_INTEGER uVar1;
    LARGE_INTEGER puVar2;
    byte *pbVar3;
    int iVar4;
    uint uVar5;
    uchar *puVar6;
    uint uVar7;
    undefined4 *puVar8;
    uint uVar9;
    uint uVar10;
    byte in_AF;
    byte in_TF;
    byte in_IF;
    byte in_NT;
    byte in_AC;
    byte in_VIF;
    byte in_VIP;
    byte in_ID;
    LARGE_INTEGER lVar11;
    LARGE_INTEGER lVar12;
    LARGE_INTEGER uVar13;
    byte local_167[4];
    byte local_163;
    undefined local_160[21];
    int local_14b;
    uint local_13f;
    undefined4 local_13b;
    undefined8 local_127;
    LARGE_INTEGER local_11b;
    uint local_117;
    uint local_113;
    int local_10f;
    uint uStack_68;
    undefined4 *puStack_64;
    undefined4 local_38;
    LARGE_INTEGER local_34;
    undefined4 local_24;
    undefined4 local_20;
    uchar local_14[4];
    undefined local_10;
    undefined local_f;
    undefined local_e;
    undefined local_d;
    // uint local_8;

    puVar8 = &local_38;
    for (iVar4 = 0xd; iVar4 != 0; iVar4 += -1) {
        *puVar8 = 0xcccccccc;
        puVar8 = puVar8 + 1;
    }
    // local_8 = __security_cookie ^ (uint) &stack0xfffffffc;
    local_14[0] = '\t'; //todo 0x71B793 ?
    local_14[1] = 0x99;
    local_14[2] = 0x8a;
    local_14[3] = 0x7b;
    local_10 = 0xfe;
    local_f = 0x46;
    local_e = 0xc2;
    local_d = 0xf0;
    local_24 = 0;
    local_20 = 0;
    local_34.u.HighPart = 0;
    hexdump("data", local_14, 8);
    puStack_64 = puVar8;
    // uStack_68 = (uint) (in_NT & 1) * 0x4000 | (uint) SCARRY4((int) &stack0xffffffb0, 0xc) * 0x800 |
    //             (uint) (in_IF & 1) * 0x200 | (uint) (in_TF & 1) * 0x100 |
    //             (uint) ((int) &stack0xffffffbc < 0) * 0x80 |
    //             (uint) (&stack0x00000000 == (undefined *) 0x44) * 0x40 | (uint) (in_AF & 1) * 0x10 |
    //             (uint) ((POPCOUNT((uint) &stack0xffffffbc & 0xff) & 1U) == 0) * 4 |
    //             (uint) ((undefined *) 0xfffffff3 < &stack0xffffffb0) | (uint) (in_ID & 1) * 0x200000 |
    //             (uint) (in_VIP & 1) * 0x100000 | (uint) (in_VIF & 1) * 0x80000 |
    //             (uint) (in_AC & 1) * 0x40000;
    puVar6 = local_14;
    pbVar3 = local_167;
    for (iVar4 = 8; iVar4 != 0; iVar4 += -1) {
        *pbVar3 = *puVar6;
        puVar6 = puVar6 + 1;
        pbVar3 = pbVar3 + 1;
    }
    // local_11b = local_34;
    local_34.u.HighPart = 0x71b793;
    lVar11 = mul(local_167[1] + 0xf366, 0, 0x1302, 0);
    local_13f = (uint)  lVar11.QuadPart >> 0x20;
    lVar11.QuadPart += (ulonglong) (uint) local_167[1];
    lVar11 = mul( lVar11.u.HighPart,  ( lVar11.QuadPart >> 0x20), 0x71b793, 0);
    local_127.QuadPart = lVar11.QuadPart + 0x7cff86;
    uVar9 = (uint) local_167[2];
    local_13f = (local_127.u.HighPart >> 0x12);
    local_10f = ((local_127.u.HighPart >> 0x12) >> 0x20);
    local_34 = local_127;
    lVar11 = div(local_127.u.LowPart, (uint) (local_127.QuadPart >> 0x20), 0x6381be9a, 0);
    lVar11.QuadPart += (local_10f << 32 | local_13f);
    lVar11 = mul( lVar11.u.HighPart, (uint) ( lVar11.QuadPart >> 0x20), 2, 0);
    lVar12= mul(uVar9 + 0xf366, 0, 0x1634, 0);
    local_13f = (uint) ( lVar12.QuadPart >> 0x20);
    uVar5 =  lVar12.u.HighPart + uVar9;
    lVar12 = mul(
        uVar5 + 1,
        local_13f + ( lVar12.u.HighPart << 16 | uVar9) + (uint) (0xfffffffe < uVar5),
        local_127.u.HighPart,
        local_127.u.LowPart);
    uVar13.QuadPart = lVar12.u.HighPart + lVar11.u.HighPart + 0x2d1f65;
    local_127.u.HighPart = (uint) (uVar13.u.HighPart >> 0x20);
    local_127.QuadPart = local_11b.QuadPart;
    local_11b = uVar13;
    uVar7 = (uint) local_167[3];
    uVar5 = local_127.u.LowPart >> 0x12;
    uVar9 = local_127.u.LowPart << 0xe;
    uVar10 = local_127.u.HighPart >> 0x12;
    lVar11 = div(local_127.u.HighPart, local_127.u.LowPart, 0x6381be9a, 0); //todo  /
    lVar11.QuadPart+=  (uVar10 << 32 | (uVar5 | uVar9)) + 0x21d78d;
    lVar11 = mul( lVar11.u.HighPart, (uint) ( lVar11.QuadPart >> 0x20), 3, 0);
    lVar12 = mul(uVar7 + 0xf366, 0, 0x1968, 0);
    local_13f = (uint) ( lVar12.QuadPart >> 0x20);
    uVar5 =  lVar12.u.HighPart + uVar7;
    lVar12 = mul(uVar5 + 1, local_13f +  lVar12.u.HighPart|uVar7 + (uint) (0xfffffffe < uVar5),
                  local_127.u.HighPart, local_127.u.HighPart);
    puVar2.QuadPart = local_11b.QuadPart;
    uVar13.QuadPart = lVar12.QuadPart + lVar11.QuadPart;
    uVar7 =  uVar13.u.HighPart;
    uVar9 = 0;
    uVar1.QuadPart = uVar13.u.LowPart >> 0x12;
    uVar5 = 0x6381be9a;
    local_13b = (undefined4) (uVar1.u.HighPart >> 0x20);
    local_11b = uVar13;
    lVar11 = div( uVar13.u.LowPart,  (uVar13.QuadPart >> 0x20), uVar5, uVar9);
    local_127.u.HighPart = lVar11.QuadPart + (local_13b << 32 |  uVar1.u.HighPart);
    uVar9 = (uint) local_163;
    lVar11 = mul(uVar9 + 0xf366, 0, 0x1c9e, 0);
    uVar5 =  lVar11.QuadPart + uVar9;
    local_13f = (int) ( lVar11.QuadPart>> 0x20) + make_large_integer( lVar11.u.HighPart, uVar9).QuadPart +// todo make a macro for this?
                (uint) (0xfffffffe < uVar5);
    lVar11 = mul(uVar5 + 1, local_13f, uVar7, *(uint *) ( puVar2.u.HighPart + 4));
    lVar12 = mul( local_127.u.HighPart, local_127.u.HighPart, 4, 0);
    uVar13.QuadPart = lVar12.u.HighPart + lVar11.u.HighPart + 0xb47d9d;
    uVar5 =  (uVar13.u.HighPart >> 0x20);
    puVar2 = uVar13;
    local_10f = 399 - (int) local_160;
    puVar2.QuadPart = (uVar13.QuadPart >> 0x13) +make_large_integer(uVar5 >> 0x10, uVar5).QuadPart, (uVar5 >> 0x10) + uVar5 +( uVar13.u.HighPart & 0xfff0);
    local_14b = 7;
    do {
        iVar4 = local_14b;
        uVar5 = *(uint *) ( puVar2.u.HighPart + 4);
        local_117 = (uVar5 << 0x19 ^  puVar2.u.HighPart) >> 0x19 | (uVar5 >> 7 ^ uVar5) << 7;
        local_113 = uVar5 >> 0x19;
        lVar11 = div( puVar2.u.HighPart, uVar5, 0x6a, 0);
        local_127.QuadPart = lVar11.QuadPart + (local_113 << 32 | local_117);
        local_11b.QuadPart =   local_167[iVar4];
        lVar12 = mul(iVar4 * iVar4, iVar4 * iVar4 >> 0x1f,  local_127.u.HighPart,  (local_127.u.HighPart >> 0x20));
        local_13f = (uint) ( lVar12.QuadPart >> 0x20);
        lVar11 = mul( local_11b.u.HighPart + 1,
                     ( local_11b.u.HighPart >> 0x1f) + (uint) ((ulonglong ) 0xfffffffe < local_11b.QuadPart),
                       puVar2.u.HighPart, uVar5);
        iVar4 = local_14b;
        lVar11.QuadPart += (local_13f << 32 |  lVar12.u.HighPart);
        local_13f = (uint) ( lVar11.QuadPart >> 0x20);
        pbVar3 = local_167 + local_14b;
        puVar2.QuadPart = lVar11.u.HighPart +  local_11b.u.HighPart *  local_11b.u.HighPart *  local_11b.u.HighPart;
        uVar5 = (uint) *pbVar3;
        pbVar3 = pbVar3 + local_10f;
        lVar11 = div( local_127.u.HighPart, local_127.u.HighPart, 0x14c9, 0); //todo
        local_14b += -1;
        puVar2.QuadPart = lVar11.u.HighPart + (int) ((int) pbVar3 * uVar5 * uVar5 * iVar4) + puVar2.u.HighPart;
    } while (-1 < local_14b);
    printf("Large integer result use ghidra\n");
    printf("%08lx  ", puVar2.u.HighPart);
    printf("%08lx  ", puVar2.u.LowPart);
}


/////////////////////////////////////////////////////////////////////////////////////


void test() {
    LARGE_INTEGER div;
    div.QuadPart = 0x0000091416154ED7 / 0x6A;
    LARGE_INTEGER mul;
    mul.QuadPart = 0x00000015ECE7BB69 * 0x31;
    printf("0x%016llX\n", div.QuadPart);
    printf("0x%016llX\n", div.u.HighPart);
    printf("0x%016llX\n", div.u.LowPart);

    printf("0x%016llX\n", mul.QuadPart);
    printf("0x%016llX\n", mul.u.HighPart);
    printf("0x%016llX\n", mul.u.LowPart);
}
*/

int main(void) {
    // test();
    asm1();
    // ghidraDecode();
    return 0;
}
