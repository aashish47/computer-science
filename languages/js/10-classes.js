class Person {
	#name;
	#age;
	constructor(name, age) {
		this.#name = name;
		this.#age = age;
	}
	getName() {
		return this.#name;
	}
	setName(name) {
		this.#name = name;
	}
	getAge() {
		return this.#age;
	}
	setAge(age) {
		this.#age = age;
	}
	toString() {
		return `name: ${this.#name} age: ${this.#age}`;
	}
}

person1 = new Person("aashish", 29);
console.log(`person1 => ${person1}`);
