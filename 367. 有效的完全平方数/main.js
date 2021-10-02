/**
 * @param {number} num
 * @return {boolean}
 */
 var isPerfectSquare = function(num) {
  if(num<2){
    return true
  }

  let left=2
  let right=num/2
  
  while(left<=right){
    let mid=Math.floor((left+right)/2)
    if(mid**2===num){
      return true
    }else if(mid**2>num){
      right=mid-1
    }else{
      left=mid+1
    }
  }

  return  false
};


console.log(isPerfectSquare(16))