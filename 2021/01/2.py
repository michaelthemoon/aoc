with open('./input') as f:
    a = int(f.readline())
    b = int(f.readline())
    c = int(f.readline())
    lines = f.readlines()


count = 0
for line in lines:
    line = int(line)
    if a+b+c < b+c+line:
        count += 1
    a, b, c = b, c, line

print(count)
