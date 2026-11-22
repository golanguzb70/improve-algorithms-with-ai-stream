# Mistakes Log

> The `reviewer` skill adds a row on **every** review. Say **"mistakes"** for a summary, **"help me improve"** for drills.

## 📊 By category
| Category | Count | Last seen | Trend |
|---|---|---|---|
| ⏱ Big-O | 7 | 2026-11-22 | 🔁 recurring |
| 🧪 Edge cases | 3 | 2026-11-22 | 🔁 recurring (but 004 + 005 had own rows ✅) |
| 🗣 Explanation | 12 | 2026-11-22 | 🔁 recurring (005 60-sec was good ✅) |
| 🏷 Pattern / brute force | 4 | 2026-11-22 | 🔁 recurring |
| ✅ Correctness | 0 | | |
| 🐹 Go idioms | 3 | 2026-11-22 | 🟢 fixed |
| 📖 Readability | 5 | 2026-11-22 | 🟢 fixed (all rows) |

**Trend:** 🔁 recurring (2+ times, still open) · 🟡 open · 🟢 fixed (not repeated in the last 3 reviews)

## 🧠 My rules (learned the hard way)
1. Constant only if the problem fixes it (26 letters → O(1)). Input lengths (m ≤ 100) stay as variables.
2. Brute force = the simplest correct idea (usually all pairs, O(n²)). Sorting is an *alternative*, not the brute force.
3. Strings in Go are immutable: sorting one needs a copy → O(n) space.
4. Edge cases go in **solution_test.go**, not just the card: 1 element, negatives, zero, extremes, duplicates, answer at the ends.
5. After every fix, update **every** place: solution.go header, card, 60-sec.
6. Use standard pattern names (Hash set, Hash map, Two pointers…). Signals = the phrase in the problem + what it implies.
7. 60-sec explanation = idea → why it beats brute force → Big-O. Short sentences.
8. Name every variable (n, m, k…), then add up **every** step: count O(n) + sort O(m log m) = O(n + m log m).

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
| 2026-11-22 | 003 Two Sum | 🧪 Edge cases | Listed edge cases in card, added no test rows | Add 2+ own rows to solution_test.go before review | 🟡 open |
| 2026-11-22 | 003 Two Sum | 🏷 Pattern / brute force | Pattern named "Index Map" | Say "Hash map (complement lookup)" | 🟢 fixed |
| 2026-11-22 | 003 Two Sum | 🗣 Explanation | 60-sec summary is one run-on; no "why it beats O(n²)" | Idea → why O(1) lookup replaces inner loop → Big-O | 🟡 open |
| 2026-11-22 | 003 Two Sum | 🗣 Explanation | "Mistake I made" left empty again | Fill it every time, 1–2 sentences | 🟢 fixed |
| 2026-11-22 | 003 Two Sum | 📖 Readability | Two passes where one works; vague names mp, e | One pass: check, then store; name it indexOf | 🟢 fixed |
| 2026-11-22 | 003 Two Sum | 🗣 Explanation | Card idea + 60-sec still describe the old two-pass code | Re-read the card after every code change | 🟡 open |
| 2026-11-22 | 003 Two Sum | 🏷 Pattern / brute force | Signals too vague: "two numbers that add up" | Signal = phrase in problem + what it implies | 🟡 open |
| 2026-11-22 | 004 Group Anagrams | ⏱ Big-O | Said O(n) time/space; dropped word length m | Input lengths stay variables: O(n·m log m), O(n·m) | 🟢 fixed |
| 2026-11-22 | 004 Group Anagrams | ⏱ Big-O | Brute force O(n³): mixed up n (words) and m (letters) | Name each variable; pairs × O(m) check = O(n²·m) | 🟢 fixed |
| 2026-11-22 | 004 Group Anagrams | 🐹 Go idioms | SortFunc + cmp.Compare; magic capacity len/4 | slices.Sort for bytes; no guessed capacities | 🟢 fixed |
| 2026-11-22 | 004 Group Anagrams | 🗣 Explanation | 60-sec ~150 words, opens with Big-O, "why" is muddled | ≤120 words; why = one lookup replaces comparing every word | 🟡 open |
| 2026-11-22 | 004 Group Anagrams | 🗣 Explanation | "Mistake I made" not an actual mistake | Name the real slip, e.g. "dropped m" | 🟢 fixed |
| 2026-11-22 | 004 Group Anagrams | 🗣 Explanation | solution.go header + 60-sec still say O(n) after fixing card | After a fix, update every place: code header, card, 60-sec | 🟢 fixed |
| 2026-11-22 | 004 Group Anagrams | 📖 Readability | Vague names (e, responseStrGroups); file not gofmt'd | word, groups; run gofmt before review | 🟢 fixed |
| 2026-11-22 | 004 Group Anagrams | 📖 Readability | Stale comments still mention old name `e` | Rename → search old name in comments too | 🟢 fixed |
| 2026-11-22 | 005 Top K Frequent | ⏱ Big-O | Said O(m log m); dropped the O(n) counting loop | Add every step: O(n + m log m) → worst O(n log n) | 🟢 fixed |
| 2026-11-22 | 005 Top K Frequent | 🗣 Explanation | "Mistake I made" left as a placeholder | Fill it every time, 1–2 sentences | 🟢 fixed |
| 2026-11-22 | 005 Top K Frequent | 🗣 Explanation | 60-sec opens with Big-O, ~130 words | Idea → why → Big-O, ≤120 words | 🟢 fixed |
| 2026-11-22 | 005 Top K Frequent | 🐹 Go idioms | `-1 * cmp.Compare(a, b)` to sort descending | Swap the args: cmp.Compare(b, a) | 🟢 fixed |
| 2026-11-22 | 005 Top K Frequent | 📖 Readability | Typo `uniquNumbers`; test file not gofmt'd | Spell-check names; gofmt -w before review | 🟢 fixed |
| 2026-11-22 | 005 Top K Frequent | 📖 Readability | Redundant checks (len>0, k checked twice); names topFrequents/response | Put the stop in the loop condition; name it buckets | 🟢 fixed |
| 2026-11-22 | 005 Top K Frequent | 🗣 Explanation | "Optimal idea" copied the tutor's hint word for word | Rewrite every hint in your own words | 🟡 open |

## 🏋️ Drills
_None yet. Say "help me improve"._
