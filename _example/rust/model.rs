

use serde::{Deserialize, Serialize};
use std::iter::Map;





/**
 * 接口标准响应
 * 
 */
#[derive(Debug, Clone)]
pub struct ApiResultBoolean {
    /**
     * 业务响应码
     * 
     */
    pub code: Option<String>,
    /**
     * 数据
     * 
     */
    pub data: Option<bool>,
    /**
     * 业务错误信息
     * 
     */
    pub err: Option<String>,
    /**
     * 业务响应信息
     * 
     */
    pub msg: Option<String>,
    
}




/**
 * 接口标准响应
 * 
 */
#[derive(Debug, Clone)]
pub struct ApiResultLong {
    /**
     * 业务响应码
     * 
     */
    pub code: Option<String>,
    /**
     * 数据
     * 
     */
    pub data: Option<u64>,
    /**
     * 业务错误信息
     * 
     */
    pub err: Option<String>,
    /**
     * 业务响应信息
     * 
     */
    pub msg: Option<String>,
    
}




/**
 * 猫
 * 
 */
#[derive(Debug, Clone)]
pub struct Cat {
    /**
     * 食物
     * 
     */
    pub foods: Option<Vec<String>>,
    /**
     * 昵称
     * 
     */
    pub nickname: Option<String>,
    
}




/**
 * 狗
 * 
 */
#[derive(Debug, Clone)]
pub struct Dog {
    /**
     * 颜色
     * 
     */
    pub color: Option<String>,
    /**
     * ID
     * 
     */
    pub id: Option<u64>,
    /**
     * 昵称
     * 
     */
    pub nickname: Option<String>,
    /**
     * 售价
     * 
     */
    pub price: Option<usize>,
    
}




/**
 * 接口标准响应
 * 
 */
#[derive(Debug, Clone)]
pub struct ApiResultListUser {
    /**
     * 业务响应码
     * 
     */
    pub code: Option<String>,
    /**
     * 数据
     * 
     */
    pub data: Option<Vec<User>>,
    /**
     * 业务错误信息
     * 
     */
    pub err: Option<String>,
    /**
     * 业务响应信息
     * 
     */
    pub msg: Option<String>,
    
}




/**
 * 接口标准响应
 * 
 */
#[derive(Debug, Clone)]
pub struct ApiResultPageDataUser {
    /**
     * 业务响应码
     * 
     */
    pub code: Option<String>,
    /**
     * 
     * 
     */
    pub data: Option<PageDataUser>,
    /**
     * 业务错误信息
     * 
     */
    pub err: Option<String>,
    /**
     * 业务响应信息
     * 
     */
    pub msg: Option<String>,
    
}




/**
 * 接口标准响应
 * 
 */
#[derive(Debug, Clone)]
pub struct ApiResultUser {
    /**
     * 业务响应码
     * 
     */
    pub code: Option<String>,
    /**
     * 
     * 
     */
    pub data: Option<User>,
    /**
     * 业务错误信息
     * 
     */
    pub err: Option<String>,
    /**
     * 业务响应信息
     * 
     */
    pub msg: Option<String>,
    
}




/**
 * 接口标准响应
 * 
 */
#[derive(Debug, Clone)]
pub struct MapResultDog {
    /**
     * 响应码
     * 
     */
    pub code: Option<String>,
    /**
     * 数据
     * 
     */
    pub data: Option<Map<String,Dog>>,
    
}




/**
 * 标准分页数据
 * 
 */
#[derive(Debug, Clone)]
pub struct PageDataDog {
    /**
     * 数据
     * 
     */
    pub entities: Option<Vec<Dog>>,
    /**
     * 页码
     * 
     */
    pub page: Option<usize>,
    /**
     * 条数
     * 
     */
    pub size: Option<usize>,
    /**
     * 总数
     * 
     */
    pub total: Option<u64>,
    
}




/**
 * 标准分页数据
 * 
 */
#[derive(Debug, Clone)]
pub struct PageDataUser {
    /**
     * 数据
     * 
     */
    pub entities: Option<Vec<User>>,
    /**
     * 页码
     * 
     */
    pub page: Option<usize>,
    /**
     * 条数
     * 
     */
    pub size: Option<usize>,
    /**
     * 总数
     * 
     */
    pub total: Option<u64>,
    
}




/**
 * 用户
 * 
 */
#[derive(Debug, Clone)]
pub struct User {
    /**
     * 年龄
     * 
     */
    pub age: Option<usize>,
    /**
     * 生日
     * 
     */
    pub birthday: Option<String>,
    /**
     * 集合
     * 
     */
    pub cats: Option<Vec<Cat>>,
    /**
     * 是否启用
     * 
     */
    pub enable: Option<bool>,
    /**
     * 引用map
     * 
     */
    pub gods: Option<Map<String,Dog>>,
    /**
     * ID
     * 
     */
    pub id: Option<u64>,
    /**
     * 基础map
     * 
     */
    pub kv: Option<Map<String,String>>,
    /**
     * 名称
     * 
     */
    pub name: Option<String>,
    /**
     * 
     * 
     */
    pub profile: Option<JsonNode>,
    /**
     * 标签
     * 
     */
    pub tags: Option<Vec<String>>,
    /**
     * 时间戳
     * 
     */
    pub timestamp: Option<u64>,
    
}




