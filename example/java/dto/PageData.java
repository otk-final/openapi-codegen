
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
public class PageData<T> implements Serializable {
    /**
     * 数据
     * 
     */
    private List<T> entities;
    /**
     * 页码
     * 
     */
    private Integer page;
    /**
     * 条数
     * 
     */
    private Integer size;
    /**
     * 总数
     * 
     */
    private Long total;
    
}
