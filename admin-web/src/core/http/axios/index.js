import axios from "axios";
import {GcsConf} from "@/config/GcsConf.js";

export class GcsAxiosHttp {
    constructor() {
        this.axiosRequestConfig = {};
        this.requestConfig()
        this.axiosInstance = axios.create(this.axiosRequestConfig);
        this.interceptorsRequest()
        this.interceptorsResponse()
    }

    get(url, config) {
        this.axiosInstance.get(url, config)
    }

    post(url, data, config) {
        return this.axiosInstance.post(url, data, config)
    }

    request(config) {
        return this.axiosInstance.request(config);
    }

    // 请求配置
    requestConfig() {
        this.axiosRequestConfig = {
            baseURL: GcsConf.env.apiBaseUrl,
            timeout: 5000,
        }
    }

    // 请求拦截
    interceptorsRequest() {
        this.axiosInstance.interceptors.request.use((config) => {
            config.data = config.data ? config.data : {}
            config.params = config.params ? config.params : {}
            return config;
        })
    }

    // 响应拦截
    interceptorsResponse() {
        this.axiosInstance.interceptors.response.use((res) => {
            return res.data;
        })
    }
}