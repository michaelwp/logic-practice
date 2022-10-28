import { flipFlop } from '../src/flip-flop';

import { expect } from 'chai';

describe('flipFlop tests', () => {
  describe('when given 45, 3, 5', () => {
    it('returns correct map', () => {
      let expected = generateExpectedMap();
      let received = flipFlop(45, 3, 5);
      
      expect(compareMaps(expected, received)).to.be.true;
    })
  })
})

function generateExpectedMap(): Map<number, string> {
  let expected = new Map<number, string>();

  expected.set(3, 'flip');
  expected.set(5, 'flop');
  expected.set(6, 'flip');
  expected.set(9, 'flip');
  expected.set(10, 'flop');
  expected.set(12, 'flip');
  expected.set(15, 'flip flop');
  expected.set(18, 'flip');
  expected.set(20, 'flop');
  expected.set(21, 'flip');
  expected.set(24, 'flip');
  expected.set(25, 'flop');
  expected.set(27, 'flip');
  expected.set(30, 'flip flop');
  expected.set(33, 'flip');
  expected.set(35, 'flop');
  expected.set(36, 'flip');
  expected.set(39, 'flip');
  expected.set(40, 'flop');
  expected.set(42, 'flip');
  expected.set(45, 'flip flop');

  return expected;
}

function compareMaps(expectedMap: Map<number, string>, receivedMap: Map<number, string>): boolean {
  if (receivedMap.size != expectedMap.size) {
    return false;
  }

  for (let key of Array.from(expectedMap.keys())) {
    let val = expectedMap.get(key);
    let testVal = receivedMap.get(key);
    if (testVal != val || (testVal === undefined && !receivedMap.has(key))) {
      return false;
    }
  }

  return true;
}
