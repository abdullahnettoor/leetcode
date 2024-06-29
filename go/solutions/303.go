package solutions

/* Question **********************************
? 303. Range Sum Query - Immutable
Given an integer array nums, handle multiple queries of the following type:

Calculate the sum of the elements of nums between indices left and right inclusive where left <= right.
Implement the NumArray class:
NumArray(int[] nums) Initializes the object with the integer array nums.
int sumRange(int left, int right) Returns the sum of the elements of nums between indices left and right inclusive (i.e. nums[left] + nums[left + 1] + ... + nums[right]).


Example 1:
Input
["NumArray", "sumRange", "sumRange", "sumRange"]
[[[-2, 0, 3, -5, 2, -1]], [0, 2], [2, 5], [0, 5]]
Output
[null, 1, -1, -3]

Explanation
NumArray numArray = new NumArray([-2, 0, 3, -5, 2, -1]);
numArray.sumRange(0, 2); // return (-2) + 0 + 3 = 1
numArray.sumRange(2, 5); // return 3 + (-5) + 2 + (-1) = -1
numArray.sumRange(0, 5); // return (-2) + 0 + 3 + (-5) + 2 + (-1) = -3


Constraints:
1 <= nums.length <= 104
-105 <= nums[i] <= 105
0 <= left <= right < nums.length
At most 104 calls will be made to sumRange.
*********************************************/

func Output303() any {
	nums := []int{1, 2, 3, 5, -6, 0, 4}
	obj := Constructor(nums)
	param_1 := obj.SumRange(0, 4)
	return param_1
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */

// * Solution -- Brute Force -- Time O(r-l) -- Space O(1)
type NumArray struct {
	arr      []int
	rangeSum int
}

func Constructor(nums []int) NumArray {
	return NumArray{arr: nums}
}

func (this *NumArray) SumRange(left int, right int) int {
	this.rangeSum = 0
	for i := left; i <= right; i++ {
		this.rangeSum += this.arr[i]
	}
	return this.rangeSum
}
