import { render } from "@testing-library/react";

import { QueryClientProvider } from "@tanstack/react-query";
import { queryClientFactory } from "../../../../../tests/query-client-factory";
import FormLabel from "./form-label";

describe("FormLabel", () => {
    const queryClient = queryClientFactory();
    const renderFormLabel = () =>
        render(
            <QueryClientProvider client={queryClient}>
                <FormLabel htmlFor="testId" label="test" />
            </QueryClientProvider>,
        );
    it("should render successfully", () => {
        const { baseElement } = renderFormLabel();
        expect(baseElement).toBeTruthy();
    });
});
