
def reverse_letters(w):
    chars = list(w)

    pointer_one = 0
    pointer_two = len(chars) - 1

    while pointer_one < pointer_two:
        chars[pointer_one], chars[pointer_two] = chars[pointer_two], chars[pointer_one]

        pointer_one += 1
        pointer_two -= 1

    return ''.join(chars)

def reverse_words(s):
    return ' '.join([reverse_letters(word) for word in s.split(" ")])
