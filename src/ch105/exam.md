1.下面代码输出什么？请简要说明。(B)
    
    var c = make(chan int)
    var a int
    
    func f() {
        a = 1
        <-c
    }
    func demo11() {
        go f()
        c <- 0
        print(a)
    }

    A. 不能编译；
    B. 输出 1；
    C. 输出 0；
    D. panic；


2.下面代码输出什么？请简要说明。 (A)

    type MyMutex struct {
        count int
        sync.Mutex
    }
    
    func main() {
        var mu MyMutex
        mu.Lock()
        var mu1 = mu
        mu.count++
        mu.Unlock()
        mu1.Lock()
        mu1.count++
        mu1.Unlock()
        fmt.Println(mu.count, mu1.count)
    }

A. 不能编译；<br>
B. 输出 1, 1；<br>
C. 输出 1, 2；<br>
D. fatal error；<br>