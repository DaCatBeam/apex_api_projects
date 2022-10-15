package webservers

// COMPOSITION PATTERN?
type WebServerContext struct {
    configStrategy configStrategy
}

// Initialzes a new WebServerContext
func NewWebServerContext() *WebServerContext {
    return &WebServerContext{}
}

// Sets the configuration strategy member
func (context *WebServerContext) SetConfigStrategy(configStrategy configStrategy) {
    context.configStrategy = configStrategy
}
