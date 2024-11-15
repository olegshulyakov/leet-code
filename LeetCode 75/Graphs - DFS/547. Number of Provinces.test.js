/*

There are n cities. Some of them are connected, while some are not. If city a is connected directly with city b, and city b is connected directly with city c, then city a is connected indirectly with city c.

A province is a group of directly or indirectly connected cities and no other cities outside of the group.

You are given an n x n matrix isConnected where isConnected[i][j] = 1 if the ith city and the jth city are directly connected, and isConnected[i][j] = 0 otherwise.

Return the total number of provinces.


Example 1:

Input: isConnected = [[1,1,0],[1,1,0],[0,0,1]]
Output: 2


Example 2:

Input: isConnected = [[1,0,0],[0,1,0],[0,0,1]]
Output: 3


Constraints:

1 <= n <= 200
n == isConnected.length
n == isConnected[i].length
isConnected[i][j] is 1 or 0.
isConnected[i][i] == 1
isConnected[i][j] == isConnected[j][i]

*/

/**
 * @param {number[][]} isConnected
 * @return {number}
 */
var findCircleNum = function (isConnected) {
    if (!isConnected || isConnected.length === 0) return 0;

    const isVisited = new Array(isConnected.length);
    const visitCity = (city) => {
        isVisited[city] = true;

        for (let neighbor = 0; neighbor < isConnected.length; neighbor++) {
            if (!isVisited[neighbor] && isConnected[city][neighbor]) {
                visitCity(neighbor);
            }
        }
    }

    let provinces = 0;
    for (let city = 0; city < isConnected.length; city++) {
        if (!isVisited[city]) {
            visitCity(city);
            provinces++;
        }
    }
    return provinces;
};

describe('547. Number of Provinces', () => {
    test.each([
        { isConnected: [], expected: 0 },
        { isConnected: [[1, 1, 0], [1, 1, 0], [0, 0, 1]], expected: 2 },
        { isConnected: [[1, 0, 0], [0, 1, 0], [0, 0, 1]], expected: 3 },
        { isConnected: [[1, 0, 0, 1], [0, 1, 1, 0], [0, 1, 1, 1], [1, 0, 1, 1]], expected: 1 },
    ])('$isConnected', ({ isConnected, expected }) => expect(findCircleNum(isConnected)).toEqual(expected));
})