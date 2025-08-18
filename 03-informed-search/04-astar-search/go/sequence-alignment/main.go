package main

import (
	"fmt"
	"strings"
)

func min(vals ...float64) float64 {
	m := vals[0]
	for _, v := range vals[1:] {
		if v < m {
			m = v
		}
	}
	return m
}

func stringalign(ainstr, binstr string, mispen, gappen, skwpen float64) {
	ain := strings.Split(ainstr, "")
	bin := strings.Split(binstr, "")
	ia := len(ain)
	ib := len(bin)

	// Inicializar matrices cost y marked
	cost := make([][]float64, ia+1)
	marked := make([][]int, ia+1)
	for i := 0; i <= ia; i++ {
		cost[i] = make([]float64, ib+1)
		marked[i] = make([]int, ib+1)
	}

	// Inicializar costos de bordes (gaps)
	cost[0][0] = 0.0
	for i := 1; i <= ia; i++ {
		cost[i][0] = cost[i-1][0] + skwpen
	}
	for j := 1; j <= ib; j++ {
		cost[0][j] = cost[0][j-1] + skwpen
	}

	// Llenar matriz de costos
	for i := 1; i <= ia; i++ {
		for j := 1; j <= ib; j++ {
			dn := cost[i-1][j] + gappen
			if j == ib {
				dn = cost[i-1][j] + skwpen
			}
			rt := cost[i][j-1] + gappen
			if i == ia {
				rt = cost[i][j-1] + skwpen
			}
			dg := cost[i-1][j-1]
			if ain[i-1] == bin[j-1] {
				dg += -1.0
			} else {
				dg += mispen
			}
			cost[i][j] = min(dn, rt, dg)
		}
	}

	i, j := ia, ib

	var aout []string
	var bout []string
	var summary []string

	// Backtracking
	for i > 0 || j > 0 {
		marked[i][j] = 1
		dn := 1e99
		rt := 1e99
		dg := 1e99

		if i > 0 {
			dn = cost[i-1][j] + gappen
			if j == ib {
				dn = cost[i-1][j] + skwpen
			}
		}
		if j > 0 {
			rt = cost[i][j-1] + gappen
			if i == ia {
				rt = cost[i][j-1] + skwpen
			}
		}
		if i > 0 && j > 0 {
			dg = cost[i-1][j-1]
			if ain[i-1] == bin[j-1] {
				dg += -1.0
			} else {
				dg += mispen
			}
		}

		if dg <= dn && dg <= rt {
			aout = append(aout, ain[i-1])
			bout = append(bout, bin[j-1])
			if ain[i-1] == bin[j-1] {
				summary = append(summary, "=")
			} else {
				summary = append(summary, "!")
			}
			i--
			j--
		} else if dn < rt {
			aout = append(aout, ain[i-1])
			bout = append(bout, " ")
			summary = append(summary, " ")
			i--
		} else {
			aout = append(aout, " ")
			bout = append(bout, bin[j-1])
			summary = append(summary, " ")
			j--
		}
		marked[i][j] = 1
	}

	// Invertir slices (porque el backtracking va del final al inicio)
	reverseStrings := func(s []string) {
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			s[i], s[j] = s[j], s[i]
		}
	}

	reverseStrings(aout)
	reverseStrings(bout)
	reverseStrings(summary)

	// Imprimir resultado
	fmt.Println("Alineamiento A:")
	fmt.Println(strings.Join(aout, ""))
	fmt.Println("Alineamiento B:")
	fmt.Println(strings.Join(bout, ""))
	fmt.Println("Resumen:")
	fmt.Println(strings.Join(summary, ""))

	// Imprimir matriz de costos con marcados
	fmt.Println("\nMatriz de costos (marcados con *):")
	header := "    " + strings.Join(bin, " ")
	fmt.Println(header)
	for row_i := 0; row_i <= ia; row_i++ {
		line := ""
		if row_i == 0 {
			line += " "
		} else {
			line += ain[row_i-1]
		}
		for col_j := 0; col_j <= ib; col_j++ {
			val := cost[row_i][col_j]
			mark := " "
			if marked[row_i][col_j] == 1 {
				mark = "*"
			}
			line += fmt.Sprintf(" %4.1f%s", val, mark)
		}
		fmt.Println(line)
	}
}

func main() {
	stringalign("ATCGTACGTA", "ATGGTCGTA", 1.0, 1.0, 1.0)
}
