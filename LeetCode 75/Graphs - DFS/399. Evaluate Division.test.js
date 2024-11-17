/*

You are given an array of variable pairs equations and an array of real numbers values, where equations[i] = [Ai, Bi] and values[i] represent the equation Ai / Bi = values[i]. Each Ai or Bi is a string that represents a single variable.

You are also given some queries, where queries[j] = [Cj, Dj] represents the jth query where you must find the answer for Cj / Dj = ?.

Return the answers to all queries. If a single answer cannot be determined, return -1.0.

Note: The input is always valid. You may assume that evaluating the queries will not result in division by zero and that there is no contradiction.

Note: The variables that do not occur in the list of equations are undefined, so the answer cannot be determined for them.


Example 1:

Input: equations = [["a","b"],["b","c"]], values = [2.0,3.0], queries = [["a","c"],["b","a"],["a","e"],["a","a"],["x","x"]]
Output: [6.00000,0.50000,-1.00000,1.00000,-1.00000]
Explanation:
Given: a / b = 2.0, b / c = 3.0
queries are: a / c = ?, b / a = ?, a / e = ?, a / a = ?, x / x = ?
return: [6.0, 0.5, -1.0, 1.0, -1.0 ]
note: x is undefined => -1.0

Example 2:

Input: equations = [["a","b"],["b","c"],["bc","cd"]], values = [1.5,2.5,5.0], queries = [["a","c"],["c","b"],["bc","cd"],["cd","bc"]]
Output: [3.75000,0.40000,5.00000,0.20000]


Example 3:

Input: equations = [["a","b"]], values = [0.5], queries = [["a","b"],["b","a"],["a","c"],["x","y"]]
Output: [0.50000,2.00000,-1.00000,-1.00000]


Constraints:

1 <= equations.length <= 20
equations[i].length == 2
1 <= A_i.length, B_i.length <= 5
values.length == equations.length
0.0 < values[i] <= 20.0
1 <= queries.length <= 20
queries[i].length == 2
1 <= C_j.length, D_j.length <= 5
A_i, B_i, C_j, D_j consist of lower case English letters and digits.

*/

/**
 * @param {string[][]} equations
 * @param {number[]} values
 * @param {string[][]} queries
 * @return {number[]}
 */
var calcEquation = function (equations, values, queries) {
    if (!equations || !values || !queries) return [];

    const evals = new Map();
    for (let index = 0; index < equations.length; index++) {
        const addToMap = (dividend, divisor, value) => {
            if (!evals.has(dividend)) evals.set(dividend, new Map());
            evals.get(dividend).set(divisor, value);
        }

        const [i, j] = equations[index];
        if (i === j) continue;

        const value = values[index];
        addToMap(i, j, value);
        addToMap(j, i, 1 / value);
    }

    const dfs = (dividend, divisor) => {
        const memorized = evals.get(dividend);
        if (memorized.has(divisor)) return memorized.get(divisor);

        for (const [key, value] of memorized) {
            const res = dfs(key, divisor);
            if (res !== -1) return value * res;
        }

        return -1;
    }

    const answers = new Array(queries.length);
    for (let index = 0; index < queries.length; index++) {
        const [i, j] = queries[index];

        if (!evals.has(i) || !evals.has(j)) answers[index] = -1;
        else if (i === j) answers[index] = 1;
        else if (evals.get(i).has(j)) answers[index] = evals.get(i).get(j);
        else answers[index] = dfs(i, j);
    }

    return answers;
};

describe('399. Evaluate Division', () => {
    test.each([
        { equations: [], values: [], queries: [], expected: [] },
        { equations: [["a", "b"], ["b", "c"]], values: [2.0, 3.0], queries: [["a", "c"], ["b", "a"], ["a", "e"], ["a", "a"], ["x", "x"]], expected: [6.00000, 0.50000, -1.00000, 1.00000, -1.00000] },
        { equations: [["a", "b"], ["b", "c"], ["bc", "cd"]], values: [1.5, 2.5, 5.0], queries: [["a", "c"], ["c", "b"], ["bc", "cd"], ["cd", "bc"]], expected: [3.75000, 0.40000, 5.00000, 0.20000] },
        { equations: [["a", "b"]], values: [0.5], queries: [["a", "b"], ["b", "a"], ["a", "c"], ["x", "y"]], expected: [0.50000, 2.00000, -1.00000, -1.00000] },
    ])('$equations, $values, $queries', ({ equations, values, queries, expected }) => expect(calcEquation(equations, values, queries)).toEqual(expected));
})