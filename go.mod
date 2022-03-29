module github.com/BillionNFTHomepage/backend

go 1.15

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/mux v1.8.0
	github.com/gorilla/websocket v1.4.2
	github.com/joho/godotenv v1.3.0
	github.com/rs/cors v1.7.0
	github.com/shopspring/decimal v0.0.0-20180709203117-cd690d0c9e24
	github.com/sirupsen/logrus v1.8.1
	google.golang.org/api v0.35.0
	google.golang.org/appengine v1.6.6
	github.com/brianosaurs/volume-practice/api/middlewares v0.0.0
	github.com/brianosaurs/volume-practice/api/responses v0.0.0
	github.com/brianosaurs/volume-practice/api/controllers v0.0.0
)

replace github.com/brianosaurus/volume-practice/api/middlewares v0.0.0 => ./api/middlewares
replace github.com/brianosaurus/volume-practice/api/responses v0.0.0 => ./api/responses
replace github.com/brianosaurus/volume-practice/api/controllers v0.0.0 => ./api/controllers
