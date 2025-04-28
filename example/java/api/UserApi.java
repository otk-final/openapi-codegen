
package example.api;

import example.dto.*;
import example.ApiClient;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.JsonNode;
import java.util.List;
import java.util.Map;
import java.util.HashMap;

/**
 * 用户接口
 */
public class UserApi {

    private final ApiClient client;

    public UserApi(ApiClient client) {
        this.client = client;
    }

	

    public static final TypeReference<Long> addResultType = new TypeReference<Long>() {
    };

	/**
	 * 添加用户
	 * 
	 */
    public Long add( User body) {
        
        return client.post("/user/add",body, addResultType);
    }

	

    public static final TypeReference<Boolean> deleteResultType = new TypeReference<Boolean>() {
    };

	/**
	 * 删除用户
	 * 
	 */
    public Boolean delete( Long id) {
        Map<String,Object> params = new HashMap<String,Object>();
        return client.delete(String.format("/user/delete/%s",id),params, deleteResultType);
    }

	

    public static final TypeReference<Long> editResultType = new TypeReference<Long>() {
    };

	/**
	 * 修改用户
	 * 
	 */
    public Long edit( Long id, User body) {
        
        return client.put(String.format("/user/edit/%s",id),body, editResultType);
    }

	

    public static final TypeReference<User> edit_1ResultType = new TypeReference<User>() {
    };

	/**
	 * 查询用户
	 * 
	 */
    public User edit_1( Long id, String keyword) {
        Map<String,Object> params = new HashMap<String,Object>();
        params.put("id",id);
        params.put("keyword",keyword);
        
        return client.get(String.format("/user/detail/%s",id),params, edit_1ResultType);
    }

	

    public static final TypeReference<List<User>> listResultType = new TypeReference<List<User>>() {
    };

	/**
	 * 用户列表
	 * 
	 */
    public List<User> list() {
        Map<String,Object> params = new HashMap<String,Object>();
        return client.get("/user/list",params, listResultType);
    }

	

    public static final TypeReference<PageData<User>> pageResultType = new TypeReference<PageData<User>>() {
    };

	/**
	 * 用户分页
	 * 
	 */
    public PageData<User> page( Integer page, Integer size) {
        Map<String,Object> params = new HashMap<String,Object>();
        params.put("page",page);
        params.put("size",size);
        
        return client.get("/user/page",params, pageResultType);
    }

	
}
