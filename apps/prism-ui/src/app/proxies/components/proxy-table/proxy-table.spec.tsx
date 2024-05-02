import { render } from "@testing-library/react";

import { QueryClientProvider } from "@tanstack/react-query";
import { queryClientFactory } from "../../../../../tests/query-client-factory";
import ProxyTable from "./proxy-table";

describe("ProxyTable", () => {
    const queryClient = queryClientFactory();
    const renderProxyTable = () =>
        render(
            <QueryClientProvider client={queryClient}>
                <ProxyTable tableData={[]} />
            </QueryClientProvider>,
        );
    it("should render successfully", () => {
        const { baseElement } = renderProxyTable();
        expect(baseElement).toBeTruthy();
    });
});
