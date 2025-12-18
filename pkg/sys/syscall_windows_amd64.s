#include "textflag.h"

#define IsC 	   $63	//0x8000000000000000 
#define FOut	   $62	//0x4000000000000000
#define IsF1 	   $61	//0x2000000000000000
#define IsF2 	   $60	//0x1000000000000000
#define IsF3 	   $59	//0x0800000000000000
#define IsF4 	   $58	//0x0400000000000000
#define IsVar2    $46   //0x0000400000000000
#define IsVar3    $45   //0x0000200000000000
#define IsVar4    $44   //0x0000100000000000
#define IsVar5    $43   //0x0000080000000000
#define IsVar6    $42   //0x0000040000000000
#define IsVar7    $41   //0x0000020000000000
#define IsVar8    $40   //0x0000010000000000
#define IsVar9    $39   //0x0000008000000000
#define IsVar10   $38   //0x0000004000000000
#define IsVar11   $37   //0x0000002000000000
#define IsVar12   $36   //0x0000001000000000
#define IsVar13   $35   //0x0000000800000000
#define IsVar14   $34   //0x0000000400000000

// -------------------------------------------------------------------------- //

// for test
// TEXT ·call3(SB),NOSPLIT, $32-48
//    MOVQ x2+24(FP), DX
//    MOVQ x3+32(FP), R8

//    MOVQ func+0(FP), AX
//    MOVQ mask+8(FP), R11
//    TESTQ $0x01, R11
//    JNZ c
//    MOVQ x1+16(FP), R10
//    SYSCALL
//    JMP retX
//    c:MOVQ x1+16(FP), CX
//       x1:TESTQ $0x02, R11
//          JZ x2
//          MOVQ x+16(FP), X0
//       x2:TESTQ $0x04, R11
//          JZ x3
//          MOVQ x2+24(FP), X1
//       x3:TESTQ $0x08, R11
//          JZ call
//          MOVQ x3+32(FP), X2
//       call: TESTQ $0x0F, SP
//             JNZ grow
//             CALL AX
//             JMP retX
//             grow:
//                SUBQ $8, SP 
//                CALL AX
//                ADDQ $8, SP
//    retX: TESTQ $0x100, mask+8(FP)
//          JNZ retF
//          MOVQ AX, r+40(FP)
//          RET
//    retF: MOVQ X0, r+40(FP)
//          RET

// -------------------------------------------------------------------------- //

#define init\
   MOVL mask+8(FP), R11\
   BTL IsC, R11\
   JC c

#define syscall\
   MOVQ x1+16(FP), R10\
   SYSCALL\
   JMP retX

#define checkFloat(mask,lbl,shift,reg)\
   BTL mask, R11\
   JC lbl\
   MOVQ x+shift(FP), reg

#define checkCall\ 
   MOVQ func+0(FP), AX\
   TESTQ $0x0F, SP\
   JNZ grow\
   CALL AX\
   JMP retX\

#define growCall\
   SUBQ $8, SP\
   CALL AX\
   ADDQ $8, SP

#define checkRet(shift)\
   BTL FOut, mask+8(FP)\
   JC retF\
   ret(AX,shift)

#define ret(reg,shift)\
   MOVQ reg, r+shift(FP)\
   RET

// -------------------------------------------------------------------------- //  

TEXT ·call3(SB), NOSPLIT, $32-48
   MOVQ x2+24(FP), DX
   MOVQ x3+32(FP), R8

   init;syscall
   c: MOVQ x1+16(FP), CX
      x1:   checkFloat(IsF1,x2,16,X0)
      x2:   checkFloat(IsF2,x3,24,X1)
      x3:   checkFloat(IsF3,call,32,X2)
      call: checkCall
      grow: growCall
   retX: checkRet(40)
   retF: ret(X0,40)

TEXT ·call6(SB), NOSPLIT, $48-72
   MOVQ x2+24(FP), DX
   MOVQ x3+32(FP), R8
   MOVQ x4+40(FP), R9
   MOVQ x5+48(FP), AX; MOVQ AX, 32(SP)
   MOVQ x6+56(FP), AX; MOVQ AX, 40(SP)

   init;syscall
   c: MOVQ x1+16(FP), CX
      x1:   checkFloat(IsF1,x2,16,X0)
      x2:   checkFloat(IsF2,x3,24,X1)
      x3:   checkFloat(IsF3,x4,32,X2)
      x4:   checkFloat(IsF4,call,40,X3)
      call: checkCall
      grow: growCall
   retX: checkRet(64)
   retF: ret(X0,64)

TEXT ·call9(SB), NOSPLIT, $72-96
   MOVQ x2+24(FP), DX
   MOVQ x3+32(FP), R8
   MOVQ x4+40(FP), R9
   MOVQ x5+48(FP), AX; MOVQ AX, 32(SP)
   MOVQ x6+56(FP), AX; MOVQ AX, 40(SP)
   MOVQ x7+64(FP), AX; MOVQ AX, 48(SP)
   MOVQ x8+72(FP), AX; MOVQ AX, 56(SP)
   MOVQ x9+80(FP), AX; MOVQ AX, 64(SP)
   
   init;syscall
   c: MOVQ x1+16(FP), CX
      x1:   checkFloat(IsF1,x2,16,X0)
      x2:   checkFloat(IsF2,x3,24,X1)
      x3:   checkFloat(IsF3,x4,32,X2)
      x4:   checkFloat(IsF4,call,40,X3)
      call: checkCall
      grow: growCall
   retX: checkRet(88)
   retF: ret(X0,88)

TEXT ·call12(SB), NOSPLIT, $96-120
   MOVQ x2+24(FP), DX
   MOVQ x3+32(FP), R8
   MOVQ x4+40(FP), R9
   MOVQ x5+48(FP), AX; MOVQ AX, 32(SP)
   MOVQ x6+56(FP), AX; MOVQ AX, 40(SP)
   MOVQ x7+64(FP), AX; MOVQ AX, 48(SP)
   MOVQ x8+72(FP), AX; MOVQ AX, 56(SP)
   MOVQ x9+80(FP), AX; MOVQ AX, 64(SP)
   MOVQ x10+88(FP), AX; MOVQ AX, 72(SP)
   MOVQ x11+96(FP), AX; MOVQ AX, 80(SP)
   MOVQ x12+104(FP), AX; MOVQ AX, 88(SP)

   init;syscall
   c: MOVQ x1+16(FP), CX
      x1:   checkFloat(IsF1,x2,16,X0)
      x2:   checkFloat(IsF2,x3,24,X1)
      x3:   checkFloat(IsF3,x4,32,X2)
      x4:   checkFloat(IsF4,call,40,X3)
      call: checkCall
      grow: growCall
   retX: checkRet(112)
   retF: ret(X0,112)

TEXT ·call15(SB), NOSPLIT, $120-144
   MOVQ x2+24(FP), DX
   MOVQ x3+32(FP), R8
   MOVQ x4+40(FP), R9
   MOVQ x5+48(FP), AX; MOVQ AX, 32(SP)
   MOVQ x6+56(FP), AX; MOVQ AX, 40(SP)
   MOVQ x7+64(FP), AX; MOVQ AX, 48(SP)
   MOVQ x8+72(FP), AX; MOVQ AX, 56(SP)
   MOVQ x9+80(FP), AX; MOVQ AX, 64(SP)
   MOVQ x10+88(FP), AX; MOVQ AX, 72(SP)
   MOVQ x11+96(FP), AX; MOVQ AX, 80(SP)
   MOVQ x12+104(FP), AX; MOVQ AX, 88(SP)
   MOVQ x13+112(FP), AX; MOVQ AX, 96(SP)
   MOVQ x14+120(FP), AX; MOVQ AX, 104(SP)
   MOVQ x15+128(FP), AX; MOVQ AX, 112(SP)

   init;syscall
   c: MOVQ x1+16(FP), CX
      x1:   checkFloat(IsF1,x2,16,X0)
      x2:   checkFloat(IsF2,x3,24,X1)
      x3:   checkFloat(IsF3,x4,32,X2)
      x4:   checkFloat(IsF4,call,40,X3)
      call: checkCall
      grow: growCall
   retX: checkRet(136)
   retF: ret(X0,136)