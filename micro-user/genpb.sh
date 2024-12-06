cd dto
protoc -I ~/go/pkg/mod/github.com/srikrsna/protoc-gen-gotag@v0.6.2 -I . --go_out=. --go-grpc_out=. user.proto
protoc -I ~/go/pkg/mod/github.com/srikrsna/protoc-gen-gotag@v0.6.2 -I . --plugin=protoc-gen-gotag=/home/msipc/go/bin/protoc-gen-gotag --gotag_out=. user.proto
