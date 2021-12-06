with open('./input') as f:
    firstLine = f.readline()
    lines = f.readlines()

count = 0
prev = firstLine
for line in lines:
    if int(line) > int(prev):
        count += 1
    prev = line

print(count)
