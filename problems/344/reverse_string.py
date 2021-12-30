# Reverse using two pointers technique.


def reverse_string(string):
    return list(reversed(string))


def reverse_string_two_pointers(string):
    pointer_one = 0
    pointer_two = len(string) - 1

    while pointer_one < pointer_two:
        string[pointer_one], string[pointer_two] = (
            string[pointer_two],
            string[pointer_one],
        )

        pointer_one += 1
        pointer_two -= 1

    return string
