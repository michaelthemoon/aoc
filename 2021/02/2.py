horizontal_position = 0
depth = 0
aim = 0

with open("./input") as f:
    lines = f.readlines()
    for line in lines:
        (instruction, movement) = line.split(" ")
        movement = int(movement)

        if instruction == "forward":
            horizontal_position += movement
            depth += movement * aim
        if instruction == "down":
            aim += movement
        if instruction == "up":
            aim -= movement
    print(horizontal_position*depth)
