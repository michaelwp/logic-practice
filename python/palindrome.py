def isPalindrome(str):
    str = str.lower()
    if len(str) < 2:
        return True
    if str[0] != str[len(str) - 1]:
        return False
    return isPalindrome(str[1 : len(str) - 1])
