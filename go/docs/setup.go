package docs

func Setup() {
	// programmatically set swagger info
	SwaggerInfo.Title = "<% .ProjectName %> API"
	SwaggerInfo.Description = "This is the cleanarc <% .ProjectName %> API."
	SwaggerInfo.Version = "1.0"
	SwaggerInfo.Host = "localhost:8080"
	SwaggerInfo.BasePath = "/api/v1"
	SwaggerInfo.Schemes = []string{"http", "https"}
}
