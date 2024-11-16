/*

There are n cities numbered from 0 to n - 1 and n - 1 roads such that there is only one way to travel between two different cities (this network form a tree). Last year, The ministry of transport decided to orient the roads in one direction because they are too narrow.

Roads are represented by connections where connections[i] = [ai, bi] represents a road from city ai to city bi.

This year, there will be a big event in the capital (city 0), and many people want to travel to this city.

Your task consists of reorienting some roads such that each city can visit the city 0. Return the minimum number of edges changed.

It's guaranteed that each city can reach city 0 after reorder.


Example 1:

Input: n = 6, connections = [[0,1],[1,3],[2,3],[4,0],[4,5]]
Output: 3
Explanation: Change the direction of edges show in red such that each node can reach the node 0 (capital).


Example 2:

Input: n = 5, connections = [[1,0],[1,2],[3,2],[3,4]]
Output: 2
Explanation: Change the direction of edges show in red such that each node can reach the node 0 (capital).


Example 3:

Input: n = 3, connections = [[1,0],[2,0]]
Output: 0


Constraints:

2 <= n <= 5 * 10^4
connections.length == n - 1
connections[i].length == 2
0 <= a_i, b_i <= n - 1
a_i != b_i

*/

/**
 * @param {number} n
 * @param {number[][]} connections
 * @return {number}
 */
var minReorder = function (n, connections) {
    let flips = 0;
    const visited = new Array(n).fill(false);
    const visitCityAndNearby = (city) => {
        visited[city] = true;

        for (const [a, b] of connections) {
            if (visited[a] && visited[b]) continue;

            if (city === a) {
                flips++;
                visitCityAndNearby(b);
            } else if (city === b) {
                visitCityAndNearby(a);
            }
        }
    }

    visitCityAndNearby(0);
    return flips;
};

describe('1466. Reorder Routes to Make All Paths Lead to the City Zero', () => {
    test.each([
        { n: 6, connections: [[0, 1], [1, 3], [2, 3], [4, 0], [4, 5]], expected: 3 },
        { n: 3, connections: [[1, 0], [2, 0]], expected: 0 },
        { n: 3, connections: [[1, 0], [2, 0]], expected: 0 },
        { n: 6, connections: [[4, 5], [0, 1], [1, 3], [2, 3], [4, 0]], expected: 3 },
    ])('$n, $connections', ({ n, connections, expected }) => expect(minReorder(n, connections)).toEqual(expected));
})