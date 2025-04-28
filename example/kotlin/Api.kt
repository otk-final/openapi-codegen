package com.demo;

import com.demo.ApiClient;
import java.io.Serializable
import com.google.gson.reflect.TypeToken


/**
 *
 */
data class ApiResult<T>(
    //业务响应码
    var code: String,
    //数据
    var data: T,
    //业务错误信息
    var err: String,
    //业务响应信息
    var msg: String,

    ) : Serializable

/**
 *
 */
data class PageData<T>(
    //数据
    var entities: List<T>,
    //页码
    var page: Int,
    //条数
    var size: Int,
    //总数
    var total: Long,

    ) : Serializable

/**
 *
 */
data class MapResult<T>(
    //响应码
    var code: String,
    //数据
    var data: Map<String, T>,

    ) : Serializable

/**
 * 猫
 */
data class Cat(
    //食物
    var foods: List<String>,
    //昵称
    var nickname: String,

    ) : Serializable

/**
 * 狗
 */
data class Dog(
    //颜色
    var color: String,
    //ID
    var id: Long,
    //昵称
    var nickname: String,
    //售价
    var price: Int,

    ) : Serializable

/**
 * 用户
 */
data class User(
    //年龄
    var age: Int,
    //生日
    var birthday: String,
    //集合
    var cats: List<Cat>,
    //是否启用
    var enable: Boolean,
    //引用map
    var gods: Map<String, Dog>,
    //ID
    var id: Long,
    //基础map
    var kv: Map<String, String>,
    //名称
    var name: String,
    //
    var profile: Any,
    //标签
    var tags: List<String>,
    //时间戳
    var timestamp: Long,

    ) : Serializable

/**
 * 动物接口
 */
class AnimalController(private var client: ApiClient) {


    var addCatResultType = object : TypeToken<Int>() {}.type

    /**
     * 响应long
     *
     */
    suspend fun addCat(body: Cat): Result<Int> {
        return client.post("/animal/cat/add", addCatResultType, body)
    }


    var addDogResultType = object : TypeToken<Dog>() {}.type

    /**
     * 路径+Body
     *
     */
    suspend fun addDog(kind: String, body: Dog): Result<Dog> {
        return client.post("/animal/add/%s/dog".format(kind), addDogResultType, body)
    }


    var catsResultType = object : TypeToken<List<Cat>>() {}.type

    /**
     * 响应list
     *
     */
    suspend fun cats(): Result<List<Cat>> {
        val params = HashMap<String, Any>()
        return client.get("/animal/cats", catsResultType, params)
    }


    var dogResultType = object : TypeToken<Dog>() {}.type

    /**
     * 响应对象
     *
     */
    suspend fun dog(id: Int): Result<Dog> {
        val params = HashMap<String, Any>()
        return client.get("/animal/dog/%s".format(id), dogResultType, params)
    }


    var mapCatsResultType = object : TypeToken<Map<String, Cat>>() {}.type

    /**
     * 响应map
     * 其他
     */
    suspend fun mapCats(): Result<Map<String, Cat>> {
        val params = HashMap<String, Any>()
        return client.get("/animal/map/cats", mapCatsResultType, params)
    }


    var mapDogResultType = object : TypeToken<Dog>() {}.type

    /**
     * 响应Map泛型
     *
     */
    suspend fun mapDog(): Result<Dog> {
        val params = HashMap<String, Any>()
        return client.get("/animal/map/dogs", mapDogResultType, params)
    }


    var pageDogsResultType = object : TypeToken<Dog>() {}.type

    /**
     * 响应Array泛型
     *
     */
    suspend fun pageDogs(): Result<Dog> {
        val params = HashMap<String, Any>()
        return client.get("/animal/page/dogs", pageDogsResultType, params)
    }


    var searchDogResultType = object : TypeToken<Dog>() {}.type

    /**
     * 路径+Query
     *
     */
    suspend fun searchDog(color: String, kind: String, name: String): Result<Dog> {
        val params = HashMap<String, Any>()
        params["color"] = color
        params["kind"] = kind
        params["name"] = name

        return client.get("/animal/query/%s/dog".format(kind), searchDogResultType, params)
    }


}

/**
 * 其他接口
 */
class OtherApi(private var client: ApiClient) {


    var getResultType = object : TypeToken<String>() {}.type

    /**
     * 响应基础类型
     *
     */
    suspend fun get(): Result<String> {
        val params = HashMap<String, Any>()
        return client.get("/other/long", getResultType, params)
    }


    var getContentResultType = object : TypeToken<Int>() {}.type

    /**
     * 响应基础类型
     *
     */
    suspend fun getContent(): Result<Int> {
        val params = HashMap<String, Any>()
        return client.get("/other/string", getContentResultType, params)
    }


    var multiplePathVariableResultType = object : TypeToken<List<Dog>>() {}.type

    /**
     * 多路径参数
     *
     */
    suspend fun multiplePathVariable(id: String, type: String): Result<List<Dog>> {
        val params = HashMap<String, Any>()
        return client.get(
            "/other/boolean/%s/%s".format(type, id), multiplePathVariableResultType, params
        )
    }


}

/**
 * 用户接口
 */
class UserApi(private var client: ApiClient) {


    var addResultType = object : TypeToken<Long>() {}.type

    /**
     * 添加用户
     *
     */
    suspend fun add(body: User): Result<Long> {
        return client.post("/user/add", addResultType, body)
    }


    var deleteResultType = object : TypeToken<Boolean>() {}.type

    /**
     * 删除用户
     *
     */
    suspend fun delete(id: Long): Result<Boolean> {
        val params = HashMap<String, Any>()
        return client.delete("/user/delete/%s".format(id), deleteResultType, params)
    }


    var editResultType = object : TypeToken<Long>() {}.type

    /**
     * 修改用户
     *
     */
    suspend fun edit(id: Long, body: User): Result<Long> {
        return client.put("/user/edit/%s".format(id), editResultType, body)
    }


    var edit_1ResultType = object : TypeToken<User>() {}.type

    /**
     * 查询用户
     *
     */
    suspend fun edit_1(id: Long, keyword: String): Result<User> {
        val params = HashMap<String, Any>()
        params["id"] = id
        params["keyword"] = keyword

        return client.get("/user/detail/%s".format(id), edit_1ResultType, params)
    }


    var listResultType = object : TypeToken<List<User>>() {}.type

    /**
     * 用户列表
     *
     */
    suspend fun list(): Result<List<User>> {
        val params = HashMap<String, Any>()
        return client.get("/user/list", listResultType, params)
    }


    var pageResultType = object : TypeToken<PageData<User>>() {}.type

    /**
     * 用户分页
     *
     */
    suspend fun page(page: Int, size: Int): Result<PageData<User>> {
        val params = HashMap<String, Any>()
        params["page"] = page
        params["size"] = size

        return client.get("/user/page", pageResultType, params)
    }


}
