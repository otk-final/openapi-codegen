




/**
 * 接口标准响应
 * 
 */
public class ApiResultBoolean
{
    /**
     * 业务响应码
     * 
     */
    public string? Code { get; set; }
    /**
     * 数据
     * 
     */
    public bool? Data { get; set; }
    /**
     * 业务错误信息
     * 
     */
    public string? Err { get; set; }
    /**
     * 业务响应信息
     * 
     */
    public string? Msg { get; set; }
    
}




/**
 * 接口标准响应
 * 
 */
public class ApiResultLong
{
    /**
     * 业务响应码
     * 
     */
    public string? Code { get; set; }
    /**
     * 数据
     * 
     */
    public long? Data { get; set; }
    /**
     * 业务错误信息
     * 
     */
    public string? Err { get; set; }
    /**
     * 业务响应信息
     * 
     */
    public string? Msg { get; set; }
    
}




/**
 * 猫
 * 
 */
public class Cat
{
    /**
     * 食物
     * 
     */
    public List<string>? Foods { get; set; }
    /**
     * 昵称
     * 
     */
    public string? Nickname { get; set; }
    
}




/**
 * 狗
 * 
 */
public class Dog
{
    /**
     * 颜色
     * 
     */
    public string? Color { get; set; }
    /**
     * ID
     * 
     */
    public long? Id { get; set; }
    /**
     * 昵称
     * 
     */
    public string? Nickname { get; set; }
    /**
     * 售价
     * 
     */
    public int? Price { get; set; }
    
}




/**
 * 接口标准响应
 * 
 */
public class ApiResultListUser
{
    /**
     * 业务响应码
     * 
     */
    public string? Code { get; set; }
    /**
     * 数据
     * 
     */
    public List<User>? Data { get; set; }
    /**
     * 业务错误信息
     * 
     */
    public string? Err { get; set; }
    /**
     * 业务响应信息
     * 
     */
    public string? Msg { get; set; }
    
}




/**
 * 接口标准响应
 * 
 */
public class ApiResultPageDataUser
{
    /**
     * 业务响应码
     * 
     */
    public string? Code { get; set; }
    /**
     * 
     * 
     */
    public PageDataUser? Data { get; set; }
    /**
     * 业务错误信息
     * 
     */
    public string? Err { get; set; }
    /**
     * 业务响应信息
     * 
     */
    public string? Msg { get; set; }
    
}




/**
 * 接口标准响应
 * 
 */
public class ApiResultUser
{
    /**
     * 业务响应码
     * 
     */
    public string? Code { get; set; }
    /**
     * 
     * 
     */
    public User? Data { get; set; }
    /**
     * 业务错误信息
     * 
     */
    public string? Err { get; set; }
    /**
     * 业务响应信息
     * 
     */
    public string? Msg { get; set; }
    
}




/**
 * 接口标准响应
 * 
 */
public class MapResultDog
{
    /**
     * 响应码
     * 
     */
    public string? Code { get; set; }
    /**
     * 数据
     * 
     */
    public Dictionary<string,Dog>? Data { get; set; }
    
}




/**
 * 标准分页数据
 * 
 */
public class PageDataDog
{
    /**
     * 数据
     * 
     */
    public List<Dog>? Entities { get; set; }
    /**
     * 页码
     * 
     */
    public int? Page { get; set; }
    /**
     * 条数
     * 
     */
    public int? Size { get; set; }
    /**
     * 总数
     * 
     */
    public long? Total { get; set; }
    
}




/**
 * 标准分页数据
 * 
 */
public class PageDataUser
{
    /**
     * 数据
     * 
     */
    public List<User>? Entities { get; set; }
    /**
     * 页码
     * 
     */
    public int? Page { get; set; }
    /**
     * 条数
     * 
     */
    public int? Size { get; set; }
    /**
     * 总数
     * 
     */
    public long? Total { get; set; }
    
}




/**
 * 用户
 * 
 */
public class User
{
    /**
     * 年龄
     * 
     */
    public int? Age { get; set; }
    /**
     * 生日
     * 
     */
    public string? Birthday { get; set; }
    /**
     * 集合
     * 
     */
    public List<Cat>? Cats { get; set; }
    /**
     * 是否启用
     * 
     */
    public bool? Enable { get; set; }
    /**
     * 引用map
     * 
     */
    public Dictionary<string,Dog>? Gods { get; set; }
    /**
     * ID
     * 
     */
    public long? Id { get; set; }
    /**
     * 基础map
     * 
     */
    public Dictionary<string,string>? Kv { get; set; }
    /**
     * 名称
     * 
     */
    public string? Name { get; set; }
    /**
     * 
     * 
     */
    public JsonNode? Profile { get; set; }
    /**
     * 标签
     * 
     */
    public List<string>? Tags { get; set; }
    /**
     * 时间戳
     * 
     */
    public long? Timestamp { get; set; }
    
}




