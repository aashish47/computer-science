const greet = (name) => console.log(`hello ${name}`);

const userGreeting = (name, greet) => {
	greet(name);
};

userGreeting("aashish", greet);

const getBread = (succeess, failure) => {
	setTimeout(() => {
		try {
			succeess("got the bread...");
		} catch {
			failure("bread not available!!!");
		}
	}, 2000);
};

const getToppings = (succeess, failure) => {
	setTimeout(() => {
		try {
			succeess("getting toppings...");
		} catch {
			failure("toppings not available!!");
		}
	}, 2000);
};

const arrageItems = (succeess, failure) => {
	setTimeout(() => {
		try {
			succeess("arranging items...");
		} catch {
			failure("arranging items failed!!");
		}
	}, 2000);
};

const cook = (succeess, failure) => {
	setTimeout(() => {
		try {
			succeess("cooking...");
		} catch {
			failure("cooking failed!!");
		}
	}, 2000);
};

const serve = (succeess, failure) => {
	setTimeout(() => {
		try {
			throw new Error();
			succeess("serving pizza...");
		} catch {
			failure("serving pizza failed!!");
		}
	}, 2000);
};

function createPizza() {
	console.log("creating pizza...");
	getBread(
		(bread) => {
			console.log(bread);
			getToppings(
				(toppings) => {
					console.log(toppings);
					arrageItems(
						(items) => {
							console.log(items);
							cook(
								(cook) => {
									console.log(cook);
									serve(
										(result) => console.log(result),
										(serveErr) => console.log(serveErr),
									);
								},
								(cookingErr) => console.log(cookingErr),
							);
						},
						(itemsErr) => console.log(itemsErr),
					);
				},
				(toppingsErr) => console.log(toppingsErr),
			);
		},
		(breadErr) => console.log(breadErr),
	);
}

createPizza();
