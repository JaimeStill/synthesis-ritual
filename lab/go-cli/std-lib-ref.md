# Go Standard Library Quick Reference

## Core Packages

### `fmt` - Formatted I/O

```go
fmt.Println("Hello, world")             // Print with newline
fmt.Printf("Value: %v, Type: %T\n", x, x) // Formatted print
fmt.Sprintf("Value: %v", x)             // Return formatted string

// Common format verbs
// %v - default value format
// %+v - struct with field names
// %#v - Go syntax representation
// %T - type
// %d - decimal integer
// %s - string
// %f - float
// %t - boolean
// %p - pointer
```

### `strings` - String Manipulation

```go
strings.Contains(s, substr)          // Check if string contains substring
strings.HasPrefix(s, prefix)         // Check if string starts with prefix
strings.HasSuffix(s, suffix)         // Check if string ends with suffix
strings.Index(s, substr)             // Index of first instance (-1 if not found)
strings.Join([]string, sep)          // Join string slice with separator
strings.Split(s, sep)                // Split string into slice by separator
strings.ToLower(s)                   // Convert to lowercase
strings.ToUpper(s)                   // Convert to uppercase
strings.TrimSpace(s)                 // Trim whitespace
strings.ReplaceAll(s, old, new)      // Replace all occurrences
```

### `strconv` - String Conversions

```go
strconv.Itoa(i)                      // Convert int to string
strconv.Atoi(s)                      // Convert string to int (with error)
strconv.ParseFloat(s, bitSize)       // Parse string to float64
strconv.ParseBool(s)                 // Parse string to bool
strconv.FormatFloat(f, fmt, prec, bitSize) // Format float to string
```

### `errors` - Error Handling

```go
errors.New("error message")          // Create new error
errors.Is(err, targetErr)            // Check if err or its chain is targetErr
errors.As(err, &target)              // Check if err or chain matches target type
fmt.Errorf("error: %w", err)         // Wrap error with additional context
```

### `reflect` - Reflection

```go
reflect.TypeOf(x)                   // Get type of x
reflect.ValueOf(x)                  // Get reflect.Value of x
v.Kind()                            // Get underlying kind (int, struct, etc.)
v.NumField()                        // Number of fields in struct
v.Field(i)                          // Get i-th field
```

## I/O and Files

### `io` - Basic I/O Interfaces

```go
io.Copy(dst, src)                   // Copy from src to dst
io.ReadAll(r)                       // Read until EOF
io.WriteString(w, s)                // Write string to writer
io.MultiReader(readers...)          // Combine multiple readers
io.MultiWriter(writers...)          // Write to multiple writers
```

### `os` - Operating System Functionality

```go
os.Open(filename)                   // Open file for reading
os.Create(filename)                 // Create or truncate file
os.OpenFile(name, flag, perm)       // Open with options
os.Remove(filename)                 // Delete file
os.Mkdir(name, perm)                // Create directory
os.MkdirAll(path, perm)             // Create directory with parents
os.Getwd()                          // Get working directory
os.Chdir(dir)                       // Change directory
os.Environ()                        // Environment variables as slice
os.Getenv("NAME")                   // Get environment variable
os.Args                             // Command line arguments
```

### `path/filepath` - File Path Manipulation

```go
filepath.Join(elem...)              // Join path elements
filepath.Abs(path)                  // Absolute path
filepath.Base(path)                 // Last element of path
filepath.Dir(path)                  // Directory portion
filepath.Ext(path)                  // File extension
filepath.IsAbs(path)                // Check if path is absolute
filepath.Walk(root, fn)             // Walk directory tree
filepath.Glob(pattern)              // File name pattern matching
```

### `bufio` - Buffered I/O

```go
bufio.NewReader(r)                  // Create buffered reader
reader.ReadLine()                   // Read line
reader.ReadString(delim)            // Read until delimiter
reader.ReadBytes(delim)             // Read until delimiter as bytes
bufio.NewScanner(r)                 // Create scanner for tokens
scanner.Scan()                      // Advance to next token
scanner.Text()                      // Get current token as string
scanner.Bytes()                     // Get current token as bytes
bufio.NewWriter(w)                  // Create buffered writer
writer.WriteString(s)               // Write string
writer.Flush()                      // Flush buffered data
```

## Data Structures and Algorithms

### `sort` - Sorting

```go
sort.Ints(a)                        // Sort int slice
sort.Strings(a)                     // Sort string slice
sort.Float64s(a)                    // Sort float64 slice
sort.Slice(slice, less func(i, j int) bool) // Sort with custom comparator
sort.SliceStable(slice, less func(i, j int) bool) // Stable sort with custom comparator
sort.Sort(data interface{})         // Sort any type implementing sort.Interface
```

### `container/list` - Doubly Linked List

```go
l := list.New()                     // Create new list
l.PushBack(value)                   // Add to back
l.PushFront(value)                  // Add to front
l.Remove(element)                   // Remove element
l.Len()                             // Get length

// Iterate over list
for e := l.Front(); e != nil; e = e.Next() {
    // e.Value is the value
}
```

### `container/heap` - Heap (Priority Queue)

```go
// Implement heap.Interface
type IntHeap []int
func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

// Usage
h := &IntHeap{2, 1, 5}
heap.Init(h)
heap.Push(h, 3)
min := heap.Pop(h).(int)
```

### `container/ring` - Circular List

```go
r := ring.New(5)                    // Create new ring of size 5
r.Value = 1                         // Set value
r = r.Next()                        // Move to next
r.Do(func(p interface{}){})         // Apply function to each element
```

## Concurrency

### `sync` - Synchronization Primitives

```go
var mu sync.Mutex
mu.Lock()                           // Lock mutex
mu.Unlock()                         // Unlock mutex

var rwmu sync.RWMutex               // Reader/writer mutex
rwmu.RLock()                        // Read lock
rwmu.RUnlock()                      // Release read lock
rwmu.Lock()                         // Write lock
rwmu.Unlock()                       // Release write lock

var wg sync.WaitGroup
wg.Add(delta)                       // Add delta to counter
wg.Done()                           // Decrement counter
wg.Wait()                           // Wait for counter to reach zero

var once sync.Once
once.Do(func(){})                   // Execute function exactly once

var cond *sync.Cond = sync.NewCond(&sync.Mutex{})
cond.Wait()                         // Release lock and wait
cond.Signal()                       // Wake one waiter
cond.Broadcast()                    // Wake all waiters

var pool sync.Pool
pool.Put(x)                         // Put object in pool
x := pool.Get()                     // Get object from pool

var m sync.Map
m.Store(key, value)                 // Store value
m.Load(key)                         // Load value
m.Delete(key)                       // Delete entry
m.LoadOrStore(key, value)           // Load or store if not exists
m.Range(func(key, value interface{}) bool {}) // Iterate
```

### `sync/atomic` - Atomic Operations

```go
atomic.AddInt64(&i, delta)          // Add delta atomically
atomic.CompareAndSwapInt64(&i, old, new) // CAS operation
atomic.LoadInt64(&i)                // Load atomically
atomic.StoreInt64(&i, v)            // Store atomically
```

### `context` - Request Context

```go
ctx := context.Background()             // Empty context
ctx := context.TODO()                   // Placeholder context
ctx, cancel := context.WithCancel(ctx)  // With cancellation
defer cancel()                          // Cancel when done
ctx, cancel := context.WithTimeout(ctx, time.Second) // With timeout
ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second)) // With deadline
ctx = context.WithValue(ctx, key, value) // With key-value
value := ctx.Value(key)                 // Get value from context
<-ctx.Done()                            // Channel closed on cancellation
ctx.Err()                               // Error on cancellation
```

## Networking

### `net/http` - HTTP Client and Server

```go
// Client
resp, err := http.Get(url)          // GET request
resp, err := http.Post(url, contentType, body) // POST request
resp, err := http.PostForm(url, values) // POST form
client := &http.Client{}
req, err := http.NewRequest("GET", url, nil) // Create request
req.Header.Add("User-Agent", "golang")
resp, err := client.Do(req)         // Send request
defer resp.Body.Close()             // Always close response body
body, err := io.ReadAll(resp.Body)  // Read response body

// Server
http.HandleFunc("/", handler)       // Register handler
http.Handle("/", handler)           // Register Handler interface
http.ListenAndServe(":8080", nil)   // Start server
http.ListenAndServeTLS(addr, certFile, keyFile, nil) // Start HTTPS server

func handler(w http.ResponseWriter, r *http.Request) {
    r.URL.Path                      // Request path
    r.Method                        // HTTP method
    r.Header                        // Request headers
    r.FormValue("key")              // Form value
    r.ParseForm()                   // Parse form
    r.Form                          // Form values
    r.Body                          // Request body
    
    w.Header().Set("Content-Type", "application/json") // Response header
    w.WriteHeader(http.StatusOK)    // Set status code
    fmt.Fprintf(w, "Hello, %s", name) // Write response
}
```

### `net` - Basic Network I/O

```go
conn, err := net.Dial("tcp", "localhost:8080") // Connect to server
listener, err := net.Listen("tcp", ":8080")    // Create listener
conn, err := listener.Accept()                 // Accept connection
```

### `net/url` - URL Parsing

```go
u, err := url.Parse("http://example.com/path?q=value") // Parse URL
u.Scheme                           // URL scheme (http)
u.Host                             // Host (example.com)
u.Path                             // Path (/path)
u.RawQuery                         // Raw query (q=value)
u.Query().Get("q")                 // Query parameter
```

## Encoding/Decoding

### `encoding/json` - JSON Processing

```go
// Struct tags
type Person struct {
    Name    string `json:"name"`
    Age     int    `json:"age,omitempty"`
    Address string `json:"-"` // Ignore field
}

// Marshal (encode)
data, err := json.Marshal(v)       // Convert to JSON
data, err := json.MarshalIndent(v, prefix, indent) // Pretty print

// Unmarshal (decode)
err := json.Unmarshal(data, &v)    // Parse JSON into v

// Streaming
dec := json.NewDecoder(r)
err := dec.Decode(&v)              // Decode from reader
enc := json.NewEncoder(w)
err := enc.Encode(v)               // Encode to writer
```

### `encoding/xml` - XML Processing

```go
// Struct tags
type Person struct {
    Name    string `xml:"name"`
    Age     int    `xml:"age,attr"` // Attribute
    Address string `xml:",omitempty"`
}

// Marshal (encode)
data, err := xml.Marshal(v)        // Convert to XML
data, err := xml.MarshalIndent(v, prefix, indent) // Pretty print

// Unmarshal (decode)
err := xml.Unmarshal(data, &v)     // Parse XML into v

// Streaming
dec := xml.NewDecoder(r)
err := dec.Decode(&v)              // Decode from reader
enc := xml.NewEncoder(w)
err := enc.Encode(v)               // Encode to writer
```

### `encoding/csv` - CSV Processing

```go
// Reading
r := csv.NewReader(file)
r.Comma = ';'                      // Set delimiter
r.Comment = '#'                    // Set comment character
record, err := r.Read()            // Read one record
records, err := r.ReadAll()        // Read all records

// Writing
w := csv.NewWriter(file)
w.Comma = ';'                      // Set delimiter
w.Write(record)                    // Write one record
w.WriteAll(records)                // Write all records
w.Flush()                          // Flush pending writes
```

### `encoding/base64` - Base64 Encoding

```go
str := base64.StdEncoding.EncodeToString(data) // Encode
data, err := base64.StdEncoding.DecodeString(str) // Decode
```

## Time and Date

### `time` - Time and Duration

```go
now := time.Now()                  // Current time
t := time.Date(2023, time.January, 1, 12, 30, 0, 0, time.UTC) // Create time
t.Year(), t.Month(), t.Day()       // Get components
t.Hour(), t.Minute(), t.Second()   // Time components
t.Format("2006-01-02 15:04:05")    // Format time (use reference time!)
t.Format(time.RFC3339)             // Standard format
t, err := time.Parse("2006-01-02", "2023-01-01") // Parse time

duration := 5 * time.Second        // Create duration
time.Sleep(duration)               // Sleep for duration
t.Add(duration)                    // Add duration
t.Sub(t2)                          // Subtract times

timer := time.NewTimer(duration)   // Create timer
<-timer.C                          // Wait for timer
timer.Stop()                       // Stop timer

ticker := time.NewTicker(duration) // Create ticker
for t := range ticker.C {          // Receive ticks
    // t is the current time
}
ticker.Stop()                      // Stop ticker
```

## Testing

### `testing` - Unit Tests and Benchmarks

```go
// In *_test.go files:
func TestXxx(t *testing.T) {
    t.Log("Log message")           // Log message
    t.Logf("Formatted %s", "message") // Formatted log
    if got != want {
        t.Errorf("Got %v, want %v", got, want) // Report failure
    }
    t.Fail()                       // Mark test as failed
    t.FailNow()                    // Fail and stop immediately
    t.Skip("Skipping test")        // Skip test
    t.Run("SubTest", func(t *testing.T) { /* nested test */ })
}

// Table-driven tests
testCases := []struct {
    name   string
    input  int
    want   string
}{
    {"positive", 1, "one"},
    {"negative", -1, "minus one"},
}
for _, tc := range testCases {
    t.Run(tc.name, func(t *testing.T) {
        got := functionToTest(tc.input)
        if got != tc.want {
            t.Errorf("Got %v, want %v", got, tc.want)
        }
    })
}

// Benchmarks
func BenchmarkXxx(b *testing.B) {
    for i := 0; i < b.N; i++ {
        // Code to benchmark
    }
}
```

### `testing/quick` - Property-Based Testing

```go
f := func(a, b int) bool {
    // Property that should always hold
    return (a+b) == (b+a)
}
if err := quick.Check(f, nil); err != nil {
    t.Error(err)
}
```

## CLI and Environment

### `flag` - Command-Line Flag Parsing

```go
var (
    name = flag.String("name", "default", "help message")
    age  = flag.Int("age", 0, "help message")
    verbose = flag.Bool("verbose", false, "help message")
)

// Parse flags
flag.Parse()

// Access values
fmt.Println(*name, *age, *verbose)

// Remaining arguments
flag.Args()
```

### `log` - Logging

```go
log.Println("Standard logger")      // Standard logger
log.Printf("Formatted %s", "message") // Formatted logging
log.Fatal("Fatal error")            // Log and exit with status 1
log.Fatalf("Fatal %s", "error")     // Formatted fatal
log.Panic("Panic message")          // Log and panic

// Custom logger
logger := log.New(os.Stderr, "PREFIX: ", log.LstdFlags)
logger.Println("Custom logger")
```

## Utility Packages

### `math` - Mathematical Functions

```go
math.Abs(x)                         // Absolute value
math.Sqrt(x)                        // Square root
math.Pow(x, y)                      // x^y
math.Sin(x), math.Cos(x), math.Tan(x) // Trigonometric
math.Log(x), math.Log10(x)          // Logarithms
math.Max(x, y), math.Min(x, y)      // Maximum, minimum
math.Ceil(x), math.Floor(x)         // Ceiling, floor
math.Round(x)                       // Round to nearest
```

### `math/rand` - Random Numbers

```go
rand.Seed(time.Now().UnixNano())    // Seed with current time
rand.Int()                          // Random int
rand.Intn(n)                        // Random int [0,n)
rand.Float64()                      // Random float64 [0.0,1.0)
rand.Perm(n)                        // Random permutation of [0,n)
rand.Shuffle(n, swap func(i, j int)) // Shuffle slice

// Crypto-secure random numbers
buf := make([]byte, 16)
_, err := rand.Read(buf)            // Fill with random bytes
```

### `regexp` - Regular Expressions

```go
re, err := regexp.Compile(`pattern`) // Compile regexp
matched := re.MatchString(s)        // Check if string matches
matches := re.FindString(s)         // Find first match
allMatches := re.FindAllString(s, -1) // Find all matches
submatch := re.FindStringSubmatch(s) // Find first match with subgroups
re.ReplaceAllString(s, repl)        // Replace all matches
```

## Special-Purpose Packages

### `html/template` and `text/template` - Templates

```go
// Define template
tmpl, err := template.New("name").Parse("Hello, {{.Name}}!")

// Execute template
data := struct{ Name string }{"Alice"}
err := tmpl.Execute(os.Stdout, data)
```

### `database/sql` - Database Access

```go
db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/dbname")
defer db.Close()

// Query
rows, err := db.Query("SELECT id, name FROM users WHERE id = ?", 1)
defer rows.Close()
for rows.Next() {
    var id int
    var name string
    err := rows.Scan(&id, &name)
    // Use id and name
}
err := rows.Err()

// Exec
result, err := db.Exec("INSERT INTO users (name) VALUES (?)", "Alice")
id, err := result.LastInsertId()
affected, err := result.RowsAffected()

// Prepared statement
stmt, err := db.Prepare("SELECT name FROM users WHERE id = ?")
defer stmt.Close()
rows, err := stmt.Query(1)

// Transaction
tx, err := db.Begin()
_, err = tx.Exec("INSERT INTO users (name) VALUES (?)", "Alice")
err = tx.Commit()
// or
tx.Rollback()
```

### `crypto` - Cryptographic Functions

```go
// SHA-256 hash
h := sha256.New()
h.Write([]byte("hello world"))
hash := h.Sum(nil)
hashString := fmt.Sprintf("%x", hash)

// HMAC
h := hmac.New(sha256.New, []byte("secret key"))
h.Write([]byte("message"))
hmac := h.Sum(nil)

// AES encryption
block, err := aes.NewCipher(key)
ciphertext := make([]byte, aes.BlockSize+len(plaintext))
iv := ciphertext[:aes.BlockSize]
_, err := io.ReadFull(rand.Reader, iv)
stream := cipher.NewCFBEncrypter(block, iv)
stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

// AES decryption
stream = cipher.NewCFBDecrypter(block, iv)
stream.XORKeyStream(plaintext, ciphertext[aes.BlockSize:])
```

### `compress` - Compression

```go
// Gzip compression
var buf bytes.Buffer
gzw := gzip.NewWriter(&buf)
_, err := gzw.Write(data)
gzw.Close()

// Gzip decompression
gzr, err := gzip.NewReader(&buf)
decompressed, err := io.ReadAll(gzr)
gzr.Close()
```

## Tips from C# to Go

1. Use **explicit error handling** rather than exceptions
   ```go
   result, err := someFunc()
   if err != nil {
       // Handle error
       return err
   }
   // Use result
   ```

2. Use **defer** for cleanup (similar to using/IDisposable)
   ```go
   file, err := os.Open("file.txt")
   if err != nil {
       return err
   }
   defer file.Close() // Will be called when function returns
   ```

3. Remember that slices and maps have **reference semantics**
   ```go
   // Modifying slice inside function affects original
   func modify(s []int) {
       s[0] = 999 // Original slice is modified
   }
   ```

4. Use **composition over inheritance**
   ```go
   // Instead of inheritance, embed structs
   type Animal struct {
       Name string
   }

   type Dog struct {
       Animal         // Embed Animal
       BreedName string
   }
   ```

5. Use **interfaces** for polymorphism
   ```go
   type Speaker interface {
       Speak() string
   }

   // Any type implementing Speak() can be used
   func MakeTalk(s Speaker) {
       fmt.Println(s.Speak())
   }
   ```
