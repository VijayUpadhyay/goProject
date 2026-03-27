# Go Learning Project

A comprehensive collection of Go programming examples organized by topic. Each directory contains a single `main.go` file with multiple examples demonstrating different concepts.

## 📁 Project Structure

```
goProject/
├── arrays/              # Array operations and capacity
├── bytes/               # Byte, string, and rune operations
├── channel/             # Channel operations (11 examples)
├── closure/             # Closure examples
├── collections/         # Map operations
├── compare/             # Comparison operations
├── conditional/         # If statements, loops, switches (4 examples)
├── datastructures/      # Arrays, slices, maps, structs (12 examples)
├── datatypes/           # Variable declarations and type conversions (4 examples)
├── errorhandling/       # Defer, panic, recover (3 examples)
├── formatters/          # String formatting
├── functions/           # Function values, closures, variadic (5 examples)
├── generics/            # Generic programming
├── goroutines/          # Goroutine examples (6 examples)
├── Images/              # Image processing (3 examples)
├── Interfaces/          # Interface implementations (15 examples)
├── interview/           # Interview questions (2 examples)
├── io/                  # File I/O operations (6 examples)
├── loops/               # Loop constructs
├── math/                # Math operations
├── methods/             # Methods and receivers (4 examples)
├── pointers/            # Pointer operations (2 examples)
├── programs/            # Algorithms (factorial, fibonacci, etc.)
├── refactor/            # Refactoring examples (13 examples)
├── reflect/             # Reflection
├── regex/               # Regular expressions
├── select/              # Select statements
├── server/              # Server examples
├── slices/              # Slice operations (16 examples)
├── sort/                # Sorting examples
├── string/              # String operations (6 examples)
├── struct/              # Struct examples (4 examples)
├── switch/              # Switch statements
├── syncChecks/          # Synchronization primitives
├── time/                # Time operations
├── udemy/               # Udemy course examples (2 examples)
├── user_input/          # User input handling
├── waitgroup/           # WaitGroup examples
├── _projects/           # Larger projects
│   ├── BIUtils/         # Business Intelligence utilities
│   └── businessInsights/# Business insights parser
└── tools/               # Consolidation scripts
```

## 🚀 Getting Started

### Prerequisites
- Go 1.21 or higher

### Running Examples

Each directory contains a `main.go` file with multiple examples. To run all examples in a directory:

```bash
# Run all interface examples
cd Interfaces
go run main.go

# Run all channel examples
cd channel
go run main.go

# Run all goroutine examples
cd goroutines
go run main.go
```

### Running Individual Examples

Each consolidated `main.go` file contains numbered example functions. You can modify the main function to run specific examples:

```go
func main() {
    // Run only specific examples
    example1_basicInterface()
    example5_StringerInterface()
}
```

## 📚 Topics Covered

### Core Concepts
- ✅ Variables and data types
- ✅ Arrays, slices, and maps
- ✅ Structs and methods
- ✅ Pointers
- ✅ Functions and closures

### Advanced Concepts
- ✅ Interfaces (15 examples)
- ✅ Goroutines and concurrency (6 examples)
- ✅ Channels (11 examples)
- ✅ Error handling
- ✅ Generics
- ✅ Reflection

### Practical Skills
- ✅ File I/O operations
- ✅ JSON marshaling/unmarshaling
- ✅ String manipulation
- ✅ Regular expressions
- ✅ Synchronization primitives

## 🔧 Project Organization

This project was recently restructured to follow Go best practices:
- **One main function per directory**: Each directory now contains a single `main.go` file
- **Multiple examples per file**: Related examples are organized as separate functions
- **Clear naming**: Example functions are named `example1_description`, `example2_description`, etc.
- **Modular structure**: Easy to run all examples or select specific ones

## 📖 Learning Path

Recommended order for beginners:

1. **Basics**: `datatypes/` → `conditional/` → `loops/`
2. **Data Structures**: `arrays/` → `slices/` → `collections/` → `datastructures/`
3. **Functions**: `functions/` → `closure/`
4. **OOP Concepts**: `struct/` → `methods/` → `Interfaces/`
5. **Concurrency**: `goroutines/` → `channel/` → `waitgroup/` → `syncChecks/`
6. **Advanced**: `generics/` → `reflect/` → `errorhandling/`
7. **Practical**: `io/` → `string/` → `regex/` → `programs/`

## 🛠️ Tools

### Consolidation Script
The `tools/consolidate.sh` script was used to restructure the project:
```bash
cd tools
./consolidate.sh ..
```

This script:
- Finds all directories with multiple Go files
- Consolidates them into a single `main.go`
- Renames each original `main()` function to `exampleN_filename()`
- Creates a new `main()` that calls all examples
- Removes the original files

## 📝 Code Quality

### Current Status
- ✅ Organized structure (one main per directory)
- ✅ Go modules initialized (`go.mod`)
- ⚠️ Some deprecated code (e.g., `io/ioutil`)
- ⚠️ Missing tests
- ⚠️ Some hardcoded paths

### Improvements Needed
1. Replace deprecated `io/ioutil` with `os` package
2. Add unit tests for all examples
3. Remove hardcoded file paths
4. Add proper error handling
5. Implement context for long-running operations
6. Add documentation comments

## 🎯 Next Steps

1. **Add Testing**: Create `*_test.go` files for each directory
2. **Fix Deprecations**: Update to use current Go APIs
3. **Add HTTP Examples**: Create REST API examples
4. **Database Integration**: Add SQL/NoSQL examples
5. **Add Logging**: Implement structured logging
6. **Context Usage**: Add context examples for cancellation

## 📄 License

This is a personal learning project.

## 🤝 Contributing

This is a personal learning repository. Feel free to fork and use for your own learning!

## 📚 Resources

- [Official Go Documentation](https://go.dev/doc/)
- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go Tour](https://go.dev/tour/)

---

**Note**: This project is continuously evolving as I learn more Go concepts and best practices.
