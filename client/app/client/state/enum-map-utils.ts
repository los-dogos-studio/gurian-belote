import { plainToInstance, type TransformFnParams } from "class-transformer";

export const enumKeyMap = <T, U>(value: TransformFnParams) => {
	const map = new Map<T, U>();
	value.value.forEach((v: U, k: string) => {
		const intKey = parseInt(k, 10);
		if (!isNaN(intKey)) {
			map.set(intKey as unknown as T, v);
		} else {
			console.warn(`Invalid key: ${k}`);
		}
	});
	return map;
};

export const enumKeyMapToClassValue = <T, U>(
	type: new (...args: any[]) => U,
) => {
	return (value: TransformFnParams) => {
		const map = new Map<T, U>();
		value.value.forEach((v: U, k: string) => {
			const intKey = parseInt(k, 10);
			if (!isNaN(intKey)) {
				map.set(intKey as unknown as T, plainToInstance(type, v));
			} else {
				console.warn(`Invalid key: ${k}`);
			}
		});
		return map;
	};
};
