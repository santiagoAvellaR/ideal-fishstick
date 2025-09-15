// Package utils proporciona utilidades comunes para el proyecto
package utils

import (
	"fmt"
	"strings"
	"time"
)

// MathUtils contiene utilidades matemáticas
type MathUtils struct{}

// Max retorna el valor máximo entre dos enteros
func (m MathUtils) Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Min retorna el valor mínimo entre dos enteros
func (m MathUtils) Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Factorial calcula el factorial de un número
func (m MathUtils) Factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * m.Factorial(n-1)
}

// IsPrime verifica si un número es primo
func (m MathUtils) IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// StringUtils contiene utilidades para manejo de strings
type StringUtils struct{}

// Reverse invierte una cadena de texto
func (s StringUtils) Reverse(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// WordCount cuenta las palabras en una cadena
func (s StringUtils) WordCount(str string) int {
	words := strings.Fields(strings.TrimSpace(str))
	return len(words)
}

// Capitalize capitaliza la primera letra de cada palabra
func (s StringUtils) Capitalize(str string) string {
	return strings.Title(strings.ToLower(str))
}

// Contains verifica si una cadena contiene una subcadena
func (s StringUtils) Contains(str, substr string) bool {
	return strings.Contains(strings.ToLower(str), strings.ToLower(substr))
}

// TimeUtils contiene utilidades para manejo de tiempo
type TimeUtils struct{}

// CurrentTime retorna la hora actual formateada
func (t TimeUtils) CurrentTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// FormatDuration formatea una duración en formato legible
func (t TimeUtils) FormatDuration(d time.Duration) string {
	hours := int(d.Hours())
	minutes := int(d.Minutes()) % 60
	seconds := int(d.Seconds()) % 60

	if hours > 0 {
		return fmt.Sprintf("%dh %dm %ds", hours, minutes, seconds)
	} else if minutes > 0 {
		return fmt.Sprintf("%dm %ds", minutes, seconds)
	}
	return fmt.Sprintf("%ds", seconds)
}

// IsWeekend verifica si una fecha cae en fin de semana
func (t TimeUtils) IsWeekend(date time.Time) bool {
	weekday := date.Weekday()
	return weekday == time.Saturday || weekday == time.Sunday
}

// DaysBetween calcula los días entre dos fechas
func (t TimeUtils) DaysBetween(date1, date2 time.Time) int {
	diff := date2.Sub(date1)
	return int(diff.Hours() / 24)
}

// Funciones públicas (exportadas) - inician con mayúscula
// Estas funciones están disponibles fuera del paquete

// GetMathUtils retorna una instancia de MathUtils
func GetMathUtils() MathUtils {
	return MathUtils{}
}

// GetStringUtils retorna una instancia de StringUtils
func GetStringUtils() StringUtils {
	return StringUtils{}
}

// GetTimeUtils retorna una instancia de TimeUtils
func GetTimeUtils() TimeUtils {
	return TimeUtils{}
}

// PrintSeparator imprime una línea separadora
func PrintSeparator(char string, length int) {
	fmt.Println(strings.Repeat(char, length))
}

// Funciones privadas (no exportadas) - inician con minúscula
// Estas funciones NO están disponibles fuera del paquete

// helper function para validación interna
func isValidString(s string) bool {
	return strings.TrimSpace(s) != ""
}

// helper function para formatear números
func formatNumber(n int) string {
	return fmt.Sprintf("%,d", n)
}
