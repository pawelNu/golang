# Loops in Go

The basic loop in Go is written in standard C-like syntax:

```go
for INITIAL; CONDITION; AFTER{
  // do something
}
```

`INITIAL` is run once at the beginning of the loop and can create
variables within the scope of the loop.

`CONDITION` is checked before each iteration. If the condition doesn't pass
then the loop breaks.

`AFTER` is run after each iteration.

For example:

```go
for i := 0; i < 10; i++ {
  fmt.Println(i)
}
// Prints 0 through 9
```

## Assignment

At Textio we have a dynamic formula for determining how much a batch of bulk messages costs to send.

### Complete the `bulkSend()` function

This function should return the total cost (as a `float64`) to send a batch of `numMessages` messages. Each message costs `1.0`, plus an additional fee. The fee structure is:

* 1st message: `1.0 + 0.00`
* 2nd message: `1.0 + 0.01`
* 3rd message: `1.0 + 0.02`
* 4th message: `1.0 + 0.03`

Use a loop to calculate the total cost and return it.

Here’s a comprehensive overview of loop syntax in Go (Golang):

### 1. Basic `for` Loop Syntax

The simplest form of a `for` loop looks like this:

```go
for i := 0; i < 10; i++ {
    fmt.Println(i)
}
```

### 2. Infinite Loop

You can define an infinite loop using just `for` without conditions:

```go
for {
    // This will run indefinitely
    fmt.Println("This will never end!")
    break // Use break to exit the loop
}
```

### 3. Loop with a Condition

You can also use a `for` loop with a condition to achieve behavior similar to `while`:

```go
i := 0
for i < 10 {
    fmt.Println(i)
    i++
}
```

### 4. Loop with `range`

The `for` loop can also be used with `range` to iterate over elements of arrays, slices, maps, and channels:

#### Iterating over a Slice:

```go
slice := []int{1, 2, 3, 4, 5}
for index, value := range slice {
    fmt.Printf("Index: %d, Value: %d\n", index, value)
}
```

#### Iterating over a Map:

```go
m := map[string]int{"a": 1, "b": 2, "c": 3}
for key, value := range m {
    fmt.Printf("Key: %s, Value: %d\n", key, value)
}
```

#### Iterating over a String:

```go
str := "hello"
for index, char := range str {
    fmt.Printf("Index: %d, Character: %c\n", index, char)
}
```

### 5. Using `continue` and `break`

You can use `continue` to skip to the next iteration of the loop or `break` to exit the loop:

#### Using `continue`:

```go
for i := 0; i < 10; i++ {
    if i%2 == 0 {
        continue // Skip even numbers
    }
    fmt.Println(i) // Print only odd numbers
}
```

#### Using `break`:

```go
for i := 0; i < 10; i++ {
    if i == 5 {
        break // Exit the loop when i equals 5
    }
    fmt.Println(i)
}
```

### Summary

In Go, you can use the `for` loop in various configurations to iterate through data. Here’s a summary of the available loop syntaxes:

1. **Basic loop with initialization, condition, and increment**:
   ```go
   for i := 0; i < 10; i++ { }
   ```

2. **Infinite loop**:
   ```go
   for { }
   ```

3. **Loop with a condition (similar to `while`)**:
   ```go
   for condition { }
   ```

4. **Iterating with `range` over slices, maps, or strings**:
   ```go
   for index, value := range collection { }
   ```

5. **Using `continue` and `break`**:
   ```go
   for { 
       if condition { continue }
       if condition { break }
   }
   ```

This covers all the loop syntax in Go! Let me know if you need further clarification or examples.
