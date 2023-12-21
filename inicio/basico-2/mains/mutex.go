package mains

import (
	"fmt"
	"sync"
)

/*
	sync.Mutex
var VARIAVEL sync.Mutex: Bloqueio mutuamente exclusivo, que é utilizado para controlar o acesso a um recurso
compartilhado, nesse caso a varável i

	Recurso compartilhado:
No contexto desse exercício, nosso recurso compartilhado é o i
Imaginemos que temos várias chamadas para => go incrementsI(&wg, &VARIAVEL, &i), cada uma delas seria uma goroutine
separada, todas compartilhando o acesso à variável i

	VARIAVEL.Lock()
Utilizamos .Lock() para adiquirir o bloqueio. Quando uma goroutine chama VARIAVEL.Lock(), ela ganha propriedade
exclusiva do bloqueio VARIAVEL. Se outras goroutines tentarem adquirir o bloqueio chamado VARIAVEL.Lock(), elas
serão bloqueadas até que a goroutine que possui o boqueio o libere

	VARIAVEL.Unlock()
Utilizamos VARIAVEL.Unlock() para liberar o bloqueio. Quando uma goroutine termina de usar o recurso compartilhado
(i) ela DEVE chamar o VARIAVEL.Unlock() para que possamos liberar o bloqueio. Assim outras goroutines adquiram o
bloqueio e acessem o recurso compartilhado
*/

func mutex() {
	// grupo de espera
	var wg sync.WaitGroup
	// quantidade de funções que estaremos esperando, nesse caso uma, incrementsI()
	wg.Add(2)

	// definimos m para utilizar .Lock e .Unlock()
	var m sync.Mutex
	i := 0

	// passamos os endereços de wg (WaitGroup), m (Mutex), e nossa variável de incremento
	go incrementsI(&wg, &m, &i) // deve printar pro volta de 10000
	go incrementsI(&wg, &m, &i) // deve printar 20000

	// esperamos todas terminar
	wg.Wait() 
}

/* 
func increments que recebe como parametros os ponteiros das variáveis */
func incrementsI(wg *sync.WaitGroup, m *sync.Mutex, i *int) {
  for x := 0; x < 10000; x++ {
    m.Lock() // travamos o recuros para apenas uma goroutine por vez
    *i++ 
    m.Unlock()
  }

  wg.Done()

  m.Lock()
  fmt.Println(*i)
  m.Unlock()
}