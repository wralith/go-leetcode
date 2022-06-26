package leetcode

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {

	lowest := 0
	highest := n

	for highest > lowest {
		half := lowest + (highest-lowest)/2
		if isBadVersion(half) {
			highest = half
		} else {
			lowest = half + 1
		}
	}
	return lowest
}

func isBadVersion(n int) bool {
	if n >= 4 {
		return true
	} else {
		return false
	}
}
