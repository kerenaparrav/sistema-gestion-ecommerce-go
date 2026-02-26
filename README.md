Sistema de Gestión de E-commerce en Go
📌 Descripción

Este proyecto consiste en el desarrollo de una API RESTful para la gestión de productos en un sistema de comercio electrónico, implementada en Go (Golang).

El sistema aplica arquitectura por capas, principios de Programación Orientada a Objetos adaptados a Go y concurrencia mediante goroutines, permitiendo gestionar productos a través de servicios web en formato JSON.

Objetivo

Aplicar buenas prácticas de desarrollo backend utilizando:

Arquitectura modular por capas

Structs e interfaces en Go

Servicios web con net/http

Serialización JSON

Concurrencia con goroutines y sync.WaitGroup

Control de versiones con Git y GitHub

Arquitectura

El sistema está organizado en las siguientes capas:

Domain: Define las estructuras principales (Product).

Repository: Maneja el almacenamiento en memoria.

Service: Contiene la lógica de negocio.

Cmd/API: Configura el servidor HTTP y los endpoints.

Endpoints principales

La API permite:

Verificar estado con /ping

Crear productos (POST /productos)

Listar productos (GET /productos)

Consultar producto por ID (GET /productos/{id})

Actualizar producto (PUT /productos/{id})

Eliminar producto (DELETE /productos/{id})

Demostrar concurrencia con /productos/cargar-demo

El endpoint de carga demo crea productos en paralelo utilizando goroutines.

Ejecución

Desde la carpeta src:

go run ./cmd/api

La API se ejecuta en:

http://localhost:8080

Tecnologías

Go (Golang)
net/http
JSON
Goroutines
Git & GitHub

Autor

Keren Parra
