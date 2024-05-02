import { render, screen } from "@testing-library/react";
import { expect } from "vitest";

import { QueryClientProvider } from "@tanstack/react-query";
import ProxyForm from "./proxy-form";

import { queryClientFactory } from "../../../../../tests/query-client-factory";

describe("ProxyForm", () => {
    const queryClient = queryClientFactory();
    const renderProxyForm = () =>
        render(
            <QueryClientProvider client={queryClient}>
                <ProxyForm />
            </QueryClientProvider>,
        );

    it("should render the basic fields", () => {
        renderProxyForm();
        expect(
            screen.getByRole("textbox", { name: /name/i }),
        ).toBeInTheDocument();
        expect(
            screen.getByRole("textbox", { name: /path/i }),
        ).toBeInTheDocument();
        expect(
            screen.getByRole("textbox", { name: /source/i }),
        ).toBeInTheDocument();
    });
});
