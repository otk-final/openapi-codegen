

using System;
using System.Net.Http;
using System.Text;
using System.Text.Json;
using System.Threading.Tasks;
using System.Collections.Generic;
using System.Web;

public interface IApiClient
{
    public Task<T> Get<T>(string path, Dictionary<string, dynamic> parameters) where T : class;

    public Task<T> Post<T, P>(string path, P parameters) where T : class where P : class;

    public Task<T> Put<T, P>(string path, P parameters) where T : class where P : class;

    public Task<T> Delete<T, P>(string path, P parameters) where T : class where P : class;
}

public class DefaultApiClient : IApiClient
{
    private readonly HttpClient _httpClient;
    private readonly JsonSerializerOptions _jsonOptions;

    public DefaultApiClient(HttpClient httpClient)
    {
        _httpClient = httpClient;
        _jsonOptions = new JsonSerializerOptions
        {
            PropertyNamingPolicy = JsonNamingPolicy.CamelCase
        };
    }

    public async Task<T> Get<T>(string path, Dictionary<string, dynamic> parameters) where T : class
    {
        var uriBuilder = new UriBuilder(_httpClient.BaseAddress + path);
        var query = HttpUtility.ParseQueryString(string.Empty);

        foreach (var kvp in parameters)
        {
            query[kvp.Key] = kvp.Value.ToString();
        }

        uriBuilder.Query = query.ToString();
        var response = await _httpClient.GetAsync(uriBuilder.ToString());
        return await DeserializeResponse<T>(response);
    }

    public async Task<T> Post<T, P>(string path, P parameters) where T : class where P : class
    {
        var content = SerializeContent(parameters);
        var response = await _httpClient.PostAsync(path, content);
        return await DeserializeResponse<T>(response);
    }

    public async Task<T> Put<T, P>(string path, P parameters) where T : class where P : class
    {
        var content = SerializeContent(parameters);
        var response = await _httpClient.PutAsync(path, content);
        return await DeserializeResponse<T>(response);
    }

    public async Task<T> Delete<T, P>(string path, P parameters) where T : class where P : class
    {
        var request = new HttpRequestMessage
        {
            Method = HttpMethod.Delete,
            RequestUri = new Uri(_httpClient.BaseAddress + path),
            Content = SerializeContent(parameters)
        };
        var response = await _httpClient.SendAsync(request);
        return await DeserializeResponse<T>(response);
    }

    private HttpContent SerializeContent<T>(T obj)
    {
        var json = JsonSerializer.Serialize(obj, _jsonOptions);
        return new StringContent(json, Encoding.UTF8, "application/json");
    }

    private async Task<T> DeserializeResponse<T>(HttpResponseMessage response) where T : class
    {
        response.EnsureSuccessStatusCode();
        var json = await response.Content.ReadAsStringAsync();
        return JsonSerializer.Deserialize<T>(json, _jsonOptions);
    }
}