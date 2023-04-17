package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

// Choice states
const (
	WorkWithSLAE          = 1
	SolveByGaussianMethod = 2
	SolveByJacobiMethod   = 3
	MethodsComparison     = 4
	Exit                  = 5
)

// Subchoice states
const (
	ManualGeneration = 1
	RandomGeneration = 2
	DisplayMatrix    = 3
	SubExit          = 4
)

const (
	NumberOfOrders = 11
)

func main() {
	clearConsole()

	var A [][]float64
	var b, x []float64

	for {
		fmt.Print(menu())

		fmt.Print("Определите действие: ")
		var choice int
		fmt.Scanln(&choice)

		clearConsole()

		switch choice {

		case WorkWithSLAE:
			clearConsole()
			fmt.Println("Вы выбрали работу с СЛАУ")
			isWorking := true
			for isWorking {
				fmt.Print(menuForSystems())

				fmt.Print("Определите действие: ")
				var subchoice int
				fmt.Scanln(&subchoice)

				clearConsole()
				switch subchoice {

				// Задание значений руками
				case ManualGeneration:
					fmt.Print("Определите порядок: ")
					var order int
					fmt.Scanln(&order)

					A = make([][]float64, order)
					for i := 0; i < order; i++ {
						A[i] = make([]float64, order)
						for j := 0; j < order; j++ {
							fmt.Printf("Определите коэффициент A[%d][%d]: ", i, j)
							fmt.Scanln(&A[i][j])
							fmt.Println()
						}
					}

					b = make([]float64, order)
					for i := 0; i < order; i++ {
						fmt.Printf("Определите свободный член b[%d]: ", i)
						fmt.Scanln(&b[i])
						fmt.Println()
					}

				// Задание значений случайным образом
				case RandomGeneration:
					fmt.Print("Определите порядок: ")
					var order int
					fmt.Scanln(&order)

					fmt.Print("Определите левую границу интервала: ")
					var left float64
					fmt.Scanln(&left)

					fmt.Print("Определите правую границу интервала: ")
					var right float64
					fmt.Scanln(&right)

					A = GenerateDiagonallyDominantMatrixWithInterval(order, left, right)
					b = GenerateVectorWithInterval(order, left, right)

				// Отображение матрицы в консоли
				case DisplayMatrix:
					if len(A) == 0 {
						fmt.Println("СЛАУ не задана")
					} else {
						fmt.Println("Текущая матрица: ")
						printMatrix(A)
					}

				// Выход из консоли
				case SubExit:
					isWorking = false

				// Ошибка ввода
				default:
					fmt.Println("Неопределенное действие. Попробуйте еще раз.")
				}
			}

		case SolveByGaussianMethod:
			if len(A) == 0 {
				fmt.Println("СЛАУ не задана")
			} else {
				fmt.Println("Решение СЛАУ методом Гаусса")
				x, _ = SolveGaussian(A, b)
				if !isArrayEmpty(x) {
					fmt.Println("Результат решения: ")
					printArray(x)
				} else {
					fmt.Println("СЛАУ не совместна")
				}
			}

		case SolveByJacobiMethod:
			if len(A) == 0 {
				fmt.Println("СЛАУ не задана")
			} else {
				fmt.Println("Решение СЛАУ методом Якоби")
				fmt.Print("Задайте точность вычисления: ")
				var eps float64
				fmt.Scanln(&eps)
				x, _ = SolveJacobi(A, b, eps, 100000)
				if !isArrayEmpty(x) {
					fmt.Println("Результат решения: ")
					printArray(x)
				} else {
					fmt.Println("СЛАУ не совместна")
				}
			}

		case MethodsComparison:
			fmt.Println("Сравнение алгоритмов Гаусса и Якоби")

			fmt.Print("Определите левую границу интервала: ")
			var left float64
			fmt.Scanln(&left)

			fmt.Print("Определите правую границу интервала: ")
			var right float64
			fmt.Scanln(&right)

			data := make([][]string, NumberOfOrders)
			for i := 0; i < NumberOfOrders; i++ {
				order := 100 + i*10

				A = GenerateDiagonallyDominantMatrixWithInterval(order, left, right)
				b = GenerateVectorWithInterval(order, left, right)

				x1, _ := SolveGaussian(A, b)
				x2, _ := SolveJacobi(A, b, 1e-6, 100000)

				data[i] = []string{
					strconv.FormatInt(int64(order), 10),
					measureTime(SolveGaussian, A, b).String(),
					measureTime(SolveJacobi, A, b, 1e-6, 100000).String(),
					strconv.FormatFloat(subtractArrays(x1, x2), 'f', 4, 64),
				}
			}

			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"Порядок", "Метод Гаусса Время", "Метод Якоби Время", "Разность"})
			table.SetAlignment(tablewriter.ALIGN_LEFT)
			table.AppendBulk(data)
			table.Render()

		case Exit:
			fmt.Println("Выход из программы")
			os.Exit(0)

		default:
			fmt.Println("Неопределенное действие. Попробуйте еще раз.")
		}
	}
}
