import BeloteApp from "~/components/game/BeloteApp";
import type { Route } from "./+types/home";

export function meta({ }: Route.MetaArgs) {
	return [
		{ title: "Gurian Belote" },
		{ name: "Gurian Belote", content: "Welcome to Gurian Belote!" },
	];
}

export default function Home() {
	return <BeloteApp />;
}
