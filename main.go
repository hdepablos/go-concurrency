package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	/*
		****************************************************************
			Concurrency golang
				1.- Básica: branch feature/concurrency-basic
					Concurrencia básica:
		****************************************************************
	*/

	startTime := start()

	timestamps := make([]string, 10)

	// Declarar un grupo asyncrono
	var wg sync.WaitGroup

	for i := range timestamps {
		// Se realiza esto para controlar el index
		idx := i
		// Agregar una tarea al grupo
		wg.Add(1)

		// esta es la función que permite el paralelismo
		go func() {
			// Ejecutar la function en paralelo
			defer wg.Done()

			// llamar a la función
			timeStamp(idx, startTime)
		}()
	}

	// Espera a que todos las tareas terminen
	wg.Wait()

	end(startTime)
}

func timeStamp(i int, startTime time.Time) {
	time.Sleep(time.Microsecond * 200)
	t := time.Now().Sub(startTime).String()
	fmt.Printf("TimeStamp %v: %s \n", i, t)
}

func start() time.Time {
	fmt.Println("Start ... !!!")
	return time.Now()
}

func end(startTime time.Time) {
	duration := time.Now().Sub(startTime).String()

	fmt.Println("duration")
	fmt.Println(duration)
	fmt.Println("End")
}
