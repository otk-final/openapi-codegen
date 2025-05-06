
import Alamofire
import Foundation

/**
 * 接口标准响应
 */
public struct ApiResultBoolean: Codable {
    // 业务响应码
    var code: String?
    // 数据
    var data: Bool?
    // 业务错误信息
    var err: String?
    // 业务响应信息
    var msg: String?
}

/**
 * 接口标准响应
 */
public struct ApiResultListUser: Codable {
    // 业务响应码
    var code: String?
    // 数据
    var data: [User]?
    // 业务错误信息
    var err: String?
    // 业务响应信息
    var msg: String?
}

/**
 * 接口标准响应
 */
public struct ApiResultLong: Codable {
    // 业务响应码
    var code: String?
    // 数据
    var data: Int64?
    // 业务错误信息
    var err: String?
    // 业务响应信息
    var msg: String?
}

/**
 * 接口标准响应
 */
public struct ApiResultPageDataUser: Codable {
    // 业务响应码
    var code: String?
    //
    var data: PageDataUser?
    // 业务错误信息
    var err: String?
    // 业务响应信息
    var msg: String?
}

/**
 * 接口标准响应
 */
public struct ApiResultUser: Codable {
    // 业务响应码
    var code: String?
    //
    var data: User?
    // 业务错误信息
    var err: String?
    // 业务响应信息
    var msg: String?
}

/**
 * 猫
 */
public struct Cat: Codable {
    // 食物
    var foods: [String]?
    // 昵称
    var nickname: String?
}

/**
 * 狗
 */
public struct Dog: Codable {
    // 颜色
    var color: String?
    // ID
    var id: Int64?
    // 昵称
    var nickname: String?
    // 售价
    var price: Int?
}

/**
 * 接口标准响应
 */
public struct MapResultDog: Codable {
    // 响应码
    var code: String?
    // 数据
    var data: [String: Dog]?
}

/**
 * 标准分页数据
 */
public struct PageDataDog: Codable {
    // 数据
    var entities: [Dog]?
    // 页码
    var page: Int?
    // 条数
    var size: Int?
    // 总数
    var total: Int64?
}

/**
 * 标准分页数据
 */
public struct PageDataUser: Codable {
    // 数据
    var entities: [User]?
    // 页码
    var page: Int?
    // 条数
    var size: Int?
    // 总数
    var total: Int64?
}

/**
 * 用户
 */
public struct User: Codable {
    // 年龄
    var age: Int?
    // 生日
    var birthday: String?
    // 集合
    var cats: [Cat]?
    // 是否启用
    var enable: Bool?
    // 引用map
    var gods: [String: Dog]?
    // ID
    var id: Int64?
    // 基础map
    var kv: [String: String]?
    // 名称
    var name: String?
    //
    var profile: any?
    // 标签
    var tags: [String]?
    // 时间戳
    var timestamp: Int64?
}

/**
 * 动物接口
 */
class AnimalController {
    var client: ApiClient

    init(client: ApiClient) {
        self.client = client
    }

    /**
     * 响应long
     */
    func addCat(body: Cat) async throws -> Int {
        return try await client.post("/animal/cat/add", body, Int.self)
    }

    /**
     * 路径+Body
     */
    func addDog(kind: String, body: Dog) async throws -> Dog {
        return try await client.post(String(format: "/animal/add/%@/dog", kind), body, Dog.self)
    }

    /**
     * 响应list
     */
    func cats() async throws -> [Cat] {
        var params: [String: Any?] = [:]
        return try await client.get("/animal/cats", params, [Cat].self)
    }

    /**
     * 响应对象
     */
    func dog(id: Int) async throws -> Dog {
        var params: [String: Any?] = [:]
        return try await client.get(String(format: "/animal/dog/%@", id), params, Dog.self)
    }

    /**
     * 响应map
     */
    func mapCats() async throws -> [String: Cat] {
        var params: [String: Any?] = [:]
        return try await client.get("/animal/map/cats", params, [String: Cat].self)
    }

    /**
     * 响应Map泛型
     */
    func mapDog() async throws -> MapResultDog {
        var params: [String: Any?] = [:]
        return try await client.get("/animal/map/dogs", params, MapResultDog.self)
    }

    /**
     * 响应Array泛型
     */
    func pageDogs() async throws -> PageDataDog {
        var params: [String: Any?] = [:]
        return try await client.get("/animal/page/dogs", params, PageDataDog.self)
    }

    /**
     * 路径+Query
     */
    func searchDog(color: String, kind: String, name: String) async throws -> Dog {
        var params: [String: Any?] = [:]
        params["color"] = color
        params["name"] = name

        return try await client.get(String(format: "/animal/query/%@/dog", kind), params, Dog.self)
    }
}

/**
 * 其他接口
 */
class OtherApi {
    var client: ApiClient

    init(client: ApiClient) {
        self.client = client
    }

    /**
     * 响应基础类型
     */
    func get() async throws -> String {
        var params: [String: Any?] = [:]
        return try await client.get("/other/long", params, String.self)
    }

    /**
     * 响应基础类型
     */
    func getContent() async throws -> Int {
        var params: [String: Any?] = [:]
        return try await client.get("/other/string", params, Int.self)
    }

    /**
     * 多路径参数
     */
    func multiplePathVariable(id: String, type: String) async throws -> [Dog] {
        var params: [String: Any?] = [:]
        return try await client.get(String(format: "/other/boolean/%@/%@", type, id), params, [Dog].self)
    }
}

/**
 * 用户接口
 */
class UserApi {
    var client: ApiClient

    init(client: ApiClient) {
        self.client = client
    }

    /**
     * 添加用户
     */
    func add(body: User) async throws -> ApiResultLong {
        return try await client.post("/user/add", body, ApiResultLong.self)
    }

    /**
     * 删除用户
     */
    func delete(id: Int64) async throws -> ApiResultBoolean {
        var params: [String: Any?] = [:]
        return try await client.delete(String(format: "/user/delete/%@", id), params, ApiResultBoolean.self)
    }

    /**
     * 修改用户
     */
    func edit(id: Int64, body: User) async throws -> ApiResultLong {
        return try await client.put(String(format: "/user/edit/%@", id), body, ApiResultLong.self)
    }

    /**
     * 查询用户
     */
    func edit_1(id: Int64, keyword: String) async throws -> ApiResultUser {
        var params: [String: Any?] = [:]
        params["keyword"] = keyword

        return try await client.get(String(format: "/user/detail/%@", id), params, ApiResultUser.self)
    }

    /**
     * 用户列表
     */
    func list() async throws -> ApiResultListUser {
        var params: [String: Any?] = [:]
        return try await client.get("/user/list", params, ApiResultListUser.self)
    }

    /**
     * 用户分页
     */
    func page(page: Int, size: Int) async throws -> ApiResultPageDataUser {
        var params: [String: Any?] = [:]
        params["page"] = page
        params["size"] = size

        return try await client.get("/user/page", params, ApiResultPageDataUser.self)
    }
}
