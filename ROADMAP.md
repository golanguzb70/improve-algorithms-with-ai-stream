# 🧠 Algorithms Roadmap: 8 Weeks to Interview-Ready

**Goal:** Solve all **NeetCode 150** problems in **Go** in **8 weeks** (2+ hours a day).
**Success looks like:**
- You name the pattern within 5 minutes.
- You explain the approach out loud before you code.
- You state the Big-O without guessing.

---

## 1. The Core Loop (use it for every problem)

| # | Step | What you do |
|---|---|---|
| 1 | **Understand** | Restate the problem in your own words: inputs, outputs, constraints. |
| 2 | **Match** | Name the pattern from the signals (see the [cheat sheet](#4-pattern-recognition-cheat-sheet)). |
| 3 | **Plan** | Write the approach in 3–5 bullets, plus the Big-O. **No code yet.** |
| 4 | **Code** | Write it in Go. Keep it clean. |
| 5 | **Test** | Try these edge cases: empty, 1 element, duplicates, negatives, max size. |
| 6 | **Explain** | Give a 60-second summary out loud: problem → idea → complexity. |

⏱ **Timebox:** stuck for 25 minutes → read the hint or ask the AI tutor → solve the problem again tomorrow, from scratch.

> **Can't start a problem?** Always begin with brute force. Say it, give its Big-O, then ask: *"What work am I repeating?"* Removing that repeated work leads you to the optimal solution.

---

## 2. Problem Card (`problems/<slug>/card.md`, created for you by "start <problem>")

```md
## <Problem name>: <Pattern>
Signals:      what in the problem told me the pattern
Brute force:  ...                      O(?)
Optimal idea: ...
Complexity:   time O(?), space O(?)
Edge cases:   ...
Mistake I made: ...
```

---

## 3. 8-Week Schedule

| Week | NeetCode topics | # | Focus skill |
|---|---|---|---|
| 1 | Big-O basics · Arrays & Hashing · Two Pointers | 14 | Big-O for **every** solution |
| 2 | Sliding Window · Stack · Binary Search | 20 | Brute force → optimize |
| 3 | Linked List · Trees (part 1: DFS/BFS basics) | 18 | Thinking recursively |
| 4 | Trees (part 2) · Tries · Heap / Priority Queue | 18 | Go `container/heap` |
| 5 | Backtracking · Graphs | 22 | Drawing the state tree or graph |
| 6 | Advanced Graphs · 1-D DP | 18 | DP: state → recurrence → base case |
| 7 | 2-D DP · Greedy · Intervals | 25 | Explaining *why* greedy works |
| 8 | Math & Geometry · Bit Manipulation · **Mock interviews** · Redo misses | 15 | Full interview simulation |
| | **Total** | **150** | |

**Rule:** if a Hard problem blocks you for more than 40 minutes, skip it, mark it ⏭, and come back to it in week 8.

### Daily session (2 hours)
| Time | Activity |
|---|---|
| 10 min | Redo old misses (spaced repetition: after **1, 3 and 7 days**) |
| 90 min | 3 new problems with the Core Loop |
| 15 min | Write the Problem Cards |
| 5 min | Update the tracker |

**Every Sunday:** a 45-minute AI mock interview, then review your weakest patterns.

---

## 4. Pattern Recognition Cheat Sheet

| If you see… | Think… |
|---|---|
| "Have I seen this before?", counting, duplicates | **Hash map / set** |
| Sorted array, pair or triplet with a target sum | **Two pointers** |
| "Contiguous subarray/substring", longest or shortest window | **Sliding window** |
| Matching brackets, "next greater element" | **Stack / monotonic stack** |
| Sorted input, or "minimum value that works" | **Binary search** (on the input or on the answer) |
| Cycle, middle node, reorder a list | **Fast & slow pointers** |
| Tree: depth, path, subtree | **DFS (recursion)** |
| Tree level by level, shortest steps | **BFS (queue)** |
| Prefix / word search | **Trie** |
| "Top K", "K-th largest", merging K lists | **Heap** |
| "All combinations / permutations / subsets" | **Backtracking** |
| Grid, islands, "connected", dependencies | **BFS/DFS · Union-Find · Topological sort** |
| Shortest path with weights | **Dijkstra** |
| "Count ways", "min/max cost", choices that overlap | **Dynamic programming** |
| A locally best choice is always safe | **Greedy** |
| Overlapping ranges, meetings | **Sort by start → intervals** |
| "Without extra space", single number, powers of 2 | **Bit manipulation (XOR, &, >>)** |

---

## 5. Big-O Cheat Sheet

**Read the constraints → they tell you the target complexity:**

| n up to | Target |
|---|---|
| 10 | O(n!) |
| 20 | O(2ⁿ) |
| 500 | O(n³) |
| 10⁴ | O(n²) |
| 10⁵ – 10⁶ | O(n log n) or O(n) |
| > 10⁸ | O(log n) or O(1) |

**Common Go operation costs:**
- map get/set: O(1) on average
- `append`: amortized O(1)
- `sort.Slice`: O(n log n)
- heap push/pop: O(log n)
- slicing `s[a:b]`: O(1)
- `copy`: O(n)
- string concatenation in a loop: O(n²). Use `strings.Builder` instead.

**Space:** count your extra memory (maps, slices, recursion depth). The input itself doesn't count.

---

## 6. Go Toolkit

```go
// Set
seen := map[int]struct{}{}
seen[x] = struct{}{}
_, ok := seen[x]

// Stack
st := []int{}
st = append(st, x)            // push
top := st[len(st)-1]          // peek
st = st[:len(st)-1]           // pop

// Queue (BFS)
q := []int{start}
for len(q) > 0 {
    cur := q[0]; q = q[1:]
    _ = cur
}

// Sort
sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })

// Min-heap (container/heap)
type MinHeap []int
func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() any {
    old := *h; n := len(old); x := old[n-1]; *h = old[:n-1]; return x
}
// use: h := &MinHeap{}; heap.Push(h, 5); heap.Pop(h).(int)

// Extremes
best := math.MaxInt   // or math.MinInt
```

---

## 7. AI in Three Roles (just ask in plain words)

The roles are defined in `.claude/skills/`. Run `claude` in this repo and say:

| Say | Role | What it does |
|---|---|---|
| "lesson" / "teach me stack" | 📚 **Lesson** | A short interactive course on the week's topic: idea → templates → worked example → Go warm-up → quiz. See `lessons/` |
| "start two sum" / "next problem" | 📂 **Start Problem** | Creates `problems/003-two-sum/` (README with the problem and tasks, `solution.go`, tests, card) and marks it 🔄 in `PROGRESS.md` |
| "help" / "I'm stuck" | 🧑‍🏫 **Tutor** | Asks guiding questions and gives one hint level at a time. Never shows the solution unless you say *"show me the solution"* |
| "review" | 🔍 **Reviewer** | Runs the tests, adds edge cases, checks Big-O and Go style, grades your card, updates `PROGRESS.md` |
| "let's have an interview" | 🎤 **Interviewer** | Gives an unseen problem and runs a 45-minute mock interview, then saves a report to `interviews/` |

⚠️ Don't ask for the answer. That skips the exact skill you're trying to build.

---

## 8. Stream Format (~60 min)

1. **Recap (5 min):** yesterday's problems and what you learned
2. **Live solve (35 min):** the Core Loop, thinking out loud
3. **AI review (10 min):** paste your solution and discuss the feedback
4. **60-second explanation (5 min):** problem → pattern → idea → Big-O
5. **Tracker update (5 min):** show your progress publicly

---

## 9. Milestones

- [ ] **Week 2:** Give the correct Big-O for every solution, without help
- [ ] **Week 4:** Solve 2 of 3 unseen Easy/Medium problems in under 30 minutes each
- [ ] **Week 6:** Name the pattern within 5 minutes for 80% of new problems
- [ ] **Week 8:** Pass 3 AI mock interviews (score ≥ 7/10) with a clear verbal explanation

---

## 10. Progress Tracker

See **[PROGRESS.md](PROGRESS.md)**. The tutor, reviewer and interviewer update it automatically.

Legend: ✅ solved alone · 🟡 solved with a hint · ❌ had to read the solution · ⏭ skipped · 🎤 mock interview

---

## 11. Resources

- NeetCode roadmap and videos: https://neetcode.io/roadmap
- LeetCode, to practice and submit: https://leetcode.com
- Big-O reference: https://www.bigocheatsheet.com
