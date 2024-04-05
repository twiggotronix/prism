import { Input, Option, Select, Typography } from "@material-tailwind/react";
import { useEffect } from "react";
import { useProxyForm } from "../../hooks/use-proxy-form.hook";
import { Method, type Proxy } from "../../models/proxy";
import styles from "./proxy-form.module.scss";

export interface ProxyFormProps {
    proxy?: Proxy;
    cancelFn?: () => void;
    savedFn?: () => void;
}

export function ProxyForm({ proxy, cancelFn, savedFn }: ProxyFormProps) {
    const {
        formData,
        setFormData,
        handleChange,
        updateProxy,
        renderFormActionButtons,
    } = useProxyForm();

    useEffect(() => {
        setFormData(proxy);
    }, [proxy, setFormData]);

    return (
        <div className={styles["container"]}>
            <form
                onSubmit={e => {
                    e.preventDefault();
                    e.stopPropagation();
                    updateProxy(proxy);
                    if (savedFn != null) {
                        savedFn();
                    }
                }}
            >
                <div className="mb-1 flex flex-col gap-6">
                    <Typography
                        variant="h6"
                        color="blue-gray"
                        className="-mb-3"
                    >
                        Name
                    </Typography>
                    <Input
                        placeholder="Name"
                        size="md"
                        value={formData.name}
                        className="!border-t-blue-gray-200 focus:!border-t-gray-900"
                        labelProps={{
                            className: "before:content-none after:content-none",
                        }}
                        name="name"
                        onChange={handleChange}
                        crossOrigin={undefined}
                    ></Input>
                    <Typography
                        variant="h6"
                        color="blue-gray"
                        className="-mb-3"
                    >
                        Method
                    </Typography>
                    <Select
                        value={formData.method}
                        label="method"
                        name="method"
                        onChange={value =>
                            handleChange({
                                target: {
                                    name: "method",
                                    value: value ?? Method.Get,
                                },
                            })
                        }
                    >
                        <Option value="GET">GET</Option>
                        <Option value="POST">POST</Option>
                        <Option value="PUT">PUT</Option>
                        <Option value="DELETE">DELETE</Option>
                    </Select>
                    <Typography
                        variant="h6"
                        color="blue-gray"
                        className="-mb-3"
                    >
                        Path
                    </Typography>
                    <Input
                        placeholder="Path"
                        size="md"
                        value={formData.path}
                        className=" !border-t-blue-gray-200 focus:!border-t-gray-900"
                        labelProps={{
                            className: "before:content-none after:content-none",
                        }}
                        name="path"
                        onChange={handleChange}
                        crossOrigin={undefined}
                    ></Input>
                    <Typography
                        variant="h6"
                        color="blue-gray"
                        className="-mb-3"
                    >
                        Source
                    </Typography>
                    <Input
                        placeholder="Source"
                        size="md"
                        value={formData.source}
                        className=" !border-t-blue-gray-200 focus:!border-t-gray-900"
                        labelProps={{
                            className: "before:content-none after:content-none",
                        }}
                        name="source"
                        onChange={handleChange}
                        crossOrigin={undefined}
                    ></Input>
                </div>
                <div className="mt-10 text-right">
                    {renderFormActionButtons(cancelFn)}
                </div>
            </form>
        </div>
    );
}

export default ProxyForm;
