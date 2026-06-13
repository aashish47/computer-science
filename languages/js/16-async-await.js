import {
	arrangeItems,
	cook,
	getBread,
	getToppings,
	serve,
} from "./15-promise.js";

const createPizza = async () => {
	try {
		const breadResult = await getBread();
		const toppingsResult = await getToppings(breadResult);
		const itemsResult = await arrangeItems(toppingsResult);
		const cookResult = await cook(itemsResult);
		const finalResult = await serve(cookResult);
	} catch (err) {
		console.log(err);
	}
};

createPizza();
