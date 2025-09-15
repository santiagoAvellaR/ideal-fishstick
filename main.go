package main

import (
	"fmt"
	"strconv"
)

// Constantes globales
const (
	AppName    = "Ideal Fishstick Go Examples"
	Version    = "1.0.0"
	MaxRetries = 3
)

// Variables globales
var (
	globalCounter int
	isDebug       = false
)

func main() {
	fmt.Println("=== Bienvenido a", AppName, "v"+Version, "===")
	fmt.Println()

	// Llamar a todas las funciones de ejemplo
	basicSyntaxExamples()
	functionExamples()
	structAndMethodExamples()
	interfaceExamples()
	errorHandlingExamples()
	concurrencyExample()
}

// basicSyntaxExamples demuestra la sintaxis básica de Go
func basicSyntaxExamples() {
	fmt.Println("=== SINTAXIS BÁSICA ===")

	// Declaración de variables
	var name string = "Santiago"
	var age int = 25
	var height float64 = 1.75
	var isStudent bool = true

	// Declaración corta
	city := "Bogotá"
	population := 8_000_000

	// Múltiples variables
	var x, y, z int = 1, 2, 3

	fmt.Printf("Nombre: %s, Edad: %d años\n", name, age)
	fmt.Printf("Altura: %.2f m, Estudiante: %t\n", height, isStudent)
	fmt.Printf("Ciudad: %s, Población: %d\n", city, population)
	fmt.Printf("Coordenadas: x=%d, y=%d, z=%d\n", x, y, z)

	// Arrays y slices
	var numbers [5]int = [5]int{1, 2, 3, 4, 5}
	languages := []string{"Go", "Python", "JavaScript", "Java"}

	fmt.Printf("Números: %v\n", numbers)
	fmt.Printf("Lenguajes: %v\n", languages)

	// Maps
	grades := make(map[string]int)
	grades["Matemáticas"] = 95
	grades["Física"] = 88
	grades["Programación"] = 100

	fmt.Printf("Calificaciones: %v\n", grades)

	// Estructuras de control
	if age >= 18 {
		fmt.Printf("%s es mayor de edad\n", name)
	} else {
		fmt.Printf("%s es menor de edad\n", name)
	}

	// Switch
	switch city {
	case "Bogotá":
		fmt.Println("Capital de Colombia")
	case "Medellín":
		fmt.Println("Ciudad de la eterna primavera")
	default:
		fmt.Println("Otra ciudad")
	}

	// For loops
	fmt.Print("Conteo: ")
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// Range
	fmt.Print("Lenguajes favoritos: ")
	for index, lang := range languages {
		fmt.Printf("%d:%s ", index+1, lang)
	}
	fmt.Println()
	fmt.Println()
}

// functionExamples demuestra diferentes tipos de funciones
func functionExamples() {
	fmt.Println("=== EJEMPLOS DE FUNCIONES ===")

	// Función simple
	result := add(10, 20)
	fmt.Printf("10 + 20 = %d\n", result)

	// Función con múltiples valores de retorno
	quotient, remainder := divide(17, 5)
	fmt.Printf("17 ÷ 5 = %d con residuo %d\n", quotient, remainder)

	// Función con valores de retorno nombrados
	area, perimeter := calculateRectangle(5, 3)
	fmt.Printf("Rectángulo 5x3: Área=%d, Perímetro=%d\n", area, perimeter)

	// Función variádica
	sum := sumAll(1, 2, 3, 4, 5)
	fmt.Printf("Suma de 1,2,3,4,5 = %d\n", sum)

	// Función anónima
	multiply := func(a, b int) int {
		return a * b
	}
	fmt.Printf("6 × 7 = %d\n", multiply(6, 7))

	// Closure
	counter := createCounter()
	fmt.Printf("Contador: %d\n", counter())
	fmt.Printf("Contador: %d\n", counter())
	fmt.Printf("Contador: %d\n", counter())

	fmt.Println()
}

// add suma dos números enteros
func add(a, b int) int {
	return a + b
}

// divide realiza división entera y retorna cociente y residuo
func divide(dividend, divisor int) (int, int) {
	quotient := dividend / divisor
	remainder := dividend % divisor
	return quotient, remainder
}

// calculateRectangle calcula área y perímetro de un rectángulo
func calculateRectangle(width, height int) (area int, perimeter int) {
	area = width * height
	perimeter = 2 * (width + height)
	return // retorno desnudo usando variables nombradas
}

// sumAll suma todos los números pasados como argumentos
func sumAll(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// createCounter demuestra closures
func createCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// =============== ESTRUCTURAS Y MÉTODOS ===============

// Person representa una persona
type Person struct {
	Name    string
	Age     int
	Email   string
	Address Address
}

// Address representa una dirección
type Address struct {
	Street  string
	City    string
	Country string
}

// String implementa el método String para Person
func (p Person) String() string {
	return fmt.Sprintf("%s (%d años) - %s", p.Name, p.Age, p.Email)
}

// GetFullAddress retorna la dirección completa
func (p Person) GetFullAddress() string {
	return fmt.Sprintf("%s, %s, %s", p.Address.Street, p.Address.City, p.Address.Country)
}

// HaveBirthday incrementa la edad de la persona
func (p *Person) HaveBirthday() {
	p.Age++
}

// UpdateEmail cambia el email de la persona
func (p *Person) UpdateEmail(newEmail string) {
	p.Email = newEmail
}

// Rectangle representa un rectángulo
type Rectangle struct {
	Width  float64
	Height float64
}

// Area calcula el área del rectángulo
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter calcula el perímetro del rectángulo
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Scale escalifica el rectángulo por un factor
func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

// structAndMethodExamples demuestra structs y métodos
func structAndMethodExamples() {
	fmt.Println("=== ESTRUCTURAS Y MÉTODOS ===")

	// Crear una persona
	person := Person{
		Name:  "Ana García",
		Age:   28,
		Email: "ana@example.com",
		Address: Address{
			Street:  "Calle 123 #45-67",
			City:    "Medellín",
			Country: "Colombia",
		},
	}

	fmt.Println("Persona:", person)
	fmt.Println("Dirección:", person.GetFullAddress())

	// Modificar usando métodos con puntero
	person.HaveBirthday()
	person.UpdateEmail("ana.garcia@newmail.com")

	fmt.Printf("Después del cumpleaños: %s, nuevo email: %s\n", person.Name, person.Email)
	fmt.Printf("Nueva edad: %d años\n", person.Age)

	// Trabajar con rectángulos
	rect := Rectangle{Width: 5.0, Height: 3.0}
	fmt.Printf("Rectángulo: %.1fx%.1f\n", rect.Width, rect.Height)
	fmt.Printf("Área: %.2f, Perímetro: %.2f\n", rect.Area(), rect.Perimeter())

	rect.Scale(2.0)
	fmt.Printf("Después de escalar x2: %.1fx%.1f\n", rect.Width, rect.Height)
	fmt.Printf("Nueva área: %.2f, Nuevo perímetro: %.2f\n", rect.Area(), rect.Perimeter())

	fmt.Println()
}

// =============== INTERFACES ===============

// Shape define la interfaz para formas geométricas
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Circle representa un círculo
type Circle struct {
	Radius float64
}

// Area implementa la interfaz Shape para Circle
func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

// Perimeter implementa la interfaz Shape para Circle
func (c Circle) Perimeter() float64 {
	return 2 * 3.14159 * c.Radius
}

// Triangle representa un triángulo
type Triangle struct {
	Base   float64
	Height float64
	Side1  float64
	Side2  float64
	Side3  float64
}

// Area implementa la interfaz Shape para Triangle
func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

// Perimeter implementa la interfaz Shape para Triangle
func (t Triangle) Perimeter() float64 {
	return t.Side1 + t.Side2 + t.Side3
}

// printShapeInfo imprime información de cualquier forma
func printShapeInfo(s Shape, name string) {
	fmt.Printf("%s - Área: %.2f, Perímetro: %.2f\n", name, s.Area(), s.Perimeter())
}

// interfaceExamples demuestra el uso de interfaces
func interfaceExamples() {
	fmt.Println("=== INTERFACES ===")

	// Crear diferentes formas
	rect := Rectangle{Width: 4.0, Height: 3.0}
	circle := Circle{Radius: 2.5}
	triangle := Triangle{
		Base:   6.0,
		Height: 4.0,
		Side1:  5.0,
		Side2:  5.0,
		Side3:  6.0,
	}

	// Usar polimorfismo con interfaces
	shapes := []Shape{rect, circle, triangle}
	names := []string{"Rectángulo", "Círculo", "Triángulo"}

	for i, shape := range shapes {
		printShapeInfo(shape, names[i])
	}

	// Calcular área total
	var totalArea float64
	for _, shape := range shapes {
		totalArea += shape.Area()
	}
	fmt.Printf("Área total de todas las formas: %.2f\n", totalArea)

	fmt.Println()
}

// =============== MANEJO DE ERRORES ===============

// Calculator simula una calculadora con manejo de errores
type Calculator struct {
	Name string
}

// Divide realiza división con manejo de errores
func (c Calculator) Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("no se puede dividir %g por cero", a)
	}
	return a / b, nil
}

// SquareRoot calcula raíz cuadrada con manejo de errores
func (c Calculator) SquareRoot(n float64) (float64, error) {
	if n < 0 {
		return 0, fmt.Errorf("no se puede calcular la raíz cuadrada de %g (número negativo)", n)
	}
	// Aproximación simple de raíz cuadrada
	if n == 0 {
		return 0, nil
	}

	x := n
	for i := 0; i < 10; i++ {
		x = (x + n/x) / 2
	}
	return x, nil
}

// parseAndDouble convierte string a número y lo duplica
func parseAndDouble(s string) (int, error) {
	num, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("error al convertir '%s' a número: %w", s, err)
	}
	return num * 2, nil
}

// errorHandlingExamples demuestra manejo de errores
func errorHandlingExamples() {
	fmt.Println("=== MANEJO DE ERRORES ===")

	calc := Calculator{Name: "MiCalculadora"}

	// División exitosa
	result, err := calc.Divide(10.0, 3.0)
	if err != nil {
		fmt.Printf("Error en división: %v\n", err)
	} else {
		fmt.Printf("10 ÷ 3 = %.4f\n", result)
	}

	// División por cero (error)
	result, err = calc.Divide(5.0, 0.0)
	if err != nil {
		fmt.Printf("Error en división: %v\n", err)
	} else {
		fmt.Printf("5 ÷ 0 = %.4f\n", result)
	}

	// Raíz cuadrada exitosa
	sqrt, err := calc.SquareRoot(16.0)
	if err != nil {
		fmt.Printf("Error en raíz cuadrada: %v\n", err)
	} else {
		fmt.Printf("√16 = %.4f\n", sqrt)
	}

	// Raíz cuadrada de número negativo (error)
	sqrt, err = calc.SquareRoot(-4.0)
	if err != nil {
		fmt.Printf("Error en raíz cuadrada: %v\n", err)
	} else {
		fmt.Printf("√(-4) = %.4f\n", sqrt)
	}

	// Conversión de string exitosa
	doubled, err := parseAndDouble("25")
	if err != nil {
		fmt.Printf("Error en conversión: %v\n", err)
	} else {
		fmt.Printf("'25' × 2 = %d\n", doubled)
	}

	// Conversión de string inválida (error)
	doubled, err = parseAndDouble("abc")
	if err != nil {
		fmt.Printf("Error en conversión: %v\n", err)
	} else {
		fmt.Printf("'abc' × 2 = %d\n", doubled)
	}

	fmt.Println()
}

// =============== CONCURRENCIA BÁSICA ===============

// concurrencyExample demuestra concurrencia básica con goroutines
func concurrencyExample() {
	fmt.Println("=== CONCURRENCIA BÁSICA ===")
	fmt.Println("Nota: Este es un ejemplo muy básico. Go tiene capacidades de concurrencia mucho más avanzadas.")

	// Canal para comunicación
	done := make(chan bool)

	// Goroutine que cuenta números
	go func() {
		fmt.Print("Contando en goroutine: ")
		for i := 1; i <= 5; i++ {
			fmt.Printf("%d ", i)
		}
		fmt.Println()
		done <- true
	}()

	// Trabajo en la goroutine principal
	fmt.Print("Contando en main: ")
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// Esperar a que termine la goroutine
	<-done
	fmt.Println("¡Ambos conteos completados!")
	fmt.Println()
}
