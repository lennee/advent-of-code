"""
Because your neighbors keep defeating you in the holiday house decorating contest year after year, you've decided to deploy one million lights in a 1000x1000 grid.

Furthermore, because you've been especially nice this year, Santa has mailed you instructions on how to display the ideal lighting configuration.

Lights in your grid are numbered from 0 to 999 in each direction; the lights at each corner are at 0,0, 0,999, 999,999, and 999,0. The instructions include whether to turn on, turn off, or toggle various inclusive ranges given as coordinate pairs. Each coordinate pair represents opposite corners of a rectangle, inclusive; a coordinate pair like 0,0 through 2,2 therefore refers to 9 lights in a 3x3 square. The lights all start turned off.

To defeat your neighbors this year, all you have to do is set up your lights by doing the instructions Santa sent you in order.

For example:

    turn on 0,0 through 999,999 would turn on (or leave on) every light.
    toggle 0,0 through 999,0 would toggle the first line of 1000 lights, turning off the ones that were on, and turning on the ones that were off.
    turn off 499,499 through 500,500 would turn off (or leave off) the middle four lights.

After following the instructions, how many lights are lit?
"""

def get_problem_input():
  with open("./day_6/lit_lights_input", "r") as f:
    return [item.split(" ") for item in f.read().split("\n")]

class Light:

    def __init__(self): self.val = 0

    def turnOn(self): self.val = 1

    def turnOff(self): self.val = 0

    def toggle(self): self.val = 1 if self.val == 0 else 0

    def process(self, command):
        if command == "turn on":
            self.turnOn()
        elif command == "turn off":
            self.turnOff()
        elif command == "toggle":
            self.toggle()

class Board:

    def __init__(self, x, y):
        self.board = [[Light() for i in range(y)] for j in range(x)]

    def process_command(self, input_arr):
        command = input_arr[:2].join(" ") if input_arr == "turn" else "toggle"
        point_1 = sorted([int(i) - 1 for i in input_arr[-3].split(",")])
        point_2 = sorted([int(i) - 1 for i in input_arr[-1].split(",")])
        i = point_1[0]
        j = point_1[1]
        while i <= point_2[0]:
            while j <= point_2[1]:
                self.board[i][j].process(command)
                j += 1
            j = point_1[1]
            i += 1

    def num_on(self):
        count = 0
        for i in self.board:
            for j in i:
                count += j.val
        return count

if __name__ == "__main__":
    board = Board(1000,1000)
    for i, command in enumerate(get_problem_input()):
        print(f'processing command {i}...')
        board.process_command(command)
    print(board.num_on())