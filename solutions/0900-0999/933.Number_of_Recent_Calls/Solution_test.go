package main

import (
	"testing"
)

/*

You have a RecentCounter class which counts the number of recent requests within a certain time frame.

Implement the RecentCounter class:

RecentCounter() Initializes the counter with zero recent requests.
int ping(int t) Adds a new request at time t, where t represents some time in milliseconds, and returns the number of requests that has happened in the past 3000 milliseconds (including the new request). Specifically, return the number of requests that have happened in the inclusive range [t - 3000, t].
It is guaranteed that every call to ping uses a strictly larger value of t than the previous call.


Example 1:

Input
["RecentCounter", "ping", "ping", "ping", "ping"]
[[], [1], [100], [3001], [3002]]
Output
[null, 1, 2, 3, 3]

Explanation
RecentCounter recentCounter = new RecentCounter();
recentCounter.ping(1);     // requests = [1], range is [-2999,1], return 1
recentCounter.ping(100);   // requests = [1, 100], range is [-2900,100], return 2
recentCounter.ping(3001);  // requests = [1, 100, 3001], range is [1,3001], return 3
recentCounter.ping(3002);  // requests = [1, 100, 3001, 3002], range is [2,3002], return 3


Constraints:

1 <= t <= 10^9
Each test case will call ping with strictly increasing values of t.
At most 10^4 calls will be made to ping.

*/

// RecentCounter maintains a queue of recent ping times and counts
// pings within the last 3000 milliseconds.
type RecentCounter struct {
	queue []int // Queue to store ping times
}

// NewRecentCounter creates and initializes a new RecentCounter instance.
// This serves as the constructor for the RecentCounter struct.
//
// Returns:
//
//	*RecentCounter - Pointer to the newly created RecentCounter instance
func Constructor() RecentCounter {
	return RecentCounter{
		queue: make([]int, 0),
	}
}

// Ping adds a new ping time and returns the number of pings
// that occurred in the past 3000 milliseconds (inclusive).
// It maintains the queue by removing outdated pings before adding the new one.
//
// Parameters:
//
//	t int - Current time in milliseconds
//
// Returns:
//
//	int - Number of pings in the time range [t-3000, t]
func (rc *RecentCounter) Ping(t int) int {
	// Remove all ping times that are older than (t - 3000)
	// This maintains a sliding window of relevant pings
	for len(rc.queue) > 0 && rc.queue[0] < (t-3000) {
		rc.queue = rc.queue[1:] // Remove first element
	}

	// Add the current ping time to the queue
	rc.queue = append(rc.queue, t)

	// Return the count of recent pings
	return len(rc.queue)
}

func TestRecentCounter(t *testing.T) {
	testCases := []struct {
		t    []int
		want []int
	}{
		{t: []int{1, 100, 3001, 3002}, want: []int{1, 2, 3, 3}},
	}
	for _, tc := range testCases {
		t.Run("933. Number of Recent Calls", func(t *testing.T) {
			recentCounter := Constructor()
			for i, timer := range tc.t {
				out := recentCounter.Ping(timer)
				if out != tc.want[i] {
					t.Errorf("Ping(%d) = %v, want: %v", timer, out, tc.want)
				}
			}
		})
	}
}
