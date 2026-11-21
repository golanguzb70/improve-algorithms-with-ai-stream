# Lesson 00 · Big-O: How Fast Is My Code?

⏱ ~30 min · Week 1 · Do this lesson before any problem

## 🎯 After this lesson you can
- Give the time and space complexity of any solution you write
- Read a problem's constraints and know which complexity you're aiming for

---

## 1. The idea in one line
Big-O tells you **how the work grows as the input grows**. It ignores constants and keeps only the biggest term.

```
3n² + 5n + 100  →  O(n²)      (n² wins as n gets big)
2n + 7          →  O(n)
50              →  O(1)
```

## 2. The ladder (fastest → slowest)

| Big-O | Name | Typical code | n = 10⁵ → ops |
|---|---|---|---|
| O(1) | constant | map lookup, `a[i]` | 1 |
| O(log n) | logarithmic | binary search, halving | ~17 |
| O(n) | linear | one loop | 10⁵ ✅ |
| O(n log n) | linearithmic | sorting | ~1.7·10⁶ ✅ |
| O(n²) | quadratic | two nested loops | 10¹⁰ ❌ too slow |
| O(2ⁿ) | exponential | all subsets | ∞ ❌ |
| O(n!) | factorial | all permutations | ∞ ❌ |

> Rule of thumb: a computer does about **10⁸ simple operations per second**. Interview limits are usually around 1 second.

## 3. How to count: 4 rules

**Rule 1: a loop multiplies by its size.**
```go
for i := 0; i < n; i++ { ... }        // O(n)
```

**Rule 2: nested loops multiply.**
```go
for i := 0; i < n; i++ {
    for j := 0; j < n; j++ { ... }    // O(n·n) = O(n²)
}
```

**Rule 3: loops one after another add, and you keep the bigger one.**
```go
for ... n { }      // O(n)
for ... n { }      // O(n)
// total O(2n) = O(n)
```

**Rule 4: halving gives log n.**
```go
for n > 1 { n /= 2 }                  // O(log n)
```

## 4. Hidden costs in Go ⚠️
These lines *look* like O(1), but some aren't:

| Code | Cost |
|---|---|
| `m[k]`, `m[k] = v`, `delete(m, k)` | O(1) on average |
| `append(s, x)` | O(1) amortized |
| `s[a:b]` (slicing) | O(1): no copy |
| `copy(dst, src)` | O(n) |
| `sort.Slice(s, ...)` | O(n log n) |
| `strings.Contains(s, sub)` | O(n·m) worst case |
| `s += "x"` inside a loop | O(n) each time → O(n²) total. Use `strings.Builder` |
| `string(bytes)` / `[]byte(str)` | O(n) copy |

## 5. Space complexity
Count the **extra** memory you create. The input itself doesn't count.
- a few variables → O(1)
- a map or slice holding up to n items → O(n)
- recursion → O(depth), because every call sits on the stack

## 6. Constraints → target complexity 🎯

| n up to | Aim for |
|---|---|
| ≤ 10 | O(n!) is fine |
| ≤ 20 | O(2ⁿ) |
| ≤ 500 | O(n³) |
| ≤ 10⁴ | O(n²) |
| ≤ 10⁶ | O(n log n) or O(n) |
| > 10⁸ | O(log n) or O(1) |

**Always read the constraints first.** They tell you how clever you need to be.

---

## 🧪 Worked example
What's the complexity of this?
```go
func hasPairWithSum(nums []int, target int) bool {
    sort.Ints(nums)                    // A
    for i := range nums {              // B
        for j := i + 1; j < len(nums); j++ {   // C
            if nums[i]+nums[j] == target { return true }
        }
    }
    return false
}
```
**Think out loud:**
- A: sorting → O(n log n)
- B + C: nested loops → O(n²)
- Total: O(n log n + n²) → **O(n²)** time. No extra structures → **O(1)** space (ignoring the sort's internals).
- The sort is wasted here: the nested loops dominate anyway.

---

## ✍️ Drills: answer out loud first, then check

**1.**
```go
for i := 0; i < n; i++ {
    for j := 0; j < 10; j++ { ... }
}
```
<details><summary>Answer</summary>O(n). The inner loop is a constant 10, so it's 10n → O(n).</details>

**2.**
```go
for i := 1; i < n; i *= 2 { ... }
```
<details><summary>Answer</summary>O(log n). i doubles each step.</details>

**3.**
```go
for i := 0; i < n; i++ {
    for j := i; j < n; j++ { ... }
}
```
<details><summary>Answer</summary>O(n²). That's n + (n-1) + … + 1 = n(n+1)/2 → O(n²).</details>

**4.**
```go
seen := map[int]bool{}
for _, x := range nums { seen[x] = true }
```
<details><summary>Answer</summary>O(n) time, O(n) space.</details>

**5.** A problem says `1 ≤ n ≤ 10⁵`. Is an O(n²) solution OK?
<details><summary>Answer</summary>No. 10¹⁰ operations is far too slow. Aim for O(n log n) or O(n).</details>

**6.**
```go
func f(n int) int {
    if n <= 1 { return n }
    return f(n-1) + f(n-2)
}
```
<details><summary>Answer</summary>O(2ⁿ) time, because each call branches into 2. O(n) space for the recursion depth. (You'll fix this with DP in week 6.)</details>

---

## ✅ Done when
- [ ] You answered 5 of the 6 drills correctly before opening the answer
- [ ] You can explain Rule 4 (why halving gives log n) in one sentence

➡️ Next: [Lesson 01 · Arrays & Hashing](../01-arrays-hashing/README.md)
