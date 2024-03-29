import { StrictMode } from "react";
import * as ReactDOM from "react-dom/client";

import { ThemeProvider } from "@material-tailwind/react";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { RouterProvider, createBrowserRouter } from "react-router-dom";
import Home from "./app/pages/home/home";

const root = ReactDOM.createRoot(
    document.getElementById("root") as HTMLElement,
);
const router = createBrowserRouter([
    {
        path: "/",
        element: <Home />,
    },
]);

const queryClient = new QueryClient();
root.render(
    <StrictMode>
        <ThemeProvider>
            <QueryClientProvider client={queryClient}>
                <RouterProvider router={router} />
            </QueryClientProvider>
        </ThemeProvider>
    </StrictMode>,
);
