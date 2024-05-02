import { render } from "@testing-library/react";

import { QueryClientProvider } from "@tanstack/react-query";
import { queryClientFactory } from "../../../../../tests/query-client-factory";
import Proxies from "./proxies";

describe("Proxies", () => {
    const queryClient = queryClientFactory();
    const renderProxies = () =>
        render(
            <QueryClientProvider client={queryClient}>
                <Proxies />
            </QueryClientProvider>,
        );
    it("should render successfully", () => {
        const { baseElement } = renderProxies();
        expect(baseElement).toBeTruthy();
    });
});
