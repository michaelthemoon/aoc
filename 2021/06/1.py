with open("./input") as f:
    fishies = f.readline().rstrip("\n").split(",")

counts = {
    "0": 0,
    "1": 0,
    "2": 0,
    "3": 0,
    "4": 0,
    "5": 0,
    "6": 0,
    "7": 0,
    "8": 0,
}

for fish in fishies:
    counts[fish] += 1

for i in range(80):
    tmp = counts["0"]
    counts["0"] = counts["1"]
    counts["1"] = counts["2"]
    counts["2"] = counts["3"]
    counts["3"] = counts["4"]
    counts["4"] = counts["5"]
    counts["5"] = counts["6"]
    counts["6"] = counts["7"] + tmp
    counts["7"] = counts["8"]
    counts["8"] = tmp

total = 0
for count in counts.values():
    total += count

print(total)
