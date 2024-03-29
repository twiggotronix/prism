import type { AppConfiguration } from "./app-config";

export const DevAppConfig: AppConfiguration = {
    baseApi: process.env.BASE_API ?? "http://localhost:9090/api",
};
