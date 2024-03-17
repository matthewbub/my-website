import { useState } from "react";
import viteLogo from "/vite.svg";
import biomeLogo from "./assets/biome.svg";
import reactLogo from "./assets/react.svg";
import tailwindLogo from "./assets/tailwind.svg";

function App() {
	const [count, setCount] = useState(0);

	return (
		<div className="w-full">
			<div className="max-w-2xl mx-auto mt-40">
				<div className="flex flex-wrap justify-center space-x-4 py-2">
					<a href="https://vitejs.dev" target="_blank" rel="noreferrer">
						<img src={viteLogo} className="logo" alt="Vite logo" />
					</a>
					<a href="https://react.dev" target="_blank" rel="noreferrer">
						<img src={reactLogo} className="logo react" alt="React logo" />
					</a>
					<a href="https://tailwindcss.com" target="_blank" rel="noreferrer">
						<img src={tailwindLogo} className="h-[32px] logo" alt="Tailwind CSS logo" />
					</a>
					<a href="https://biome.dev" target="_blank" rel="noreferrer">
						<img src={biomeLogo} className="h-[32px]" alt="Tailwind CSS logo" />
					</a>
				</div>
				<h1 className="text-4xl font-bold text-center text-stone-100 mb-4">Vite + React + TailwindCSS + Biome</h1>
				<div className="text-center">
					<button className="border border-blue-600 rounded bg-blue-600 text-white px-4 py-2" type="button" onClick={() => setCount((count) => count + 1)}>
						count is {count}
					</button>
					<p className="text-xl mt-6 text-stone-400">
						Edit <code>src/App.tsx</code> and save to test HMR
					</p>
				</div>
				<div className="bg-yellow-500 rounded-lg py-2 mt-8">
					<p className="text-xl text-black font-medium text-center">
						Click on the logos to learn more
					</p>
				</div>
			</div>
		</div>
	);
}

export default App;
