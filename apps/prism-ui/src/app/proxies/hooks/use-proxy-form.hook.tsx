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

    const updateProxy = (proxy: Proxy): void => {
        setProxyMutation.mutateAsync({ proxy });
    };

    return {
        formData,
        setFormData,
        updateProxy,
    };
};
