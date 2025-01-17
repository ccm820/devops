class PlatformClient {
    String baseUrl
    String apiKey

    PlatformClient(String baseUrl, String apiKey) {
        this.baseUrl = baseUrl
        this.apiKey = apiKey
    }

    def callApi(String endpoint, String method, Map requestBody = [:]) {
        try {
            def response = httpRequest(
                url: "${baseUrl}${endpoint}",
                httpMode: method,
                contentType: 'APPLICATION_JSON',
                requestBody: requestBody ? groovy.json.JsonOutput.toJson(requestBody) : null,
                customHeaders: [[name: 'Authorization', value: "Bearer ${apiKey}"]]
            )
            return response.content
        } catch (Exception e) {
            error "Failed to call API: ${e.message}"
        }
    }
}
