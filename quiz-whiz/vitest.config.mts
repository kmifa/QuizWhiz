import react from "@vitejs/plugin-react";
import tsconfigPaths from "vite-tsconfig-paths";
import { defineConfig } from "vitest/config";

export default defineConfig({
	plugins: [react(), tsconfigPaths()],
	test: {
		globals: true, // This is needed by @testing-library to be cleaned up after each test
		include: ["src/**/*.spec.{js,jsx,ts,tsx}"],
		coverage: {
			include: ["src/**/*"],
			exclude: ["src/**/*.stories.{js,jsx,ts,tsx}", "**/*.d.ts"],
			reporter: ["html"],
		},
		environmentMatchGlobs: [["**/*.spec.tsx", "jsdom"]],
		setupFiles: ["./vitest-setup.ts"],
	},
});
