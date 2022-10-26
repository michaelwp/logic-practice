import { isPalindrome } from '../src/palindrome'

import { expect } from 'chai'

describe('isPalindrome tests', () => {
  describe('when given a string of length 0', () => {
    it('returns true', () => {
      expect(isPalindrome('')).to.be.true;
    })
  })

  describe('when given a string of length 1', () => {
    it('returns true', () => {
      expect(isPalindrome('a')).to.be.true;
    })
  })

  describe('when given string "aa"', () => {
    it('returns true', () => {
      expect(isPalindrome('aa')).to.be.true;
    })
  })

  describe('when given string "aba"', () => {
    it('returns true', () => {
      expect(isPalindrome('aba')).to.be.true;
    })
  })

  describe('when given string "abc"', () => {
    it('returns false', () => {
      expect(isPalindrome('abc')).to.be.false;
    })
  })

  describe('when given string "tacocat"', () => {
    it('returns true', () => {
      expect(isPalindrome('tacocat')).to.be.true;
    })
  })

  describe('when given string "racecar"', () => {
    it('returns true', () => {
      expect(isPalindrome('racecar')).to.be.true;
    })
  })

  describe('when given string "notpalindrome"', () => {
    it('returns false', () => {
      expect(isPalindrome('notpalindrome')).to.be.false;
    })
  })
})
