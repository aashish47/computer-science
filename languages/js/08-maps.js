const map = new Map();
map.set("apple", "red");
map.set("orange", "orange");
console.log(map);
const arr = Array.from(map);
arr.forEach(([key, value]) => console.log(`${key}: ${value}`));
const obj = Object.fromEntries(arr);
console.log(obj);
