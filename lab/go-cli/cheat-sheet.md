# Go Cheat Sheet for C# Developers

## Core Data Structures

| C# | Go | Notes |
|---|---|---|
| `int`, `string`, `bool`, etc. | `int`, `string`, `bool`, etc. | Go has similar primitive types but with fewer variations |
| `var x = value` | `x := value` | Type inference in Go using `:=` (short declaration) |
| `null` | `nil` | Go's equivalent to null |
| `List<T>` | `[]Type` (slice) | Slices are Go's dynamic arrays |
| `Dictionary<K, V>` | `map[KeyType]ValueType` | Maps are Go's equivalent to dictionaries |
| `HashSet<T>` | No direct equivalent | Use `map[Type]bool` or `map[Type]struct{}` |
| `Tuple<T1, T2>` | No direct equivalent | Use a struct with unnamed or named fields |
| `class Person { ... }` | `type Person struct { ... }` | Go uses structs instead of classes |
| `enum` | `const` and `iota` | Go uses constants with iota for enum-like constructs |
| `interface IService { ... }` | `type Service interface { ... }` | Go interfaces are implicitly implemented |

### Slices (Go's Dynamic Arrays)

```go
// Declaration and initialization
var names []string                 // Declare empty slice
names := []string{"Alice", "Bob"}  // Declare and initialize 
names := make([]string, 5)         // Create with initial capacity

// Operations
names = append(names, "Charlie")   // Add element (similar to Add() in C#)
slice := names[1:3]                // Slicing (getting a sub-slice)
copy(dest, source)                 // Copy elements

// Length and capacity
len(names)                         // Length (similar to Count in C#)
cap(names)                         // Capacity (internal array size)
```

### Maps (Go's Dictionaries)

```go
// Declaration and initialization
var users map[string]int                      // Declare nil map
users := make(map[string]int)                 // Initialize empty map
users := map[string]int{"Alice": 28, "Bob": 25} // Declare and initialize

// Operations
users["Charlie"] = 30                         // Add or update
age, exists := users["Alice"]                 // Check existence (like TryGetValue)
delete(users, "Bob")                          // Remove entry
```

### Custom Types and Structs

```go
// Define a struct (similar to a C# class)
type Person struct {
    ID        int
    FirstName string
    LastName  string
    Age       int
    Address   *Address      // Pointer to another struct
}

// Embedded structs (similar to inheritance, but composition)
type Employee struct {
    Person                  // Embedded struct (composition)
    EmployeeID string
    Department string
}

// Initialize
alice := Person{ID: 1, FirstName: "Alice", LastName: "Smith", Age: 30}
// or
var bob Person
bob.FirstName = "Bob"
```

## Pointers and Value Semantics

### C# vs Go Memory Model

| C# | Go | Notes |
|---|---|---|
| Reference types (classes) | Pointer types (`*Type`) | Both reference a memory location |
| Value types (structs) | Non-pointer types | Both copy the entire value when assigned |
| `ref` parameters | Pointer parameters (`*Type`) | Both allow modifying the original value |
| Boxing/unboxing | Type assertions | Converting between interface and concrete types |

### Pointers in Go

```go
// Declare a variable
x := 42

// Create a pointer to x
p := &x            // & operator creates a pointer (like ref in C#)

// Dereference a pointer
fmt.Println(*p)    // * operator accesses the value (42)

// Change value through pointer
*p = 100           // x is now 100

// Pointer to struct
type Person struct {
    Name string
    Age  int
}

person := Person{Name: "Alice", Age: 30}
pointerToPerson := &person

// Access fields through pointer (both work)
fmt.Println((*pointerToPerson).Name)  // Explicit dereference
fmt.Println(pointerToPerson.Name)     // Implicit dereference (syntactic sugar)
```

### When to Use Pointers in Go

#### Use Pointers When:

1. **You need to modify the passed value**
   ```go
   func incrementAge(p *Person) {
       p.Age++  // Modifies the original Person
   }
   ```

2. **The struct is large** (for performance)
   ```go
   // Efficient when Person has many fields
   func processPerson(p *Person) {
       // Work with p
   }
   ```

3. **To represent optional values** (nil pointers)
   ```go
   var optionalPerson *Person  // nil by default
   if condition {
       optionalPerson = &Person{Name: "Bob"}
   }
   ```

4. **Implementing interfaces with methods that modify the receiver**
   ```go
   type Counter interface {
       Increment()
       GetCount() int
   }

   type SimpleCounter struct {
       count int
   }

   // Must use pointer receiver to modify count
   func (c *SimpleCounter) Increment() {
       c.count++
   }

   func (c *SimpleCounter) GetCount() int {
       return c.count
   }
   ```

#### Use Values When:

1. **The type is small** (basic types, small structs)
   ```go
   func square(x int) int {
       return x * x
   }
   ```

2. **You want immutability**
   ```go
   func processPersonImmutably(p Person) {
       // Creates a local copy, original unchanged
       p.Age++
   }
   ```

3. **The type is already a reference type** (maps, slices, channels)
   ```go
   // No need for pointer - slice header is already a reference
   func addElement(slice []int, element int) {
       slice = append(slice, element)  // Might not work as expected!
   }

   // Better to return the new slice
   func addElement(slice []int, element int) []int {
       return append(slice, element)
   }
   ```

### Go vs C# Mental Model

**C# Mental Model**:

* Classes are reference types (heap-allocated, passed by reference)
* Structs are value types (stack-allocated by default, passed by value)
* Use `ref` or `out` to pass value types by reference

**Go Mental Model**:

* Everything is passed by value by default
* Even pointer variables are passed by value (their value is an address)
* Slices, maps, and channels contain internal pointers (reference-like behavior)
* Use `*Type` for pointer types to modify the original or save memory

### Common Gotchas for C# Developers

1. **Slice Append Behavior**
   ```go
   // This might not work as expected
   func addToSlice(slice []int, item int) {
       slice = append(slice, item)  // Only modifies local copy
   }

   // Instead, return the new slice
   func addToSlice(slice []int, item int) []int {
       return append(slice, item)
   }
   ```

2. **Method Receivers**
   ```go
   type Counter struct { count int }

   // Value receiver - modifications don't affect original
   func (c Counter) IncrementWrong() {
       c.count++  // Only affects local copy
   }

   // Pointer receiver - modifications affect original
   func (c *Counter) Increment() {
       c.count++  // Modifies original
   }
   ```

3. **Nil is not the same as C#'s null**
   ```go
   // This works! You can call methods on nil pointers
   var c *Counter = nil
   // If the method handles nil properly:
   func (c *Counter) SafeIncrement() {
       if c == nil {
           return
       }
       c.count++
   }
   ```

4. **Interface Values and nil**
   ```go
   // A nil pointer assigned to an interface is NOT a nil interface!
   var p *Person = nil
   var i interface{} = p

   fmt.Println(p == nil)  // true
   fmt.Println(i == nil)  // false (!!)
   ```

### Performance Considerations

1. **Stack vs Heap Allocation**
   * Go's compiler can escape small values to stack (optimization)
   * Use pointers for large structs to avoid copying
   * Benchmark before optimizing - Go's copy speed is often fast

2. **Cache-Friendly Code**
   * Value types can lead to better cache locality
   * Contiguous memory (arrays, slices) is faster than pointer chasing

3. **Method Call Performance**
   * Value receivers may cause copying (potentially expensive)
   * Pointer receivers avoid copying but may cause indirection

### Best Practices

1. **Be Consistent**
   * If a type has any pointer receiver methods, consider making all methods pointer receivers
   * Document whether functions take ownership of passed values

2. **Consider API Design**
   * Return new values instead of modifying inputs when appropriate
   * Use pointers for "output parameters" (like C#'s `out`)

3. **Avoid Premature Optimization**
   * Start with clear semantics
   * Use profiling to identify actual bottlenecks

4. **Mind the Garbage Collector**
   * Fewer allocations mean less GC pressure
   * Reuse objects when appropriate

## Object-Oriented Patterns in Go

### Constructor Pattern

```go
// C# constructor pattern:
// public Person(string name, int age) { ... }

// Go constructor pattern (factory function)
func NewPerson(firstName string, lastName string, age int) *Person {
    return &Person{
        FirstName: firstName,
        LastName:  lastName,
        Age:       age,
    }
}

// Usage
person := NewPerson("Alice", "Smith", 30)
```

### Methods (Function Receivers)

```go
// Value receiver (like instance method)
func (p Person) FullName() string {
    return p.FirstName + " " + p.LastName
}

// Pointer receiver (can modify the struct)
func (p *Person) SetAge(age int) {
    p.Age = age
}

// Usage
name := person.FullName()
person.SetAge(31)
```

### Interfaces

```go
// Define an interface
type Greeter interface {
    Greet() string
}

// Implement the interface 
// (implicit - no "implements" keyword needed)
func (p Person) Greet() string {
    return "Hello, my name is " + p.FullName()
}

// Interface usage
func SayHello(g Greeter) {
    fmt.Println(g.Greet())
}

// Use it
SayHello(person)  // Person implicitly implements Greeter
```

### Composition over Inheritance

```go
// In C#: class Manager : Employee { ... }

// In Go, use embedding for composition
type Address struct {
    Street  string
    City    string
    Country string
}

type Person struct {
    Name    string
    Address Address  // Composition by field
}

// Or embed directly for "inheritance-like" behavior
type Employee struct {
    Person           // Embedded (fields and methods promoted)
    EmployeeID string
}

// Usage - direct access to embedded fields
employee := Employee{}
employee.Name = "Alice"  // Directly access Person's field
```

### Functional Options Pattern (for flexible constructors)

```go
// Options type
type PersonOption func(*Person)

// Option functions
func WithAge(age int) PersonOption {
    return func(p *Person) {
        p.Age = age
    }
}

func WithAddress(addr Address) PersonOption {
    return func(p *Person) {
        p.Address = &addr
    }
}

// Constructor with options
func NewPersonWithOptions(firstName, lastName string, opts ...PersonOption) *Person {
    p := &Person{
        FirstName: firstName,
        LastName:  lastName,
    }
    
    // Apply all options
    for _, opt := range opts {
        opt(p)
    }
    
    return p
}

// Usage
person := NewPersonWithOptions(
    "Alice", 
    "Smith",
    WithAge(30),
    WithAddress(Address{Street: "123 Main St"}),
)
```

## Concurrency and Parallelism

### Goroutines vs Tasks

```csharp
// C# async/await
async Task DoWorkAsync() {
    await Task.Delay(1000);
    Console.WriteLine("Work done!");
}

// Calling it
await DoWorkAsync();
```

```go
// Go goroutines
func doWork() {
    time.Sleep(time.Second)
    fmt.Println("Work done!")
}

// Calling it
go doWork()  // Fire and forget

// If you need to wait:
var wg sync.WaitGroup
wg.Add(1)
go func() {
    defer wg.Done()
    doWork()
}()
wg.Wait()
```

### Channels (Go's Communication Primitive)

```go
// Create channel
ch := make(chan string)     // Unbuffered channel
ch := make(chan string, 10) // Buffered channel (capacity 10)

// Send value to channel
ch <- "hello"  // Blocks if channel is full

// Receive from channel
msg := <-ch    // Blocks if channel is empty

// Close channel
close(ch)

// Range over channel (until closed)
for msg := range ch {
    fmt.Println(msg)
}

// Select (like switch for channels)
select {
case msg := <-ch1:
    // Use msg from ch1
case ch2 <- "response":
    // Sent to ch2
case <-time.After(time.Second):
    // Timeout after 1 second
default:
    // Non-blocking option
}
```

### Context (for cancellation and timeouts)

```go
// Create context with cancellation
ctx, cancel := context.WithCancel(context.Background())
defer cancel()  // Always call cancel to release resources

// Create context with timeout
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

// Pass context to function
func doWorkWithContext(ctx context.Context) error {
    select {
    case <-ctx.Done():
        return ctx.Err() // Context cancelled or timed out
    case <-time.After(2 * time.Second):
        return nil // Work completed
    }
}
```

### Common Concurrency Patterns

#### Worker Pool

```go
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
    defer wg.Done()
    for job := range jobs {
        // Process job
        result := job * 2
        results <- result
    }
}

// Create pools
jobs := make(chan int, 100)
results := make(chan int, 100)

// Start workers
var wg sync.WaitGroup
for w := 1; w <= 3; w++ {
    wg.Add(1)
    go worker(w, jobs, results, &wg)
}

// Send jobs
for j := 1; j <= 9; j++ {
    jobs <- j
}
close(jobs)

// Wait for all workers
go func() {
    wg.Wait()
    close(results)
}()

// Collect results
for result := range results {
    fmt.Println(result)
}
```

#### Fan-out, Fan-in

```go
func fanOut(input <-chan int, n int) []<-chan int {
    outputs := make([]<-chan int, n)
    for i := 0; i < n; i++ {
        outputs[i] = processData(input)
    }
    return outputs
}

func fanIn(inputs ...<-chan int) <-chan int {
    output := make(chan int)
    var wg sync.WaitGroup
    
    for _, input := range inputs {
        wg.Add(1)
        go func(ch <-chan int) {
            defer wg.Done()
            for val := range ch {
                output <- val
            }
        }(input)
    }
    
    go func() {
        wg.Wait()
        close(output)
    }()
    
    return output
}
```

### sync Package (Mutex, WaitGroup, Once)

```go
// Mutex for protecting shared data
var (
    mu      sync.Mutex
    counter int
)

func incrementCounter() {
    mu.Lock()
    defer mu.Unlock()
    counter++
}

// RWMutex for read-heavy workloads
var (
    rwMu  sync.RWMutex
    cache map[string]string
)

// Read operation
func getFromCache(key string) string {
    rwMu.RLock()
    defer rwMu.RUnlock()
    return cache[key]
}

// Write operation
func updateCache(key, value string) {
    rwMu.Lock()
    defer rwMu.Unlock()
    cache[key] = value
}

// Once for one-time initialization
var (
    instance *Singleton
    once     sync.Once
)

func GetInstance() *Singleton {
    once.Do(func() {
        instance = &Singleton{}
    })
    return instance
}
```

## Common C# to Go Translations

| C# Pattern | Go Pattern |
|---|---|---|
| `class` with properties | `struct` with exported fields |
| Constructor | Factory function (`NewXxx`) |
| Inheritance | Embedding (composition) |
| Properties with getters/setters | Public fields or methods |
| Interfaces (explicit) | Interfaces (implicit) |
| Exception handling | Error checking and propagation |
| Async/await | Goroutines and channels |
| LINQ | Slice operations and functions |
| Events | Channels or callbacks |
| Delegates | Function types |
| Generics | Built-in generic types (Go 1.18+) |
| Extension methods | Not available (use regular functions) |
| Using statement | defer statement |
| Attributes | struct tags |

## Key Differences to Remember

1. Go uses **composition over inheritance**
2. Go has no classes, only structs and interfaces
3. Go interfaces are implemented implicitly
4. Go has no exceptions, uses error values instead
5. Go concurrency is based on CSP (Communicating Sequential Processes)
6. Goroutines are lightweight compared to threads or Tasks
7. Channels are the primary communication mechanism
8. Go uses pointers but has no pointer arithmetic
9. Go has garbage collection but no constructors/destructors
10. Go favors explicit over implicit behavior
