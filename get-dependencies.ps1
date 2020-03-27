# powershell script to install go dependencies
# untested
go get -u   google.golang.org/grpc `
            google.golang.org/grpc/codes `
            google.golang.org/grpc/status `
	    gopkg.in/yaml.v2 `
            google.golang.org/grpc/metadata `
            google.golang.org/grpc/reflection `
            google.golang.org/grpc/examples/helloworld/helloworld `
            golang.org/x/net/context `
            golang.org/x/oauth2/google `
            github.com/golang/protobuf/proto `
	    github.com/golang/protobuf/protoc-gen-go
