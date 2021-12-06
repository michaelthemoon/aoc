parse_lines = lambda f: [l.rstrip("\n") for l in f]
with open("./input") as f:
    lines = parse_lines(f.readlines())

counts = {}
totalRecords = 0
binaryLength = len(lines[0])

for i in range(binaryLength):
    counts[i] = 0

for line in lines:
    totalRecords += 1
    for i, s in enumerate(line):
        counts[i] += int(s)

gammaRate = 0
epsilonRate = 0
for k, v in counts.items():
    # We have more 1s than zeros
    actualBinaryPosition = binaryLength - k - 1
    if v > totalRecords / 2:
        gammaRate += 2**actualBinaryPosition
    else:
        epsilonRate += 2**actualBinaryPosition

print(gammaRate * epsilonRate)
