### 1. Using `sort.Slice`

- Signature:
  ```go
  sort.Slice(slice interface{}, less func(i, j int) bool)
  ```
- Sorts the slice **in place** according to `less`, which returns `true` if the element at index `i` should be ordered **before** the element at `j`.

### 2. Ascending / Descending order

```go
// Ascending:
sort.Slice(arr, func(i, j int) bool {
    return arr[i] < arr[j]
})

// Descending:
sort.Slice(arr, func(i, j int) bool {
    return arr[i] > arr[j]
})
```

### 3. Sorting custom structs

If you have a slice of structs, you can sort by any field:

```go
type Person struct {
    Name string
    Age  int
}

people := []Person{...}

sort.Slice(people, func(i, j int) bool {
    return people[i].Age < people[j].Age // sort by Age ascending
})
```

---

### 4. Sorting paired slices together

Sometimes you want to sort two related slices in lock-step (e.g., `red` and `blue` arrays) by the values in one slice. You can do this by defining a custom type, or you can sort indices:

Example: sort by red speeds descending, and reorder blue to match:

```go
type Pair struct {
    Red, Blue int
}

pairs := make([]Pair, len(red))
for i := range red {
    pairs[i] = Pair{red[i], blue[i]}
}

sort.Slice(pairs, func(i, j int) bool {
    return pairs[i].Red > pairs[j].Red
})

// Then extract sorted values back if needed
for i := range pairs {
    red[i], blue[i] = pairs[i].Red, pairs[i].Blue
}
```

---

### 5. Other built-in sorting functions

- `sort.Ints([]int)` sorts ascending ints.
- `sort.Strings([]string)` sorts ascending strings.
- `sort.Float64s([]float64)` sorts ascending floats.

If you want descending, you need `sort.Slice` or to sort ascending then reverse manually.
