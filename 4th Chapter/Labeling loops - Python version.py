samples = ["hello", "apple_π!"]

for sample in samples:
    for char in sample:
        print(char)
        if char == "l":
            break
    print()
