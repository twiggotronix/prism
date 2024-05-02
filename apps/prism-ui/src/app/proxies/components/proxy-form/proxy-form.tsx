import { faClose, faSave } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { DevTool } from "@hookform/devtools";
import { Button } from "@material-tailwind/react";
import { FormProvider, useForm } from "react-hook-form";
import SelectInput from "../../../common/form/select-input/select-input";
import TextInput from "../../../common/form/text-input/text-input";
import { useProxyMutation } from "../../hooks/use-proxy-mutation.hook";
import { Method, type Proxy } from "../../models/proxy";
import styles from "./proxy-form.module.scss";

export interface ProxyFormProps {
    proxy?: Proxy;
    cancelFn?: () => void;
    savedFn?: () => void;
}
type ProxyForm = {
    name: string;
    path: string;
    method: string;
    source: string;
};

export function ProxyForm({ proxy, cancelFn, savedFn }: ProxyFormProps) {
    const { setProxyMutation } = useProxyMutation();

    const form = useForm<ProxyForm>({
        mode: "onBlur",
        defaultValues: {
            name: proxy?.name,
            path: proxy?.path,
            method: proxy?.method ?? Method.Get,
            source: proxy?.source,
        },
    });
    const onSubmit = form.handleSubmit(formData => {
        setProxyMutation.mutateAsync({
            proxy: {
                ...formData,
                id: proxy?.id,
                method: formData.method as Method,
            },
        });
        if (savedFn != null) {
            savedFn();
        }
    });

    return (
        <div className={styles["container"]}>
            <DevTool control={form.control} placement="top-right" />
            <FormProvider {...form}>
                <form
                    noValidate
                    autoComplete="off"
                    onSubmit={e => {
                        e.preventDefault();
                    }}
                >
                    <div className={"mb-1 flex flex-col gap-6"}>
                        <TextInput
                            id={"proxy-name"}
                            placeholder={"Name"}
                            label={"Name"}
                            name={"name"}
                            validation={{
                                required: {
                                    value: true,
                                    message: "Please enter a name",
                                },
                            }}
                        ></TextInput>
                        <SelectInput
                            id={"proxy-method"}
                            label={"Method"}
                            name={"method"}
                            options={[
                                { value: "GET", label: "GET" },
                                { value: "POST", label: "POST" },
                                { value: "PUT", label: "PUT" },
                                { value: "DELETE", label: "DELETE" },
                            ]}
                            validation={{
                                required: {
                                    value: true,
                                    message: "Please select a method",
                                },
                            }}
                        />

                        <TextInput
                            id={"proxy-path"}
                            label={"Path"}
                            placeholder={"Path"}
                            name={"path"}
                            validation={{
                                required: {
                                    value: true,
                                    message: "Please enter a path",
                                },
                            }}
                        ></TextInput>

                        <TextInput
                            id={"proxy-source"}
                            type={"url"}
                            label={"Source"}
                            placeholder={"Source"}
                            name={"source"}
                            validation={{
                                required: {
                                    value: true,
                                    message: "Please enter a source",
                                },
                                pattern: {
                                    value: /^(https?:\/\/)[-a-zA-Z0-9@:.]{2,256}\.[a-z]{2,6}\/?[-a-zA-Z0-9@:%_+.~#?&/=]*$/gi,
                                    message: "Please enter a valid URL",
                                },
                            }}
                        ></TextInput>
                    </div>
                    <div className={"mt-10 text-right"}>
                        <Button
                            onClick={onSubmit}
                            variant="outlined"
                            disabled={!form.formState.isValid}
                        >
                            <FontAwesomeIcon
                                icon={faSave}
                                size={"lg"}
                                className={"mr-1"}
                            />
                            save
                        </Button>
                        <Button onClick={cancelFn} className={"ml-3"}>
                            <FontAwesomeIcon
                                icon={faClose}
                                size={"lg"}
                                className={"mr-1"}
                            />
                            Close
                        </Button>
                    </div>
                </form>
            </FormProvider>
        </div>
    );
}

export default ProxyForm;
