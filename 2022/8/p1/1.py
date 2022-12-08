
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
res = 0

def is_tree_visible(trees, row, column):
    # edges
    if column == 0 or row == 0 or column == len_column - 1 or row == len_row - 1:
        return 1

    # look left from tree to edge
    # if all trees from left to source tree are low, we return 1 (its visible)
    if all(trees[row][c] < trees[row][column] for c in range(column)):
        return 1

    # look right from tree to edge
    # if all trees from source to right are low, we return 1 (its visible)
    if all(trees[row][c] < trees[row][column] for c in range(column+1, len_column)):
        return 1    

    # look up from tree to edge
    if all(trees[r][column] < trees[row][column] for r in range(row)):
        return 1

    # look down from tree to edge
    if all(trees[r][column] < trees[row][column] for r in range(row+1, len_row)):
        return 1

    return 0

for row in range(len_row):
    for column in range(len_column):
        res += is_tree_visible(trees, row, column)

print(res)