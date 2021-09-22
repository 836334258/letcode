/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number}
 */
var searchInsert = function (nums, target) {
  let max = nums.length - 1
  let min = 0

  while (min <= max) {
    let mid = Math.floor((max + min) / 2)
    if (target == nums[mid]) {
      return mid
    } else if (target < nums[mid]) {
      max = mid - 1
    } else {
      min = mid + 1
    }
  }



  nums.splice(nums.length - 1, 0, target)
  console.log(nums.sort())
  return nums.sort((a,b)=>a-b).findIndex((element) => {
    return element === target
  });
};

const nums = [3,5,7,9,10]
const target = 8


console.log(searchInsert(nums, target))