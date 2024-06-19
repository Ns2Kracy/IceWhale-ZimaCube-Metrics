import {
	Table,
	TableBody,
	TableCell,
	TableColumn,
	TableHeader,
	TableRow,
} from "@nextui-org/react";
import type { components } from "../api/openapi";

export default function ZimaCube(props: {
	metrics: components["schemas"]["Metric"][];
}) {
	return (
		<Table aria-labelledby="IceWhale ZimaCube Metrics">
			<TableHeader>
				<TableColumn>服务名称</TableColumn>
				<TableColumn>当前 CPU</TableColumn>
				<TableColumn>平均 CPU</TableColumn>
				<TableColumn>最大 CPU</TableColumn>
				<TableColumn>当前内存</TableColumn>
				<TableColumn>平均内存</TableColumn>
				<TableColumn>最大内存</TableColumn>
				<TableColumn>运行时间</TableColumn>
			</TableHeader>
			<TableBody>
				{props.metrics.map((item) => (
					<TableRow key={item.name}>
						<TableCell>{item.name}</TableCell>
						<TableCell>{item.cpu}</TableCell>
						<TableCell>{item.avg_cpu}</TableCell>
						<TableCell>{item.max_cpu}</TableCell>
						<TableCell>{item.mem}</TableCell>
						<TableCell>{item.avg_mem}</TableCell>
						<TableCell>{item.max_mem}</TableCell>
						<TableCell>{item.uptime}</TableCell>
					</TableRow>
				))}
			</TableBody>
		</Table>
	);
}
