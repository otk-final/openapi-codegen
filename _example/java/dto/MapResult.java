
package example.dto;

import com.fasterxml.jackson.databind.JsonNode;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.io.Serializable;
import java.util.List;
import java.util.Map;


/**
 * 
 * 
 */
@Data
@NoArgsConstructor
@AllArgsConstructor
public class MapResult<T> implements Serializable {
    /**
     * 响应码
     * 
     */
    private String code;
    /**
     * 数据
     * 
     */
    private Map<String,T> data;
    
}
