function add(x, y) {
	return `${this.time}: ${x + y}`;
}

const addBind = add.bind({ time: new Date().toLocaleTimeString() }, 5, 7);
console.log(addBind());

const addApply = add.apply({ time: new Date().toLocaleTimeString() }, [10, 23]);
console.log(addApply);

console.log(add.call({ time: new Date().toLocaleTimeString() }, 23, 11));
