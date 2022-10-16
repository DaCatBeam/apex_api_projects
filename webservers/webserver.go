package webservers

// COMPOSITION PATTERN?
type WebServerContext struct {
    // Factory concerned with creating Config objects from files
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
