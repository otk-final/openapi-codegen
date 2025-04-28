import {ApiClient} from "./client"

/**
 *
 */
export interface ApiResult<T> {
    //业务响应码
    code?: string
    //数据
    data?: T
    //业务错误信息
    err?: string
    //业务响应信息
    msg?: string

}

/**
 *
 */
export interface PageData<T> {
    //数据
    entities?: T[]
    //页码
    page?: number
    //条数
    size?: number
    //总数
    total?: number

}

/**
 *
 */
export interface MapResult<T> {
    //响应码
    code?: string
    //数据
    data?: Record<string, T>

}

/**
 * 猫
 */
export interface Cat {
    //食物
    foods?: string[]
    //昵称
    nickname?: string

}

/**
 * 狗
 */
export interface Dog {
    //颜色
    color?: string
    //ID
    id?: number
    //昵称
    nickname?: string
    //售价
    price?: number

}

/**
 * 用户
 */
export interface User {
    //年龄
    age?: number
    //生日
    birthday?: string
    //集合
    cats?: Cat[]
    //是否启用
    enable?: boolean
    //引用map
    gods?: Record<string, Dog>
    //ID
    id?: number
    //基础map
    kv?: Record<string, string>
    //名称
    name?: string
    //
    profile?: any
    //标签
    tags?: string[]
    //时间戳
    timestamp?: number

}

/**
 * 动物接口
 */
export const AnimalController = {

    /**
     *
     * 响应long
     */
    addCat: async (body: Cat) => {
        return await ApiClient.post("/animal/cat/add", body) as number
    },


    /**
     *
     * 路径+Body
     */
    addDog: async (kind: string, body: Dog) => {
        return await ApiClient.post(`/animal/add/${kind}/dog`, body) as Dog
    },


    /**
     *
     * 响应list
     */
    cats: async () => {
        let params = {};
        return await ApiClient.get("/animal/cats", params) as Cat[]
    },


    /**
     *
     * 响应对象
     */
    dog: async (id: number) => {
        let params = {};
        return await ApiClient.get(`/animal/dog/${id}`, params) as Dog
    },


    /**
     * 其他
     * 响应map
     */
    mapCats: async () => {
        let params = {};
        return await ApiClient.get("/animal/map/cats", params) as Record<string, Cat>
    },


    /**
     *
     * 响应Map泛型
     */
    mapDog: async () => {
        let params = {};
        return await ApiClient.get("/animal/map/dogs", params) as MapResult<Dog>
    },


    /**
     *
     * 响应Array泛型
     */
    pageDogs: async () => {
        let params = {};
        return await ApiClient.get("/animal/page/dogs", params) as PageData<Dog>
    },


    /**
     *
     * 路径+Query
     */
    searchDog: async (color: string, kind: string, name: string) => {
        let params = {color: color, name: name};
        return await ApiClient.get(`/animal/query/${kind}/dog`, params) as Dog
    },


}

/**
 * 其他接口
 */
export const OtherApi = {

    /**
     *
     * 响应基础类型
     */
    get: async () => {
        let params = {};
        return await ApiClient.get("/other/long", params) as string
    },


    /**
     *
     * 响应基础类型
     */
    getContent: async () => {
        let params = {};
        return await ApiClient.get("/other/string", params) as number
    },


    /**
     *
     * 多路径参数
     */
    multiplePathVariable: async (id: string, type: string) => {
        let params = {};
        return await ApiClient.get(`/other/boolean/${type}/${id}`, params) as Dog[]
    },


}

/**
 * 用户接口
 */
export const UserApi = {

    /**
     *
     * 添加用户
     */
    add: async (body: User) => {
        return await ApiClient.post("/user/add", body) as ApiResult<number>
    },


    /**
     *
     * 删除用户
     */
    delete: async (id: number) => {
        let params = {};
        return await ApiClient.delete(`/user/delete/${id}`, params) as ApiResult<boolean>
    },


    /**
     *
     * 修改用户
     */
    edit: async (id: number, body: User) => {
        return await ApiClient.put(`/user/edit/${id}`, body) as ApiResult<number>
    },


    /**
     *
     * 查询用户
     */
    edit_1: async (id: number, keyword: string) => {
        let params = {keyword: keyword};
        return await ApiClient.get(`/user/detail/${id}`, params) as ApiResult<User>
    },


    /**
     *
     * 用户列表
     */
    list: async () => {
        let params = {};
        return await ApiClient.get("/user/list", params) as ApiResult<User[]>
    },


    /**
     *
     * 用户分页
     */
    page: async (page: number, size: number) => {
        let params = {page: page, size: size};
        return await ApiClient.get("/user/page", params) as ApiResult<PageData<User>>
    },


}
