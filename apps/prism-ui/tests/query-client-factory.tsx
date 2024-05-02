import { QueryClient } from "@tanstack/react-query";

export const queryClientFactory = () =>
    new QueryClient({
        defaultOptions: {
            queries: {
                retry: false,
            },
        },
    });
