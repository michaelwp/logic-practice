export function isPalindrome(phrase: string): boolean {
  if (phrase.length < 2) {
    return true;
  }

  if (phrase[0] != phrase[phrase.length - 1]) {
    return false;
  }

  return isPalindrome(phrase.substring(1, phrase.length - 1))
}
