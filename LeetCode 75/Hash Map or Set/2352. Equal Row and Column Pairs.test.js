/*

Given a 0-indexed n x n integer matrix grid, return the number of pairs (ri, cj) such that row ri and column cj are equal.

A row and column pair is considered equal if they contain the same elements in the same order (i.e., an equal array).



Example 1:


Input: grid = [[3,2,1],[1,7,6],[2,7,7]]
Output: 1
Explanation: There is 1 equal row and column pair:
- (Row 2, Column 1): [2,7,7]
Example 2:


Input: grid = [[3,1,2,2],[1,4,4,5],[2,4,2,2],[2,4,2,2]]
Output: 3
Explanation: There are 3 equal row and column pairs:
- (Row 0, Column 0): [3,1,2,2]
- (Row 2, Column 2): [2,4,2,2]
- (Row 3, Column 2): [2,4,2,2]


Constraints:

n == grid.length == grid[i].length
1 <= n <= 200
1 <= grid[i][j] <= 10^5

*/

/**
 * @param {number[][]} grid
 * @return {number}
 */
var equalPairs = function (grid) {
    let pairCnt = 0;
    const rows = new Set();
    const cols = new Set();
    for (let i = 0; i < grid.length; i++) {
        const row = [];
        const column = [];
        for (let j = 0; j < grid.length; j++) {
            row.push(grid[j][i]);
            column.push(grid[i][j]);
        }
        rows.add(row);
        cols.add(column);
    }

    for (const row of rows) {
        for (const col of cols) {
            if (row.some((e, i) => e !== col[i])) continue;
            pairCnt++;
        }
    }

    return pairCnt;
};

describe('2352. Equal Row and Column Pairs', () => {
    test.each([
        { grid: [[3, 2, 1], [1, 7, 6], [2, 7, 7]], expected: 1 },
        { grid: [[3, 1, 2, 2], [1, 4, 4, 5], [2, 4, 2, 2], [2, 4, 2, 2]], expected: 3 },
    ])('$grid', ({ grid, expected }) => expect(equalPairs(grid)).toEqual(expected));
})