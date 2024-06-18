import { URL, fileURLToPath } from "node:url";
import react from "@vitejs/plugin-react";

import { type CommonServerOptions, defineConfig } from "vite";
import { ui } from "../raw/usr/share/casaos/modules/zimacube-metrics.json";

const devBase = ui.entry.substring(0, ui.entry.lastIndexOf("/"));
const prodBase = "./";

// Enable the proxy for the development server.
const useProxy = false;
// Enable the global proxy for the development server.
const useGlobalProxy = true;

// Custom proxy settings
const proxy = {
	// "^/chat.*": { target: "http://10.0.0.65:8001", changeOrigin: true },
} as CommonServerOptions["proxy"];

// globalProxyTarget:
// The target server for the global proxy. except the custom proxy settings.
const globalProxyTarget = "http://10.0.0.85";

const globalProxy = (() => {
	const result = ({} as CommonServerOptions["proxy"]) || {}; // Initialize result as an empty object
	const excludePaths = [
		devBase,
		...Object.keys(proxy || {}).map((key) =>
			key.replace("^", "").replace(".*", ""),
		),
	].join("|");
	result[`^(?!${excludePaths})/.*`] = {
		target: globalProxyTarget,
		changeOrigin: true,
		ws: true,
	};
	return result;
})();

const isDevMode = process.env.NODE_ENV === "development";
const base = isDevMode ? devBase : prodBase;

// https://vitejs.dev/config/
export default defineConfig({
	base,
	build: {
		outDir: "../raw/usr/share/casaos/www/modules/zimacube-metrics",
	},
	plugins: [react()],
	resolve: {
		alias: {
			"@": fileURLToPath(new URL("./src", import.meta.url)),
		},
	},
	server: {
		proxy: {
			...(useGlobalProxy ? globalProxy : {}),
			...(useProxy ? proxy : {}),
		},
	},
});
