# 🧠 Curso de Introdução ao Desenvolvimento Backend com Go

Este repositório contém o material didático e os códigos desenvolvidos durante o curso **Introdução ao Desenvolvimento Backend com Golang**, ministrado no âmbito do programa **OxeTech**.

O curso tem como objetivo apresentar os fundamentos da linguagem **Go (Golang)** e conduzir o aluno à construção de **APIs e servidores backend completos**, passando por conceitos de concorrência, persistência, autenticação e boas práticas.

---

## 🎯 Objetivo do curso

Ao final do curso, o aluno será capaz de:

- Compreender os **fundamentos da linguagem Go**
- Estruturar projetos backend utilizando **`net/http`**, **Gin** e **GORM**
- Aplicar conceitos de **concorrência** e **comunicação entre goroutines**
- Realizar **operações CRUD completas** com banco de dados relacional
- Implementar **autenticação, middlewares e testes básicos**

---

## 🏗️ Estrutura do curso

O curso está dividido em **5 módulos**, totalizando **20 aulas** (2 horas cada).

---

### 🧩 **Módulo 1 — Fundamentos da Linguagem Go**  
**Aulas 1 a 4**

> Primeiros passos com Go: sintaxe, controle de fluxo, funções e coleções.

| Aula | Tema | Conteúdo |
|------|------|----------|
| 1 | Apresentação e ambiente | Instalação, configuração, Hello World |
| 2 | Variáveis e controle de fluxo | Tipos, `if`, `switch`, operadores |
| 3 | Loops e funções | `for`, `range`, funções com retorno |
| 4 | Coleções | Arrays, slices, maps, `make`, `append` |

---

### 🧱 **Módulo 2 — Structs, Ponteiros e Abstração**  
**Aulas 5 a 8**

> Criação de tipos compostos, uso de métodos, interfaces e concorrência com goroutines.

| Aula | Tema | Conteúdo |
|------|------|----------|
| 5 | Structs e composição | Structs aninhadas, modelagem de dados |
| 6 | Ponteiros e métodos | Passagem por referência e receptores |
| 7 | Interfaces e polimorfismo | Contratos, `interface{}`, type switch |
| 8 | Concorrência com Goroutines | `go func()`, `WaitGroup`, channels, select |

---

### 🌐 **Módulo 3 — Introdução ao Backend com Go**  
**Aulas 9 a 11**

> Construindo o primeiro servidor web com Go e entendendo o ciclo de uma API.

| Aula | Tema | Conteúdo |
|------|------|----------|
| 9 | Fundamentos de backend e HTTP | O que é backend, `net/http`, rotas |
| 10 | JSON e CRUD em memória | `encoding/json`, status codes, handlers |
| 11 | Persistência com SQL direto | `database/sql`, SQLite/Postgres, queries |

---

### ⚙️ **Módulo 4 — Framework Gin e ORM GORM**  
**Aulas 12 a 15**

> Estruturando APIs reais com framework, ORM e boas práticas de validação.

| Aula | Tema | Conteúdo |
|------|------|----------|
| 12 | Framework Gin + SQL | Rotas, middlewares e binding automático |
| 13 | ORM com GORM | Migrations, CRUD, filtros e paginação |
| 14 | Validações e erros | Binding com validação, erros estruturados |
| 15 | Relacionamentos e joins | `hasOne`, `hasMany`, `belongsTo`, preload |

---

### 🔐 **Módulo 5 — Middlewares, Autenticação e Extras**  
**Aulas 16 a 20**

> Implementando autenticação, autorização e testes automatizados no backend Go.

| Aula | Tema | Conteúdo |
|------|------|----------|
| 16 | Middlewares e autenticação | JWT, controle de acesso e logging |
| 17 | Organização do projeto | Estrutura de pacotes e boas práticas |
| 18 | Testes em Go | `testing`, mocks e `httptest` |
| 19 | Deploy e ambiente | `.env`, build, Docker e Nginx |
| 20 | Encerramento e projeto final | Mini API completa e revisão geral |

---

## 💻 Pré-requisitos

- Conhecimentos básicos em lógica de programação  
- Noções gerais de HTTP e JSON  
- **Go 1.22+** instalado ([https://go.dev/dl/](https://go.dev/dl/))  
- Editor de código (VSCode recomendado)
--- 

Cada pasta contém exemplos e exercícios correspondentes às aulas do módulo.

---

## 🚀 Projeto Final

Ao final do curso, os alunos desenvolvem uma **API REST completa** em Go, com:

- CRUD de usuários ou produtos  
- Persistência com GORM  
- Middleware de autenticação (JWT)  
- Estrutura modularizada (`handlers`, `services`, `models`, `routes`)  
- Testes básicos e deploy containerizado  

---

## 📚 Licença

Este projeto está licenciado sob a **MIT License**.  
Sinta-se livre para utilizar, modificar e compartilhar o conteúdo para fins educacionais.

---

### 👨‍🏫 Instrutor
**Ulpio Netto**  
Instrutor Backend — Programa OxeTech  
[LinkedIn](https://linkedin.com/in/ulpionetto) · [GitHub](https://github.com/Ulpio)

---

> 💡 *O Go é simples, direto e poderoso — e agora você vai ver o porquê ele é uma das linguagens mais queridas no mundo do backend.*
