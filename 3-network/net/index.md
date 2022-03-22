在协程没有流行以前，传统的网络编程中，同步阻塞是性能底下的代名词。各种基于 epoll 的异步非阻塞的模型虽然提高了性能，但是基于回调函数的编程方式却非常不符合人的的直线思维模式。开发出来的代码的也不那么容易被人理解。

Golang 的出现，可以说是将协程编程模式推向了一个高潮。这种新的编程方式达到了一个折中的状态，既兼顾了同步编程方式的简单易用，也在底层通过协程和epoll的配合避免了线程切换的性能损耗。换句话说就是既简单，性能又还不挺错。

golang 的 net 包通过结合使用协程和 epoll。封装出了一种在应用层看起来是同步阻塞的编程模型，但是底层实际上仍然是基于 epoll 的异步非阻塞来运行的。