package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Student representa un estudiante (struct con tags JSON)
type Student struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Age      int      `json:"age"`
	Courses  []string `json:"courses"`
	GPA      float64  `json:"gpa"`
	IsActive bool     `json:"is_active"`
}

// Course representa un curso
type Course struct {
	Code       string    `json:"code"`
	Title      string    `json:"title"`
	Credits    int       `json:"credits"`
	Students   []Student `json:"students"`
	Instructor string    `json:"instructor"`
}

func main() {
	fmt.Println("=== EJEMPLO DE JSON EN GO ===")
	fmt.Println()

	// Crear estudiantes
	students := []Student{
		{
			ID:       1,
			Name:     "Ana Pérez",
			Age:      20,
			Courses:  []string{"PROG101", "MATH201", "PHYS101"},
			GPA:      3.8,
			IsActive: true,
		},
		{
			ID:       2,
			Name:     "Carlos López",
			Age:      22,
			Courses:  []string{"PROG101", "MATH201"},
			GPA:      3.5,
			IsActive: true,
		},
		{
			ID:       3,
			Name:     "María García",
			Age:      19,
			Courses:  []string{"PROG101", "PHYS101"},
			GPA:      3.9,
			IsActive: false,
		},
	}

	// Crear un curso
	course := Course{
		Code:       "PROG101",
		Title:      "Introducción a la Programación",
		Credits:    3,
		Students:   students,
		Instructor: "Dr. Santiago Avella",
	}

	// Marshaling: struct a JSON
	fmt.Println("=== MARSHALING (Struct -> JSON) ===")

	// JSON básico
	studentJSON, err := json.Marshal(students[0])
	if err != nil {
		log.Fatalf("Error al convertir estudiante a JSON: %v", err)
	}
	fmt.Printf("Estudiante en JSON (compacto):\n%s\n\n", studentJSON)

	// JSON con indentación
	studentJSONPretty, err := json.MarshalIndent(students[0], "", "  ")
	if err != nil {
		log.Fatalf("Error al convertir estudiante a JSON: %v", err)
	}
	fmt.Printf("Estudiante en JSON (formato bonito):\n%s\n\n", studentJSONPretty)

	// Curso completo con indentación
	courseJSON, err := json.MarshalIndent(course, "", "  ")
	if err != nil {
		log.Fatalf("Error al convertir curso a JSON: %v", err)
	}
	fmt.Printf("Curso completo en JSON:\n%s\n\n", courseJSON)

	// Unmarshaling: JSON a struct
	fmt.Println("=== UNMARSHALING (JSON -> Struct) ===")

	// JSON de ejemplo como string
	jsonString := `{
		"id": 4,
		"name": "Pedro Ramírez",
		"age": 21,
		"courses": ["MATH201", "PHYS101", "CHEM101"],
		"gpa": 3.7,
		"is_active": true
	}`

	// Convertir JSON a struct
	var newStudent Student
	err = json.Unmarshal([]byte(jsonString), &newStudent)
	if err != nil {
		log.Fatalf("Error al convertir JSON a struct: %v", err)
	}

	fmt.Printf("Nuevo estudiante desde JSON:\n")
	fmt.Printf("ID: %d\n", newStudent.ID)
	fmt.Printf("Nombre: %s\n", newStudent.Name)
	fmt.Printf("Edad: %d años\n", newStudent.Age)
	fmt.Printf("Cursos: %v\n", newStudent.Courses)
	fmt.Printf("GPA: %.1f\n", newStudent.GPA)
	fmt.Printf("Activo: %t\n", newStudent.IsActive)
	fmt.Println()

	// Trabajar con maps para JSON dinámico
	fmt.Println("=== JSON DINÁMICO CON MAPS ===")

	var dynamicData map[string]interface{}
	err = json.Unmarshal([]byte(jsonString), &dynamicData)
	if err != nil {
		log.Fatalf("Error al convertir JSON a map: %v", err)
	}

	fmt.Println("Datos dinámicos desde JSON:")
	for key, value := range dynamicData {
		fmt.Printf("%s: %v (tipo: %T)\n", key, value, value)
	}
	fmt.Println()

	// Array de estudiantes en JSON
	fmt.Println("=== ARRAY DE ESTUDIANTES ===")
	studentsJSON, err := json.MarshalIndent(students, "", "  ")
	if err != nil {
		log.Fatalf("Error al convertir estudiantes a JSON: %v", err)
	}
	fmt.Printf("Todos los estudiantes:\n%s\n", studentsJSON)

	fmt.Println("=== ¡Ejemplos de JSON completados! ===")
}
