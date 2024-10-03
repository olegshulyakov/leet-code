/*

We are given an array asteroids of integers representing asteroids in a row.

For each asteroid, the absolute value represents its size, and the sign represents its direction (positive meaning right, negative meaning left). Each asteroid moves at the same speed.

Find out the state of the asteroids after all collisions. If two asteroids meet, the smaller one will explode. If both are the same size, both will explode. Two asteroids moving in the same direction will never meet.


Example 1:

Input: asteroids = [5,10,-5]
Output: [5,10]
Explanation: The 10 and -5 collide resulting in 10. The 5 and 10 never collide.


Example 2:

Input: asteroids = [8,-8]
Output: []
Explanation: The 8 and -8 collide exploding each other.


Example 3:

Input: asteroids = [10,2,-5]
Output: [10]
Explanation: The 2 and -5 collide resulting in -5. The 10 and -5 collide resulting in 10.


Constraints:

2 <= asteroids.length <= 10^4
-1000 <= asteroids[i] <= 1000
asteroids[i] != 0

*/

/**
 * @param {number[]} asteroids
 * @return {number[]}
 */
var asteroidCollision = function (asteroids) {
    const stack = [];
    for (let i = 0; i < asteroids.length; i++) {
        if (stack.length === 0) {
            stack.push(asteroids[i]);
            continue;
        }

        let last = stack.pop();
        if (last > 0 && asteroids[i] < 0) {
            while (stack.length > 0 && last > 0 && last < -asteroids[i]) {
                last = stack.pop();
            }

            if (last > 0) {
                if (last > -asteroids[i]) {
                    stack.push(last);
                } else if (last < -asteroids[i]) {
                    stack.push(asteroids[i]);
                }
            } else {
                stack.push(last);
                stack.push(asteroids[i]);
            }
        }
        else {
            stack.push(last);
            stack.push(asteroids[i]);
        }
    }
    return stack;
};

describe('735. Asteroid Collision', () => {
    test.each([
        { asteroids: [], expected: [] },
        { asteroids: [5, 10, -5], expected: [5, 10] },
        { asteroids: [8, -8], expected: [] },
        { asteroids: [10, 2, -5], expected: [10] },
        { asteroids: [-2, -1, 1, 2], expected: [-2, -1, 1, 2] },
        { asteroids: [-2, -2, 1, -2], expected: [-2, -2, -2] },
        { asteroids: [-2, -2, 1, -1], expected: [-2, -2] },
        { asteroids: [-2, 2, 1, -2], expected: [-2] },
        { asteroids: [1, -2, -2, -2], expected: [-2, -2, -2] },
    ])('$asteroids', ({ asteroids, expected }) => expect(asteroidCollision(asteroids)).toEqual(expected));
})