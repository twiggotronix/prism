import { faClose, faSave } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { Button } from "@material-tailwind/react";
import { FormProvider, useForm } from "react-hook-form";
import SelectInput from "../../../common/form/select-input/select-input";
import TextInput from "../../../common/form/text-input/text-input";
import { useProxyMutation } from "../../hooks/use-proxy-mutation.hook";
import { type Proxy } from "../../models/proxy";
import styles from "./proxy-form.module.scss";

export interface ProxyFormProps {
    proxy?: Proxy;
    cancelFn?: () => void;
    savedFn?: () => void;
}

export function ProxyForm({ proxy, cancelFn, savedFn }: ProxyFormProps) {
    const { setProxyMutation } = useProxyMutation();

    const form = useForm<Proxy>();
    const onSubmit = form.handleSubmit(formData => {
        setProxyMutation.mutateAsync({ proxy: { ...formData, id: proxy?.id } });
        if (savedFn != null) {
            savedFn();
        }
    });

    return (
        <div className={styles["container"]}>
            <FormProvider {...form}>
                <form
                    noValidate
                    autoComplete="off"
                    onSubmit={e => {
                        e.preventDefault();
                    }}
                >
                    <div className="mb-1 flex flex-col gap-6">
                        <TextInput
                            id={"proxy-name"}
                            placeholder="Name"
                            label="Name"
                            name="name"
                            defaultValue={proxy?.name ?? ""}
                            validation={{
                                required: {
                                    value: true,
                                    message: "A name is required",
                                },
                            }}
                        ></TextInput>
                        <SelectInput
                            defaultValue={proxy?.method}
                            id="method"
                            label="Method"
                            name="name"
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
                            label="Path"
                            placeholder="Path"
                            defaultValue={proxy?.path ?? ""}
                            name="path"
                            validation={{
                                required: {
                                    value: true,
                                    message: "Please input a path",
                                },
                            }}
                        ></TextInput>

                        <TextInput
                            id={"proxy-source"}
                            label="Source"
                            placeholder="Source"
                            defaultValue={proxy?.source ?? ""}
                            name="source"
                            validation={{
                                required: {
                                    value: true,
                                    message: "Please input a source",
                                },
                            }}
                        ></TextInput>
                    </div>
                    <div className="mt-10 text-right">
                        <Button onClick={onSubmit} variant="outlined">
                            <FontAwesomeIcon
                                icon={faSave}
                                size={"lg"}
                                className="mr-1"
                            />
                            save
                        </Button>
                        <Button onClick={cancelFn} className="ml-3">
                            <FontAwesomeIcon
                                icon={faClose}
                                size={"lg"}
                                className="mr-1"
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
