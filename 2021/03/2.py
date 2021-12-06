parse_lines = lambda f: [l.rstrip("\n") for l in f]
with open("./input") as f:
    lines = parse_lines(f.readlines())

def convertToDecimal(binaryString):
    result = 0
    for i, s in enumerate(binaryString):
        if s == "1":
            result += 2**(len(binaryString)-i-1)
    return result


def filter(lines, currCharIndex, inverse):
    if len(lines) == 1:
        return lines[0]

    count = 0
    for line in lines:
        if line[currCharIndex] == "1":
            count += 1

    totalLength = len(lines)

    x = 1 if  count >= totalLength / 2 else 0
    if inverse:
        x += 1
    char = "1" if x % 2 != 0 else "0"

    nextLines = []
    for line in lines:
        if line[currCharIndex] == char:
            nextLines.extend([line])

    return filter(nextLines, currCharIndex+1, inverse)


print(convertToDecimal(filter(lines, 0, False)) * convertToDecimal(filter(lines, 0, True)))
