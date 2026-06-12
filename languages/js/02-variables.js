console.log(a); // var is hosited and initalized as undefined
var a = 5;
// console.log(b); // let is hosited but uninitialized and throws error
let b = 10;

// let b = 20; cannot redeclare let
const c = "apples";

console.log(a, b, c);
console.log(typeof a, typeof b, typeof c);
