package golanggoroutine

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
channel := make(chan string)
defer close (channel)
// channel <- "Wahyu" //mengirim data ke channel
go func ()  {
	time.Sleep(2 * time.Second)
	channel <- "Wahyu Nurdian"
	fmt.Println("Selesai mengirim data ke channel")
}()
data := <-channel
fmt.Println(data)
time.Sleep(5 * time.Second)
}

func GiveMeResponse(channel chan string)  {
	time.Sleep(2* time.Second)
	channel <- "Wahyu Nurdian"
}

func TestChannelAsParameter(t * testing.T)  {
	channel := make(chan string)
defer close (channel)
go GiveMeResponse(channel)

data := <-channel
fmt.Println(data)
time.Sleep(5 * time.Second)

}


func OnlyIn( channel chan<- string)  {
	time.Sleep(2 * time.Second)
	channel <- "Wahyu Nurdian"

	fmt.Println("Selesai mengirim data ke channel")
}

func OnlyOut( channel <-chan string)  {
	data := <-channel
fmt.Println(data)

time.Sleep(5 * time.Second)
}

func TestInOutChannel( t *testing.T)  {
	channel := make(chan string)
defer close (channel)
go OnlyIn(channel)
go OnlyOut(channel)

time.Sleep(5 * time.Second)
}

func TestBufferedChannel(t *testing.T)  {
	channel := make(chan string, 3 )
	defer close(channel)
	

	go func ()  {
		channel <- "Wahyu"
		channel <- "Nurd"
		channel <- "Dian"
	}()

	
go func ()  {
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)
}()
time.Sleep( 2 *time.Second)
	fmt.Println("Selesai")
}

func TestRangeChannel(t *testing.T)  {
	channel := make(chan string)

	go func ()  {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke" + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Menerima data", data)
	}
	fmt.Println("Selesai")
}