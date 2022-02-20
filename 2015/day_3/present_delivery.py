"""
Santa is delivering presents to an infinite two-dimensional grid of houses.

He begins by delivering a present to the house at his starting location, and then an elf at the North Pole calls him via radio and tells him where to move next. Moves are always exactly one house to the north (^), south (v), east (>), or west (<). After each move, he delivers another present to the house at his new location.

However, the elf back at the north pole has had a little too much eggnog, and so his directions are a little off, and Santa ends up visiting some houses more than once. How many houses receive at least one present?

For example:

    > delivers presents to 2 houses: one at the starting location, and one to the east.
    ^>v< delivers presents to 4 houses in a square, including twice to the house at his starting/ending location.
    ^v^v^v^v^v delivers a bunch of presents to some very lucky children at only 2 houses.


"""

from collections import defaultdict

def get_problem_input() -> str:
  with open("present_delivery_input", "r") as f:
    return f.readlines()[0]

def update_index(cur_index, move):
  index = cur_index
  if move == ">": return (index[0] + 1, index[1])
  if move == "<": return (index[0] - 1, index[1])
  if move == "^": return (index[0], index[1] + 1)
  if move == "v": return (index[0], index[1] - 1)
  return index

def deliver_presents_alternate(delivery_input):
  deliveries = defaultdict(int)
  santa_index = (0, 0)
  deliveries[f'{santa_index[0]},{santa_index[1]}'] += 2
  for i in delivery_input:
    santa_index = update_index(santa_index, i)
    deliveries[f'{santa_index[0]},{santa_index[1]}'] += 1

  return len(deliveries.keys())

def deliver_presents_alternate(delivery_input):
  deliveries = defaultdict(int)
  santa_index = (0, 0)
  robo_index = (0, 0)
  deliveries[f'{santa_index[0]},{santa_index[1]}'] += 2
  robo = False
  for i in delivery_input:
    index = robo_index if robo else santa_index
    index = update_index(index, i)
    deliveries[f'{index[0]},{index[1]}'] += 1
    if robo:
      robo_index = index
    else:
      santa_index = index

    robo = not robo

  return len(deliveries.keys())


if __name__ == "__main__":
  print(deliver_presents_alternate(get_problem_input()))