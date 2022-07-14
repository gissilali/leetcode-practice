/**
 * @param {string} s
 * @return {boolean}
 */
var isPalindrome = function(s) {
      const reversedString = s
    .trim()
    .replace(/\s/g, "")
    .replace(/[^a-zA-Z0-9 ]/g, "")
    .split("")
    .reverse()
    .join("")
    .toLowerCase();
  
  return (
    reversedString ==
    s
      .trim()
      .replace(/\s/g, "")
      .replace(/[^a-zA-Z0-9 ]/g, "")
      .split("")
      .join("")
      .toLowerCase()
  );
};