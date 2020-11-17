package main

import (
	"fmt"
	"math"
)

type Graph [][]float64

func (g Graph) FloydWarshall() Graph {
	n := len(g)
	weights := *g.Weights()

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				weights[i][j] = math.Min(weights[i][j], weights[i][k]+weights[k][j])
			}
		}
	}

	return weights
}

func (g Graph) Dijkstra() Graph {
	n := len(g)
	weights := *g.Weights()

	// К - стартовая точка
	for k := 0; k < n; k++ {
		visited := make([]bool, n) // Посещенные вершины

		// Помечаем все вершин как не посещенные
		for i := 0; i < n; i++ {
			visited[i] = false
		}

		// Прохожусь по всем строкам
		for i := 0; i < n; i++ {
			min := math.Inf(+1)
			current := 0 // текущая наименьшая

			// Прохожусь по всем столбцам
			for j := 0; j < n; j++ {
				if visited[j] == false && weights[k][j] < min {
					current = j
					min = weights[k][j]
				}
			}

			if min == math.Inf(+1) {
				break
			}

			// Прохожусь по всем столбцам
			for j := 0; j < n; j++ {
				if visited[j] == false {
					weights[k][j] = math.Min(weights[k][j], weights[k][current]+weights[current][j])
				}
			}

			visited[current] = true
		}
	}

	return weights
}

// Weights возвращает граф весов, где 0
// (отсутствие ребра между двумя вершинами)
// заменяется за положительную бесконечность
func (g Graph) Weights() *Graph {
	n := len(g)
	w := make(Graph, n)
	for i := range w {
		w[i] = make([]float64, n)
		copy(w[i], g[i])
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j && w[i][j] == 0 {
				w[i][j] = math.Inf(+1)
				continue
			}
		}
	}

	return &w
}

func (g Graph) Print() {
	n := len(g)

	for i := 0; i < n; i++ {

		for j := 0; j < n; j++ {
			if g[i][j] == math.Inf(+1) {
				fmt.Printf("    ∞")
			} else {
				fmt.Printf("%*v", 5, g[i][j])
			}
		}
		fmt.Println()
	}

}

func main() {
	// Online: https://play.golang.org/p/DZf9RywFG_Y

	g := Graph{
		/*  i */
		/*  0 */ {0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 3, 0},
		/*  1 */ {15, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		/*  2 */ {7, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		/*  3 */ {0, 0, 0, 0, 0, 0, 0, 0, 0, 9, 0, 0, 0, 0, 0},
		/*  4 */ {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 0},
		/*  5 */ {0, 0, 0, 0, 3, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0},
		/*  6 */ {0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
		/*  7 */ {0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0},
		/*  8 */ {0, 13, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		/*  9 */ {0, 0, 0, 0, 0, 0, 0, 0, 11, 0, 0, 0, 0, 0, 0},
		/* 10 */ {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5},
		/* 11 */ {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 11, 0, 0, 0, 0},
		/* 12 */ {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 7, 0},
		/* 13 */ {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 9, 0, 0, 0},
		/* 14 */ {0, 0, 0, 7, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	fmt.Println("FloydWarshall:")
	g.FloydWarshall().Print()
	fmt.Println()
	fmt.Println("Dijkstra:")
	g.Dijkstra().Print()
}
