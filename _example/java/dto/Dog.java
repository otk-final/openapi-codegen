
package example.dto;

import com.fasterxml.jackson.databind.JsonNode;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.io.Serializable;
import java.util.List;
import java.util.Map;


/**
 * 狗
 * 
 */
@Data
@NoArgsConstructor
@AllArgsConstructor
public class Dog implements Serializable {
    /**
     * 颜色
     * 
     */
    private String color;
    /**
     * ID
     * 
     */
    private Long id;
    /**
     * 昵称
     * 
     */
    private String nickname;
    /**
     * 售价
     * 
     */
    private Integer price;
    
}
