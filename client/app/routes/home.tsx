import type { Route } from "./+types/home";
import Game from "~/game/Game";

export function meta({}: Route.MetaArgs) {
  return [
    { title: "Gurian Belote" },
    { name: "Gurian Belote", content: "Welcome to Gurian Belote!" },
  ];
}

export default function Home() {
  return <Game />;
}
