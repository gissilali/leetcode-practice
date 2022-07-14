/**
 * @param {number[]} nums
 * @return {boolean}
 */
var containsDuplicate = function(nums) {
    const duplicateChecker = [];
    for(let i = 0; i < nums.length; i++) {
        if(duplicateChecker.includes(nums[i])) {
            return true;
        } else {
            duplicateChecker.push(nums[i]);
        }
    }
    
    return false;
};