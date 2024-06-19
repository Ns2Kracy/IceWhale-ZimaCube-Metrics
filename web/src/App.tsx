import {
	QueryClient,
	QueryClientProvider,
	useQuery,
} from "@tanstack/react-query";
import type { ParamsOption, RequestBodyOption } from "openapi-fetch";

import client from "./api";
import type { paths } from "./api/openapi";
import ZimaCube from "./components/ZimaCube";

type UseQueryOptions<T> = ParamsOption<T> &
	RequestBodyOption<T> & {
		reactQuery?: {
			enabled: boolean;
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
	const { data } = getMetrics({
		reactQuery: {
			enabled: true,
		},
	});

	const metric = data?.data || [];

	return <ZimaCube metrics={metric} />;
}

function App() {
	const reactQueryClient = new QueryClient({
		defaultOptions: {
			queries: {
				refetchInterval: 1000,
				refetchIntervalInBackground: false,
			},
		},
	});

	return (
		<QueryClientProvider client={reactQueryClient}>
			<Metrics />
		</QueryClientProvider>
	);
}

export default App;
