class Person:

	def __init__(self, name: str, age: int):
		self.name = name
		self.age = age

	def __repr__(self) -> str:
		return f"Person('{self.name}', {self.age})"



def edit_attempt(a_string: str, a_number: int, a_person: Person) -> None:
	a_string = "Look, I'm a new string!"
	a_number *= 2
	a_person.name = "Juanito"
	a_person.age = 100


def edit_dict(a_dict: dict) -> None:
	a_dict[2] = "dos"
	a_dict[3] = "tres"
	a_dict[100] = "cien"
	del a_dict[1]


def edit_list(a_list: list[int]) -> None:
	for index, value in enumerate(a_list):
		a_list[index] = value * value
	a_list.append(1234567890)


if __name__ == "__main__":
	s = "Hello"
	n = 25
	p = Person("Israel", 29)
	a_dict = {0: "zero", 1: "one", 2: "two"}
	a_list = [1, 2, 3, 4, 5]


	print(f"String: {s}\nNumber: {n}\nPerson: {p}\n\n")
	# String: Hello
	# Number: 25
	# Person: Person('Israel', 29)

	edit_attempt(s, n, p)

	print(f"String: {s}\nNumber: {n}\nPerson: {p}\n\n")
	# String: Hello
	# Number: 25
	# Person: Person('Juanito', 100)

	print(a_dict)
	# {0: 'zero', 1: 'one', 2: 'two'}
	edit_dict(a_dict)
	print(a_dict)
	# {0: 'zero', 2: 'dos', 3: 'tres', 100: 'cien'}

	print(a_list)
	# [1, 2, 3, 4, 5]
	edit_list(a_list)
	print(a_list)
	# [1, 4, 9, 16, 25, 1234567890]
