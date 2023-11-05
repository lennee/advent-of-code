"""
--- Day 5: Doesn't He Have Intern-Elves For This? ---

Santa needs help figuring out which strings in his text file are naughty or nice.

A nice string is one with all of the following properties:

    It contains at least three vowels (aeiou only), like aei, xazegov, or aeiouaeiouaeiou.
    It contains at least one letter that appears twice in a row, like xx, abcdde (dd), or aabbccdd (aa, bb, cc, or dd).
    It does not contain the strings ab, cd, pq, or xy, even if they are part of one of the other requirements.

For example:

    ugknbfddgicrmopn is nice because it has at least three vowels (u...i...o...), a double letter (...dd...), and none of the disallowed substrings.
    aaa is nice because it has at least three vowels and a double letter, even though the letters used by different rules overlap.
    jchzalrnumimnmhp is naughty because it has no double letter.
    haegwjzuvuyypxyu is naughty because it contains the string xy.
    dvszwmarrgswjxmb is naughty because it contains only one vowel.

How many strings are nice?

--- Part Two ---

Realizing the error of his ways, Santa has switched to a better model of determining whether a string is naughty or nice. None of the old rules apply, as they are all clearly ridiculous.

Now, a nice string is one with all of the following properties:

    It contains a pair of any two letters that appears at least twice in the string without overlapping, like xyxy (xy) or aabcdefgaa (aa), but not like aaa (aa, but it overlaps).
    It contains at least one letter which repeats with exactly one letter between them, like xyx, abcdefeghi (efe), or even aaa.

For example:

    qjhvhtzxzqqjkmpb is nice because is has a pair that appears twice (qj) and a letter that repeats with exactly one letter between them (zxz).
    xxyxx is nice because it has a pair that appears twice and a letter that repeats with one between, even though the letters used by each rule overlap.
    uurcxstgmygtbstg is naughty because it has a pair (tg) but no repeat with a single letter between them.
    ieodomkazucvgmuy is naughty because it has a repeating letter with one between (odo), but no pair that appears twice.

How many strings are nice under these new rules?

"""
from collections import defaultdict
import numpy as np

letters = defaultdict(int)
letters['a']= 1
letters['e']= 1
letters['i']= 1
letters['o']= 1
letters['u']= 1

illegal_combos = ["ab", "cd", "pq", "xy"]

def get_problem_input():
  with open("find_nice_input", "r") as f:
    return np.array(f.read().split("\n"))

def check_is_nice(s):
  prev_char = ""
  has_double = False
  vowel_count = 0

  for char in s:
    if char == prev_char:
      has_double = True
    else:
      if f'{prev_char}{char}' in illegal_combos:
        return False

    prev_char = char
    vowel_count += letters[char]

  return has_double and vowel_count > 2

def get_nice_count(arr):
  return np.array([check_is_nice(i) for i in arr], dtype=bool).sum()



"""
Realizing the error of his ways, Santa has switched to a better model of determining whether a string is naughty or nice. None of the old rules apply, as they are all clearly ridiculous.

Now, a nice string is one with all of the following properties:

    It contains a pair of any two letters that appears at least twice in the string without overlapping, like xyxy (xy) or aabcdefgaa (aa), but not like aaa (aa, but it overlaps).
    It contains at least one letter which repeats with exactly one letter between them, like xyx, abcdefeghi (efe), or even aaa.

For example:

    qjhvhtzxzqqjkmpb is nice because is has a pair that appears twice (qj) and a letter that repeats with exactly one letter between them (zxz).
    xxyxx is nice because it has a pair that appears twice and a letter that repeats with one between, even though the letters used by each rule overlap.
    uurcxstgmygtbstg is naughty because it has a pair (tg) but no repeat with a single letter between them.
    ieodomkazucvgmuy is naughty because it has a repeating letter with one between (odo), but no pair that appears twice.

How many strings are nice under these new rules?
"""

def check_is_nice_update(s):
  pairs = defaultdict(int)
  has_double = False
  has_repeat = False
  prev = 0

  for i, char in enumerate(s):


    if pairs[f'{s[i-1]}{char}'] >= 1 and i-prev >= 2:
      has_double = True
      prev = i

    if i < len(s) - 2 and s[i-1] == s[i+1]:
      has_repeat = True

    if pairs[f'{s[i-1]}{char}'] >= 1 and i-prev < 2:
      continue

    pairs[f'{s[i-1]}{char}'] += 1



  print(pairs)
  return has_double and has_repeat

def get_nice_count(arr):
  return np.array([check_is_nice_update(i) for i in arr], dtype=bool).sum()

if __name__ == "__main__":

  print(get_nice_count(get_problem_input()))