package main

import "fmt"
import "time"

func main() {

	i := 2
	fmt.Print("Escribe ", i , " como ")

	//ejemplo de switch simple
	switch i{
		case 1: 
			fmt.Println("uno")
		case 2: 
			fmt.Println("dos")
		case 3: 
			fmt.Println("tres")
	}

	//se pueden evaluar varias condiciones separadas con ',' que puedan responder una misma respuesta
	switch time.Now().Weekday(){
	case time.Saturday, time.Sunday:
		fmt.Println("Es fin de semana")
	default:
		fmt.Println("Es día de semana")
	}

	//emplear switch sin expresión alguna puede servir como alternativa a if-else
	t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Println("Antes del mediodía")
    case t.Hour() == 12 && t.Minute() == 00:
    	fmt.Println("Mediodía")
    default:
        fmt.Println("Después del mediodía")
    }

}