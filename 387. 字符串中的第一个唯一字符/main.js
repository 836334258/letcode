/**
 * @param {string} s
 * @return {number}
 */
 var firstUniqChar = function(s) {
  const m =new Map()

  for(const [k,v] of Array.from(s).entries()){
    let tmp=m.get(v)|| 0
    m.set(v,++tmp)
  }

  for(const [k,v] of m){
    if(v==1) return s.indexOf(k)
  }

  return -1
};


let s = "leetcode"

console.log(firstUniqChar(s))