




/**
 * 动物接口
 */
public class AnimalController
{

    private IApiClient client;

    public AnimalController(IApiClient client) {
        this.client = client;
    }

	
	/**
	 * 响应long
	 * 
	 */
    public Task<int> addCat( Cat body) {
        
        return client.Post<,>("/animal/cat/add",body);
    }

	
	/**
	 * 路径+Body
	 * 
	 */
    public Task<Dog> addDog( string kind, Dog body) {
        
        return client.Post<,>(string.Format("/animal/add/{0}/dog",kind),body);
    }

	
	/**
	 * 响应list
	 * 
	 */
    public Task<List<Cat>> cats() {
        var param = new Dictionary<string,dynamic>();
        return client.Get<List<Cat>>("/animal/cats",param);
    }

	
	/**
	 * 响应对象
	 * 
	 */
    public Task<Dog> dog( int id) {
        var param = new Dictionary<string,dynamic>();
        return client.Get<Dog>(string.Format("/animal/dog/{0}",id),param);
    }

	
	/**
	 * 响应map
	 * 其他
	 */
    public Task<Dictionary<string,Cat>> mapCats() {
        var param = new Dictionary<string,dynamic>();
        return client.Get<Dictionary<string,Cat>>("/animal/map/cats",param);
    }

	
	/**
	 * 响应Map泛型
	 * 
	 */
    public Task<MapResultDog> mapDog() {
        var param = new Dictionary<string,dynamic>();
        return client.Get<MapResultDog>("/animal/map/dogs",param);
    }

	
	/**
	 * 响应Array泛型
	 * 
	 */
    public Task<PageDataDog> pageDogs() {
        var param = new Dictionary<string,dynamic>();
        return client.Get<PageDataDog>("/animal/page/dogs",param);
    }

	
	/**
	 * 路径+Query
	 * 
	 */
    public Task<Dog> searchDog( string color, string kind, string name) {
        var param = new Dictionary<string,dynamic>();
        param["color"] = color;
        param["kind"] = kind;
        param["name"] = name;
        
        return client.Get<Dog>(string.Format("/animal/query/{0}/dog",kind),param);
    }

	
}



/**
 * 其他接口
 */
public class OtherApi
{

    private IApiClient client;

    public OtherApi(IApiClient client) {
        this.client = client;
    }

	
	/**
	 * 响应基础类型
	 * 
	 */
    public Task<string> get() {
        var param = new Dictionary<string,dynamic>();
        return client.Get<string>("/other/long",param);
    }

	
	/**
	 * 响应基础类型
	 * 
	 */
    public Task<int> getContent() {
        var param = new Dictionary<string,dynamic>();
        return client.Get<int>("/other/string",param);
    }

	
	/**
	 * 多路径参数
	 * 
	 */
    public Task<List<Dog>> multiplePathVariable( string id, string type) {
        var param = new Dictionary<string,dynamic>();
        return client.Get<List<Dog>>(string.Format("/other/boolean/{0}/{1}",type,id),param);
    }

	
}



/**
 * 用户接口
 */
public class UserApi
{

    private IApiClient client;

    public UserApi(IApiClient client) {
        this.client = client;
    }

	
	/**
	 * 添加用户
	 * 
	 */
    public Task<ApiResultLong> add( User body) {
        
        return client.Post<,>("/user/add",body);
    }

	
	/**
	 * 删除用户
	 * 
	 */
    public Task<ApiResultBoolean> delete( long id) {
        var param = new Dictionary<string,dynamic>();
        return client.Delete<,>(string.Format("/user/delete/{0}",id),param);
    }

	
	/**
	 * 修改用户
	 * 
	 */
    public Task<ApiResultLong> edit( long id, User body) {
        
        return client.Put<,>(string.Format("/user/edit/{0}",id),body);
    }

	
	/**
	 * 查询用户
	 * 
	 */
    public Task<ApiResultUser> edit_1( long id, string keyword) {
        var param = new Dictionary<string,dynamic>();
        param["id"] = id;
        param["keyword"] = keyword;
        
        return client.Get<ApiResultUser>(string.Format("/user/detail/{0}",id),param);
    }

	
	/**
	 * 用户列表
	 * 
	 */
    public Task<ApiResultListUser> list() {
        var param = new Dictionary<string,dynamic>();
        return client.Get<ApiResultListUser>("/user/list",param);
    }

	
	/**
	 * 用户分页
	 * 
	 */
    public Task<ApiResultPageDataUser> page( int page, int size) {
        var param = new Dictionary<string,dynamic>();
        param["page"] = page;
        param["size"] = size;
        
        return client.Get<ApiResultPageDataUser>("/user/page",param);
    }

	
}











