golang中使用的汇编是

将测试代码编译成汇编
```shell
# go tool compile -S -N -l main.go > main.s
# GOOS=linux GOARCH=amd64 go tool compile -S -N -l main.go > main.s
```

可以看到编译后的汇编有对 runtime.morestack_noctxt 的调用
```shell
//file:main.S
# runtime.morestack_noctxt(SB)
```

其中 morestack_noctxt 是在 runtime 包中定义的一个汇编函数
```
//file:asm_amd64.s
TEXT runtime·morestack_noctxt(SB),NOSPLIT|NOFRAME,$0-0
	MOVD	RSP, RSP
	MOVW	$0, R26
	B runtime·morestack(SB)
	
TEXT runtime·morestack(SB),NOSPLIT|NOFRAME,$0-0
    ......
    // 在 m->g0 栈上调用 newstack.
    MOVQ	m_g0(BX), BX
	MOVQ	BX, g(CX)
	MOVQ	(g_sched+gobuf_sp)(BX), SP
	CALL	runtime·newstack(SB)
	CALL	runtime·abort(SB)	// crash if newstack returns
	RET
```

runtime·morestack 做完校验和赋值操作后会切换到 G0 调用 runtime·newstack来完成扩容的操作

newstack 就定义在 go 代码中了

```go
//file:runtime/stack.go
func newstack() {
	......
	//新栈大小是老栈的两倍
    oldsize := gp.stack.hi - gp.stack.lo
    newsize := oldsize * 2
	
	//进行栈的申请和拷贝
    copystack(gp, newsize)
}
```

在 copystack 中完成新栈的分配，将旧栈拷贝到新栈，

```go
//file:runtime/stack.go
func copystack(gp *g, newsize uintptr) {
	//分配新的栈空间
	new := stackalloc(uint32(newsize))

	// 将原栈中的内存拷贝到新栈中
	memmove(unsafe.Pointer(new.hi-ncopy), unsafe.Pointer(old.hi-ncopy), ncopy)

	// 将 G 上的栈引用切换成新栈
	gp.stack = new
	gp.stackguard0 = new.lo + _StackGuard // NOTE: might clobber a preempt request
	gp.sched.sp = new.hi - used
	gp.stktopsp += adjinfo.delta

	//释放原栈内存
	stackfree(old)
}
```

至此，栈扩张完毕！