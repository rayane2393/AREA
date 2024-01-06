module API/Discord

go 1.21.1

replace API/call_API => ../../call_API

replace API/DB_management => ../../../DB_management

require API/call_API v0.0.0-00010101000000-000000000000

require (
	API/DB_management v0.0.0-00010101000000-000000000000 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/lib/pq v1.10.9 // indirect
)
