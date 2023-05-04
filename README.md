# --Data Stack Structure--

This is a simple implementation of a stack data structure in Go. The stack supports the following functions:

- `Push(item interface{})`: adds a new item to the top of the stack
- `Pop() interface{}`: removes the top item from the stack and returns it
- `Peek() interface{}`: returns the top item from the stack without removing it
- `Empty() bool`: returns true if the stack is empty, false otherwise

## Example Usage

To create a new stack, simply call the `NewStack()` function:

```go
stack := NewStack()
```

### Push()

To add an item to the stack, call the `Push()` function and pass the item as an argument:

```go
stack.Push(1)
stack.Push(2)
stack.Push(3)
```

The stack now contains the items [1, 2, 3].

### Pop()

To remove the top item from the stack, call the `Pop()` function:

```go
item := stack.Pop()
```

This removes the item 3 from the stack and returns it.

### Peek()

To get the top item from the stack without removing it, call the `Peek()` function:

```go
item := stack.Peek()
```

This returns the item 2, which is the top item on the stack.

### Empty()

To check if the stack is empty, call the `Empty()` function:

```go
if stack.Empty() {
    fmt.Println("The stack is empty")
} else {
    fmt.Println("The stack is not empty")
}
```

This will print "The stack is not empty", since we have added items to the stack.

That's it! This simple stack implementation should be useful for a wide range of applications in Go.
