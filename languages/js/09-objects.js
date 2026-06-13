person1 = {
	name: "aashish",
	age: 29,
};

console.log(person1);

function Person(name, age) {
	this.name = name;
	this.age = age;
}

person2 = new Person("agarwal", 29);
console.log(person2);

const { name, age } = person2;
console.log(name, age);
