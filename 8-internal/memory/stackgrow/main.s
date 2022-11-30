"".main STEXT size=103 args=0x0 locals=0x28 funcid=0x0 align=0x0
	0x0000 00000 (main.go:3)	TEXT	"".main(SB), ABIInternal, $40-0
	0x0000 00000 (main.go:3)	CMPQ	SP, 16(R14)
	0x0004 00004 (main.go:3)	PCDATA	$0, $-2
	0x0004 00004 (main.go:3)	JLS	94
	0x0006 00006 (main.go:3)	PCDATA	$0, $-1
	0x0006 00006 (main.go:3)	SUBQ	$40, SP
	0x000a 00010 (main.go:3)	MOVQ	BP, 32(SP)
	0x000f 00015 (main.go:3)	LEAQ	32(SP), BP
	0x0014 00020 (main.go:3)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0014 00020 (main.go:3)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0014 00020 (main.go:4)	MOVQ	$1, "".a+24(SP)
	0x001d 00029 (main.go:4)	MOVQ	$2, "".b+16(SP)
	0x0026 00038 (main.go:5)	MOVQ	"".a+24(SP), AX
	0x002b 00043 (main.go:5)	MOVL	$2, BX
	0x0030 00048 (main.go:5)	PCDATA	$1, $0
	0x0030 00048 (main.go:5)	CALL	"".add1(SB)
	0x0035 00053 (main.go:6)	MOVQ	"".b+16(SP), BX
	0x003a 00058 (main.go:6)	MOVQ	"".a+24(SP), AX
	0x003f 00063 (main.go:6)	NOP
	0x0040 00064 (main.go:6)	CALL	"".add2(SB)
	0x0045 00069 (main.go:7)	MOVQ	"".b+16(SP), BX
	0x004a 00074 (main.go:7)	MOVQ	"".a+24(SP), AX
	0x004f 00079 (main.go:7)	CALL	"".add3(SB)
	0x0054 00084 (main.go:8)	MOVQ	32(SP), BP
	0x0059 00089 (main.go:8)	ADDQ	$40, SP
	0x005d 00093 (main.go:8)	RET
	0x005e 00094 (main.go:8)	NOP
	0x005e 00094 (main.go:3)	PCDATA	$1, $-1
	0x005e 00094 (main.go:3)	PCDATA	$0, $-2
	0x005e 00094 (main.go:3)	NOP
	0x0060 00096 (main.go:3)	CALL	runtime.morestack_noctxt(SB)
	0x0065 00101 (main.go:3)	PCDATA	$0, $-1
	0x0065 00101 (main.go:3)	JMP	0
	0x0000 49 3b 66 10 76 58 48 83 ec 28 48 89 6c 24 20 48  I;f.vXH..(H.l$ H
	0x0010 8d 6c 24 20 48 c7 44 24 18 01 00 00 00 48 c7 44  .l$ H.D$.....H.D
	0x0020 24 10 02 00 00 00 48 8b 44 24 18 bb 02 00 00 00  $.....H.D$......
	0x0030 e8 00 00 00 00 48 8b 5c 24 10 48 8b 44 24 18 90  .....H.\$.H.D$..
	0x0040 e8 00 00 00 00 48 8b 5c 24 10 48 8b 44 24 18 e8  .....H.\$.H.D$..
	0x0050 00 00 00 00 48 8b 6c 24 20 48 83 c4 28 c3 66 90  ....H.l$ H..(.f.
	0x0060 e8 00 00 00 00 eb 99                             .......
	rel 49+4 t=7 "".add1+0
	rel 65+4 t=7 "".add2+0
	rel 80+4 t=7 "".add3+0
	rel 97+4 t=7 runtime.morestack_noctxt+0
"".add1 STEXT nosplit size=77 args=0x10 locals=0x28 funcid=0x0 align=0x0
	0x0000 00000 (main.go:10)	TEXT	"".add1(SB), NOSPLIT|ABIInternal, $40-16
	0x0000 00000 (main.go:10)	SUBQ	$40, SP
	0x0004 00004 (main.go:10)	MOVQ	BP, 32(SP)
	0x0009 00009 (main.go:10)	LEAQ	32(SP), BP
	0x000e 00014 (main.go:10)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x000e 00014 (main.go:10)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x000e 00014 (main.go:10)	FUNCDATA	$5, "".add1.arginfo1(SB)
	0x000e 00014 (main.go:10)	MOVQ	AX, "".x+48(SP)
	0x0013 00019 (main.go:10)	MOVQ	BX, "".y+56(SP)
	0x0018 00024 (main.go:10)	MOVQ	$0, "".~r0(SP)
	0x0020 00032 (main.go:11)	MOVUPS	X15, ""..autotmp_3+12(SP)
	0x0026 00038 (main.go:11)	MOVUPS	X15, ""..autotmp_3+16(SP)
	0x002c 00044 (main.go:11)	LEAQ	""..autotmp_3+12(SP), CX
	0x0031 00049 (main.go:11)	TESTB	AL, (CX)
	0x0033 00051 (main.go:11)	JMP	53
	0x0035 00053 (main.go:12)	MOVQ	"".x+48(SP), AX
	0x003a 00058 (main.go:12)	ADDQ	"".y+56(SP), AX
	0x003f 00063 (main.go:12)	MOVQ	AX, "".~r0(SP)
	0x0043 00067 (main.go:12)	MOVQ	32(SP), BP
	0x0048 00072 (main.go:12)	ADDQ	$40, SP
	0x004c 00076 (main.go:12)	RET
	0x0000 48 83 ec 28 48 89 6c 24 20 48 8d 6c 24 20 48 89  H..(H.l$ H.l$ H.
	0x0010 44 24 30 48 89 5c 24 38 48 c7 04 24 00 00 00 00  D$0H.\$8H..$....
	0x0020 44 0f 11 7c 24 0c 44 0f 11 7c 24 10 48 8d 4c 24  D..|$.D..|$.H.L$
	0x0030 0c 84 01 eb 00 48 8b 44 24 30 48 03 44 24 38 48  .....H.D$0H.D$8H
	0x0040 89 04 24 48 8b 6c 24 20 48 83 c4 28 c3           ..$H.l$ H..(.
"".add2 STEXT size=165 args=0x10 locals=0xd8 funcid=0x0 align=0x0
	0x0000 00000 (main.go:15)	TEXT	"".add2(SB), ABIInternal, $216-16
	0x0000 00000 (main.go:15)	LEAQ	-88(SP), R12
	0x0005 00005 (main.go:15)	CMPQ	R12, 16(R14)
	0x0009 00009 (main.go:15)	PCDATA	$0, $-2
	0x0009 00009 (main.go:15)	JLS	134
	0x000b 00011 (main.go:15)	PCDATA	$0, $-1
	0x000b 00011 (main.go:15)	SUBQ	$216, SP
	0x0012 00018 (main.go:15)	MOVQ	BP, 208(SP)
	0x001a 00026 (main.go:15)	LEAQ	208(SP), BP
	0x0022 00034 (main.go:15)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0022 00034 (main.go:15)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0022 00034 (main.go:15)	FUNCDATA	$5, "".add2.arginfo1(SB)
	0x0022 00034 (main.go:15)	MOVQ	AX, "".x+224(SP)
	0x002a 00042 (main.go:15)	MOVQ	BX, "".y+232(SP)
	0x0032 00050 (main.go:15)	MOVQ	$0, "".~r0(SP)
	0x003a 00058 (main.go:16)	MOVUPS	X15, ""..autotmp_3+8(SP)
	0x0040 00064 (main.go:16)	LEAQ	""..autotmp_3+16(SP), DI
	0x0045 00069 (main.go:16)	PCDATA	$0, $-2
	0x0045 00069 (main.go:16)	DUFFZERO	$299
	0x0058 00088 (main.go:16)	PCDATA	$0, $-1
	0x0058 00088 (main.go:16)	LEAQ	""..autotmp_3+8(SP), CX
	0x005d 00093 (main.go:16)	TESTB	AL, (CX)
	0x005f 00095 (main.go:16)	NOP
	0x0060 00096 (main.go:16)	JMP	98
	0x0062 00098 (main.go:17)	MOVQ	"".x+224(SP), AX
	0x006a 00106 (main.go:17)	ADDQ	"".y+232(SP), AX
	0x0072 00114 (main.go:17)	MOVQ	AX, "".~r0(SP)
	0x0076 00118 (main.go:17)	MOVQ	208(SP), BP
	0x007e 00126 (main.go:17)	ADDQ	$216, SP
	0x0085 00133 (main.go:17)	RET
	0x0086 00134 (main.go:17)	NOP
	0x0086 00134 (main.go:15)	PCDATA	$1, $-1
	0x0086 00134 (main.go:15)	PCDATA	$0, $-2
	0x0086 00134 (main.go:15)	MOVQ	AX, 8(SP)
	0x008b 00139 (main.go:15)	MOVQ	BX, 16(SP)
	0x0090 00144 (main.go:15)	CALL	runtime.morestack_noctxt(SB)
	0x0095 00149 (main.go:15)	MOVQ	8(SP), AX
	0x009a 00154 (main.go:15)	MOVQ	16(SP), BX
	0x009f 00159 (main.go:15)	PCDATA	$0, $-1
	0x009f 00159 (main.go:15)	NOP
	0x00a0 00160 (main.go:15)	JMP	0
	0x0000 4c 8d 64 24 a8 4d 3b 66 10 76 7b 48 81 ec d8 00  L.d$.M;f.v{H....
	0x0010 00 00 48 89 ac 24 d0 00 00 00 48 8d ac 24 d0 00  ..H..$....H..$..
	0x0020 00 00 48 89 84 24 e0 00 00 00 48 89 9c 24 e8 00  ..H..$....H..$..
	0x0030 00 00 48 c7 04 24 00 00 00 00 44 0f 11 7c 24 08  ..H..$....D..|$.
	0x0040 48 8d 7c 24 10 48 89 6c 24 f0 48 8d 6c 24 f0 e8  H.|$.H.l$.H.l$..
	0x0050 00 00 00 00 48 8b 6d 00 48 8d 4c 24 08 84 01 90  ....H.m.H.L$....
	0x0060 eb 00 48 8b 84 24 e0 00 00 00 48 03 84 24 e8 00  ..H..$....H..$..
	0x0070 00 00 48 89 04 24 48 8b ac 24 d0 00 00 00 48 81  ..H..$H..$....H.
	0x0080 c4 d8 00 00 00 c3 48 89 44 24 08 48 89 5c 24 10  ......H.D$.H.\$.
	0x0090 e8 00 00 00 00 48 8b 44 24 08 48 8b 5c 24 10 90  .....H.D$.H.\$..
	0x00a0 e9 5b ff ff ff                                   .[...
	rel 80+4 t=7 runtime.duffzero+299
	rel 145+4 t=7 runtime.morestack_noctxt+0
"".add3 STEXT size=165 args=0x10 locals=0x1398 funcid=0x0 align=0x0
	0x0000 00000 (main.go:20)	TEXT	"".add3(SB), ABIInternal, $5016-16
	0x0000 00000 (main.go:20)	MOVQ	SP, R12
	0x0003 00003 (main.go:20)	PCDATA	$0, $-2
	0x0003 00003 (main.go:20)	SUBQ	$4888, R12
	0x000a 00010 (main.go:20)	JCS	131
	0x000c 00012 (main.go:20)	CMPQ	R12, 16(R14)
	0x0010 00016 (main.go:20)	JLS	131
	0x0012 00018 (main.go:20)	PCDATA	$0, $-1
	0x0012 00018 (main.go:20)	SUBQ	$5016, SP
	0x0019 00025 (main.go:20)	MOVQ	BP, 5008(SP)
	0x0021 00033 (main.go:20)	LEAQ	5008(SP), BP
	0x0029 00041 (main.go:20)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0029 00041 (main.go:20)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0029 00041 (main.go:20)	FUNCDATA	$5, "".add3.arginfo1(SB)
	0x0029 00041 (main.go:20)	MOVQ	AX, "".x+5024(SP)
	0x0031 00049 (main.go:20)	MOVQ	BX, "".y+5032(SP)
	0x0039 00057 (main.go:20)	MOVQ	$0, "".~r0(SP)
	0x0041 00065 (main.go:21)	MOVUPS	X15, ""..autotmp_3+8(SP)
	0x0047 00071 (main.go:21)	LEAQ	""..autotmp_3+16(SP), DI
	0x004c 00076 (main.go:21)	MOVL	$624, CX
	0x0051 00081 (main.go:21)	XORL	AX, AX
	0x0053 00083 (main.go:21)	REP
	0x0054 00084 (main.go:21)	STOSQ
	0x0056 00086 (main.go:21)	LEAQ	""..autotmp_3+8(SP), DX
	0x005b 00091 (main.go:21)	TESTB	AL, (DX)
	0x005d 00093 (main.go:21)	JMP	95
	0x005f 00095 (main.go:22)	MOVQ	"".x+5024(SP), AX
	0x0067 00103 (main.go:22)	ADDQ	"".y+5032(SP), AX
	0x006f 00111 (main.go:22)	MOVQ	AX, "".~r0(SP)
	0x0073 00115 (main.go:22)	MOVQ	5008(SP), BP
	0x007b 00123 (main.go:22)	ADDQ	$5016, SP
	0x0082 00130 (main.go:22)	RET
	0x0083 00131 (main.go:22)	NOP
	0x0083 00131 (main.go:20)	PCDATA	$1, $-1
	0x0083 00131 (main.go:20)	PCDATA	$0, $-2
	0x0083 00131 (main.go:20)	MOVQ	AX, 8(SP)
	0x0088 00136 (main.go:20)	MOVQ	BX, 16(SP)
	0x008d 00141 (main.go:20)	CALL	runtime.morestack_noctxt(SB)
	0x0092 00146 (main.go:20)	MOVQ	8(SP), AX
	0x0097 00151 (main.go:20)	MOVQ	16(SP), BX
	0x009c 00156 (main.go:20)	PCDATA	$0, $-1
	0x009c 00156 (main.go:20)	NOP
	0x00a0 00160 (main.go:20)	JMP	0
	0x0000 49 89 e4 49 81 ec 18 13 00 00 72 77 4d 3b 66 10  I..I......rwM;f.
	0x0010 76 71 48 81 ec 98 13 00 00 48 89 ac 24 90 13 00  vqH......H..$...
	0x0020 00 48 8d ac 24 90 13 00 00 48 89 84 24 a0 13 00  .H..$....H..$...
	0x0030 00 48 89 9c 24 a8 13 00 00 48 c7 04 24 00 00 00  .H..$....H..$...
	0x0040 00 44 0f 11 7c 24 08 48 8d 7c 24 10 b9 70 02 00  .D..|$.H.|$..p..
	0x0050 00 31 c0 f3 48 ab 48 8d 54 24 08 84 02 eb 00 48  .1..H.H.T$.....H
	0x0060 8b 84 24 a0 13 00 00 48 03 84 24 a8 13 00 00 48  ..$....H..$....H
	0x0070 89 04 24 48 8b ac 24 90 13 00 00 48 81 c4 98 13  ..$H..$....H....
	0x0080 00 00 c3 48 89 44 24 08 48 89 5c 24 10 e8 00 00  ...H.D$.H.\$....
	0x0090 00 00 48 8b 44 24 08 48 8b 5c 24 10 0f 1f 40 00  ..H.D$.H.\$...@.
	0x00a0 e9 5b ff ff ff                                   .[...
	rel 142+4 t=7 runtime.morestack_noctxt+0
go.cuinfo.packagename. SDWARFCUINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
""..inittask SNOPTRDATA size=24
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 00 00 00                          ........
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
"".add1.arginfo1 SRODATA static dupok size=5
	0x0000 00 08 08 08 ff                                   .....
"".add2.arginfo1 SRODATA static dupok size=5
	0x0000 00 08 08 08 ff                                   .....
"".add3.arginfo1 SRODATA static dupok size=5
	0x0000 00 08 08 08 ff                                   .....
