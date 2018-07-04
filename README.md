Short example to demo middlewares in grpc. Demo includes both server side middleware and client side middleware.

To run the demo
----------------

    $ mkdir $GOPATH/src/github.com/sagarrakshe
    $ cd $GOPATH/src/github.com/sagarrakshe
    $ git clone https://github.com/sagarrakshe/grpc-middleware-example
    $ dep ensure (assuming you have `dep`)

Run server

    $ cd server
    $ go run main.go

Run client

    $ cd client
    $ go run main.go

For client side middleware

    $ cd interceptor_client
    $ go run main.go
