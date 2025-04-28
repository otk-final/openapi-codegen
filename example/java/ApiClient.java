

package example;

import com.fasterxml.jackson.core.type.TypeReference;

import java.util.Map;

public interface ApiClient  {

    <T> T get(String path, Map<String, Object> params, TypeReference<T> resultType);

    <P, T> T put(String path, P body, TypeReference<T> resultType);

    <P, T> T post(String path, P params, TypeReference<T> resultType);

    <P, T> T delete(String path, P params, TypeReference<T> resultType);
}
