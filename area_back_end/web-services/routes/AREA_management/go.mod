module API/AREA_management

go 1.21.1

replace API/call_API => ../call_API

replace API/API_key => ../../API_key

replace API/usefull_functions => ../../usefull_functions

replace API/Spotify => ./Spotify

replace API/Email => ./Email

replace API/Lol => ./Lol

replace API/Meteo => ./Meteo

replace API/Github => ./Github

replace API/Tft => ./Tft

replace API/Time => ./Time

replace API/Steam => ./Steam

replace API/Discord => ./Discord

replace API/DB_management => ../../DB_management

require (
	API/API_key v0.0.0-00010101000000-000000000000
	API/DB_management v0.0.0-00010101000000-000000000000
	API/Discord v0.0.0-00010101000000-000000000000
	API/Email v0.0.0-00010101000000-000000000000
	API/Github v0.0.0-00010101000000-000000000000
	API/Lol v0.0.0-00010101000000-000000000000
	API/Meteo v0.0.0-00010101000000-000000000000
	API/Spotify v0.0.0-00010101000000-000000000000
	API/Steam v0.0.0-00010101000000-000000000000
	API/Tft v0.0.0-00010101000000-000000000000
	API/Time v0.0.0-00010101000000-000000000000
	API/call_API v0.0.0-00010101000000-000000000000
	API/usefull_functions v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.9.1
)

require (
	github.com/bytedance/sonic v1.9.1 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20221115062448-fe3a3abad311 // indirect
	github.com/gabriel-vasile/mimetype v1.4.2 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.14.0 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/cpuid/v2 v2.2.4 // indirect
	github.com/leodido/go-urn v1.2.4 // indirect
	github.com/lib/pq v1.10.9 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.0.8 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.11 // indirect
	golang.org/x/arch v0.3.0 // indirect
	golang.org/x/crypto v0.9.0 // indirect
	golang.org/x/net v0.10.0 // indirect
	golang.org/x/sys v0.8.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
