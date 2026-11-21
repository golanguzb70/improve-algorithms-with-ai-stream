---
name: start-problem
description: Prepares the full workspace to solve a NeetCode 150 problem. Use when the user says "start <problem>", "I'm going to solve <problem>", "let's do <problem>", "next problem", "start 012", or "what's next?". Creates problems/<ID>-<slug>/ with README (statement + task checklist), solution.go stub, solution_test.go with examples, and card.md, then marks it in progress in PROGRESS.md.
---

# Start Problem

Set up everything the user needs to start solving **immediately**. You prepare the workspace. You **never** write the solution logic.

## 1. Identify the problem
`PROGRESS.md` is the master list of all 150 problems, with IDs 001–150.
- By name or ID ("start two sum", "start 003"): find that row. Match names fuzzily.
- "next" or "what's next?": take the first ⬜ row in ID order. Before it, offer any row whose "Redo on" date is today or earlier. Run `date +%F` to get today.
- Not in the list: ask whether to add it as an extra problem (ID `X01`, `X02`, …) at the bottom, under "Extras".
- If `problems/<ID>-<slug>/` already exists, **don't overwrite it**. Show its status and open tasks, and ask whether they want a fresh redo. If yes, rename the old folder to `<ID>-<slug>/attempt-1/`. Don't delete anything.

From the row, take: ID, title, slug (from the LeetCode URL), difficulty, week, and the category (the `###` heading above the row).

**Lesson check:** if the lesson for this category isn't ✅ in the **📚 Lessons** table, say so in one line: "📚 Lesson NN <topic> isn't done yet. Say **lesson** first (recommended), or **continue** to start anyway." Wait for their answer. If Lesson 00 Big-O isn't done, mention it too.

## 2. Create `problems/<ID>-<slug>/`
For example `problems/003-two-sum/`. It holds four files:

### `README.md`
Copy it from `templates/README.md` and fill every `{{…}}` placeholder (`{{folder}}` = `<ID>-<slug>`). Then write:
- **Problem:** a clear statement **in your own words** (don't paste LeetCode text verbatim).
- **Examples:** 2–3 examples, each with a short "Why".
- **Constraints:** the real ones (sizes and value ranges). They hint at the target Big-O.
- Don't name the pattern or the approach. Discovering it is the task.

### `solution.go`
```go
package solution

// <Title>
// <link>
//
// Time:  O(?)
// Space: O(?)
func <leetcodeFuncName>(<exact LeetCode params in Go>) <return type> {
	// TODO: your solution
	return <zero value>
}
```
- Every folder uses `package solution`.
- Use the exact LeetCode Go signature and function name, so the code can be pasted into LeetCode.
- **Linked list or tree problems:** import `github.com/golanguzb70/improve-algorithms-with-ai-stream/ds` and add `type ListNode = ds.ListNode` / `type TreeNode = ds.TreeNode`.
- **Problems with other custom nodes** (Clone Graph's `Node`, Copy List's random `Node`): define the struct exactly as LeetCode does.
- **Design problems** (Min Stack, LRU Cache, Trie, Twitter, …): write the struct, a `Constructor(...)` and the method stubs, matching LeetCode.
- **In-place problems** (no return value): keep the signature, and the test checks the mutated input.
- It must compile: `go vet` must pass.

### `solution_test.go`
- Use table-driven tests with the README examples, plus a commented-out block where the user adds their own edge cases.
- Use `ds.NewList`, `ds.ListToSlice`, `ds.NewTree`, `ds.TreeToSlice` and `ds.FindNode` for list and tree problems.
- If any valid answer is accepted (any order, any valid pair), normalize first: sort the result, or check the property instead of exact equality.
- Design problems: write a test that replays the LeetCode operations sequence.
- The tests **should fail** against the stub. That's expected, and it's the user's starting point.

### `card.md`
Copy it from `templates/card.md` and fill `{{id}}` and `{{title}}`. Leave the rest for the user.

## 3. Verify
Run `go vet ./problems/<ID>-<slug>/`. It must pass. Then run `go test ./problems/<ID>-<slug>/`. The tests must compile and fail. Fix any compile errors in the stub or tests.

## 4. Update `PROGRESS.md`
- Set the row's Status to 🔄 and Started to today (`date +%F`).
- Update the dashboard: the In progress count, Current week (from the row) and Next up (the next ⬜ after this one).

## 5. Tell the user (short, ready to read on stream)
```
📂 problems/003-two-sum/  ·  🟢 Easy · Week 1 · Arrays & Hashing
Files: README.md (problem + tasks) · solution.go · solution_test.go · card.md

✅ Tasks
1. Understand → card.md
2. Signals + pattern
3. Brute force + Big-O
4. Plan optimal + Big-O
5. Code solution.go
6. go test ./problems/003-two-sum/  + add edge cases
7. 60-sec explanation
8. Say "review"

First question: <one question that gets them restating the problem>
```
Then stop and let them work. Don't start tutoring unless they ask.
