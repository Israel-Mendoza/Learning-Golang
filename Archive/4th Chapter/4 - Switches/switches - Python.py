# A list of strings
words = ["a", "cow", "smile", "gopher", "octopus", "anthropologist"]

for word in words:
    # Declating a variable called 
    # "size" and "switching" through it:
    # In Python, there is no fall through!
    match size := len(word):
        case 1 | 2 | 3 | 4:
            print(f"{word!r} is a short word!")
        case 5:
            print(f"{word!r} is exactly the right length: {size}")
        case 6 | 7 | 8 | 9: # Empty case, nothing happens
            pass
        case other:
            print(f"{word!r} is a long word! Its length is {size}")
# 'a' is a short word!
# 'cow' is a short word!
# 'smile' is exactly the right length: 5
# 'anthropologist' is a long word! Its length is 14


for i in range(10):
    if i % 2 == 0:
        print(f"{i} is even")
    elif i % 3 == 0:
        print(f"{i} is divisible by 3 but not 2")
    elif i % 7 == 0:
	    print("exit the loop!")
	    break
    else:
        print(f"{i} is boring")