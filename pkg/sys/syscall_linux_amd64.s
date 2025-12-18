#include "textflag.h"

#define IsC 	   $63	//0x8000000000000000 
#define FOut	   $62	//0x4000000000000000
#define IsF1 	   $61	//0x2000000000000000
#define IsF2 	   $60	//0x1000000000000000
#define IsF3 	   $59	//0x0800000000000000
#define IsF4 	   $58	//0x0400000000000000
#define IsF5 	   $57	//0x0200000000000000
#define IsF6 	   $56	//0x0100000000000000
#define IsF7 	   $55	//0x0080000000000000
#define IsF8 	   $54	//0x0040000000000000
#define IsF9 	   $53	//0x0020000000000000
#define IsF10 	   $52	//0x0010000000000000
#define IsF11 	   $51	//0x0008000000000000
#define IsF12 	   $50	//0x0004000000000000
#define IsF13 	   $49	//0x0002000000000000
#define IsF14 	   $48   //0x0001000000000000

#define x1 		$0	//00000001
#define x2 		$1	//00000010
#define x3 		$2	//00000100
#define x4 		$3	//00001000
#define x5 		$4	//00010000
#define x6 		$5	//00100000
#define f1 		$0	//00000001
#define f2 		$1	//00000010
#define f3 		$2	//00000100
#define f4 		$3	//00001000
#define f5 		$4	//00010000
#define f6 		$5	//00100000
#define f7 		$6	//01000000
#define f8 		$7	//10000000

// -------------------------------------------------------------------------- //

// for test
// TEXT ·call3(SB),NOSPLIT, $64-48
//    MOVQ mask+8(FP), R11
//    BTQ IsC, R11
//    JC c

//    MOVQ func+0(FP), AX
//    MOVQ x+16(FP), DI
//    MOVQ x+24(FP), SI
//    MOVQ x+32(FP), DX
//    SYSCALL
//    JMP ret

//    c:
//       MOVQ R12, X12
//       MOVQ R13, X13
//       MOVQ R14, X14
//       MOVQ R15, X15
//       LEAQ 16(SP), R12
//       LEAQ 40(SP), R13
//       LEAQ 0(SP), R10
//       MOVB $0, R14B
//       MOVB $0, R15B

//       v1: 
//          MOVQ x+16(FP), AX
//          BTQ IsF1, R11
//          JC v1f
//          MOVQ AX, (R12)
//          ADDQ $8, R12
//          SHLB $1, R14B
//          INCB R14B
//          JMP v2
//       v1f:
//          MOVQ AX, (R13)
//          ADDQ $8, R13
//          SHLB $1, R15B
//          INCB R15B
//          INCB R11B
//          JMP v2

//       v2: 
//          MOVQ x+24(FP), AX
//          BTQ IsF2, R11
//          JC v2f
//          MOVQ AX, (R12)
//          ADDQ $8, R12
//          SHLB $1, R14B
//          INCB R14B
//          JMP v3
//       v2f:
//          MOVQ AX, (R13)
//          ADDQ $8, R13
//          SHLB $1, R15B
//          INCB R15B
//          INCB R11B
//          JMP v3

//       v3: 
//          MOVQ x+32(FP), AX
//          BTQ IsF3, R11
//          JC v3f
//          MOVQ AX, (R12)
//          ADDQ $8, R12
//          SHLB $1, R14B
//          INCB R14B
//          JMP regs
//       v3f:
//          MOVQ AX, (R13)
//          ADDQ $8, R13
//          SHLB $1, R15B
//          INCB R15B
//          INCB R11B
//          JMP regs

//       regs:
//          LEAQ 16(SP), R12
//          BTW x1, R14
//          JNC regf
//          MOVQ 0(R12), DI
//          BTW x2, R14
//          JNC regf
//          MOVQ 8(R12), SI
//          BTW x3, R14
//          JNC regf
//          MOVQ 16(R12), DX
//       regf:
//          LEAQ 40(SP), R13
//          BTW f1, R15
//          JNC call
//          MOVQ 0(R13), X0
//          BTW f2, R15
//          JNC call
//          MOVQ 8(R13), X1
//          BTW f3, R15
//          JNC call
//          MOVQ 16(R13), X2

//       call:
//          MOVB R11B, AX
//          MOVQ X12, R12
//          MOVQ X13, R13
//          MOVQ X14, R14
//          MOVQ X15, R15
//          MOVQ func+0(FP), R10
//          TESTQ $0x0F, SP
//          JNZ grow
//          CALL R10
//          JMP ret
//       grow:
//          MOVQ SP, BP
//          SUBQ $8, SP
//          CALL R10
//          MOVQ BP, SP      

//    ret: 
//          BTQ FOut, mask+8(FP)
//          JC retf
//          MOVQ AX, r+40(FP)
//          RET
//    retf: MOVQ X0, r+40(FP)
//          RET

// -------------------------------------------------------------------------- //

// Архитектура	Volatile регистры	Назначение
// x64 Windows	RAX, RCX, RDX, R8, R9, R10, R11, XMM0-XMM5	Можно использовать без сохранения
// x64 Linux	RAX, RDI, RSI, RDX, RCX, R8, R9, R10, R11, XMM0-XMM15	Можно использовать свободно
// 2. Non-volatile (Callee-saved) регистры - нужно сохранять
// Эти регистры должны быть сохранены если вы их изменяете:

// Архитектура	Non-volatile регистры
// x64 Windows	RBX, RBP, RDI, RSI, RSP, R12-R15, XMM6-XMM15
// x64 Linux	RBX, RBP, R12-R15

// R10- указатель на начало переменных в SP
// R11- значение маски
// R12- указатель на начало целых чисел в SP 
// R13- указатель на начало чисел с точкой в SP 
// R14- счетчик количества целых чисел параметров для регистров (только 1 байт)
// R15- счетчик количества чисел с точкой параметров для регистров (только 1 байт)
// X12- бекап R12
// X13- бекап R13
// X14- бекап R14
// X15- бекап R15

#define init\
   MOVQ mask+8(FP), R11\
   BTQ IsC, R11\
   JC c

#define syscall\
   MOVQ func+0(FP), AX\
   MOVQ x+16(FP), DI\
   MOVQ x+24(FP), SI\
   MOVQ x+32(FP), DX\
   MOVQ x+40(FP), R10\
   MOVQ x+48(FP), R8\
   MOVQ x+56(FP), R9\
   SYSCALL\
   JMP ret

#define C(shiftX,shiftF)\
   MOVQ R12, X12\
   MOVQ R13, X13\
   MOVQ R14, X14\
   MOVQ R15, X15\
   MOVB $0, R14B\
   MOVB $0, R15B\
   LEAQ shiftX(SP), R12\
   LEAQ shiftF(SP), R13\
   LEAQ 0(SP), R10\
   TESTQ $0x0F, SP\
   JZ v1\
   SUBQ $8, R10

#define v(maskF,shiftFP,lblF,lblNext)\
   MOVQ x+shiftFP(FP), AX\
   BTQ maskF, R11\
   JC lblF\
   MOVQ AX, (R12)\
   ADDQ $8, R12\
   SHLB $1, R14B\
   INCB R14B\
   JMP lblNext\
lblF:\
   MOVQ AX, (R13)\
   ADDQ $8, R13\
   SHLB $1, R15B\
   INCB R15B\
   INCB R11B\
   JMP lblNext

#define vSP(maskF,shiftFP,lblF,lblSP,lblNext)\
   MOVQ x+shiftFP(FP), AX\
   BTQ maskF, R11\
   JC lblF\
   BTW x6, R14\
   JC lblSP\
   MOVQ AX, (R12)\
   ADDQ $8, R12\
   SHLB $1, R14B\
   INCB R14B\
   JMP lblNext\
lblF:\
   BTW f8, R15\
   JC lblSP\
   MOVQ AX, (R13)\
   ADDQ $8, R13\
   SHLB $1, R15B\
   INCB R15B\
   INCB R11B\
   JMP lblNext\
lblSP:\
   MOVQ AX, (R10)\
   ADDQ $8, R10\
   JMP lblNext

#define sp(shiftFP,lblNext)\
   MOVQ x+shiftFP(FP), AX\
   MOVQ AX, (R10)\
   ADDQ $8, R10\
   JMP lblNext

#define _setReg(_shiftSP, _regSP, reg, mask, regMask, lbl)\
   BTW mask, regMask\
   JNC lbl\
   MOVQ _shiftSP(_regSP), reg

#define setRegs(shiftX,shiftF, lbl)\
   LEAQ shiftX(SP), R12\
   _setReg(0,R12,DI,x1,R14,regf)\
   _setReg(8,R12,SI,x2,R14,regf)\
   _setReg(16,R12,DX,x3,R14,regf)\
   _setReg(24,R12,CX,x4,R14,regf)\
   _setReg(32,R12,R8,x5,R14,regf)\
   _setReg(40,R12,R9,x6,R14,regf)\
regf:\
   LEAQ shiftF(SP), R13\
   _setReg(0,R13,X0,f1,R15,lbl)\
   _setReg(8,R13,X1,f2,R15,lbl)\
   _setReg(16,R13,X2,f3,R15,lbl)\
   _setReg(24,R13,X3,f4,R15,lbl)\
   _setReg(32,R13,X4,f5,R15,lbl)\
   _setReg(40,R13,X5,f6,R15,lbl)\
   _setReg(48,R13,X6,f7,R15,lbl)\
   _setReg(56,R13,X7,f8,R15,lbl)

#define checkCall\ 
   MOVB R11B, AX\
   MOVQ X12, R12\
   MOVQ X13, R13\
   MOVQ X14, R14\
   MOVQ X15, R15\
   MOVQ func+0(FP), R10\
   TESTQ $0x0F, SP\
   JNZ grow\
   CALL R10\
   JMP ret\
grow:\
   MOVQ SP, BP\
   SUBQ $8, SP\
   CALL R10\
   MOVQ BP, SP  

#define checkRet(shift)\
   BTQ FOut, mask+8(FP)\
   JC retf\
   r(AX,shift)\
retf:\
   r(X0,shift)

#define r(reg,shift)\ 
   MOVQ reg, res+shift(FP)\
   RET
   
// -------------------------------------------------------------------------- //

TEXT ·call3(SB), NOSPLIT, $64-48
   init

   MOVQ func+0(FP), AX
   MOVQ x+16(FP), DI
   MOVQ x+24(FP), SI
   MOVQ x+32(FP), DX
   SYSCALL
   JMP ret

   c: C(16,40)
      v1:   v(IsF1,16,v1f,v2) 
      v2:   v(IsF2,24,v2f,v3)
      v3:   v(IsF3,32,v3f,regs)
      regs: setRegs(16,40,call)
      call: checkCall
   ret: checkRet(40)

TEXT ·call6(SB), NOSPLIT, $144-72
   init
   syscall
   c: C(40,88)
      v1:   v(IsF1,16,v1f,v2) 
      v2:   v(IsF2,24,v2f,v3)
      v3:   v(IsF3,32,v3f,v4)
      v4:   v(IsF4,40,v4f,v5)
      v5:   v(IsF5,48,v5f,v6)
      v6:   v(IsF6,56,v6f,regs)
      regs: setRegs(40,88,call)
      call: checkCall
   ret: checkRet(64)

TEXT ·call9(SB), NOSPLIT, $176-96
   init
   syscall
   c: C(64,112)
      v1:   v(IsF1,16,v1f,v2)
      v2:   v(IsF2,24,v2f,v3)
      v3:   v(IsF3,32,v3f,v4)
      v4:   v(IsF4,40,v4f,v5)
      v5:   v(IsF5,48,v5f,v6)
      v6:   v(IsF6,56,v6f,v7)
      v7:   vSP(IsF7,64,v7f,v7sp,v8)
      v8:   vSP(IsF8,72,v8f,v8sp,v9)
      v9:   vSP(IsF9,80,v9f,v9sp,regs)
      regs: setRegs(64,112,call)         
      call: checkCall
   ret: checkRet(88)

TEXT ·call12(SB), NOSPLIT, $208-120
   init
   syscall
   c: C(88,136)
      v1:   v(IsF1,16,v1f,v2)
      v2:   v(IsF2,24,v2f,v3)
      v3:   v(IsF3,32,v3f,v4)
      v4:   v(IsF4,40,v4f,v5)
      v5:   v(IsF5,48,v5f,v6)
      v6:   v(IsF6,56,v6f,v7)
      v7:   vSP(IsF7,64,v7f,v7sp,v8)
      v8:   vSP(IsF8,72,v8f,v8sp,v9)
      v9:   vSP(IsF9,80,v9f,v9sp,v10)
      v10:  vSP(IsF10,88,v10f,v10sp,v11)
      v11:  vSP(IsF11,96,v11f,v11sp,v12)
      v12:  vSP(IsF12,104,v12f,v12sp,regs)
      regs: setRegs(88,136,call)
      call: checkCall
   ret: checkRet(112)

TEXT ·call15(SB), NOSPLIT, $224-144
   init
   syscall
   c: C(112,160)
      v1:   v(IsF1,16,v1f,v2)
      v2:   v(IsF2,24,v2f,v3)
      v3:   v(IsF3,32,v3f,v4)
      v4:   v(IsF4,40,v4f,v5)
      v5:   v(IsF5,48,v5f,v6)
      v6:   v(IsF6,56,v6f,v7)
      v7:   vSP(IsF7,64,v7f,v7sp,v8)
      v8:   vSP(IsF8,72,v8f,v8sp,v9)
      v9:   vSP(IsF9,80,v9f,v9sp,v10)
      v10:  vSP(IsF10,88,v10f,v10sp,v11)
      v11:  vSP(IsF11,96,v11f,v11sp,v12)
      v12:  vSP(IsF12,104,v12f,v12sp,v13)
      v13:  vSP(IsF13,112,v13f,v13sp,v14)
      v14:  vSP(IsF14,120,v14f,v14sp,v15)
      v15:  sp(128,regs)
      regs: setRegs(112,160,call)
      call: checkCall
   ret: checkRet(136)
