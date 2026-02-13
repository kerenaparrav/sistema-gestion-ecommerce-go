Sistema de Gestión de E-commerce en Go

Descripción

Este proyecto consiste en el desarrollo de un sistema de gestión de productos para un e-commerce utilizando el lenguaje de programación Go (Golang).
El sistema funciona por consola y permite registrar productos, listarlos y validar los datos ingresados por el usuario.
Se aplicó una arquitectura por capas para organizar correctamente la lógica del programa.

Objetivo

Aplicar buenas prácticas de programación en Go mediante:
Entrada de datos por consola
Validación de errores
Separación de responsabilidades (arquitectura por capas)
Uso de módulos en Go
Control de versiones con Git y GitHub

Arquitectura del proyecto

El sistema está dividido en tres capas principales:

Domain

Contiene las estructuras principales del sistema.

Ejemplo: Product

Repository

Encargado del almacenamiento de datos en memoria.

Ejemplo:

Guardar productos

Obtener productos registrados

Service

Contiene la lógica del negocio.

Ejemplo:

Registrar productos

Listar productos

Cmd/App

Punto de entrada del programa (main.go) y menú interactivo.

Funcionalidades

El sistema permite:
Registrar productos
Listar productos registrados
Validar entradas del usuario
Manejo de errores en consola

Ejecución del programa

Ubicarse dentro de la carpeta src:

cd src
go run ./cmd/app

Ejemplo de uso
--- Menú Principal ---
1. Registrar producto
2. Listar productos
3. Salir

Ingrese ID: 1
Ingrese nombre: Arroz
Ingrese precio: 2.50

Producto registrado correctamente

Tecnologías utilizadas
Go (Golang)
Git
GitHub
Programación por capas

Autor
Keren Parra
