package main

import (
	"fmt"

	"./process"
)

func showCounter(c chan bool, flag bool) {
	for {
		c <- flag
	}
}
func hideCounter(c chan bool) {
	for {
		c <- false
	}
}
func searchByID(list []*process.Process, id int) {

	for _, value := range list {
		if value.Id == id {
			value.Kill()
			fmt.Printf(" El proceso %d  fue terminado exitosamente\n", value.Id)
		}
	}
}
func main() {
	list := []*process.Process{}
	flag := false
	id := 1
	for {
		fmt.Println("Adinistrador de procesos en golang")
		fmt.Println("Selecciona una opcion")
		fmt.Println("1.- Registrar un Proceso")
		fmt.Println("2.- Mostrar procesos")
		fmt.Println("3.- Terminar proceso")
		fmt.Println("4.- salir")
		var opc int
		fmt.Scanln(&opc)
		switch opc {
		case 1:
			pros := &process.Process{Id: id, Counter: 0, Iterate: true, Chanel: &flag}
			go pros.Start()
			list = append(list, pros)
			id++
		case 2:
			flag = !flag
		case 3:
			var idx int
			fmt.Println("Ingrese id del proceso a terminar:")
			fmt.Scan(&idx)
			searchByID(list, idx)
		case 4:
			for _, v := range list {
				v.Kill()
			}
			fmt.Println("Todos los procesos fueron terminado exitosamente")
			break
		default:
			fmt.Println("La opcion seleccionada es invalida")
		}
	}

}
