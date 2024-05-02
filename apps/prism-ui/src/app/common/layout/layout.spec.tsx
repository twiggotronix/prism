import { render } from "@testing-library/react";

import { QueryClientProvider } from "@tanstack/react-query";
import { RouterProvider, createMemoryRouter } from "react-router-dom";
import { queryClientFactory } from "../../../../tests/query-client-factory";
import { routes } from "../../routes/routes";

describe("Layout", () => {
    const queryClient = queryClientFactory();
    const renderLayout = () => {
        const router = createMemoryRouter(routes);

        router.navigate(routes[0].path as string);
        return render(
            <QueryClientProvider client={queryClient}>
                <RouterProvider router={router} />
            </QueryClientProvider>,
        );
    };
    it("should render successfully", () => {
        const { baseElement } = renderLayout();
        expect(baseElement).toBeTruthy();
    });
});
