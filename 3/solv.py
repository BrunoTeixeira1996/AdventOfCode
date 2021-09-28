from math import prod



# traversing the map at the top-left corner of the map and following a slope of right 3 and down 1
def first_part(aux: list) -> int:
	right = 3
	counter = 0
	list_len = len(aux[0])

	# iterates over the lix from 0 to len(aux) 1 by 1
	for i in range(0, len(aux), 1):
		"""
		# goes down the list (aux[i]), and goes 3 to the side (aux[i][(i * right)])
		# but using % operator we can reach the end of the row and go back to the begining
		# example : 
		# i = 0, aux[0][0]
		# i = 1, aux[1][3]
		# i = 2, aux[2][6]
		# ...
		# i = 10, aux[10][30]
		# i = 11, aux[11][2]
		# i = 12, aux[12][5]
		# i = 13, aux[13][8]
		"""
		letter = aux[i][(i * right) % list_len]
		if letter == "#":
			counter += 1

	return counter


def second_part(aux: list, right: int, down: int) -> int:
	counter = 0
	list_len = len(aux[0])

	for i in range(0, len(aux)):
		# if i * down > len(aux) we reached the end of the list of lists since now we have a down 2 in the slopes
		if i * down > len(aux):
			break

		# aux[i*down] because now we have a down 2 or more
		letter = aux[i * down][(i * right) % list_len]
		if letter == "#":
			counter += 1

	return counter


if __name__ == '__main__':
	aux = []
	with open("input", "r") as f:
		for line in f:
			aux.append(line.strip())

	print(first_part(aux))

	slopes = [(1,1), (3,1), (5,1), (7,1), (1,2)]
	print(prod([second_part(aux, slope[0], slope[1]) for slope in slopes]))