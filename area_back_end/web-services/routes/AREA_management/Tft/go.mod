module API/Tft

go 1.21.1

replace API/DB_management => ../../../DB_management

replace API/call_API => ../../call_API

require (
	API/DB_management v0.0.0-00010101000000-000000000000
	API/call_API v0.0.0-00010101000000-000000000000
)

require (
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/lib/pq v1.10.9 // indirect
)
