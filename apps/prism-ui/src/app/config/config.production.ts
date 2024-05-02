import type { AppConfiguration } from "./app-config";

export const ProdAppConfig: AppConfiguration = {
    baseApi: process.env.BASE_API ?? "",
};
