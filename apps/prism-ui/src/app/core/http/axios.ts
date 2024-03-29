import axios from "axios";
import { AppConfig } from "../../config/config";

export const axiosInstance = axios.create({
    baseURL: AppConfig.baseApi,
});
