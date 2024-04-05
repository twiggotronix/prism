import { axiosInstance } from "../../core/http/axios";
import type { Proxy } from "../models/proxy";

export const ProxyService = {
    getProxies: (): Promise<Proxy[]> =>
        axiosInstance.get<Proxy[]>(`/proxies`).then(response => response.data),
    updateProxy: (proxy: Proxy) =>
        axiosInstance
            .post<Proxy>(`/proxies`, proxy)
            .then(response => response.data),
    deleteProxy: (proxy: Proxy) =>
        axiosInstance.delete<Proxy>(`/proxies/${proxy.id}`),
};
