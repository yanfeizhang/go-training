"".main STEXT size=71 args=0x0 locals=0x18 funcid=0x0 align=0x0
	0x0000 00000 (main.go:3)	TEXT	"".main(SB), ABIInternal, $24-0
	0x0000 00000 (main.go:3)	CMPQ	SP, 16(R14)
	0x0004 00004 (main.go:3)	PCDATA	$0, $-2
	0x0004 00004 (main.go:3)	JLS	59
	0x0006 00006 (main.go:3)	PCDATA	$0, $-1
	0x0006 00006 (main.go:3)	SUBQ	$24, SP
	0x000a 00010 (main.go:3)	MOVQ	BP, 16(SP)
	0x000f 00015 (main.go:3)	LEAQ	16(SP), BP
	0x0014 00020 (main.go:3)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0014 00020 (main.go:3)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0014 00020 (main.go:4)	MOVQ	$1, "".n+8(SP)
	0x001d 00029 (main.go:5)	MOVL	$1, AX
	0x0022 00034 (main.go:5)	PCDATA	$1, $0
	0x0022 00034 (main.go:5)	CALL	"".func1(SB)
	0x0027 00039 (main.go:6)	MOVQ	"".n+8(SP), AX
	0x002c 00044 (main.go:6)	CALL	"".func2(SB)
	0x0031 00049 (main.go:7)	MOVQ	16(SP), BP
	0x0036 00054 (main.go:7)	ADDQ	$24, SP
	0x003a 00058 (main.go:7)	RET
	0x003b 00059 (main.go:7)	NOP
	0x003b 00059 (main.go:3)	PCDATA	$1, $-1
	0x003b 00059 (main.go:3)	PCDATA	$0, $-2
	0x003b 00059 (main.go:3)	NOP
	0x0040 00064 (main.go:3)	CALL	runtime.morestack_noctxt(SB)
	0x0045 00069 (main.go:3)	PCDATA	$0, $-1
	0x0045 00069 (main.go:3)	JMP	0
	0x0000 49 3b 66 10 76 35 48 83 ec 18 48 89 6c 24 10 48  I;f.v5H...H.l$.H
	0x0010 8d 6c 24 10 48 c7 44 24 08 01 00 00 00 b8 01 00  .l$.H.D$........
	0x0020 00 00 e8 00 00 00 00 48 8b 44 24 08 e8 00 00 00  .......H.D$.....
	0x0030 00 48 8b 6c 24 10 48 83 c4 18 c3 0f 1f 44 00 00  .H.l$.H......D..
	0x0040 e8 00 00 00 00 eb b9                             .......
	rel 35+4 t=7 "".func1+0
	rel 45+4 t=7 "".func2+0
	rel 65+4 t=7 runtime.morestack_noctxt+0
"".func1 STEXT size=143 args=0x8 locals=0xd8 funcid=0x0 align=0x0
	0x0000 00000 (main.go:9)	TEXT	"".func1(SB), ABIInternal, $216-8
	0x0000 00000 (main.go:9)	LEAQ	-88(SP), R12
	0x0005 00005 (main.go:9)	CMPQ	R12, 16(R14)
	0x0009 00009 (main.go:9)	PCDATA	$0, $-2
	0x0009 00009 (main.go:9)	JLS	120
	0x000b 00011 (main.go:9)	PCDATA	$0, $-1
	0x000b 00011 (main.go:9)	SUBQ	$216, SP
	0x0012 00018 (main.go:9)	MOVQ	BP, 208(SP)
	0x001a 00026 (main.go:9)	LEAQ	208(SP), BP
	0x0022 00034 (main.go:9)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0022 00034 (main.go:9)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0022 00034 (main.go:9)	FUNCDATA	$5, "".func1.arginfo1(SB)
	0x0022 00034 (main.go:9)	MOVQ	AX, "".n+224(SP)
	0x002a 00042 (main.go:9)	MOVQ	$0, "".~r0(SP)
	0x0032 00050 (main.go:10)	MOVUPS	X15, ""..autotmp_2+8(SP)
	0x0038 00056 (main.go:10)	LEAQ	""..autotmp_2+16(SP), DI
	0x003d 00061 (main.go:10)	PCDATA	$0, $-2
	0x003d 00061 (main.go:10)	NOP
	0x0040 00064 (main.go:10)	DUFFZERO	$299
	0x0053 00083 (main.go:10)	PCDATA	$0, $-1
	0x0053 00083 (main.go:10)	LEAQ	""..autotmp_2+8(SP), CX
	0x0058 00088 (main.go:10)	TESTB	AL, (CX)
	0x005a 00090 (main.go:10)	JMP	92
	0x005c 00092 (main.go:11)	MOVQ	"".n+224(SP), AX
	0x0064 00100 (main.go:11)	MOVQ	AX, "".~r0(SP)
	0x0068 00104 (main.go:11)	MOVQ	208(SP), BP
	0x0070 00112 (main.go:11)	ADDQ	$216, SP
	0x0077 00119 (main.go:11)	RET
	0x0078 00120 (main.go:11)	NOP
	0x0078 00120 (main.go:9)	PCDATA	$1, $-1
	0x0078 00120 (main.go:9)	PCDATA	$0, $-2
	0x0078 00120 (main.go:9)	MOVQ	AX, 8(SP)
	0x007d 00125 (main.go:9)	NOP
	0x0080 00128 (main.go:9)	CALL	runtime.morestack_noctxt(SB)
	0x0085 00133 (main.go:9)	MOVQ	8(SP), AX
	0x008a 00138 (main.go:9)	PCDATA	$0, $-1
	0x008a 00138 (main.go:9)	JMP	0
	0x0000 4c 8d 64 24 a8 4d 3b 66 10 76 6d 48 81 ec d8 00  L.d$.M;f.vmH....
	0x0010 00 00 48 89 ac 24 d0 00 00 00 48 8d ac 24 d0 00  ..H..$....H..$..
	0x0020 00 00 48 89 84 24 e0 00 00 00 48 c7 04 24 00 00  ..H..$....H..$..
	0x0030 00 00 44 0f 11 7c 24 08 48 8d 7c 24 10 0f 1f 00  ..D..|$.H.|$....
	0x0040 48 89 6c 24 f0 48 8d 6c 24 f0 e8 00 00 00 00 48  H.l$.H.l$......H
	0x0050 8b 6d 00 48 8d 4c 24 08 84 01 eb 00 48 8b 84 24  .m.H.L$.....H..$
	0x0060 e0 00 00 00 48 89 04 24 48 8b ac 24 d0 00 00 00  ....H..$H..$....
	0x0070 48 81 c4 d8 00 00 00 c3 48 89 44 24 08 0f 1f 00  H.......H.D$....
	0x0080 e8 00 00 00 00 48 8b 44 24 08 e9 71 ff ff ff     .....H.D$..q...
	rel 75+4 t=7 runtime.duffzero+299
	rel 129+4 t=7 runtime.morestack_noctxt+0
"".func2 STEXT nosplit size=67 args=0x8 locals=0x28 funcid=0x0 align=0x0
	0x0000 00000 (main.go:14)	TEXT	"".func2(SB), NOSPLIT|ABIInternal, $40-8
	0x0000 00000 (main.go:14)	SUBQ	$40, SP
	0x0004 00004 (main.go:14)	MOVQ	BP, 32(SP)
	0x0009 00009 (main.go:14)	LEAQ	32(SP), BP
	0x000e 00014 (main.go:14)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x000e 00014 (main.go:14)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x000e 00014 (main.go:14)	FUNCDATA	$5, "".func2.arginfo1(SB)
	0x000e 00014 (main.go:14)	MOVQ	AX, "".n+48(SP)
	0x0013 00019 (main.go:14)	MOVQ	$0, "".~r0(SP)
	0x001b 00027 (main.go:15)	MOVUPS	X15, ""..autotmp_2+12(SP)
	0x0021 00033 (main.go:15)	MOVUPS	X15, ""..autotmp_2+16(SP)
	0x0027 00039 (main.go:15)	LEAQ	""..autotmp_2+12(SP), CX
	0x002c 00044 (main.go:15)	TESTB	AL, (CX)
	0x002e 00046 (main.go:15)	JMP	48
	0x0030 00048 (main.go:16)	MOVQ	"".n+48(SP), AX
	0x0035 00053 (main.go:16)	MOVQ	AX, "".~r0(SP)
	0x0039 00057 (main.go:16)	MOVQ	32(SP), BP
	0x003e 00062 (main.go:16)	ADDQ	$40, SP
	0x0042 00066 (main.go:16)	RET
	0x0000 48 83 ec 28 48 89 6c 24 20 48 8d 6c 24 20 48 89  H..(H.l$ H.l$ H.
	0x0010 44 24 30 48 c7 04 24 00 00 00 00 44 0f 11 7c 24  D$0H..$....D..|$
	0x0020 0c 44 0f 11 7c 24 10 48 8d 4c 24 0c 84 01 eb 00  .D..|$.H.L$.....
	0x0030 48 8b 44 24 30 48 89 04 24 48 8b 6c 24 20 48 83  H.D$0H..$H.l$ H.
	0x0040 c4 28 c3                                         .(.
go.cuinfo.packagename. SDWARFCUINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
""..inittask SNOPTRDATA size=24
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 00 00 00                          ........
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
"".func1.arginfo1 SRODATA static dupok size=3
	0x0000 00 08 ff                                         ...
"".func2.arginfo1 SRODATA static dupok size=3
	0x0000 00 08 ff                                         ...
