# Lesson 01 · Arrays & Hashing

⏱ ~45 min · Week 1 · Problems 001–009

## 🎯 After this lesson you can
- Use a map to turn "search again" (O(n)) into "look up" (O(1))
- Recognize the 4 hashing patterns: **set, counter, index map, grouping key**
- Use prefix sums and prefix products

---

## 1. The idea in one line
**Trade memory for speed.** Instead of scanning the array again and again, remember what you've seen in a map.

```
Brute force:  for each x → scan the whole array again   O(n²)
Hashing:      for each x → ask the map "seen it?"       O(n)
```

## 2. Go maps in 30 seconds
```go
m := map[string]int{}        // create
m["go"]++                    // missing keys start at the zero value (0)
v, ok := m["rust"]           // ok == false if missing
delete(m, "go")
for k, v := range m { }      // ⚠️ the order is RANDOM
len(m)
```
- Keys must be comparable: `int`, `string`, `byte`, arrays like `[26]int`, and structs. **Slices can't be keys.**
- Set: `map[int]struct{}` (no memory for the values) or `map[int]bool` (easier to read).

## 3. The 4 patterns

### Pattern A: Set ("have I seen this?")
```go
seen := map[int]bool{}
for _, x := range nums {
    if seen[x] { /* found a repeat */ }
    seen[x] = true
}
```

### Pattern B: Counter ("how many of each?")
```go
count := map[rune]int{}
for _, ch := range s { count[ch]++ }
```
💡 For lowercase letters only, use an array. It's faster and comparable:
```go
var count [26]int
for _, ch := range s { count[ch-'a']++ }
// two [26]int arrays can be compared with ==
```

### Pattern C: Index map ("where did I see this?")
```go
idx := map[int]int{}         // value → index
for i, x := range nums {
    if j, ok := idx[something(x)]; ok { /* use j and i */ }
    idx[x] = i
}
```

### Pattern D: Grouping key ("which things belong together?")
Build a **key** that is the same for every item in a group, then use `map[key][]item`.
```go
groups := map[string][]string{}
for _, w := range words {
    key := makeKey(w)        // e.g. its length, first letter, a sorted form...
    groups[key] = append(groups[key], w)
}
```

## 4. Bonus pattern: prefix sums
Precompute the running totals once, then answer "sum of a[l..r]" in O(1).
```go
prefix := make([]int, len(a)+1)
for i, x := range a { prefix[i+1] = prefix[i] + x }
sum := prefix[r+1] - prefix[l]   // sum of a[l..r]
```
The same idea works with **products**, and **from the right side** too (suffix).

## 5. Signals: when to think "hashing"

| The problem says… | Think |
|---|---|
| "duplicate", "unique", "distinct" | Set |
| "anagram", "frequency", "most common" | Counter |
| "two numbers that…", "pair", "complement" | Index map |
| "group", "categorize" | Grouping key |
| "sum of a subarray", "everything except i" | Prefix / suffix |
| Brute force has a nested loop that *searches* | Replace the inner loop with a map |

## 6. Complexity
- Map insert/lookup: O(1) on average → a single pass is **O(n) time, O(n) space**
- Sorting instead of hashing: O(n log n) time, but O(1) extra space. It's a valid trade-off, so mention it in interviews.

## 7. Common mistakes ⚠️
- Relying on map iteration order (it's random in Go)
- Using a slice as a map key (it won't compile). Convert it to a string or an array.
- Forgetting `ok` when the zero value is a valid answer (`idx[x]` returns 0 when the key is missing!)
- Inserting into the map **before** checking it, so an element matches itself

---

## 🧪 Worked example: First Unique Character (LeetCode 387, not in our list)
> Return the index of the first character in `s` that appears exactly once, or -1.
> `"leetcode"` → 0 · `"loveleetcode"` → 2 · `"aabb"` → -1

1. **Understand:** we need the first index whose character has a count of 1.
2. **Signals:** "appears exactly once" → **counter**.
3. **Brute force:** for each i, count s[i] across the whole string → O(n²).
4. **Optimal:** pass 1 counts every char; pass 2 returns the first i with count 1.
```go
func firstUniqChar(s string) int {
    var count [26]int
    for _, ch := range s { count[ch-'a']++ }
    for i, ch := range s {
        if count[ch-'a'] == 1 { return i }
    }
    return -1
}
```
5. **Complexity:** O(n) time, O(1) space (26 is a constant).
6. **60-second explanation:** *"I count every letter in one pass, using a 26-slot array since the input is lowercase. In a second pass I return the first index whose count is 1. Two linear passes give O(n) time, and the array is fixed-size, so the space is O(1)."*

---

## 🏋️ Warm-up exercise
Open `exercise.go`, implement the 3 functions, then run:
```bash
go test ./lessons/01-arrays-hashing/
```

## ✍️ Quiz

**1.** Why is `[26]int` sometimes better than `map[rune]int`?
<details><summary>Answer</summary>It's faster (no hashing), uses fixed O(1) space, and two arrays can be compared with <code>==</code> or used as a map key.</details>

**2.** You need to group words that have the same letters. What could the key be?
<details><summary>Answer</summary>The sorted word (as a string), or its <code>[26]int</code> letter count.</details>

**3.** What's wrong here?
```go
idx := map[int]int{}
for i, x := range nums {
    idx[x] = i
    if j, ok := idx[target-x]; ok { return []int{j, i} }
}
```
<details><summary>Answer</summary>It inserts before checking, so when <code>target-x == x</code> it matches itself. Check first, then insert.</details>

**4.** You need the sum of `a[2..5]` for many queries. Which tool do you use?
<details><summary>Answer</summary>Prefix sums: <code>prefix[6] - prefix[2]</code>, O(1) per query.</details>

---

## ✅ Done when
- [ ] All tests in `exercise_test.go` pass
- [ ] Quiz: 3 of 4 correct
- [ ] You can name the 4 hashing patterns from memory

➡️ Next: [Lesson 02 · Two Pointers](../02-two-pointers/README.md) · Practice: problems **001–009**
