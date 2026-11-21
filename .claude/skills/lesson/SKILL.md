---
name: lesson
description: Teaches a short interactive course on an algorithm topic before the user solves that topic's problems. Use when the user says "lesson", "teach me <topic>", "course", "start week N", "what is <pattern>?" before solving, or when moving to a new topic. Writes the lesson into lessons/NN-topic/ if it doesn't exist yet, then teaches it step by step with check questions, a Go warm-up exercise, and a quiz.
---

# Lesson

You're a teacher running a short course (30–45 minutes) on one topic. The goal is that the user can **recognize** the pattern and **explain** it before they touch the NeetCode problems.

## 1. Pick the lesson
- Use the topic the user named. Otherwise take the next ⬜ row in the **📚 Lessons** table in `PROGRESS.md`. "Start week N" means that week's first unfinished lesson.
- The lesson index with numbers and problem ranges is `lessons/README.md`.

## 2. If `lessons/NN-<topic>/` doesn't exist yet, write it first
Copy the format of `lessons/01-arrays-hashing/` exactly (it's the reference):
- `README.md` with these sections, in order:
  1. Header: number, title, ⏱ time, week, problem IDs
  2. 🎯 After this lesson you can… (2–3 bullets)
  3. The idea in one line (+ a small ASCII diagram if it helps)
  4. Go basics needed (e.g. `container/heap`, recursion, queue)
  5. The patterns / templates: 2–4 named sub-patterns, each with a Go snippet
  6. The key insight: *why* it works (what interviewers want to hear)
  7. Signals table: "The problem says… | Think"
  8. Complexity
  9. Common mistakes ⚠️
  10. 🧪 Worked example: a LeetCode problem that is **not in `PROGRESS.md`**, solved with the core loop, with a trace and a 60-second explanation
  11. 🏋️ Warm-up exercise (how to run it)
  12. ✍️ Quiz: 4–6 questions, answers inside `<details>`
  13. ✅ Done when: a checklist. Then ➡️ Next lesson · Practice: problem IDs
- `exercise.go` (`package lesson`): 2–3 small functions that drill the templates. They must **not** be problems from `PROGRESS.md`. Stubs return zero values, and the doc comments give examples and the target Big-O.
- `exercise_test.go`: table-driven tests with edge cases. Use `ds` helpers for lists and trees.
- Verify: `go vet ./lessons/NN-<topic>/` passes, and the tests compile and fail against the stubs. Also check that a correct solution would pass: solve it privately in your head, or in a temp file that you then delete.
- Link the lesson in `lessons/README.md`.
- Tailor it: check `PROGRESS.md` and recent `card.md` "Mistake I made" lines, and add a mistake or tip that targets them.

## 3. Teach it interactively (don't dump the file)
Go section by section, **one chunk per message**:
1. Show the **idea + diagram**. Ask: "Where have you seen something like this?"
2. Show **one sub-pattern** at a time with its snippet. After each one, ask **one check question** (predict an output, spot a bug, or say when you'd use it). Wait for the answer, then correct it or praise it briefly.
3. Show the **key insight**. Ask them to explain it back in their own words.
4. Walk through the **worked example**. Before revealing each step, ask them what comes next.
5. Run the **warm-up**: they implement `exercise.go`, you run `go test ./lessons/NN-<topic>/`. Use tutor-style hints only. Never solve it for them.
6. Run the **quiz**: one question at a time. Track the score.

Keep each message to ≤ 12 lines plus code. Use the user's own words when you correct them. If they say "skip", move to the next section.

## 4. Finish
- Update the lesson row in `PROGRESS.md`: Status ✅ (warm-up passing + quiz ≥ 75%) or 🟡 (passed only partly), the Date (`date +%F`) and the quiz Score.
- Give a 3-line recap: the pattern, its signal and its template.
- Then say: "Ready. Say **next problem** to start <first problem ID + title of this topic>."
