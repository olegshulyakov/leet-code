package main

import (
	"testing"
)

/*

In the world of Dota2, there are two parties: the Radiant and the Dire.

The Dota2 senate consists of senators coming from two parties. Now the Senate wants to decide on a change in the Dota2 game. The voting for this change is a round-based procedure. In each round, each senator can exercise one of the two rights:

Ban one senator's right: A senator can make another senator lose all his rights in this and all the following rounds.
Announce the victory: If this senator found the senators who still have rights to vote are all from the same party, he can announce the victory and decide on the change in the game.
Given a string senate representing each senator's party belonging. The character 'R' and 'D' represent the Radiant party and the Dire party. Then if there are n senators, the size of the given string will be n.

The round-based procedure starts from the first senator to the last senator in the given order. This procedure will last until the end of voting. All the senators who have lost their rights will be skipped during the procedure.

Suppose every senator is smart enough and will play the best strategy for his own party. Predict which party will finally announce the victory and change the Dota2 game. The output should be "Radiant" or "Dire".


Example 1:

Input: senate = "RD"
Output: "Radiant"
Explanation:
The first senator comes from Radiant and he can just ban the next senator's right in round 1.
And the second senator can't exercise any rights anymore since his right has been banned.
And in round 2, the first senator can just announce the victory since he is the only guy in the senate who can vote.


Example 2:

Input: senate = "RDD"
Output: "Dire"
Explanation:
The first senator comes from Radiant and he can just ban the next senator's right in round 1.
And the second senator can't exercise any rights anymore since his right has been banned.
And the third senator comes from Dire and he can ban the first senator's right in round 1.
And in round 2, the third senator can just announce the victory since he is the only guy in the senate who can vote.


Constraints:

n == senate.length
1 <= n <= 10^4
senate[i] is either 'R' or 'D'.

*/

// predictPartyVictory determines which party will win in the Dota2 Senate voting process.
// Each senator can ban an opposing senator in a round-based procedure.
// The goal is to determine which party will finally announce the victory.
//
// The algorithm simulates the voting process:
// 1. Each senator takes turns in a queue-like fashion
// 2. When it's a senator's turn, they ban the first opposing senator
// 3. The banning senator goes to the back of the queue for the next round
// 4. This continues until only one party remains
//
// Parameters:
//
//	senate string - String representing senators where 'R' = Radiant, 'D' = Dire
//
// Returns:
//
//	string - Winning party ("Radiant" or "Dire")
func predictPartyVictory(senate string) string {
	// Convert string to slice for easier manipulation
	queue := []rune(senate)
	indexOpposite := 0

	// Continue until only one senator remains or no opposing senator found
	for len(queue) > 1 && indexOpposite > -1 {
		// Get the current senator (first in queue)
		senator := queue[0]
		// Remove the first senator from queue
		queue = queue[1:]

		// Find the first opposing senator
		indexOpposite = -1
		for i, s := range queue {
			if s != senator {
				indexOpposite = i
				break
			}
		}

		// If opposing senator found, remove them from the queue
		if indexOpposite > -1 {
			// Remove the opposing senator
			queue = append(queue[:indexOpposite], queue[indexOpposite+1:]...)
		}

		// Add the current senator back to the end of the queue for next round
		queue = append(queue, senator)
	}

	// Return the winning party
	if queue[0] == 'R' {
		return "Radiant"
	}
	return "Dire"
}

func TestPredictPartyVictory(t *testing.T) {
	testCases := []struct {
		senate string
		want   string
	}{
		{senate: "RD", want: "Radiant"},
		{senate: "RDD", want: "Dire"},
		{senate: "DDRRR", want: "Dire"},
		{senate: "RRR", want: "Radiant"},
	}
	for _, tc := range testCases {
		t.Run("649. Dota2 Senate", func(t *testing.T) {
			out := predictPartyVictory(tc.senate)
			if out != tc.want {
				t.Errorf("predictPartyVictory(%v) = %v, expected %v", tc.senate, out, tc.want)
			}
		})
	}
}
