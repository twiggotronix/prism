import { render } from "@testing-library/react";

import { QueryClientProvider } from "@tanstack/react-query";
import { RouterProvider, createMemoryRouter } from "react-router-dom";
import { queryClientFactory } from "../../../../tests/query-client-factory";
import { routes } from "../../routes/routes";

describe("Home", () => {
    const queryClient = queryClientFactory();
    const renderHome = () => {
        const router = createMemoryRouter(routes);

        router.navigate(routes[0].path as string);

        return render(
            <QueryClientProvider client={queryClient}>
                <RouterProvider router={router} />
            </QueryClientProvider>,
        );
    };
    it("should render Home component successfully", () => {
        const { baseElement } = renderHome();
        expect(baseElement).toBeTruthy();
    });
});
