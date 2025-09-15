package main

import (
	"fmt"
	"time"

	"github.com/santiagoAvellaR/ideal-fishstick/utils"
)

func main() {
	fmt.Println("=== EJEMPLOS USANDO EL PAQUETE UTILS ===")
	fmt.Println()

	// Ejemplos de utilidades matemáticas
	mathUtils := utils.GetMathUtils()

	fmt.Println("--- Utilidades Matemáticas ---")
	fmt.Printf("Max(15, 8) = %d\n", mathUtils.Max(15, 8))
	fmt.Printf("Min(15, 8) = %d\n", mathUtils.Min(15, 8))
	fmt.Printf("Factorial(5) = %d\n", mathUtils.Factorial(5))

	// Verificar números primos
	numbers := []int{2, 7, 15, 17, 25}
	fmt.Print("Números primos: ")
	for _, num := range numbers {
		if mathUtils.IsPrime(num) {
			fmt.Printf("%d ", num)
		}
	}
	fmt.Println()

	utils.PrintSeparator("-", 40)

	// Ejemplos de utilidades de string
	stringUtils := utils.GetStringUtils()

	fmt.Println("--- Utilidades de String ---")
	originalText := "hola mundo desde Go"
	fmt.Printf("Texto original: %s\n", originalText)
	fmt.Printf("Texto invertido: %s\n", stringUtils.Reverse(originalText))
	fmt.Printf("Capitalizado: %s\n", stringUtils.Capitalize(originalText))
	fmt.Printf("Cantidad de palabras: %d\n", stringUtils.WordCount(originalText))
	fmt.Printf("¿Contiene 'Go'?: %t\n", stringUtils.Contains(originalText, "Go"))

	utils.PrintSeparator("-", 40)

	// Ejemplos de utilidades de tiempo
	timeUtils := utils.GetTimeUtils()

	fmt.Println("--- Utilidades de Tiempo ---")
	fmt.Printf("Hora actual: %s\n", timeUtils.CurrentTime())

	// Duración de ejemplo
	duration := 3*time.Hour + 25*time.Minute + 45*time.Second
	fmt.Printf("Duración formateada: %s\n", timeUtils.FormatDuration(duration))

	// Verificar fin de semana
	now := time.Now()
	fmt.Printf("¿Es fin de semana hoy?: %t\n", timeUtils.IsWeekend(now))

	// Días entre fechas
	date1 := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	date2 := time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC)
	fmt.Printf("Días entre 2024-01-01 y 2024-12-31: %d días\n", timeUtils.DaysBetween(date1, date2))

	utils.PrintSeparator("=", 50)
	fmt.Println("¡Ejemplos completados con éxito!")
}
