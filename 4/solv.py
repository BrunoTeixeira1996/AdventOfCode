
def validate_passport_first_part(passport: str) -> bool:
    keys = ["byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"]

    for key in keys:
        if key not in passport:
            return False
    return True

def first_part(aux: list) -> int:
    counter = 0
    current_passport = ""
    valid_passports = []

    for line in aux:
        # building the current passport until we find a blank word
        if line != "":
            current_passport += " "  + line
        # if we reach a blank word we know we read all the passport key:values
        else:
            # validade the passport, True is valid, False is not
            if validate_passport_first_part(current_passport):
                valid_passports.append(current_passport)
                counter += 1

            # reseting the passport string
            current_passport = ""

    # checks the last passport since in the last string there is no blank word
    if validate_passport_first_part(current_passport):
        valid_passports.append(current_passport)
        counter += 1

    return counter, valid_passports


"""
====================================================================================================
"""
def validate_byr(birth_year: str) -> bool:
    if len(birth_year) == 4:
        if 1920 <= int(birth_year) <= 2002:
            return True

    return False

def validate_iyr(issue_year: str) -> bool:
    if len(issue_year) == 4:
        if 2010 <= int(issue_year) <= 2020:
            return True

    return False

def validate_eyr(expiration_year: str) -> bool:
    if len(expiration_year) == 4:
        if 2020 <= int(expiration_year) <= 2030:
            return True

    return False

def validate_hgt(height: str) -> bool:
    # grab the first 2 chars of the string
    unit = height[-2:]
    valid_units = ["in", "cm"]

    # checks if there is a "in" or a "cm"
    if unit not in valid_units:
        return False

    if unit == "cm":
        if 150 <= int(height[:-2]) <= 193:
            return True
    elif unit == "in":
        if 59 <= int(height[:-2]) <= 76:
            return True

    return False

def validate_hcl(hair_color: str) -> bool:
    # grabs the # if exit
    hashtag = hair_color[0]

    if hashtag != "#":
        return False

    if len(hair_color[1:]) != 6:
        return False

    

    return True

def validate_ecl(eye_color: str) -> bool:
    valid_eye_colors = ["amb", "blu", "brn", "gry", "grn", "hzl", "oth"]

    if eye_color not in valid_eye_colors:
        return False

    return True

def validate_pid(passport_id: str) -> bool:
    if len(passport_id) != 9:
        return False

    return True

def validate_passport_second_part(aux: list) -> bool:
    passports = aux.split()
    data = {}

    for passport in passports:
        key = passport[:3]
        value = passport[4:]
        data[key] = value

    if not validate_byr(data["byr"]): return False
    if not validate_iyr(data["iyr"]): return False
    if not validate_eyr(data["eyr"]): return False
    if not validate_hgt(data["hgt"]): return False
    if not validate_hcl(data["hcl"]): return False
    if not validate_ecl(data["ecl"]): return False
    if not validate_pid(data["pid"]): return False
    
    return True

def second_part(valid_passports: list) -> int:
    counter = 0

    for valid_passport in valid_passports:
        if validate_passport_second_part(valid_passport):
            counter += 1

    return counter


if __name__ == '__main__':
    
    with open("input", "r") as f:
        aux = f.readlines()
        aux = [ line.strip() for line in aux]

    total, valid_passports = first_part(aux)
    print(total)

    print(second_part(valid_passports))



        