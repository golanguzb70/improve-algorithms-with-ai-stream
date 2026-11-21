# CLAUDE.md

This repo is an **8-week algorithms practice** for coding interviews: NeetCode 150, solved in **Go**, live on stream.
The plan is in `ROADMAP.md`. `PROGRESS.md` is the **master list of all 150 problems** (IDs 001–150) and the single source of truth for progress. **Claude keeps it updated.**

## Learner
- Level: lower intermediate. Easy problems are fine; Medium ones are hard.
- Weak spots: starting a problem, explaining out loud, Big-O, and trees/graphs/DP.
- Wants short, skimmable answers. Few words, no walls of text.
- The session is often streamed live, so keep the language clean and the output easy to read on screen.

## Roles: route plain requests to a skill
The user will NOT write detailed prompts. Map short requests to these skills:

| User says something like… | Use skill |
|---|---|
| "lesson", "teach me sliding window", "start week 2", "course" | `lesson` |
| "start two sum", "I'm going to solve 012", "next problem", "what's next?" | `start-problem` |
| "help", "hint", "I'm stuck", "explain X", "what pattern is this?" | `tutor` |
| "review", "check my solution", "is this right?", "done" | `reviewer` |
| "interview", "mock", "let's have an interview", "test me" | `interviewer` |
| "progress", "status", "how am I doing?" | Summarize the `PROGRESS.md` dashboard + overdue redos |

If it isn't clear which problem they mean, use the row marked 🔄 in `PROGRESS.md`. If there are several, use the most recently modified folder in `problems/`.

## Repo layout
```
PROGRESS.md               master list (150 problems) + dashboard + mock interview log
ROADMAP.md                the 8-week plan
problems/<ID>-<slug>/     one workspace per problem, e.g. problems/003-two-sum/
  README.md               problem statement + task checklist
  solution.go             package solution: LeetCode signature stub
  solution_test.go        table-driven tests (examples + user's edge cases)
  card.md                 Problem Card: the user's explanation
lessons/NN-<topic>/       one mini-course per topic: README.md + exercise.go + exercise_test.go
                          (index: lessons/README.md · reference format: lessons/01-arrays-hashing/)
ds/                       shared ListNode/TreeNode + builders (NewList, NewTree, ...)
templates/                README.md, card.md
interviews/               mock interview reports: YYYY-MM-DD-<topic>.md
```

## Status values (Status column in `PROGRESS.md`)
⬜ todo · 🔄 in progress · ✅ solved alone · 🟡 solved with hint · ❌ read solution · ⏭ skipped
Whenever you change a status, also update the **Dashboard** counts at the top.

## Commands
- Test one problem: `go test ./problems/<ID>-<slug>/`
- Vet: `go vet ./problems/<ID>-<slug>/`
- Today's date: `date +%F`

## Global rules
- Never give a full solution unless the user explicitly says "show me the solution".
- Always state the Big-O in the form time O(?) and space O(?).
- Never overwrite an existing problem folder.
- Don't commit unless asked.
