import {
	QueryClient,
	QueryClientProvider,
	useQuery,
} from "@tanstack/react-query";
import type { ParamsOption, RequestBodyOption } from "openapi-fetch";
import { useEffect, useState } from "react";

import client from "./api";
import type { paths } from "./api/openapi";
import ZimaCube from "./components/ZimaCube";

type UseQueryOptions<T> = ParamsOption<T> &
	RequestBodyOption<T> & {
		// add your custom options here
		reactQuery?: {
			enabled: boolean; // Note: React Query type’s inference is difficult to apply automatically, hence manual option passing here
			// add other React Query options as needed
		};
	};

function getMetrics({ params, body }: UseQueryOptions<paths["/"]["get"]>) {
	return useQuery({
		queryKey: ["getMetrics", params, body],
		queryFn: async ({ signal }) => {
			const { data } = await client.GET("/", { params, body, signal });
			return data;
		},
	});
}

function Metrics() {
	const { data, error, isLoading } = getMetrics({
		reactQuery: {
			enabled: true,
		},
	});

	if (isLoading) return <div>Loading...</div>;
	if (error) return <div>Error: {error.message}</div>;

	return <ZimaCube metrics={data} />;
}

function App() {
	const [reactQueryClient] = useState(
		new QueryClient({
			defaultOptions: {
				queries: {
					networkMode: "offlineFirst", // keep caches as long as possible
					refetchOnWindowFocus: false, // don’t refetch on window focus
				},
			},
		}),
	);

	return (
		<QueryClientProvider client={reactQueryClient}>
			<Metrics />
		</QueryClientProvider>
	);
}

export default App;
