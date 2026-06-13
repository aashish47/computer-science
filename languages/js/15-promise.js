const getBread = () => {
	return new Promise((resolve, reject) => {
		setTimeout(() => {
			resolve("Bread 🍞");
		}, 2000);
	});
};

const getToppings = (bread) => {
	return new Promise((resolve, reject) => {
		setTimeout(() => {
			resolve(`${bread} + Toppings 🧀`);
		}, 2000);
	});
};

const arrangeItems = (pizzaWithToppings) => {
	return new Promise((resolve, reject) => {
		setTimeout(() => {
			resolve(`${pizzaWithToppings} + Arranged Layers 🍕`);
		}, 2000);
	});
};

const cook = (arrangedPizza) => {
	return new Promise((resolve, reject) => {
		setTimeout(() => {
			resolve(`${arrangedPizza} -> Oven Baked 🔥`);
		}, 2000);
	});
};

const serve = (finalPizza) => {
	return new Promise((resolve, reject) => {
		setTimeout(() => {
			resolve(`${finalPizza} -> Served! Enjoy your Pizza! 🎉`);
		}, 2000);
	});
};

const createPizza = () => {
	getBread()
		.then((breadResult) => {
			console.log(breadResult);
			return getToppings(breadResult);
		})
		.then((toppingsResult) => {
			console.log(toppingsResult);
			return arrangeItems(toppingsResult);
		})
		.then((arrangedResult) => {
			console.log(arrangedResult);
			return cook(arrangedResult);
		})
		.then((cookedResult) => {
			console.log(cookedResult);
			return serve(cookedResult);
		})
		.then((finalResult) => {
			console.log(finalResult);
		})
		.catch((err) => console.log(err));
};
createPizza();

export { arrangeItems, cook, getBread, getToppings, serve };
