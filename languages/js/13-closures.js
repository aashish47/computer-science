function adder(x) {
	return (y) => x + y;
}

add5 = adder(5);

console.log(add5(10));
