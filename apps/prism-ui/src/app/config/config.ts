import type { EnvType } from "./app-config";
import { DevAppConfig } from "./config.development";
import { ProdAppConfig } from "./config.production";

const currentEnv: EnvType = (process.env.NODE_ENV as EnvType) ?? "development";

export const AppConfig =
    currentEnv === "development" ? DevAppConfig : ProdAppConfig;
