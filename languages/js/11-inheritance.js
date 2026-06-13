const a = {
	"a": 1,
	"b": 2,
	__proto__: {
		"c": 5,
	},
};

console.log(a.__proto__, Object.getPrototypeOf(a));
console.log(Object.prototype);

const b = Object.create(a);
console.log(
	b.__proto__,
	b.__proto__.__proto__,
	b.__proto__.__proto__.__proto__,
	b.__proto__.__proto__.__proto__.__proto__,
);

/*---------------------------------------*/

function Box(val) {
	this.val = val;
}

Box.prototype.getValue = function () {
	return this.val;
};

const box = new Box(4);
console.log(box.getValue());

/*---------------------------------------*/

function Player(name) {
	this.name = name;
}

Player.prototype.shoot = function () {
	console.log(this.name, "Bang!!!");
};

player1 = new Player("aashish");
player1.shoot();
console.log(Object.getPrototypeOf(player1));

Object.setPrototypeOf(player1, { b: 2 });
console.log(player1.__proto__);

player2 = new Player("agarwal");
console.log(player2.__proto__);
console.log(player2.shoot());

/*---------------------------------------*/

const User = {
	greet() {
		console.log(`Hiiii ${this.name}`);
	},
};

const user1 = Object.create(User);
user1.name = "aashish";
user1.greet();
