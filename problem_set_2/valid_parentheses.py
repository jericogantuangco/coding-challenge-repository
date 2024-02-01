

def isValid(s: str) -> bool:
    """
    This functions checks whether the parentheses are valid or not
    """
    database = {
        '(': ')',
        '{': '}',
        '[': ']',
    }

    stack = []

    if len(s) <= 1:
        return False

    for character in s:
        if character in database.keys():
            stack.append(character)
        elif stack and database[stack[-1]] == character:
            stack.pop()
        else:
            return False

    if len(stack) > 0:
        return False

    return True

if __name__ == '__main__':

    print(isValid('()[]{}'))
    print(isValid('([)]'))
    print(isValid('()'))