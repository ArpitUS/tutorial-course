# Tutorial Course Repository

This repository contains structured, hands-on programming tutorials for three languages:

- Go (Golang)
- Python
- Rust

Each language course is organized into ten modules, progressing from beginner to intermediate level.  
Every module includes:

- A `README.md` explaining the topic and objectives  
- A starter file (`main_before` / `before_main`) for practice  
- A reference file (`main_after` / `after_main`) with the solution  

---

## Repository Overview

```text
tutorial-course/
│
├── Go_Tutorial/
│ ├── Module01/ … Module10/
│ └── README.md
│
├── Python_Tutorial/
│ ├── Module_01_Basics/ … Module_10_Automation/
│ └── README.md
│
├── Rust_Course/
│ ├── Module_01_Basics/ … Module_10_Project/
│ └── README.md
│
└── README.md
```

---

## Go Tutorial

**Goal:** Learn the fundamentals of the Go language — from setup to building concurrent CLI applications.

**Modules Overview:**

1. Hello, Go! Getting Started  
2. Variables, Constants, and Data Types  
3. Control Flow and Loops  
4. Functions  
5. Arrays, Slices, and Maps  
6. Structs and Interfaces  
7. Error Handling  
8. Concurrency with Goroutines  
9. Packages and Modules  
10. Mini Project – CLI Application  

**Example Run:**

```bash
cd Go_Tutorial/Module01
go run main_after.go
```

---

## Python Tutorial

### Goal

Build a solid understanding of Python for automation, scripting, and backend development.

### Modules Overview

1. Python Basics
2. Variables & Data Types
3. Control Flow
4. Functions
5. Data Structures
6. Strings & File Handling
7. Error Handling
8. Object-Oriented Programming
9. Modules & Packages
10. Automation & Mini Project

### Example Run

```bash
cd Rust_Course/Module_01_Basics
rustc after_main.rs && ./after_main
```

---

## Setup Instructions

### Prerequisites

| Language | Required Version | Installation |
|-----------|------------------|---------------|
| Go | 1.22+ | https://go.dev/dl/ |
| Python | 3.10+ | https://www.python.org/downloads/ |
| Rust | 1.80+ | https://rustup.rs/ |

### Recommended Editor

Visual Studio Code with:

- Go extension
- Python extension
- Rust Analyzer

---

## Learning Approach

Each module follows the same structure:

1. Read the `README.md` to understand the objectives
2. Open the starter file (`main_before` or `before_main.rs`)
3. Complete the exercises and test the output
4. Compare your code with the solution file (`main_after` or `after_main.rs`)

This format encourages practical, iterative learning and helps build muscle memory for syntax, structure, and problem-solving.

---

## Contributing

Contributions, improvements, and new tutorials are welcome.

### To contribute

#### Fork and clone this repository

```bash
git clone https://github.com/YourUsername/tutorial-course.git
```

#### Create a new branch

```bash
git checkout -b feature/new-tutorial
```

#### Commit your changes

```bash
git commit -m "feat: add new Go concurrency module examples"
```

#### Push and open a Pull Request

```bash
git push origin feature/new-tutorial
```
