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
