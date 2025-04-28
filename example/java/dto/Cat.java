
package example.dto;

import com.fasterxml.jackson.databind.JsonNode;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.io.Serializable;
import java.util.List;
import java.util.Map;


/**
 * 猫
 * 
 */
@Data
@NoArgsConstructor
@AllArgsConstructor
public class Cat implements Serializable {
    /**
     * 食物
     * 
     */
    private List<String> foods;
    /**
     * 昵称
     * 
     */
    private String nickname;
    
}
