import { describe, it } from 'node:test';
import assert from 'node:assert';

import { distanceSum } from './one';

describe('one', () => {
  it('should pass', () => {
    const distance = distanceSum(
      [3, 4, 2, 1, 3, 3],
      [4, 3, 5, 3, 9, 3],
    );
    assert.equal(distance, 11);
  });
});