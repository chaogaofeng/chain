module github.com/glodnet/chain

go 1.16

require (
	github.com/CosmWasm/wasmd v0.24.0
	github.com/cenkalti/backoff v2.2.1+incompatible
	github.com/cosmos/cosmos-sdk v0.45.4
	github.com/cosmos/ibc-go/v3 v3.0.0
	github.com/ethereum/go-ethereum v1.10.16
	github.com/fatih/color v1.13.0
	github.com/ghodss/yaml v1.0.0
	github.com/goccy/go-yaml v1.9.4
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.5.2
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/imdario/mergo v0.3.12
	github.com/otiai10/copy v1.6.0
	github.com/pelletier/go-toml v1.9.4
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.12.1
	github.com/rogpeppe/go-internal v1.6.2 // indirect
	github.com/rs/cors v1.8.2
	github.com/spf13/cast v1.4.1
	github.com/spf13/cobra v1.4.0
	github.com/spf13/viper v1.11.0
	github.com/stretchr/testify v1.7.1
	github.com/syndtr/goleveldb v1.0.1-0.20210819022825-2ae1ddf74ef7
	github.com/tendermint/starport v0.19.4
	github.com/tendermint/tendermint v0.34.19
	github.com/tendermint/tm-db v0.6.7
	github.com/tharsis/ethermint v0.14.0
	github.com/tjfoc/gmsm v1.4.0
	google.golang.org/genproto v0.0.0-20220414192740-2d67ff6cf2b4
	google.golang.org/grpc v1.45.0
	gopkg.in/yaml.v2 v2.4.0
)

replace (
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	github.com/keybase/go-keychain => github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4
	google.golang.org/grpc => google.golang.org/grpc v1.33.2
)

replace (
	github.com/CosmWasm/wasmd => ../wasmd
	github.com/cosmos/cosmos-sdk => ../cosmos-sdk
	github.com/tendermint/tendermint => ../tendermint
	//github.com/CosmWasm/wasmd => github.com/chaogaofeng/wasmd gnchain-20220505
	//github.com/cosmos/cosmos-sdk => github.com/chaogaofeng/cosmos-sdk gnchain-20220505
	//github.com/tendermint/tendermint => github.com/chaogaofeng/tendermint gnchain-20220505
)
