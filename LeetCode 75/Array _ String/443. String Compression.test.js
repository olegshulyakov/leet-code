/*

Description

*/

/**
 * @param {character[]} chars
 * @return {number}
 */
var compress = function (chars) {
    if (chars.length === 0 || chars.length === 1) return chars.length;

    let compress = chars[0];
    let lastCharacter = chars[0];
    let counter = 1;
    for (let i = 1; i < chars.length; i++) {
        const currentCharacter = chars[i];

        if (currentCharacter === lastCharacter) {
            counter++;
            continue;
        }

        if (counter > 1) {
            compress += counter;
        }
        compress += currentCharacter;
        lastCharacter = currentCharacter;
        counter = 1;
    }

    if (counter > 1) {
        compress += counter;
    }

    for (let i = 0; i < compress.length; i++) {
        chars[i] = compress[i];
    }

    return compress.length;
};

describe('443. String Compression', () => {
    test('["a", "a", "b", "b", "c", "c", "c"]', () => expect(compress(["a", "a", "b", "b", "c", "c", "c"])).toEqual(6));
    test('["a"]', () => expect(compress(["a"])).toEqual(1));
    test('["a", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b"]', () => expect(compress(["a", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b"])).toEqual(4));
})