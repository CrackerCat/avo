// Code generated by command: go run asm.go -out gp8.s -stubs stub.go. DO NOT EDIT.

#include "textflag.h"

// func GP8() uint8
TEXT ·GP8(SB), NOSPLIT, $0-1
	MOVB $0x01, AL
	MOVB $0x02, CL
	MOVB $0x03, DL
	MOVB $0x04, BL
	MOVB $0x05, BP
	MOVB $0x06, SI
	MOVB $0x07, DI
	MOVB $0x08, R8
	MOVB $0x09, R9
	MOVB $0x0a, R10
	MOVB $0x0b, R11
	MOVB $0x0c, R12
	MOVB $0x0d, R13
	MOVB $0x0e, R14
	MOVB $0x0f, R15
	MOVB $0x10, AH
	MOVB $0x11, CH
	MOVB $0x12, DH
	MOVB $0x13, BH
	ADDB CL, AL
	ADDB DL, AL
	ADDB BL, AL
	ADDB BP, AL
	ADDB SI, AL
	ADDB DI, AL
	ADDB R8, AL
	ADDB R9, AL
	ADDB R10, AL
	ADDB R11, AL
	ADDB R12, AL
	ADDB R13, AL
	ADDB R14, AL
	ADDB R15, AL
	ADDB AH, AL
	ADDB CH, AL
	ADDB DH, AL
	ADDB BH, AL
	MOVB AL, ret+0(FP)
	RET
