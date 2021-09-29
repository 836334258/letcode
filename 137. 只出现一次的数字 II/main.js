/**
 * @param {number[]} nums
 * @return {number}
 */
 var singleNumber = function(nums) {
  let res=0
  for(let i=0;i<32;i++){
    let sum=0
    for(let num of nums){
      sum+=(num>>i) & 1
    }
    res ^=(sum%3)<<i
  }
};

let nums = [2,2,3,2]
addBinary(nums,3)