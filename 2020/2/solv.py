
# returns how many passwords are valid for the first part
def first_part(aux: list) -> int:

	counter = 0
	# iterates over a list of lists
	for i in aux:
		# counts how many times does the letters appears in the password 
		char_occurr = i[3].count(i[2])

		# grab the min and max values
		mini,maxi = int(i[0]), int(i[1])

		# check if its between
		if mini <= char_occurr <= maxi:
			counter += 1

	return counter

# returns how many passwords are valid for the second part
def second_part(aux: list) -> int:

	counter = 0
	# iterates over a list of lists
	for i in aux:

		# grabs the first and the second indexes and the first and second letters
		first_index, second_index = int(i[0]), int(i[1])
		first_letter, second_letter = i[3][first_index-1], i[3][second_index-1]

		# if the letters are not the same and at least one of them is the real letter
		if first_letter != second_letter and (first_letter == i[2] or second_letter == i[2]):
			counter += 1

	return counter

if __name__ == '__main__':
	aux = []
	with open("input", "r") as f:
		for line in f:
			aux.append(line.replace("-"," ").replace(":"," ").split())
	
	print(first_part(aux))
	print(second_part(aux))
