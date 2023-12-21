package mains

import (
	"fmt"
	"sync"
	"time"
)

func SynAwaiGroup() {
	var wg sync.WaitGroup
	wg.Add(3) // três processos que serão aguardados linhas: 13, 14, 15

	go callDatabase(&wg)    // passamos a REFERENCIA de wg -> linha 10
	go callAPI(&wg)         // passamos a REFERENCIA de wg -> linha 10
	go processInternal(&wg) // passamos a REFERENCIA de wg -> linha 10
 
	wg.Wait() // Esperamos as goroutines terminarem
}

 
func callDatabase(wg *sync.WaitGroup) { // em cada função passamos o ponteiro de wg
	time.Sleep(1 * time.Second)
	fmt.Println("Finalizado callDatabase")
	wg.Done() // Quando a função for finalizada, wg.Done() como aponta para wg -> linha 10, finaliza e espera os seguintes terminarem
}
func callAPI(wg *sync.WaitGroup) {
	time.Sleep(2 * time.Second)
	fmt.Println("Finalizado callAPI")
	wg.Done()

}
func processInternal(wg *sync.WaitGroup) {
	time.Sleep(3 * time.Second)
	fmt.Println("Finalizado processInternal")
	wg.Done()

}
