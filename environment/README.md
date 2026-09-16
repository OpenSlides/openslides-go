# Environment

The environment package is used, to configure a go-service. The Package contains
functions, to access environment variables in a way, so that the configuration
for a service can automaticly be generated.

For this to work, the initalization phase is not allowed to do any IO. In other
case, the IO would also happen, then the docutmentation is build.

Therefore, a service service has to start in different phases.

The first phase ready the environment variables, but does nothing else. For
example, it could be build like this:

```go
func initService(lookup environment.Environmenter) (func(context.Context) error, error) {
	// No IO allowed in this function
	oslog.InitLog(lookup)
	someType, background := somePackage.New(lookup)

	service := func(ctx context.Context) error {
		// This is a callback. IO is allowed here.
		
		oslog.Info("Listen on %s", listenAddr)
		return http.Run(ctx, someType)
	}
	return service, nil
}
```

All functions inside initService (like `somePackage.New`) are also not allowed
to do IO. If IO is needed, it has to be done in the callback.


When this is done, the service can either be started:

```go
func run(ctx context.Context) error {
	lookup := new(environment.ForProduction)

	service, err := initService(lookup)
	if err != nil {
		return fmt.Errorf("init services: %w", err)
	}

	return service(ctx)
}
```

Or the documentation can be build:

```go
func buildDocu() error {
	lookup := new(environment.ForDocu)

	if _, err := initService(lookup); err != nil {
		return fmt.Errorf("init services: %w", err)
	}

	doc, err := lookup.BuildDoc()
	if err != nil {
		return fmt.Errorf("build doc: %w", err)
	}

	fmt.Println(doc)
	return nil
}
```

As a maxim: If the function has a `context.Context` as an argument, you can do
IO. If the function has not a `context.Context` as an argument, do not create
one with `context.Background()`. You have to do the IO elsewhere.
