module microservices/post-client

go 1.14

// For dev purposes.
// Point to our local depository instead of pulling from a remote repo
// so that we can save time running locally.
replace github.com/vnqx/microservices/post-service => ../post-service

require (
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/vnqx/microservices/post-service v0.0.0-00010101000000-000000000000 // indirect
	golang.org/x/net v0.0.0-20200707034311-ab3426394381 // indirect
	golang.org/x/sys v0.0.0-20200803210538-64077c9b5642 // indirect
	golang.org/x/text v0.3.3 // indirect
	google.golang.org/genproto v0.0.0-20200731012542-8145dea6a485 // indirect
	google.golang.org/grpc v1.31.0
	google.golang.org/protobuf v1.25.0 // indirect
)
