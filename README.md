# Go Programming Course - 6 Month Curriculum

This repository contains the code and materials from a comprehensive 6-month Go programming course that I taught to students interested in learning the Go language. As the instructor, I designed a curriculum that progressed from basic programming concepts to advanced web application development, guiding students through hands-on exercises and projects.

## Course Structure

### Week 1: Go Fundamentals
- Basic data types: booleans, integers, and floats
- Variable declaration and initialization (`var` vs `:=`)
- Constants
- Composite types: arrays, slices, and maps
- Working with slices: append, make, capacity
- Structs and anonymous structs
- Control flow: if statements, conditionals

### Week 2: Functions, Pointers, and Methods
- Command-line arguments
- Formatting and printing with the `fmt` package
- Functions: declarations, parameters, and return values
- Pointers and memory management
- Methods and receiver functions
- Interfaces and type implementations
- Home assignment: Building a book listing and search application

### Week 3: Modules, Packages, and Concurrency
- Go modules and dependency management
- Package organization and visibility
- Concurrency basics
- Goroutines and parallel execution
- Synchronization with WaitGroups
- Worker pools for concurrent task processing
- Race conditions and synchronization primitives
- `sync.Once` for one-time initialization
- Home assignment: Enhanced book management system with concurrency

### Week 4: File Operations and Database Integration
- File handling and I/O operations
- Working with JSON and other file formats
- Introduction to databases
- GORM (Go Object Relational Mapper)
- Creating models and database migrations
- Basic CRUD operations
- Complex queries and relationships
- Home assignment: Migrating the book management system to use a database

### Week 5: Web Development and REST APIs
- HTTP client operations
- Building web servers in Go
- Introduction to the `net/http` package
- Using the Gorilla Mux router
- REST API principles:
  - Client-server architecture
  - Statelessness
  - Cacheability
  - Layered systems
  - Code-on-demand
  - Uniform interfaces
- API documentation with Swagger
- Home assignment: Creating a RESTful API for the book management system

### Week 6: Advanced Web Application Development
- Project structure and organization
- Separation of concerns
- Internal vs public packages
- Configuration management
- Middleware and authentication
- Error handling and logging
- API versioning
- Documentation
- Deployment considerations
- Final project: Complete bookstore API with authentication

## Getting Started

Each week's directory contains the code examples and exercises covered during that week. To run the code examples:

1. Make sure you have Go installed (version 1.16+ recommended)
2. Navigate to the specific week's directory
3. Run the example files with `go run filename.go`

For projects with multiple files, follow the instructions in the respective week's README or comments.

## Home Assignments

Throughout the course, students worked on a progressive book management system that evolved in complexity:

1. Week 2: Simple CLI book listing and searching
2. Week 3: Enhanced CLI with additional operations and concurrency
3. Week 4: Database integration with GORM
4. Week 5: RESTful API implementation
5. Week 6: Complete web application with authentication and documentation

## Resources

- [Official Go Documentation](https://golang.org/doc/)
- [Go by Example](https://gobyexample.com/)
- [GORM Documentation](https://gorm.io/docs/)
- [Gorilla Mux Router](https://github.com/gorilla/mux)
- [Swagger Documentation](https://swagger.io/docs/)