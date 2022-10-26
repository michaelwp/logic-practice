import unittest
from palindrome import isPalindrome

class Palindrome(unittest.TestCase):
    def setUp(self):
        self.testcase1 = ''
        self.testcase2 = 'a'
        self.testcase3 = 'aa'
        self.testcase4 = 'aba'
        self.testcase5 = 'abc'
        self.testcase6 = 'tacocat'
        self.testcase7 = 'racecar'
        self.testcase8 = 'notpalindrome'
    
    def tearDown(self):
        """Executed after each test"""
        pass

    def test_is_empty_sting_palindrome(self):
        is_palindrome = isPalindrome(self.testcase1)

        self.assertEqual(is_palindrome, True)
    
    def test_is_a_palindrome(self):
        is_palindrome = isPalindrome(self.testcase2)

        self.assertEqual(is_palindrome, True)
    
    def test_is_aa_palindrome(self):
        is_palindrome = isPalindrome(self.testcase3)

        self.assertEqual(is_palindrome, True)
    
    def test_is_aba_palindrome(self):
        is_palindrome = isPalindrome(self.testcase4)

        self.assertEqual(is_palindrome, True)
    
    def test_is_abc_palindrome(self):
        is_palindrome = isPalindrome(self.testcase5)

        self.assertEqual(is_palindrome, False)
    
    def test_is_tacocat_palindrome(self):
        is_palindrome = isPalindrome(self.testcase6)

        self.assertEqual(is_palindrome, True)
    
    def test_is_racecar_palindrome(self):
        is_palindrome = isPalindrome(self.testcase7)

        self.assertEqual(is_palindrome, True)
    
    def test_is_notpalindrome_palindrome(self):
        is_palindrome = isPalindrome(self.testcase8)

        self.assertEqual(is_palindrome, False)

if __name__ == '__main__':
    unittest.main()

