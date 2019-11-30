/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/

twoSum(nums []int, target int) []int {
	m:= make(map[int][]int)
	//Adding the elements of the array to a map (hash table)
	for i:=0; i<len(nums);i++{
		_,duplicate:=m[nums[i]]
		if(duplicate){
			m[nums[i]]=append(m[nums[i]],i)
		}else{
			m[nums[i]]=[]int{i}
		}
	}
	var indices []int
	var index []int
	var exists bool

	//Looping over nums slice to find calculate complement and look it up in hash table
	for j:=0;j<len(nums);j++{
		complement:=target-nums[j]
		index,exists=m[complement]
		if(exists){
			for k:=0;k<len(index);k++{
				if(index[k]!=j){
					indices=append(indices,index[k])
				}
			}
		}
	}
	return indices
}
