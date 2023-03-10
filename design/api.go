package design

import (
	_ "ddd-goa-ent-project/design/service"
	"ddd-goa-ent-project/internal/env"

	. "goa.design/goa/v3/dsl"
	cors "goa.design/plugins/v3/cors/dsl"
)

var _ = API("ddd-goa-ent-project_box", func() {
	Title("ddd-goa-ent-project Box")
	Description("ddd-goa-ent-project Box")

	cors.Origin("*", func() {
		cors.Headers("X-Requested-With, Content-Type, Accept, Origin, Authorization, X-Api-Version")
		cors.Credentials()
		cors.Methods("GET", "POST", "OPTIONS", "PUT", "DELETE", "PATCH")
	})

	// Server describes a single process listening for client requests. The DSL
	// defines the set of services that the server hosts as well as hosts details.
	Server("ddd-goa-ent-project", func() {
		Description("ddd-goa-ent-project Box Service")

		// List the services hosted by this server.
		Services([]string{
			"healthy", "admin", "web",
		}...)

		// List the Hosts and their transport URLs.
		Host("localhost", func() {
			// Transport specific URLs, supported schemes are:
			// 'http', 'https', 'grpc' and 'grpcs' with the respective default
			// ports: 80, 443, 8080, 8443.
			URI(env.OpenAPIDomain)
			// Variable describes a URI variable.
			Variable("version", String, "API version", func() {
				// URL parameters must have a default value and/or an
				// enum validation.
				Default("v1")
			})
		})
	})
})
