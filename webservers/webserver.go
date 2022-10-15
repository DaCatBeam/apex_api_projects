package webservers

type WebServerContext struct {
    configStrategy configStrategy
}

func NewWebServerContext() *WebServerContext {
    return &WebServerContext{}
}

func (context *WebServerContext) SetConfigStrategy(configStrategy configStrategy) {
    context.configStrategy = configStrategy
}
