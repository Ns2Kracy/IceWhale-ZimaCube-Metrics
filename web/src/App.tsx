import {
	Card,
	Table,
	TableBody,
	TableCell,
	TableColumn,
	TableHeader,
	TableRow,
} from "@nextui-org/react";
import axios from "axios";

import { useEffect, useState } from "react";

const baseURL = "http://10.0.0.85";
const metricsAPI = `${baseURL}/v2/metrics/`;

interface Metrics {
	name: string;
	cpu: string;
	avg_cpu: string;
	max_cpu: string;
	mem: string;
	avg_mem: string;
	max_mem: string;
}

const DataTable = () => {
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
		const intervalId = setInterval(fetchData, 1000); // 每秒获取一次数据

		return () => clearInterval(intervalId); // 清除定时器以防止内存泄漏
	}, []);

	return (
		<Card>
			<Table aria-labelledby="IceWhale ZimaCube Metrics">
				<TableHeader>
					<TableColumn>服务名称</TableColumn>
					<TableColumn>当前 CPU</TableColumn>
					<TableColumn>平均 CPU</TableColumn>
					<TableColumn>最大 CPU</TableColumn>
					<TableColumn>当前内存</TableColumn>
					<TableColumn>平均内存</TableColumn>
					<TableColumn>最大内存</TableColumn>
				</TableHeader>
				<TableBody>
					{data.map((item: Metrics, index: number) => (
						// biome-ignore lint/suspicious/noArrayIndexKey: <explanation>
						<TableRow key={index}>
							<TableCell>{item.name}</TableCell>
							<TableCell>{item.cpu}</TableCell>
							<TableCell>{item.avg_cpu}</TableCell>
							<TableCell>{item.max_cpu}</TableCell>
							<TableCell>{item.mem}</TableCell>
							<TableCell>{item.avg_mem}</TableCell>
							<TableCell>{item.max_mem}</TableCell>
						</TableRow>
					))}
				</TableBody>
			</Table>
		</Card>
	);
};

function App() {
	return (
		<>
			<DataTable />
		</>
	);
}

export default App;
