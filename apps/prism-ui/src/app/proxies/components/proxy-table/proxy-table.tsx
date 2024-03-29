import { Typography } from "@material-tailwind/react";
import {
    createColumnHelper,
    flexRender,
    getCoreRowModel,
    useReactTable,
    type TableOptions,
} from "@tanstack/react-table";
import type { Proxy } from "../../models/proxy";
import styles from "./proxy-table.module.scss";

export interface TableProps {
    tableData: Proxy[];
}

export function ProxyTable({ tableData }: TableProps) {
    const columnHelper = createColumnHelper<Proxy>();
    const columns = [
        columnHelper.accessor("name", {
            header: "Name",
            cell: props => <p>{props.getValue()}</p>,
        }),
        columnHelper.accessor("method", {
            header: "Method",
            cell: props => <p>{props.getValue()}</p>,
            size: 10,
        }),
        columnHelper.accessor("path", {
            header: "Path",
            cell: props => <p>{props.getValue()}</p>,
        }),
        columnHelper.accessor("source", {
            header: "Source",
            cell: props => <p>{props.getValue()}</p>,
        }),
    ];
    const table = useReactTable({
        columns,
        data: tableData,
        getCoreRowModel: getCoreRowModel(),
    } as TableOptions<Proxy>);
    return (
        <div className={styles["container"]}>
            <table className="w-full min-w-[640px] table-auto">
                <thead>
                    {table.getHeaderGroups().map(headerGroup => (
                        <tr key={headerGroup.id}>
                            {headerGroup.headers.map(header => (
                                <th
                                    key={header.id}
                                    className="border-blue-gray-50 border-b py-3 px-5 text-left"
                                >
                                    <Typography
                                        variant="small"
                                        className="text-blue-gray-400 text-[11px] font-bold uppercase"
                                    >
                                        {header.column.columnDef.header?.toString()}
                                    </Typography>
                                </th>
                            ))}
                        </tr>
                    ))}
                </thead>
                <tbody>
                    {table.getRowModel().rows.map((row, key) => {
                        const className = `py-3 px-5 ${
                            key === tableData.length - 1
                                ? ""
                                : "border-b border-blue-gray-50"
                        }`;
                        return (
                            <tr>
                                {row.getVisibleCells().map(cell => (
                                    <td key={cell.id} className={className}>
                                        {flexRender(
                                            cell.column.columnDef.cell,
                                            cell.getContext(),
                                        )}
                                    </td>
                                ))}
                            </tr>
                        );
                    })}
                </tbody>
            </table>
        </div>
    );
}

export default ProxyTable;
