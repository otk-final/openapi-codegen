
import axios, {AxiosError, AxiosResponse} from "axios";

export const axiosClient = axios.create({
    baseURL: "http://localhost:8083",
    headers: {
        'Content-Type': 'application/json;charset=utf-8'
    },
    timeout: Number(60000),
    withCredentials: false
})

axiosClient.interceptors.response.use((response: AxiosResponse) => {
    return response.data
}, (error: AxiosError) => {
    return Promise.reject(error)
})

export const ApiClient = {

    get: async (path: string, params?: any) => {
        return await axiosClient.get(path,{params:params}) as any
    },

    post: async (path: string, body?: any) => {
        return await axiosClient.post(path, body) as any
    },

    put: async (path: string, body?: any) => {
        return await axiosClient.put(path, body) as any
    },

    delete: async (path: string, body?: any) => {
        return await axiosClient.delete(path, body) as any
    },
}
