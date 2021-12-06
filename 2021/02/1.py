horizontal_position = 0
depth = 0

with open("./input") as f:
    lines = f.readlines()
    for line in lines:
        (instruction, movement) = line.split(" ")
        if instruction == "forward":
            horizontal_position += int(movement)
        if instruction == "down":
            depth += int(movement)
        if instruction == "up":
            depth -= int(movement)

    print(horizontal_position * depth)
