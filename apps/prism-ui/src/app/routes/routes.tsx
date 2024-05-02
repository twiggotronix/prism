import type { RouteObject } from "react-router-dom";
import Home from "../pages/home/home";

export const routes: RouteObject[] = [
    {
        path: "/",
        element: <Home />,
    },
];
