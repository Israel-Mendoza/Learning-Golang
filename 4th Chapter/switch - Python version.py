# A list of strings
words = ["a", "cow", "smile", "gopher", "octopus", "anthropologist"]

for word in words:
    match len(word):
        case 1 | 2 | 3 | 4:
            print(f"{word!r} is a short word!")
        case 5:
            print(f"{word!r} is exactly the right length: {len(word)}")
        case 6 | 7 | 8 | 9:
            pass
        case other:
            print(f"{word!r} is a long word! Its length is {other!r}")
# 'a' is a short word!
# 'cow' is a short word!
# 'smile' is exactly the right length: 5
# 'anthropologist' is a long word! Its length is 14