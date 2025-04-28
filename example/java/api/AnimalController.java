
package example.api;

import example.dto.*;
import example.ApiClient;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.JsonNode;
import java.util.List;
import java.util.Map;
import java.util.HashMap;

/**
 * 动物接口
 */
public class AnimalController {

    private final ApiClient client;

    public AnimalController(ApiClient client) {
        this.client = client;
    }

	

    public static final TypeReference<Integer> addCatResultType = new TypeReference<Integer>() {
    };

	/**
	 * 响应long
	 * 
	 */
    public Integer addCat( Cat body) {
        
        return client.post("/animal/cat/add",body, addCatResultType);
    }

	

    public static final TypeReference<Dog> addDogResultType = new TypeReference<Dog>() {
    };

	/**
	 * 路径+Body
	 * 
	 */
    public Dog addDog( String kind, Dog body) {
        
        return client.post(String.format("/animal/add/%s/dog",kind),body, addDogResultType);
    }

	

    public static final TypeReference<List<Cat>> catsResultType = new TypeReference<List<Cat>>() {
    };

	/**
	 * 响应list
	 * 
	 */
    public List<Cat> cats() {
        Map<String,Object> params = new HashMap<String,Object>();
        return client.get("/animal/cats",params, catsResultType);
    }

	

    public static final TypeReference<Dog> dogResultType = new TypeReference<Dog>() {
    };

	/**
	 * 响应对象
	 * 
	 */
    public Dog dog( Integer id) {
        Map<String,Object> params = new HashMap<String,Object>();
        return client.get(String.format("/animal/dog/%s",id),params, dogResultType);
    }

	

    public static final TypeReference<Map<String,Cat>> mapCatsResultType = new TypeReference<Map<String,Cat>>() {
    };

	/**
	 * 响应map
	 * 其他
	 */
    public Map<String,Cat> mapCats() {
        Map<String,Object> params = new HashMap<String,Object>();
        return client.get("/animal/map/cats",params, mapCatsResultType);
    }

	

    public static final TypeReference<Dog> mapDogResultType = new TypeReference<Dog>() {
    };

	/**
	 * 响应Map泛型
	 * 
	 */
    public Dog mapDog() {
        Map<String,Object> params = new HashMap<String,Object>();
        return client.get("/animal/map/dogs",params, mapDogResultType);
    }

	

    public static final TypeReference<Dog> pageDogsResultType = new TypeReference<Dog>() {
    };

	/**
	 * 响应Array泛型
	 * 
	 */
    public Dog pageDogs() {
        Map<String,Object> params = new HashMap<String,Object>();
        return client.get("/animal/page/dogs",params, pageDogsResultType);
    }

	

    public static final TypeReference<Dog> searchDogResultType = new TypeReference<Dog>() {
    };

	/**
	 * 路径+Query
	 * 
	 */
    public Dog searchDog( String color, String kind, String name) {
        Map<String,Object> params = new HashMap<String,Object>();
        params.put("color",color);
        params.put("kind",kind);
        params.put("name",name);
        
        return client.get(String.format("/animal/query/%s/dog",kind),params, searchDogResultType);
    }

	
}
