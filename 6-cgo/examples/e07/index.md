我们通过分析一个简单的 cgo 函数调用来了解 cgo 的调用过程。

如下是一个简单的 cgo 函数调用，在Go语言中调用了C语言实现的函数 println。
```go
package main

//int sum(int a, int b) { return a+b; }
import "C"

func main() {
    println(C.sum(1, 1))
}
```

通过 cgo 命令行工具在_obj 目录生成中间文件

```
# go tool cgo main.go
```

## C.sum函数的实现
在生成的main.cgo1.go文件中，可以看到C.sum函数的实现。
```
package main

//int sum(int a, int b) { return a+b; }
import _ "unsafe"

func main() {
    println((_Cfunc_sum)(1, 1))
}
```

可以看到，C.sum函数被转换成了一个名为_Cfunc_sum的Go语言函数。
_Cfunc_sum 函数（这是一个go函数，）在 cgo 生成的 _cgo_gotypes.go 文件中定义。

```shell
//go:cgo_unsafe_args
func _Cfunc_sum(p0 _Ctype_int, p1 _Ctype_int) (r1 _Ctype_int) {
	_cgo_runtime_cgocall(_cgo_e119c51a7968_Cfunc_sum, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
		_Cgo_use(p1)
	}
	return
}
```
_Cfunc_sum 是C 函数 sum 在 Go 空间的入口。
在Cfunc_sum中，通过_cgo_runtime_cgocall 函数再间接调用 C 函数 sum。
它的参数 p0，p1 通过_Cgo_use 逃逸到了堆上。


## runtime.cgocall 函数
runtime.cgocall 函数是实现 Go 语言到 C 语言函数跨界调用的关键。其中传入的参数中，_cgo_e119c51a7968_Cfunc_sum是一个C语言实现的函数。

其中 cgocall 的源码位于 go 语言运行时的 runtime/cgocall.go。工作包括
- 做一些调度相关的准备动作
- 进行Go与C之间call ABI操作

```go
//file：runtime/cgocall.go
func cgocall(fn, arg unsafe.Pointer) int32 {
	......
	
    mp := getg().m // 获取当前 goroutine 的 M
    mp.ncgocall++  // 总 cgo 计数 +1
    mp.ncgo++      // 当前 cgo 计数 +1
    
    mp.cgoCallers[0] = 0 // 重置追踪
    
    entersyscall() // 进入系统调用,保存上下文, 标记当前 goroutine 独占 m, 跳过垃圾回收
    
    osPreemptExtEnter(mp) // 标记异步抢占, 使异步抢占逻辑失效
    
    mp.incgo = true // 修改状态
    errno := asmcgocall(fn, arg) // 真正进行方法调用的地方
    
    mp.incgo = false // 修改状态
    mp.ncgo-- // 当前 cgo 调用-1
    
    osPreemptExtExit(mp) // 恢复异步抢占
    
    exitsyscall() // 退出系统调用,恢复调度器控制
    ......
	
    // 避免 GC 过早回收
    KeepAlive(fn)
    KeepAlive(arg)
    KeepAlive(mp)
    
    return errno
}
```

在上面的源码中，有几个重要的函数
- entersyscall 函数：将当前的 M 与 P 剥离，防止 C 程序独占 M 时，阻塞 P 的调度
- asmcgocall：将栈切换到 g0 的系统栈，并执行 C 函数调用
- exitsyscall：寻找合适的 P 来运行从 C 函数返回的 Go 程，优先选择调用 C 之前依附的 P，其次选择其他空闲的 P

值得注意的是当 Go 程在调用 C 函数时，会单独占用一个系统线程。因此如果在 Go 程中并发调用 C 函数。
如果 C 函数中又存在阻塞操作，就很可能会造成 Go 程序不停的创建新的系统线程，而 Go 并不会回收系统线程，过多的线程数会拖垮整个系统


### entersyscall
我们再来看 entersyscall 函数。该函数将M与P剥离，防止系统调用阻塞P的调度，保存上下文。

```go
//file:runtime/proc.go
func entersyscall() {
	reentersyscall(getcallerpc(), getcallersp())
}
func reentersyscall(pc, sp uintptr) {
    _g_ := getg()
    ......
	
	//保存g的现场信息，rsp, rbp, rip等
	save(pc, sp)
    _g_.syscallsp = sp
    _g_.syscallpc = pc
    casgstatus(_g_, _Grunning, _Gsyscall)
	...
    //解除P与M的绑定
    pp := _g_.m.p.ptr()
    pp.m = 0
    _g_.m.oldp.set(pp) //把p记录在oldp中，等从系统调用返回时，优先绑定这个p
    _g_.m.p = 0
}
```

entersyscall 直接调用了reentersyscall函数，reentersyscall首先把现场信息保存在当前g的sched成员中，然后解除m和p的绑定关系并设置p的状态为_Psyscall.

sysmon监控线程需要依赖该状态实施抢占, sysmon线程通过 retake => handoffp
- 如果调用不超过20us则不会触发任何事件。
- 如果调用超过20us可能会导致新线程的启动 

```go
//file:runtime/proc.go
func retake(now int64) uint32 {
    for i := 0; i < len(allp); i++ {
        _p_ := allp[i]
    }
    ......

	if s == _Psyscall { 
		...
		handoffp(_p_)
    }
}
```

handoffp 方法会调用 startm 来启动一个新的 M,出来接管P。

```go
//file:runtime/proc.go
func handoffp(_p_ *p) {
	...
	startm(_p_, false)
}
```

### asmcgocall
asmcgocall 是一个汇编函数，用于调用 C 函数。
将当前栈移到系统栈去执行，因为 C 需要"无穷大"的栈，在 Go 的栈上执行 C 函数会导致栈溢出

该函数在不同平台有不同的实现，拿amd64平台为例：

```
//file:runtime/asm_amd64.s
TEXT ·asmcgocall(SB),NOSPLIT,$0-20
	MOVQ	fn+0(FP), AX
	MOVQ	arg+8(FP), BX

	MOVQ	SP, DX

	// Figure out if we need to switch to m->g0 stack.
	// We get called to create new OS threads too, and those
	// come in on the m->g0 stack already. Or we might already
	// be on the m->gsignal stack.
	// 考虑是否需要切换到 m.g0 栈
    // 也用来调用创建新的 OS 线程，这些线程已经在 m.g0 栈中了
	get_tls(CX)
	MOVQ	g(CX), DI
	CMPQ	DI, $0
	JEQ	nosave
	MOVQ	g_m(DI), R8
	MOVQ	m_gsignal(R8), SI
	CMPQ	DI, SI
	JEQ	nosave
	MOVQ	m_g0(R8), SI
	CMPQ	DI, SI
	JEQ	nosave

	// Switch to system stack.
	// 切换到系统栈
	CALL	gosave_systemstack_switch<>(SB)
	MOVQ	SI, g(CX)
	MOVQ	(g_sched+gobuf_sp)(SI), SP

	// Now on a scheduling stack (a pthread-created stack).
	// Make sure we have enough room for 4 stack-backed fast-call
	// registers as per windows amd64 calling convention.
	// 于调度栈中（pthread 新创建的栈）
    // 确保有足够的空间给四个 stack-based fast-call 寄存器
    // 为使得 windows amd64 调用服务
	SUBQ	$64, SP
	ANDQ	$~15, SP	// 为 gcc ABI 对齐
	MOVQ	DI, 48(SP)	// save g
	MOVQ	(g_stack+stack_hi)(DI), DI
	SUBQ	DX, DI
	MOVQ	DI, 40(SP)	// 保存栈深 (不能仅保存 SP，因为栈可能在回调时被复制)
	MOVQ	BX, DI		// DI = AMD64 ABI 第一个参数
	MOVQ	BX, CX		// CX = Win64 第一个参数
	CALL	AX          // 调用 fn

	// Restore registers, g, stack pointer.
	// 恢复寄存器、 g、栈指针
	get_tls(CX)
	MOVQ	48(SP), DI
	MOVQ	(g_stack+stack_hi)(DI), SI
	SUBQ	40(SP), SI
	MOVQ	DI, g(CX)
	MOVQ	SI, SP

	MOVL	AX, ret+16(FP)
	RET

nosave:
    // 在系统栈上运行，可能没有 g
    // 没有 g 的情况发生在线程创建中或线程结束中（比如 Solaris 平台上的 needm/dropm）
    // 这段代码和上面类似，但没有保存和恢复 g，且没有考虑栈的移动问题（因为我们在系统栈上，而非 goroutine 栈）
    // 如果已经在系统栈上，则上面的代码可被直接使用，在 Solaris 上会进入下面这段代码。
    // 使用这段代码来为所有 "已经在系统栈" 的调用进行服务，从而保持正确性。
    SUBQ    $64, SP
    ANDQ    $~15, SP // ABI 对齐
    MOVQ    $0, 48(SP) // 上面的代码保存了 g, 确保 debug 时可用
    MOVQ    DX, 40(SP) // 保存原始的栈指针
    MOVQ    BX, DI  // DI = AMD64 ABI 第一个参数
    MOVQ    BX, CX  // CX = Win64 第一个参数
    CALL    AX
    MOVQ    40(SP), SI // 恢复原来的栈指针
    MOVQ    SI, SP
    MOVL    AX, ret+16(FP)
    RET	
```

### exitsyscall
exitsyscall的基本思路是，
- 先尝试获取一个p（优先尝试获取前面移交出去的p），若获取到了则直接返回到用户代码继续执行用户逻辑即可；
- 否则调用mcall切换到g0栈执行exitsyscall0函数

```go
//file:runtime/proc.go
func exitsyscall() {
    _g_ := getg()

	//进入系统调用之前保存的P
	oldp := _g_.m.oldp.ptr()

	//因为在进入系统调用之前已经解除了m和p之间的绑定，所以现在需要绑定p
    if exitsyscallfast(oldp) {
		...
        // There's a cpu for us, so we can run.
		//系统调用完成，增加syscalltick计数，sysmon线程依靠它判断是否是同一次系统调用
        _g_.m.p.ptr().syscalltick++
		
        // We need to cas the status and scan before resuming...
		//casgstatus函数会处理一些垃圾回收相关的事情，我们只需知道该函数重新把g设置成_Grunning状态即可
        casgstatus(_g_, _Gsyscall, _Grunning)
        
        // 返回到用户代码继续执行
        return
    }
	
    // Call the scheduler.
    //没有绑定到p，调用mcall切换到g0栈执行exitsyscall0函数
    mcall(exitsyscall0)
    ......
}
```

exitsyscall0还是会继续尝试获取空闲的p，若还是获取不到就会调用stopm将当前线程睡眠，等待被其它线程唤醒

## c语言实现的_cgo_e119c51a7968_Cfunc_sum函数

对于函数C语言中的sum函数，经过cgo编译后生成了一个名为_cgo_main.c的文件，该文件包含了C语言函数_cgo_e119c51a7968_Cfunc_sum的实现。

```c
void
_cgo_e119c51a7968_Cfunc_sum(void *v)
{
	struct {
		int p0;
		int p1;
		int r;
		char __pad12[4];
	} __attribute__((__packed__)) *_cgo_a = v;
	char *_cgo_stktop = _cgo_topofstack();
	__typeof__(_cgo_a->r) _cgo_r;
	_cgo_tsan_acquire();
	_cgo_r = sum(_cgo_a->p0, _cgo_a->p1);
	_cgo_tsan_release();
	_cgo_a = (void*)((char*)_cgo_a + (_cgo_topofstack() - _cgo_stktop));
	_cgo_a->r = _cgo_r;
	_cgo_msan_write(&_cgo_a->r, sizeof(_cgo_a->r));
}
```

函数体各段代码含义如下
函数_cgo_e119c51a7968_Cfunc_sum接收一个参数
- 输入参数v被转化为_cgo_a参数，它是一个结构体，包含两个整数成员p0和p1，以及一个整数成员r。
- _cgo_topofstack 函数用于 C 函数调用后恢复调用栈
- _cgo_tsan_acquire 和 _cgo_tsan_release 则是用于扫描 CGO 相关的函数，是对 CGO 相关函数的指针做相关检查
- sum用于真正计算两个数的和
- 通过_cgo_a->r = _cgo_r将结果赋值给_cgo_a->r

- 更详细的细节可以参考 https://golang.org/src/cmd/cgo/doc.go 内部的代码注释和 runtime.cgocall 函数的实现

## 结论
在cgo的调用中
- cgo 调用会将当前协程栈移到系统栈
- cgo 高并发调用且阻塞超过 20 微秒时会新建线程



## 参考
- [Go与C的桥梁：CGO入门剖析与实践](https://cloud.tencent.com/developer/article/1786332?areaId=106001)
- [cgo内部机制](https://chai2010.cn/advanced-go-programming-book/ch2-cgo/ch2-05-internal.html)
- [从源码分析 Go 语言使用 cgo 导致的线程增长](https://www.cnblogs.com/t102011/p/17457120.html)
- [Golang调度器源码分析](https://cs50mu.github.io/post/golang-scheduler/)