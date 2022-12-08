
with open("input", "r") as f:
    trees = [list(map(int, line.strip())) for line in f.readlines()]


""""
[
    [3, 0, 3, 7, 3],
    [2, 5, 5, 1, 2],
    [6, 5, 3, 3, 2],
    [3, 3, 5, 4, 9],
    [3, 5, 3, 9, 0]
]
"""
len_row = len(trees)
len_column = len(trees[0])
res = []

def tree_scenic(trees, row, column):
    left = 0
    right = 0
    up = 0
    down = 0

    # edges so we don't care since multiplying with 0 will always give 0
    if column == 0 or row == 0 or column == len_column - 1 or row == len_row - 1:
        return 0
    
    # look left from tree to edge
    for c in range(1, column+1):
        if trees[row][column - c] >= trees[row][column]:
            break
    left = c

    # look right from tree to edge
    for c in range(1, len_column - column):
        if trees[row][column + c] >= trees[row][column]:
            break
    right = c

    # look up from tree to edge
    for r in range(1, row+1):
        if trees[row - r][column] >= trees[row][column]:
            break
    up = r

    # look down from tree to edge
    for r in range(1, len_row - row):
        if trees[row + r][column] >= trees[row][column]:
            break
    down = r

    return left * right * up * down


for row in range(len_row):
    for column in range(len_column):
        res.append(tree_scenic(trees, row, column)) 

print(max(res))