package golanggoroutine

import (
	"fmt"
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