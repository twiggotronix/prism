import { faClose, faSave } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { Button } from "@material-tailwind/react";
import { useState } from "react";
import { Method, type Proxy } from "../models/proxy";
import { useProxyMutation } from "./use-proxy-mutation.hook";

export const useProxyForm = () => {
    const [formData, setFormData] = useState<Omit<Proxy, "id">>({
        name: "",
        method: Method.Get,
        path: "",
        source: "",
    });
    const { setProxyMutation } = useProxyMutation();

    const handleChange = ({
        target,
    }: {
        target: { name: string; value: string };
    }) => {
        const { name, value } = target;
        setFormData(prevFormData => ({ ...prevFormData, [name]: value }));
    };

    const updateProxy = (initialProxy: Proxy): void => {
        const updatedProxy: Proxy = {
            name: formData.name,
            id: initialProxy.id,
            path: formData.path,
            method: formData.method,
            source: formData.source,
        };

        setProxyMutation.mutateAsync({ proxy: updatedProxy });
    };

    const renderFormActionButtons = (cancelFn?: () => void) => {
        return (
            <>
                <Button type="submit" variant="outlined">
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
            </>
        );
    };

    return {
        formData,
        setFormData,
        handleChange,
        updateProxy,
        renderFormActionButtons,
    };
};
