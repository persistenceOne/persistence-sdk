module github.com/persistenceOne/persistenceSDK

go 1.16

require (
	github.com/99designs/keyring v1.1.6
	github.com/CosmWasm/wasmd v0.17.0
	github.com/Shopify/sarama v1.19.0
	github.com/asaskevich/govalidator v0.0.0-20210307081110-f21760c49a8d
	github.com/bartekn/go-bip39 v0.0.0-20171116152956-a05967ea095d
	github.com/cosmos/cosmos-sdk v0.42.5
	github.com/gogo/protobuf v1.3.3
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/pkg/errors v0.9.1
	github.com/regen-network/cosmos-proto v0.3.1
	github.com/spf13/cobra v1.1.3
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/crypto v0.0.0-20191022145703-50d29ede1e15
	github.com/tendermint/tendermint v0.34.10
	github.com/tendermint/tm-db v0.6.4
	honnef.co/go/tools v0.0.1-2020.1.6
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
replace github.com/cosmos/cosmos-sdk => ../cosmos-sdk
