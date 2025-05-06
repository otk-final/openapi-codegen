
package example.dto;

import com.fasterxml.jackson.databind.JsonNode;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.io.Serializable;
import java.util.List;
import java.util.Map;


/**
 * 用户
 * 
 */
@Data
@NoArgsConstructor
@AllArgsConstructor
public class User implements Serializable {
    /**
     * 年龄
     * 
     */
    private Integer age;
    /**
     * 生日
     * 
     */
    private String birthday;
    /**
     * 集合
     * 
     */
    private List<Cat> cats;
    /**
     * 是否启用
     * 
     */
    private Boolean enable;
    /**
     * 引用map
     * 
     */
    private Map<String,Dog> gods;
    /**
     * ID
     * 
     */
    private Long id;
    /**
     * 基础map
     * 
     */
    private Map<String,String> kv;
    /**
     * 名称
     * 
     */
    private String name;
    /**
     * 
     * 
     */
    private JsonNode profile;
    /**
     * 标签
     * 
     */
    private List<String> tags;
    /**
     * 时间戳
     * 
     */
    private Long timestamp;
    
}
