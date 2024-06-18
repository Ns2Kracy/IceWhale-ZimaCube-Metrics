import {
	Accordion,
	AccordionItem,
	Table,
	TableBody,
	TableCell,
	TableColumn,
	TableHeader,
	TableRow,
} from "@nextui-org/react";
import type { components } from "../api/openapi";

export default function ZimaCube(props: {
	metrics: components["responses"]["ResponseZimaCubeMetricsOK"]["data"][];
}) {
	return (
		<Accordion>
			<AccordionItem title="ZimaCube:">
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
						{props.metrics.map(
							(item: components["responses"]["ResponseZimaCubeMetricsOK"]) => (
								<TableRow key={item.data.name}>
									<TableCell>{item.data.name}</TableCell>
									<TableCell>{item.data.currentCpu}</TableCell>
									<TableCell>{item.data.avgCpu}</TableCell>
									<TableCell>{item.data.maxCpu}</TableCell>
									<TableCell>{item.data.currentMemory}</TableCell>
									<TableCell>{item.data.avgMemory}</TableCell>
									<TableCell>{item.data.maxMemory}</TableCell>
								</TableRow>
							),
						)}
					</TableBody>
				</Table>
			</AccordionItem>
		</Accordion>
	);
}
