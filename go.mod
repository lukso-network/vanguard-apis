module github.com/prysmaticlabs/ethereumapis

go 1.13

require (
	github.com/ferranbt/fastssz v0.0.0-20210120143747-11b9eff30ea9
	github.com/gogo/protobuf v1.3.1
	github.com/golang/protobuf v1.4.2
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/prysmaticlabs/eth2-types v0.0.0-20210127031309-22cbe426eba6
	github.com/prysmaticlabs/go-bitfield v0.0.0-20210202205921-7fcea7c45dc8
	google.golang.org/genproto v0.0.0-20200528191852-705c0b31589b
	google.golang.org/grpc v1.33.1
	google.golang.org/protobuf v1.24.0
)

replace (
	github.com/ferranbt/fastssz => github.com/atif-konasl/fastssz v0.0.0-20210825113045-242c536f4f90
	github.com/prysmaticlabs/ethereumapis => ./
)
