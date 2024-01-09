package mains

import "fmt"

func Channels() {
	/* Quando criamos um channel, ele espera que seja consumido de alguma forma*/
	channel := make(chan int)

	/* Também temos o buffered channel, onde podemos definir o tamanho do buffer
	   e ele não espera que seja consumido de alguma forma
	   OBS: O tamanho do buffer deve ser maior que o numero de goroutines que irão consumir o channel */
	bufferedChannel := make(chan int, 3)

	go setNumbers(bufferedChannel)

	go setNumbers(channel)

	for v := range channel {
		fmt.Println(v)
	}
}


func setNumbers(channel chan int) {
	for i := 0; i < 10; i++ {
		channel <- i
	}
	close(channel)
	
}
