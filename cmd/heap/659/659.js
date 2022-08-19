/**
 * @param {number[]} nums
 * @return {boolean}
 */
var isPossible = function (nums) {
  const counts = {};

  for (const num of nums) {
    counts[num] = counts[num] ? counts[num] + 1 : 1;
  }
  console.log(counts)
  for (let num of nums) {
    let freq = counts[num]
    if (freq === 0) {
      continue
    }

    let count = 0
    let currentFreq = counts[num]
    while (currentFreq !== undefined && counts[num] >= freq) {
      if (currentFreq > freq) {
        freq = currentFreq
      }
      counts[num]--
      count++
      num++
      currentFreq = counts[num]
    }
    if (count < 3) {
      return false
    }
  }

  return true
};