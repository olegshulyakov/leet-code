/*

Description

*/

/**
 *
 * @param {string} s
 * @returns {string}
 */
function func(s) {
    return s;
}

describe('', () => {
    test.each([
        { nums: [], expected: [] },
    ])('$nums', ({ nums, expected }) => expect(func(nums)).toEqual(expected));
})