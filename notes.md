# Go:
    1. Speed
    2. Concurrency

Use the right programming patterns for optimal performance. (Design guidelines)

Go's execution model is the actual machine and not a virtual machine (as in the case of Java)

Variable Types: how much memory is the variable going to use?

Word: AMD64 architecture (CPU hardware) => 1 word = 64 bits
Pointers need to be of the same size as the word size in order to store addresses
int: size = 64 on a 64 bit architecture since pointers are integers too

Address size = Word size = Integer size

# Zero values
Every variable in Go must be initialized. If the programmer does not do this, the variable is intialized to its zero-value (all bits set to zero)

# Strings

**Size**: 2 words
- Word 0: pointer to a backing array of bytes. Zero values = nil
- Word 1: Size. Number of bytes. Zero value = 0

## Example:
```go
    a := "hello"
```        
- Backing array: 5 bytes backing array to hold the binary representation of "hello" 
- Word 0: A pointer that stores the starting address of the backing array
- Word 1: 5


# Casting
Go does not support type casting. Instead it requires the programmer to perform a type conversion that looks like a function call

# Struct alignment
## Example 1: Bad alignment

```go
type example struct {
    flag bool       // 1 byte
    counter int16   // 2 bytes
    pi float32      // 4 bytes
}
```
**Alignment:**
- A 2 bytes value must fall on a 2 byte alignment position (Bit 0, bit 16, bit 32 ....)
- A 4 bytes value must fall on a 2 byte alignment position (Bit 0, bit 32, bit 64 ....)

Bit positions
1. 0-7 : bool 
2. 8-15: one padded byte in order to fit the next int16 variable
3. 16-31: int16
4. 32-63: float32

> Total space: 64 bits = 8 bytes!


## Example 2: Good alignment -> Order the LARGEST type FIRST in a struct
```go    
type example struct {
    pi float32      // 4 bytes
    counter int16   // 2 bytes
    flag bool       // 1 byte
}
```
**Alignment:**
- A 2 bytes value must fall on a 2 byte alignment position (Bit 0, bit 16, bit 32 ....)
- A 4 bytes value must fall on a 2 byte alignment position (Bit 0, bit 32, bit 64 ....)

**Bit positions**
1. 0-31: float32
2. 32-47: int16
3. 48-55: bool

> Total space: 56 bits = 7 bytes! (Space savings)