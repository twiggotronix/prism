import { faAdd, faPencil, faTrash } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import {
    Button,
    Dialog,
    DialogBody,
    DialogHeader,
    Typography,
} from "@material-tailwind/react";
import {
    createColumnHelper,
    flexRender,
    getCoreRowModel,
    useReactTable,
    type Row,
    type TableOptions,
} from "@tanstack/react-table";
import { useState } from "react";
import { useProxyDeleteMutation } from "../../hooks/use-proxy-delete-mutation.hook";
import { Method, Proxy } from "../../models/proxy";
import ProxyForm from "../proxy-form/proxy-form";
import styles from "./proxy-table.module.scss";

export interface TableProps {
    tableData: Proxy[];
}

export function ProxyTable({ tableData }: TableProps) {
    const [openDialog, setOpenDialog] = useState<boolean>(false);
    const [currentRow, setCurrentRow] = useState<Row<Proxy>>();
    const { setProxyDeleteMutation } = useProxyDeleteMutation();

    const handleOpen = () => setOpenDialog(!openDialog);
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
        enableRowSelection: true,
    } as TableOptions<Proxy>);

    const closeDialog = () => {
        setCurrentRow(undefined);
        setOpenDialog(false);
    };

    return (
        <div className={styles["container"]}>
            <div className="text-right">
                <Button
                    title="Add a new proxy"
                    onClick={() => {
                        setCurrentRow(undefined);
                        setOpenDialog(true);
                    }}
                >
                    <FontAwesomeIcon icon={faAdd} />
                </Button>
            </div>
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
                            <th className="border-blue-gray-50 border-b py-3 px-5 text-left">
                                <Typography
                                    variant="small"
                                    className="text-blue-gray-400 text-[11px] font-bold uppercase"
                                >
                                    Actions
                                </Typography>
                            </th>
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
                            <tr key={row.id}>
                                {row.getVisibleCells().map(cell => (
                                    <td key={cell.id} className={className}>
                                        {flexRender(
                                            cell.column.columnDef.cell,
                                            cell.getContext(),
                                        )}
                                    </td>
                                ))}
                                <td>
                                    <Button
                                        data-ripple-light="true"
                                        data-dialog-target="dialog"
                                        className="select-none disabled:pointer-events-none"
                                        onClick={() => {
                                            setCurrentRow(row);
                                            handleOpen();
                                        }}
                                    >
                                        <FontAwesomeIcon icon={faPencil} />
                                    </Button>
                                    <Button
                                        data-ripple-light="true"
                                        data-dialog-target="dialog"
                                        className="ml-2 select-none disabled:pointer-events-none"
                                        onClick={() => {
                                            setProxyDeleteMutation.mutateAsync({
                                                proxy: row.original,
                                            });
                                        }}
                                    >
                                        <FontAwesomeIcon icon={faTrash} />
                                    </Button>
                                </td>
                            </tr>
                        );
                    })}
                </tbody>
            </table>
            <Dialog open={openDialog} handler={handleOpen}>
                <>
                    <DialogHeader>
                        {currentRow?.original.name ?? "New proxy"}
                    </DialogHeader>
                    <DialogBody>
                        <ProxyForm
                            proxy={
                                currentRow?.original ??
                                ({
                                    method: Method.Get,
                                    name: "",
                                    path: "",
                                    source: "",
                                } as Proxy)
                            }
                            cancelFn={closeDialog}
                            savedFn={closeDialog}
                        />
                    </DialogBody>
                </>
            </Dialog>
        </div>
    );
}

export default ProxyTable;
