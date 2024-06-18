import axios from "axios";
import { useEffect, useState } from "react";

import ZimaCube from "./components/Cube";

const baseURL = "http://10.0.0.85";
const metricsAPI = `${baseURL}/v2/metrics/`;

function App() {
	const [data, setData] = useState([]);

	useEffect(() => {
		const fetchData = async () => {
			try {
				const response = await axios.get(metricsAPI);
				const data = response.data;
				setData(data.data);
			} catch (error) {
				console.error("Error fetching data:", error);
			}
		};

		fetchData();
		const intervalId = setInterval(fetchData, 1000);

		return () => clearInterval(intervalId);
	}, []);

	return (
		<>
			<ZimaCube metrics={data} />
		</>
	);
}

export default App;
