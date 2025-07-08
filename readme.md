Let's dive deep into **structs in Go** — one of the most fundamental and powerful data structures in the language.

---

## 🧱 What is a Struct?

A **`struct`** in Go is a **composite data type** (i.e., it groups multiple values into a single unit). It's used to create **custom data types** that model real-world entities — like `User`, `Product`, `Book`, `Transaction`, etc.

Think of a `struct` as a **class without methods**, though we can **attach methods to structs** in Go.

---

## 🧪 Declaring a Struct

```go
type User struct {
	Name string
	Age  int
}
```

* `User` is a new custom type.
* It has two **fields**: `Name` and `Age`.

This is **definition only**. To use it, we need to create instances.

---

## 🧬 Creating Struct Instances

### 1. Using Struct Literal

```go
user := User{Name: "Skyy", Age: 29}
```

### 2. Without Field Names (order matters)

```go
user := User{"Skyy", 29}
```

> ✅ Not recommended for readability

### 3. Declare and then assign

```go
var user User
user.Name = "Skyy"
user.Age = 29
```

---

## 🔍 Accessing and Modifying Fields

```go
fmt.Println(user.Name) // "Skyy"

user.Age = 30
```

---

## 🪄 Zero Values

If you declare a struct without assigning fields:

```go
var u User
fmt.Println(u) // {"" 0}
```

* All fields get **zero values** (`""` for string, `0` for int, etc.)

---

## 🧭 Pointers to Structs

Just like with any type, we can create pointers to structs:

```go
u := &User{Name: "Skyy", Age: 29}
fmt.Println(u.Name) // Go automatically dereferences u
```

Go provides syntactic sugar: `u.Name` is the same as `(*u).Name`.

This is very useful when **passing structs to functions** or **mutating data**.

---

## 🧮 Structs with Methods

We can associate **methods** with a struct type:

```go
func (u User) Greet() {
	fmt.Println("Hello,", u.Name)
}
```

> This is a **value receiver** – works on a copy of `User`.

If we want to **mutate fields**, use a **pointer receiver**:

```go
func (u *User) UpdateName(newName string) {
	u.Name = newName
}
```

---

## 📦 Embedding Structs (Composition)

Go doesn’t support inheritance but allows **embedding**.

```go
type Address struct {
	City  string
	State string
}

type Employee struct {
	Name string
	Age  int
	Address // embedded struct
}
```

Now you can access `Employee.City` directly:

```go
emp := Employee{Name: "Skyy", Age: 29, Address: Address{"Kolkata", "WB"}}
fmt.Println(emp.City) // "Kolkata"
```

---

## 🌐 Struct Tags (Used in JSON, Database)

Struct tags add metadata used by packages like `encoding/json`, `gorm`, etc.

```go
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
```

Used with packages like:

```go
json.Marshal(user) // produces {"name":"Skyy","age":29}
```

---

## 🧾 Anonymous Structs

Quick way to declare a one-time struct:

```go
emp := struct {
	Name string
	Age  int
}{"Skyy", 29}
```

---

## 🗃️ Array/Slice of Structs

```go
users := []User{
	{Name: "Alice", Age: 25},
	{Name: "Bob", Age: 30},
}
```

---

## 🧠 Best Practices with Structs

| Tip                                | Description                                      |
| ---------------------------------- | ------------------------------------------------ |
| Use field names when initializing  | Improves readability                             |
| Use pointer receivers for mutation | Avoid unnecessary copies                         |
| Group related fields logically     | Easier to maintain and understand                |
| Use struct tags for API/database   | Interoperability with other systems              |
| Use embedded structs for reuse     | Clean, DRY design without inheritance complexity |

---

## 🧰 Real World Example

```go
type Product struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	InStock  bool    `json:"in_stock"`
}

func (p *Product) ApplyDiscount(discount float64) {
	p.Price = p.Price - (p.Price * discount / 100)
}
```

---

Let's dive deep into **structs in Go** — one of the most fundamental and powerful data structures in the language.

---

## 🧱 What is a Struct?

A **`struct`** in Go is a **composite data type** (i.e., it groups multiple values into a single unit). It's used to create **custom data types** that model real-world entities — like `User`, `Product`, `Book`, `Transaction`, etc.

Think of a `struct` as a **class without methods**, though we can **attach methods to structs** in Go.

---

## 🧪 Declaring a Struct

```go
type User struct {
	Name string
	Age  int
}
```

* `User` is a new custom type.
* It has two **fields**: `Name` and `Age`.

This is **definition only**. To use it, we need to create instances.

---

## 🧬 Creating Struct Instances

### 1. Using Struct Literal

```go
user := User{Name: "Skyy", Age: 29}
```

### 2. Without Field Names (order matters)

```go
user := User{"Skyy", 29}
```

> ✅ Not recommended for readability

### 3. Declare and then assign

```go
var user User
user.Name = "Skyy"
user.Age = 29
```

---

## 🔍 Accessing and Modifying Fields

```go
fmt.Println(user.Name) // "Skyy"

user.Age = 30
```

---

## 🪄 Zero Values

If you declare a struct without assigning fields:

```go
var u User
fmt.Println(u) // {"" 0}
```

* All fields get **zero values** (`""` for string, `0` for int, etc.)

---

## 🧭 Pointers to Structs

Just like with any type, we can create pointers to structs:

```go
u := &User{Name: "Skyy", Age: 29}
fmt.Println(u.Name) // Go automatically dereferences u
```

Go provides syntactic sugar: `u.Name` is the same as `(*u).Name`.

This is very useful when **passing structs to functions** or **mutating data**.

---

## 🧮 Structs with Methods

We can associate **methods** with a struct type:

```go
func (u User) Greet() {
	fmt.Println("Hello,", u.Name)
}
```

> This is a **value receiver** – works on a copy of `User`.

If we want to **mutate fields**, use a **pointer receiver**:

```go
func (u *User) UpdateName(newName string) {
	u.Name = newName
}
```

---

## 📦 Embedding Structs (Composition)

Go doesn’t support inheritance but allows **embedding**.

```go
type Address struct {
	City  string
	State string
}

type Employee struct {
	Name string
	Age  int
	Address // embedded struct
}
```

Now you can access `Employee.City` directly:

```go
emp := Employee{Name: "Skyy", Age: 29, Address: Address{"Kolkata", "WB"}}
fmt.Println(emp.City) // "Kolkata"
```

---

## 🌐 Struct Tags (Used in JSON, Database)

Struct tags add metadata used by packages like `encoding/json`, `gorm`, etc.

```go
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
```

Used with packages like:

```go
json.Marshal(user) // produces {"name":"Skyy","age":29}
```

---

## 🧾 Anonymous Structs

Quick way to declare a one-time struct:

```go
emp := struct {
	Name string
	Age  int
}{"Skyy", 29}
```

---

## 🗃️ Array/Slice of Structs

```go
users := []User{
	{Name: "Alice", Age: 25},
	{Name: "Bob", Age: 30},
}
```

---

## 🧠 Best Practices with Structs

| Tip                                | Description                                      |
| ---------------------------------- | ------------------------------------------------ |
| Use field names when initializing  | Improves readability                             |
| Use pointer receivers for mutation | Avoid unnecessary copies                         |
| Group related fields logically     | Easier to maintain and understand                |
| Use struct tags for API/database   | Interoperability with other systems              |
| Use embedded structs for reuse     | Clean, DRY design without inheritance complexity |

---

## 🧰 Real World Example

```go
type Product struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	InStock  bool    `json:"in_stock"`
}

func (p *Product) ApplyDiscount(discount float64) {
	p.Price = p.Price - (p.Price * discount / 100)
}
```
---

