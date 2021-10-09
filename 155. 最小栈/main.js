var MinStack = function() {
  this.arr=[]
  this.minArr=[]
};

/** 
 * @param {number} val
 * @return {void}
 */
MinStack.prototype.push = function(val) {
  this.arr.push(val)
  this.minArr.push(Math.min(val,this.getMin()??Infinity))
};

/**
 * @return {void}
 */
MinStack.prototype.pop = function() {
  this.arr.pop()
  this.minArr.pop()
};

/**
 * @return {number}
 */
MinStack.prototype.top = function() {
  return this.arr[this.arr.length-1]
};

/**
 * @return {number}
 */
MinStack.prototype.getMin = function() {
  return this.minArr[this.minArr.length-1]
};


 var obj = new MinStack()
 obj.push(val)
 obj.pop()
 var param_3 = obj.top()
 var param_4 = obj.getMin()