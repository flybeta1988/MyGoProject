module xnw.com/routes

go 1.16

replace xnw.com/utils => ../utils

replace xnw.com/api => ../api

require (
	xnw.com/api v0.0.0-00010101000000-000000000000
	xnw.com/core v0.0.0-00010101000000-000000000000
)

replace xnw.com/core => ../core

replace xnw.com/models => ../models
