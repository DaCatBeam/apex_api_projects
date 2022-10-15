package webservers

// COMPOSITION PATTERN?
type WebServerContext struct {
    ConfigObjectFactory ConfigObjectFactory
    databaseStrategy databaseStrategy
}

// Initialzes a new WebServerContext
func NewWebServerContext() *WebServerContext {
    wsc := &WebServerContext{}
    wsc.ConfigObjectFactory = ConfigObjectFactory{}

    return wsc
}

// Sets the configuration strategy member
func (context *WebServerContext) SetDatabaseStrategy(strat databaseStrategy) {
    context.databaseStrategy = strat
}
