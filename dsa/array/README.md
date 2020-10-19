# Array

- An _array_ is a fundamentally a **list of similar values**.
- Elements of an array are **stored in contagious memory** blocks.
- All the elements have to have the **same data type**. 
- Arrays have a fixed length; array length **cannot be changed** once created. 

An array has usually 3 attributes:

- A **name** to reference it.
- A **type** to fill it with.
- A **size**.

```go
intarr := [5]int{1, 2, 3, 4, 5}  // name=intarr; type=int; size=5
```

We reference an element in the array, we use the array name and the integer index (starting from 0).

To access the number 3 from `intarr` defined above, we'd use:

```go
intarr[2]  // 2 instead of 3 because index starts at 0
```    

An array at every index of another array will be two-dimensional array.

We access the multi-dimensional array the same way as one-dimensional array. Just the difference is, the first index is the column (matrices reference) and the second is row.

Consider the `name` matrix below:

| Index | 0        | 1        | 2         | 3         |
|-------|:--------:|:--------:|:---------:|:---------:|
| 0     | "Steven" | "Alex"   | "Dave"    | "Jake"    |
| 1     | "Adam"   | "Lucus"  | "Quinten" | "John"    |
| 2     | "Sean"   | "Marcus" | "Carl"    | "Jackson" |
| 3     | "Peter"  | "Cam"    | "Anthony" | "Ethan"   |


To access `Quinten`, we can do `name[2][1]`. Similarly, Peter is at `name[0][3]`.

Operations on Array:

1. Add/Insert

Takes `O(n)` when inserting in middle/start of array. That's because we have to shift each element to right.

Variations: add, addFirst, addLast, addByIndex

2. Get/Access

- Takes `O(1)` to access element.

Variations: getByIndex

3. Search/Traverse

Takes `O(n)` when using linear search. 

4. Remove/Delete

Deleting also takes O(n) because we are considering worst case. Every element at the right of the array needs to be shifted to the left.

Advantages:

- As compared to LinkedList, data in the array can be access in O(1) time.

Disadvantages:

- The size of an array cannot be changed. Space can be wasted if allocated space are not used.
- Inserting and deleting are not efficient.
