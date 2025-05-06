
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
public class ApiResult<T> implements Serializable {
    /**
     * 业务响应码
     * 
     */
    private String code;
    /**
     * 数据
     * 
     */
    private T data;
    /**
     * 业务错误信息
     * 
     */
    private String err;
    /**
     * 业务响应信息
     * 
     */
    private String msg;
    
}
