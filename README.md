# Ideal Fishstick - Estructura Básica de Go

Un proyecto de introducción a Go que demuestra la sintaxis básica, funciones, métodos, estructuras, interfaces, manejo de errores y más conceptos fundamentales del lenguaje.

## 🚀 Estructura del Proyecto

```
ideal-fishstick/
├── main.go              # Programa principal con ejemplos completos
├── go.mod               # Módulo de Go
├── utils/               # Paquete de utilidades personalizado
│   └── utils.go         # Utilidades matemáticas, string y tiempo
├── examples/            # Ejemplos específicos
│   ├── utils_example.go # Ejemplo usando el paquete utils
│   └── json_example.go  # Ejemplo de manejo de JSON
└── README.md           # Este archivo
```

## 📚 Conceptos Demostrados

### 1. **Sintaxis Básica**
- Declaración de variables (`var`, `:=`)
- Constantes (`const`)
- Tipos de datos (int, string, float64, bool)
- Arrays y slices
- Maps
- Estructuras de control (if/else, switch, for, range)

### 2. **Funciones**
- Funciones básicas
- Múltiples valores de retorno
- Valores de retorno nombrados
- Funciones variádicas (`...`)
- Funciones anónimas
- Closures

### 3. **Estructuras y Métodos**
- Definición de structs
- Métodos con receptores por valor
- Métodos con receptores por puntero
- Composición de structs
- Método String() (implementación de Stringer)

### 4. **Interfaces**
- Definición de interfaces
- Implementación implícita
- Polimorfismo
- Interface{} para tipos dinámicos

### 5. **Manejo de Errores**
- Patrón `error` estándar de Go
- Creación de errores personalizados
- Wrapping de errores con `fmt.Errorf`
- Propagación y manejo de errores

### 6. **Concurrencia Básica**
- Goroutines
- Channels
- Comunicación entre goroutines

### 7. **Organización de Paquetes**
- Creación de paquetes personalizados
- Funciones y tipos exportados/no exportados
- Importación de paquetes locales

### 8. **Manejo de JSON**
- Marshaling (struct → JSON)
- Unmarshaling (JSON → struct)
- Tags de struct
- JSON dinámico con maps

## 🏃‍♂️ Cómo Ejecutar

### Prerrequisitos
- Go 1.19 o superior instalado
- Terminal o línea de comandos

### Ejecutar el programa principal
```bash
# Clonar el repositorio
git clone https://github.com/santiagoAvellaR/ideal-fishstick.git
cd ideal-fishstick

# Ejecutar todos los ejemplos principales
go run main.go
```

### Ejecutar ejemplos específicos

#### Ejemplo del paquete utils
```bash
go run examples/utils_example.go
```

#### Ejemplo de JSON
```bash
go run examples/json_example.go
```

### Compilar programas
```bash
# Compilar programa principal
go build -o main main.go
./main

# Compilar ejemplos
go build -o utils_example examples/utils_example.go
go build -o json_example examples/json_example.go
```

## 📖 Qué Aprenderás

Al explorar este proyecto aprenderás:

1. **Sintaxis fundamental de Go**: variables, constantes, tipos de datos
2. **Funciones**: cómo definir y usar diferentes tipos de funciones
3. **Estructuras de datos**: arrays, slices, maps, structs
4. **Programación orientada a objetos**: métodos, interfaces, polimorfismo
5. **Manejo de errores**: el patrón idiomático de Go para errores
6. **Concurrencia básica**: goroutines y channels
7. **Organización de código**: paquetes y módulos
8. **Serialización**: trabajar con JSON
9. **Buenas prácticas**: convenciones de nomenclatura, documentación

## 🔍 Ejemplos Destacados

### Función con Múltiples Retornos
```go
func divide(dividend, divisor int) (int, int) {
    quotient := dividend / divisor
    remainder := dividend % divisor
    return quotient, remainder
}
```

### Struct con Métodos
```go
type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}
```

### Manejo de Errores
```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("no se puede dividir %g por cero", a)
    }
    return a / b, nil
}
```

### Interface
```go
type Shape interface {
    Area() float64
    Perimeter() float64
}
```

## 🎯 Próximos Pasos

Después de dominar estos conceptos básicos, puedes explorar:

- **Testing**: Escribir tests con `testing`
- **HTTP**: Crear servidores web con `net/http`
- **Bases de datos**: Conectar con SQL databases
- **Concurrencia avanzada**: Select statements, worker pools
- **Módulos**: Gestión avanzada de dependencias
- **Deployment**: Dockerización y deployment

## 📝 Notas Adicionales

- Go es un lenguaje compilado y tipado estáticamente
- No hay clases, solo structs con métodos
- Las interfaces se implementan implícitamente
- Go favorece la composición sobre la herencia
- La concurrencia es un ciudadano de primera clase en Go
- El manejo de errores es explícito (sin excepciones)

## 🤝 Contribuciones

¡Las contribuciones son bienvenidas! Si tienes ejemplos adicionales o mejoras, no dudes en hacer un pull request.

---

**¡Disfruta aprendiendo Go! 🐹**
