# Hashed - Custom Modular Hash Generator

This is a modular implementation of a custom hash generator in Go, featuring:
- A lightweight `xxhash64`-inspired hashing algorithm.
- Base62 encoding to shorten the hash.
- Modular, clean project structure.

## 📁 Project Structure

```
/hashmod
├── main.go
└── hash/
    ├── hash.go        
    ├── base62.go      
    └── const.go       
```

## ⚙️ How It Works

1. `xxhash64(input string)`: Processes the string input and applies bitwise rotations and multipliers to produce a 64-bit hash.
2. `base62Conversion(num uint64)`: Converts the hash to a compact base62 string.
3. `GenerateHash(input string)`: Returns a 10-character base62-encoded hash string.

## 🚀 Getting Started

### Step 1: Initialize the Module

```bash
go get github.com/anuragsh-code/hash-assignment/hash
```

### Step 2: Use Function 

```bash
hash.GenerateHash("Brevo")
```

You should see output like:

```
x9agL5GhN4
```

## 🔍 Example Usage

```go
package main

import (
    "fmt"
    "anuragsh-code/hash-assignment/hash"
)

func main() {
    fmt.Println(hasher.GenerateHash("Brevo"))
}
```