import type { TransformFnParams } from "class-transformer";

export const stringMapToIntEnumMap = <T, U>(value: TransformFnParams) => {
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
}
