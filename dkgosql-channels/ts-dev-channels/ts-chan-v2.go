package main
import "fmt"

func main(){
	var statusChan = make(chan string)


	go func(sc chan string){
		i:=0
		for i<10{
			sc <- fmt.Sprintf("init id:%d",i)
			i++
		}
		close(statusChan)
	}(statusChan)

	for {
		select{
		case out, ok := <-statusChan:
			if !ok{
				fmt.Printf("channel closed\n")	
				return
			}
			fmt.Printf("Status Channel:%s\n\n", out)
		default:
			// fmt.Printf("default\n")
		}
	} 

}