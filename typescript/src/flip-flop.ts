export function flipFlop(num: number, key1: number, key2: number): Map<number, string> {
  let result =  new Map<number, string>;

  for (let i = 1; i <= num; i++) {
    if (i % (key1 * key2) == 0) {
      result.set(i, 'flip flop');
    } else if (i % key1 == 0) {
      result.set(i, 'flip');
    } else if (i % key2 == 0) {
      result.set(i, 'flop');
    }
  }

  return result;
}
