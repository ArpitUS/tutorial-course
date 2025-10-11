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
