package main

import (
	"fmt"
)

// Función sucesores (con switch optimizable)
func sucesores(n int) []int {
	switch n {
	case 1:
		return []int{2, 4, 5}
	case 2:
		return []int{1, 3, 4, 5, 6}
	case 3:
		return []int{2, 5, 6}
	case 4:
		return []int{1, 2, 5, 7, 8}
	case 5:
		return []int{1, 2, 3, 4, 6, 7, 8, 9}
	case 6:
		return []int{2, 3, 5, 8, 9}
	case 7:
		return []int{4, 5, 8}
	case 8:
		return []int{4, 5, 6, 7, 9}
	case 9:
		return []int{5, 6, 8}
	default:
		return nil
	}
}

// Búsqueda en anchura
func anchura(nodoInicio, nodoFin int) {
	lista := []int{nodoInicio}

	for len(lista) > 0 {
		nodoActual := lista[0]
		lista = lista[1:]

		fmt.Println(nodoActual)

		if nodoActual == nodoFin {
			fmt.Println("SOLUCIÓN")
			return
		}

		temp := sucesores(nodoActual)
		fmt.Println(temp)

		if temp != nil {
			lista = append(lista, temp...)
			fmt.Println(lista)
		}
	}
	fmt.Println("NO-SOLUCIÓN")
}

// Búsqueda en profundidad
func profundidad(nodoInicio, nodoFin int) {
	lista := []int{nodoInicio}

	for len(lista) > 0 {
		nodoActual := lista[0]
		lista = lista[1:]

		fmt.Println(nodoActual)

		if nodoActual == nodoFin {
			fmt.Println("SOLUCIÓN")
			return
		}

		temp := sucesores(nodoActual)
		reverse(temp)
		fmt.Println(temp)

		if temp != nil {
			temp = append(temp, lista...)
			lista = temp
			fmt.Println(lista)
		}
	}
	fmt.Println("NO-SOLUCIÓN")
}

// Función para invertir un slice de enteros (reverse)
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// Función principal
func main() {
	fmt.Println("Anchura:")
	anchura(1, 9)

	fmt.Println("\nProfundidad:")
	profundidad(1, 9)
}
