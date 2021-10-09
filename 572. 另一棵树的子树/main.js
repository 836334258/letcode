/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {TreeNode} root
 * @param {TreeNode} subRoot
 * @return {boolean}
 */
 var isSubtree = function(root, subRoot) {
  if(!root){
    return false
  }

  return check(root,subRoot) || isSubtree(root.left,subRoot) || isSubtree(root.right,subRoot)
};


function check(a,b){
  if(!a && !b){
    return true
  }

  if(!a || !b){
    return  false
  }

  if(a.val===b.val){
    return check(a.left,b.left) && check(a.right,b.right)
  }

  return  false

}