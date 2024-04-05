import { useMutation, useQueryClient } from "@tanstack/react-query";
import type { Proxy } from "../models/proxy";
import { proxyKey } from "../proxy.key";
import { ProxyService } from "../services/proxy.service";

export const useProxyDeleteMutation = () => {
    const queryClient = useQueryClient();

    const setProxyDeleteMutation = useMutation({
        mutationFn: ({ proxy }: { proxy: Proxy }) =>
            ProxyService.deleteProxy(proxy),
        ...{
            onSuccess: () =>
                queryClient.invalidateQueries({
                    queryKey: proxyKey.getProxies,
                }),
        },
    });

    return { setProxyDeleteMutation };
};
