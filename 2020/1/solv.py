# iterate over a list and sum every number with each other
def first_part(numbers: list) -> int:
	res_part_one = 1
	for value in numbers:
		for next_value in numbers:
			if value + next_value == 2020:
				res_part_one = res_part_one * next_value
				return res_part_one

# iterate over a list and sum every number with each other
def second_part(numbers: list) -> int:
	res_part_two = 1
	for value in numbers:
		for next_value in numbers:
			for next_next_value in numbers:
				if value + next_value + next_next_value == 2020:
					res_part_two = value * next_value * next_next_value
					return res_part_two



if __name__ == '__main__':
	numbers = []
	# reads every number and appends to a list
	with open("input", "r") as f:
		for line in f:
			numbers.append(int(line))

	print(first_part(numbers))
	print(second_part(numbers))


