# Mistakes Log

> The `reviewer` skill adds a row on **every** review. Say **"mistakes"** for a summary, **"help me improve"** for drills.

## 📊 By category
| Category | Count | Last seen | Trend |
|---|---|---|---|
| ⏱ Big-O | 4 | 2026-11-21 | 🔁 recurring |
| 🧪 Edge cases | 2 | 2026-11-21 | 🔁 recurring |
| 🗣 Explanation | 3 | 2026-11-21 | 🔁 recurring |
| 🏷 Pattern / brute force | 2 | 2026-11-21 | 🔁 recurring |
| ✅ Correctness | 0 | | |
| 🐹 Go idioms | 1 | 2026-11-21 | 🟢 fixed |
| 📖 Readability | 0 | | |

**Trend:** 🔁 recurring (2+ times, still open) · 🟡 open · 🟢 fixed (not repeated in the last 3 reviews)

## 🧠 My rules (learned the hard way)
1. Bounded alphabet (26 letters) → space is O(1), not O(n).
2. Brute force = the simplest correct idea (usually all pairs, O(n²)). Sorting is an *alternative*, not the brute force.
3. Strings in Go are immutable: sorting one needs a copy → O(n) space.
4. Always list edge cases: empty, 1 element, negatives, extremes, duplicates at the ends.
5. Keep `card.md` in sync with the code after every fix.

## 📝 Log
| Date | Problem | Category | Mistake | Rule / fix | Status |
|---|---|---|---|---|---|
| 2026-11-21 | 001 Contains Duplicate | 🧪 Edge cases | Wrote "no edge cases" | Always list: 1 element, negatives, extremes, repeat at the ends | 🟢 fixed |
| 2026-11-21 | 001 Contains Duplicate | 🗣 Explanation | Said the map stores indexes; no Big-O in the 60-sec summary | Say what's stored + Big-O + why this beats alternatives | 🟢 fixed |
| 2026-11-21 | 001 Contains Duplicate | 🏷 Pattern / brute force | Called sorting the brute force | Brute force = compare all pairs, O(n²) | 🟢 fixed |
| 2026-11-21 | 001 Contains Duplicate | 🏷 Pattern / brute force | Pattern named "Duplicate finder" | Use the real name: "Hash set" | 🟢 fixed |
| 2026-11-21 | 001 Contains Duplicate | ⏱ Big-O | Missed that sort mutates input + O(log n) stack | Mention side effects and hidden stack space | 🟡 open |
| 2026-11-21 | 002 Valid Anagram | ⏱ Big-O | Space O(n) for a 26-letter map | Bounded alphabet → O(1) (or O(k)) | 🟢 fixed |
| 2026-11-21 | 002 Valid Anagram | ⏱ Big-O | Brute force said O(n³) | n chars × O(n) count = O(n²) | 🟢 fixed |
| 2026-11-21 | 002 Valid Anagram | ⏱ Big-O | Sorting a string said O(1) space | Go strings are immutable → copy to []byte = O(n) | 🟢 fixed |
| 2026-11-21 | 002 Valid Anagram | 🧪 Edge cases | Skipped adding own test cases (twice) | Add 2+ own rows to solution_test.go before review | 🟢 fixed |
| 2026-11-21 | 002 Valid Anagram | 🐹 Go idioms | Unneeded final loop over the map | If lengths match and no count goes negative, you're done | 🟢 fixed |
| 2026-11-21 | 002 Valid Anagram | 🗣 Explanation | card.md still described the removed loop | Update the card after every code change | 🟡 open |
| 2026-11-21 | 002 Valid Anagram | 🗣 Explanation | "Mistake I made" not in own words | Write it yourself, 1–2 sentences | 🟡 open |

## 🏋️ Drills
_None yet. Say "help me improve"._
