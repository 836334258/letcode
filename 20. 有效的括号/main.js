/**
 * @param {string} s
 * @return {boolean}
 */
var isValid = function (s) {
  let len = s.length
  if (len % 2 !== 0) {
    return false
  }

  const pair = new Map([
    [')', '('],
    ['}', '{'],
    [']', '['],
  ])

  const stak = []

  for (let i = 0; i < len; i++) {
    if (pair.has(s[i]) && stak.length) {
        if (stak[stak.length - 1] !== pair.get(s[i])) {
          return false
        }
        stak.pop()
     
    }else {
      stak.push(s[i])
    }
    console.log(stak)
  }
  return !stak.length
}

let s="))"
console.log(isValid(s))
