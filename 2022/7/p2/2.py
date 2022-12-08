from collections import defaultdict

with open("input", "r") as f:
    lines = f.readlines()

path = []
dirsizes = defaultdict(int)
for line in lines:
    tokens = line.strip().split(" ")
    
    if tokens[0] == "$" and tokens[1] == "cd":
        if tokens[2] == "/":
            path = ["/"]
        elif tokens[2] == "..":
            path.pop()
        else:
            path.append(tokens[2])
    elif tokens[0] == "$" and tokens[1] == "ls":
        pass
    else:
        try:
            for i in range(len(path)):
                dirsizes["".join(path[:i+1])] += int(tokens[0])
        except:
            pass

max_disk_space = 70000000
need = 30000000
need_to_free = need - (max_disk_space - dirsizes["/"])

res = min(v for v in dirsizes.values() if v >= need_to_free)

print(res)