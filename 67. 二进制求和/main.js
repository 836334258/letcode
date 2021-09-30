/**
 * @param {string} a
 * @param {string} b
 * @return {string}
 */
 var addBinary = function(a, b) {
  a='0b'+a
  b='0b'+b
  return (BigInt(a)+BigInt(b)).toString(2)
};


let a = "11", b = "1"


console.log(addBinary(a,b))