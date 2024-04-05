import { useMutation, useQueryClient } from "@tanstack/react-query";
import type { Proxy } from "../models/proxy";
import { proxyKey } from "../proxy.key";
import { ProxyService } from "../services/proxy.service";

export const useProxyMutation = () => {
    const queryClient = useQueryClient();

    const setProxyMutation = useMutation({
        mutationFn: ({ proxy }: { proxy: Proxy }) =>
            ProxyService.updateProxy(proxy),
        ...{
            onSuccess: () =>
                queryClient.invalidateQueries({
                    queryKey: proxyKey.getProxies,
                }),
        },
    });

    return { setProxyMutation };
};
