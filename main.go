package main

import (
	"fmt"
	"net"
	"sync"
)

func main(){
	for i:=21;i<1200;i++{
		address := fmt.Sprintf("20.194.168.28:%d",i)
		dial, err := net.Dial("tcp", address)
		if err !=nil{
			fmt.Printf("%s 关闭\n",address)
			continue
		}
		dial.Close()
		fmt.Printf("%s ok\n",address)
	}
}


func curr(){
	var wg sync.WaitGroup
	for i:=21;i<1200;i++{
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			address := fmt.Sprintf("20.194.168.28:%d",j)
			dial, err := net.Dial("tcp", address)
			if err !=nil{
				fmt.Printf("%s 关闭\n",address)
				return
			}
			dial.Close()
			fmt.Printf("%s ok\n",address)
		}(i)
	}
	wg.Wait()
}


func workerPool(){
	ports:=make(chan int,100)
	result:=make(chan int)
	var open []int
	for i:=0;i<cap(ports);i++{
		go worker(ports,result)
	}
	go func() {
		for i:=1;i<1024;i++{
			ports<-i
		}
	}()

	for i:=0;i<1024;i++{
		p := <-result
		if p!=0{
			open=append(open,p)
		}
	}
	close(ports)
	close(result)
}

func worker(ports,result chan int){
	for p:=range ports{
		address := fmt.Sprintf("127.0.0.1:%d",p)
		dial, err := net.Dial("tcp", address)
		if err !=nil{
			result<-0
			return
		}
		dial.Close()
		result<-p
	}
}