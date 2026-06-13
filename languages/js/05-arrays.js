const nums = [1, 2, 3, 4, 5];

nums.forEach((num) => console.log(num));
console.log(nums.map((num) => num * 10));
console.log(nums.filter((num) => num % 2 === 0));
console.log(nums.reduce((accum, num) => accum + num, 0));

console.log(nums.splice(1, 3, -1, -2, -3));
console.log(nums);

console.log(nums.slice(1, 4));

console.log([0, 0, ...nums]);
