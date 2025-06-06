

package com.demo;

import com.google.gson.Gson
import okhttp3.MediaType
import okhttp3.RequestBody
import okhttp3.ResponseBody
import retrofit2.Response
import retrofit2.Retrofit
import retrofit2.http.Body
import retrofit2.http.DELETE
import retrofit2.http.GET
import retrofit2.http.POST
import retrofit2.http.PUT
import retrofit2.http.QueryMap
import retrofit2.http.Url
import java.lang.reflect.Type

/**
 * TODO You need to implement the current interface, which can refer to Retorfit or OkHttp
 */
class ApiClient {

    suspend fun <V : Any> get(path: String, retType: Type, params: Map<String, Any>? = emptyMap()): Result<V> {
        return toResult(defaultClient.get(path, params), retType)
    }

    suspend fun <P : Any, V : Any> post(path: String, retType: Type, params: P?): Result<V> {
        return toResult(defaultClient.post(path, toBody(params)), retType)
    }

    suspend fun <P : Any, V : Any> put(path: String, retType: Type, params: P?): Result<V> {
        return toResult(defaultClient.put(path, toBody(params)), retType)
    }

    suspend fun <P : Any, V : Any> delete(path: String, retType: Type, params: P?): Result<V> {
        return toResult(defaultClient.delete(path, toBody(params)), retType)
    }

    private fun <V : Any> toResult(resp: Response<ResponseBody>, retType: Type): Result<V> {
        return if (resp.isSuccessful) {
            Result.success(defaultGson.fromJson(resp.body()!!.charStream(), retType))
        } else {
            Result.failure(Exception("Request failed with status code: ${resp.code()}"))
        }
    }

    private fun <P : Any> toBody(params: P?): RequestBody? {
        return RequestBody.create(MediaType.parse("application/json"), defaultGson.toJson(params))
    }
}


var defaultGson = Gson()
private var defaultClient =
    Retrofit.Builder().baseUrl("http://localhost:8083").build().create(RetrofitClient::class.java)

interface RetrofitClient {

    @GET
    suspend fun get(@Url path: String, @QueryMap(encoded = true) params: Map<String, @JvmSuppressWildcards Any>?): Response<ResponseBody>

    @POST
    suspend fun post(@Url path: String, @Body params: RequestBody?): Response<ResponseBody>

    @PUT
    suspend fun put(@Url path: String, @Body params: RequestBody?): Response<ResponseBody>

    @DELETE
    suspend fun delete(@Url path: String, @Body params: RequestBody?): Response<ResponseBody>
}

