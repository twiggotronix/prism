import { render } from "@testing-library/react";

import { QueryClientProvider } from "@tanstack/react-query";
import { RouterProvider, createMemoryRouter } from "react-router-dom";
import { queryClientFactory } from "../../../../../tests/query-client-factory";
import { routes } from "../../../routes/routes";
import type { BreadcrumbItem } from "../models/breadcrumb-item";

describe("Breadcrumb", () => {
    const queryClient = queryClientFactory();
    const renderBreadcrumb = (breadCrumbItems: BreadcrumbItem[]) => {
        const router = createMemoryRouter(routes);

        router.navigate("/");

        return render(
            <QueryClientProvider client={queryClient}>
                <RouterProvider router={router} />
            </QueryClientProvider>,
        );
    };
    it("should render successfully", () => {
        const { baseElement } = renderBreadcrumb([
            { label: "test", to: routes[0].path as string },
        ]);
        expect(baseElement).toBeTruthy();
    });
});
