package designMode

import (
	"sync"
	"sync/atomic"
)

type singleton struct{}

var instance *singleton

func GetInstance() *singleton {

	if instance == nil {
		instance = &singleton{} // 非线程安全的
	}

	return instance
}

// singleton by lock
var mu sync.Mutex

func GetInstanceByLock() *singleton {
	mu.Lock()
	defer mu.Unlock()

	if instance == nil {
		instance = &singleton{}
	}

	return instance
}

// double check  因为编译器优化，但是没有实例保存的状态的原子性检查。全面的技术考虑，这并不是安美的。但是已经比之前的方法好多了。
func GetInstanceByDoubleCheck() *singleton {
	if instance == nil {
		mu.Lock()
		defer mu.Unlock()

		if instance == nil {
			instance = &singleton{}
		}
	}

	return instance
}

// init
func init() {
	instance = &singleton{}
}

// atomic double check
var initialized uint32

func GetInstanceByAtomic() *singleton {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}

	mu.Lock()
	defer mu.Unlock()

	if atomic.LoadUint32(&initialized) == 0 {
		instance = &singleton{}
		atomic.StoreUint32(&initialized, 1)
	}

	return instance
}

// go 惯用的单例 方式 do once
var once sync.Once

func GetInstanceByDoOnce() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

//总结
//当涉及到并行和并发代码时，需要详细检查你的代码。始终让你的团队成员进行代码审查，因此对于这样的事情才能更容易的监督。
//
//所有的切换到Go语言的新开发者，需要明确的理解线程安全的原理才能更好的改善你的代码。即使Go语言本身通过做了很多努力允许你使用很少的并发知识来设计并发代码。仍有一些语言无法帮你处理的一些情况，你依然需要在开发代码时应用最佳的实践方法
//
//补充
//为什么不使用 init()函数声明单例？这样更简单高效
//答： 如果想要实现懒初始化(Lazy initialization)则需要考虑同步问题
//sync.Once 会保证又一次调用成功之后再返回，也就是说如果多人同时调用Do函数，则只有一个人调用成功之后，其他人才会返回。
//当调用函数Panic时，也会认为其已经调用完成，其它调用方也会返回。
