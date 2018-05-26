package main


import (
	"fmt"
	"runtime"
	"sync"
	"math/rand"
	"time"
)


var arr [][]int
var wg sync.WaitGroup
var mutex sync.Mutex

func init()  {
	rand.Seed(time.Now().UnixNano())
}

func main(){
	cpuNums := runtime.NumCPU()
	fmt.Println("CPU核心数：",cpuNums)
	runtime.GOMAXPROCS(cpuNums)

	ch := make(chan int)
	nums := 10
	wg.Add(nums)
	for i:=0;i<nums;i++{
		go work(i,ch)
	}
	ch <-1
	wg.Wait()
	fmt.Println("总数=",len(arr))
	fmt.Println(arr)
	final()
}

func final()  {
	fmt.Println("okokokok=====")
}

func work(id int,ch chan int){
	fmt.Println(id)

	var res []int
	rand.Seed(time.Now().UnixNano())
	for {
		_,ok := <-ch
		if !ok{
			fmt.Println("通道关闭")
			return
		}
		r := rand.Intn(10000)
		if r == 100{
			res = append(res,id)
		}

		if len(res) == 10{
			mutex.Lock()
			{
				arr = append(arr,res)
			}
			mutex.Unlock()
			close(ch)
			break
		}

	}
	defer wg.Done()
}

func sortArr() []int{
	var row [] int
	for i:=0;i<4;i++{
		row = append(row, i)
	}
	return row
}


