import type { Route } from "./+types/home";
import { Welcome } from "../welcome/welcome";

export function meta({}: Route.MetaArgs) {
  return [
    { title: "Gurian Belote" },
    { name: "Gurian Belote", content: "Welcome to Gurian Belote!" },
  ];
}

export default function Home() {
  return <Welcome />;
}
