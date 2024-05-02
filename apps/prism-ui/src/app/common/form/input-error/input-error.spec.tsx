import { render } from "@testing-library/react";

import { QueryClientProvider } from "@tanstack/react-query";
import { queryClientFactory } from "../../../../../tests/query-client-factory";
import InputError from "./input-error";

describe("InputError", () => {
    const queryClient = queryClientFactory();
    const renderInputError = () =>
        render(
            <QueryClientProvider client={queryClient}>
                <InputError message="a test message" />
            </QueryClientProvider>,
        );
    it("should render successfully", () => {
        const { baseElement } = renderInputError();
        expect(baseElement).toBeTruthy();
    });
});
