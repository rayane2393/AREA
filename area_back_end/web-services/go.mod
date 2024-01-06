module API/web-services

go 1.21.1

replace API/routes => ./routes

replace API/login => ./routes/login

replace API/user_management_admin => ./routes/user_management_admin

replace API/call_API => ./routes/call_API

replace API/API_key => ./API_key

replace API/DB_management => ./DB_management

replace API/AREA_management => ./routes/AREA_management

replace API/usefull_functions => ./usefull_functions

replace API/Spotify => ./routes/AREA_management/Spotify

replace API/Email => ./routes/AREA_management/Email

replace API/Lol => ./routes/AREA_management/Lol

replace API/Meteo => ./routes/AREA_management/Meteo

replace API/Github => ./routes/AREA_management/Github

replace API/Tft => ./routes/AREA_management/Tft

replace API/Steam => ./routes/AREA_management/Steam

replace API/about => ./routes/about

replace API/Time => ./routes/AREA_management/Time

replace API/Discord => ./routes/AREA_management/Discord

replace API/OAuth => ./OAuth

require (
	API/AREA_management v0.0.0-00010101000000-000000000000
	API/DB_management v0.0.0-00010101000000-000000000000
	API/routes v0.0.0-00010101000000-000000000000
	github.com/swaggo/swag v1.16.2
)

require (
	API/API_key v0.0.0-00010101000000-000000000000 // indirect
	API/Discord v0.0.0-00010101000000-000000000000 // indirect
	API/Email v0.0.0-00010101000000-000000000000 // indirect
	API/Github v0.0.0-00010101000000-000000000000 // indirect
	API/Lol v0.0.0-00010101000000-000000000000 // indirect
	API/Meteo v0.0.0-00010101000000-000000000000 // indirect
	API/OAuth v0.0.0-00010101000000-000000000000 // indirect
	API/Spotify v0.0.0-00010101000000-000000000000 // indirect
	API/Steam v0.0.0-00010101000000-000000000000 // indirect
	API/Tft v0.0.0-00010101000000-000000000000 // indirect
	API/Time v0.0.0-00010101000000-000000000000 // indirect
	API/about v0.0.0-00010101000000-000000000000 // indirect
	API/call_API v0.0.0-00010101000000-000000000000 // indirect
	API/login v0.0.0-00010101000000-000000000000 // indirect
	API/usefull_functions v0.0.0-00010101000000-000000000000 // indirect
	API/user_management_admin v0.0.0-00010101000000-000000000000 // indirect
	github.com/KyleBanks/depth v1.2.1 // indirect
	github.com/bytedance/sonic v1.10.2 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20230717121745-296ad89f973d // indirect
	github.com/chenzhuoyu/iasm v0.9.1 // indirect
	github.com/gabriel-vasile/mimetype v1.4.3 // indirect
	github.com/gin-contrib/cors v1.4.0 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/gin-gonic/gin v1.9.1 // indirect
	github.com/go-openapi/jsonpointer v0.20.0 // indirect
	github.com/go-openapi/jsonreference v0.20.2 // indirect
	github.com/go-openapi/spec v0.20.9 // indirect
	github.com/go-openapi/swag v0.22.4 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.15.5 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/cpuid/v2 v2.2.5 // indirect
	github.com/leodido/go-urn v1.2.4 // indirect
	github.com/lib/pq v1.10.9 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.1.0 // indirect
	github.com/rogpeppe/go-internal v1.9.0 // indirect
	github.com/swaggo/files v1.0.1 // indirect
	github.com/swaggo/gin-swagger v1.6.0 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.11 // indirect
	golang.org/x/arch v0.5.0 // indirect
	golang.org/x/crypto v0.14.0 // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	golang.org/x/tools v0.14.0 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
