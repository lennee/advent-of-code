"""
Here's an easy puzzle to warm you up.

Santa is trying to deliver presents in a large apartment building, but he can't find the right floor - the directions he got are a little confusing. He starts on the ground floor (floor 0) and then follows the instructions one character at a time.

An opening parenthesis, (, means he should go up one floor, and a closing parenthesis, ), means he should go down one floor.

The apartment building is very tall, and the basement is very deep; he will never find the top or bottom floors.

For example:

    (()) and ()() both result in floor 0.
    ((( and (()(()( both result in floor 3.
    ))((((( also results in floor 3.
    ()) and ))( both result in floor -1 (the first basement level).
    ))) and )())()) both result in floor -3.

To what floor do the instructions take Santa?
"""

def get_problem_input() -> str:
  with open("floors_input_2", "r") as f:
    return f.readlines()[0]

# Day 1 Problem, Part 1
def find_floor_number(paren_map: str) -> int:
  cur_floor = 0

  for i in paren_map:
    if i == "(": cur_floor += 1
    elif i == ")": cur_floor -= 1

  return cur_floor

# Day 1 Problem, Part 2
def find_first_basement_index(paren_map: str) -> int:
  cur_floor = 0

  for i, v in enumerate(paren_map):
    if v == "(": cur_floor += 1
    elif v == ")": cur_floor -= 1

    if cur_floor < 0: return i + 1

  return -1

if __name__ == "__main__":
  print(find_first_basement_index(get_problem_input()))