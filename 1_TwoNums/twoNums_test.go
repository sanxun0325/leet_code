package __TwoNums

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func twoSum(nums []int,target int)[]int{
	numsMap:=make(map[int]int,len(nums))

	for i,v:=range nums{
		 numsMap[v]=i
		 num:=target-v
		 if index,ok:=numsMap[num];ok{
		 	return []int{i,index}
		 }
	}
	return nil
}

func TestTwoSum(t *testing.T){
	nums:=[]int{2,7,11,15}
	result:=twoSum(nums,9)
	assert.NotNil(t,result)
	assert.True(t,result[0]==1)
	assert.True(t,result[1]==0)
}