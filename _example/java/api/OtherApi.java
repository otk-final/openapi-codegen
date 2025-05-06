
package example.api;

import example.dto.*;
import example.ApiClient;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.JsonNode;
import java.util.List;
import java.util.Map;
import java.util.HashMap;

/**
 * 其他接口
 */
public class OtherApi {

    private final ApiClient client;

    public OtherApi(ApiClient client) {
        this.client = client;
    }

	

    public static final TypeReference<String> getResultType = new TypeReference<String>() {
    };

	/**
	 * 响应基础类型
	 * 
	 */
    public String get() {
        Map<String,Object> params = new HashMap<String,Object>();
        return client.get("/other/long",params, getResultType);
    }

	

    public static final TypeReference<Integer> getContentResultType = new TypeReference<Integer>() {
    };

	/**
	 * 响应基础类型
	 * 
	 */
    public Integer getContent() {
        Map<String,Object> params = new HashMap<String,Object>();
        return client.get("/other/string",params, getContentResultType);
    }

	

    public static final TypeReference<List<Dog>> multiplePathVariableResultType = new TypeReference<List<Dog>>() {
    };

	/**
	 * 多路径参数
	 * 
	 */
    public List<Dog> multiplePathVariable( String id, String type) {
        Map<String,Object> params = new HashMap<String,Object>();
        return client.get(String.format("/other/boolean/%s/%s",type,id),params, multiplePathVariableResultType);
    }

	
}
