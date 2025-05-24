

use std::collections::HashMap;
use std::io::Error;
use std::iter::Map;
use crate::Client::ApiClient;
use crate::Model::*;




/**
 * 动物接口
 */
pub struct AnimalController {
    client: dyn ApiClient,
}

impl AnimalController {

	

	/**
	 * 响应long
	 * 
	 */
    async fn addCat(&self, body: Cat) -> Result<usize,Error>  {
        
        self.client.post(String::from("/animal/cat/add"),body)
    }

	

	/**
	 * 路径+Body
	 * 
	 */
    async fn addDog(&self, kind: String,body: Dog) -> Result<Dog,Error>  {
        
        self.client.post(String::from(format!("/animal/add/${kind}/dog",kind)),body)
    }

	

	/**
	 * 响应list
	 * 
	 */
    async fn cats(&self, ) -> Result<Vec<Cat>,Error>  {
        let mut params = HashMap::new();
        self.client.get(String::from("/animal/cats"),params)
    }

	

	/**
	 * 响应对象
	 * 
	 */
    async fn dog(&self, id: usize) -> Result<Dog,Error>  {
        let mut params = HashMap::new();
        self.client.get(String::from(format!("/animal/dog/${id}",id)),params)
    }

	

	/**
	 * 响应map
	 * 其他
	 */
    async fn mapCats(&self, ) -> Result<Map<String,Cat>,Error>  {
        let mut params = HashMap::new();
        self.client.get(String::from("/animal/map/cats"),params)
    }

	

	/**
	 * 响应Map泛型
	 * 
	 */
    async fn mapDog(&self, ) -> Result<MapResultDog,Error>  {
        let mut params = HashMap::new();
        self.client.get(String::from("/animal/map/dogs"),params)
    }

	

	/**
	 * 响应Array泛型
	 * 
	 */
    async fn pageDogs(&self, ) -> Result<PageDataDog,Error>  {
        let mut params = HashMap::new();
        self.client.get(String::from("/animal/page/dogs"),params)
    }

	

	/**
	 * 路径+Query
	 * 
	 */
    async fn searchDog(&self, color: String,kind: String,name: String) -> Result<Dog,Error>  {
        let mut params = HashMap::new();
        params.insert("color",color.to_string());
        params.insert("kind",kind.to_string());
        params.insert("name",name.to_string());
        
        self.client.get(String::from(format!("/animal/query/${kind}/dog",kind)),params)
    }

	
}



/**
 * 其他接口
 */
pub struct OtherApi {
    client: dyn ApiClient,
}

impl OtherApi {

	

	/**
	 * 响应基础类型
	 * 
	 */
    async fn get(&self, ) -> Result<String,Error>  {
        let mut params = HashMap::new();
        self.client.get(String::from("/other/long"),params)
    }

	

	/**
	 * 响应基础类型
	 * 
	 */
    async fn getContent(&self, ) -> Result<usize,Error>  {
        let mut params = HashMap::new();
        self.client.get(String::from("/other/string"),params)
    }

	

	/**
	 * 多路径参数
	 * 
	 */
    async fn multiplePathVariable(&self, id: String,type: String) -> Result<Vec<Dog>,Error>  {
        let mut params = HashMap::new();
        self.client.get(String::from(format!("/other/boolean/${type}/${id}",type,id)),params)
    }

	
}



/**
 * 用户接口
 */
pub struct UserApi {
    client: dyn ApiClient,
}

impl UserApi {

	

	/**
	 * 添加用户
	 * 
	 */
    async fn add(&self, body: User) -> Result<ApiResultLong,Error>  {
        
        self.client.post(String::from("/user/add"),body)
    }

	

	/**
	 * 删除用户
	 * 
	 */
    async fn delete(&self, id: u64) -> Result<ApiResultBoolean,Error>  {
        let mut params = HashMap::new();
        self.client.delete(String::from(format!("/user/delete/${id}",id)),params)
    }

	

	/**
	 * 修改用户
	 * 
	 */
    async fn edit(&self, id: u64,body: User) -> Result<ApiResultLong,Error>  {
        
        self.client.put(String::from(format!("/user/edit/${id}",id)),body)
    }

	

	/**
	 * 查询用户
	 * 
	 */
    async fn edit_1(&self, id: u64,keyword: String) -> Result<ApiResultUser,Error>  {
        let mut params = HashMap::new();
        params.insert("id",id.to_string());
        params.insert("keyword",keyword.to_string());
        
        self.client.get(String::from(format!("/user/detail/${id}",id)),params)
    }

	

	/**
	 * 用户列表
	 * 
	 */
    async fn list(&self, ) -> Result<ApiResultListUser,Error>  {
        let mut params = HashMap::new();
        self.client.get(String::from("/user/list"),params)
    }

	

	/**
	 * 用户分页
	 * 
	 */
    async fn page(&self, page: usize,size: usize) -> Result<ApiResultPageDataUser,Error>  {
        let mut params = HashMap::new();
        params.insert("page",page.to_string());
        params.insert("size",size.to_string());
        
        self.client.get(String::from("/user/page"),params)
    }

	
}











