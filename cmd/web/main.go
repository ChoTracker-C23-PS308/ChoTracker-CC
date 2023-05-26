package main

import (
	"fmt"
	aDelivery "github.com/ChoTracker-C23-PS308/ChoTracker-CC/internal/delivery/article/http"
	hDelivery "github.com/ChoTracker-C23-PS308/ChoTracker-CC/internal/delivery/history/http"
	uDelivery "github.com/ChoTracker-C23-PS308/ChoTracker-CC/internal/delivery/user/http"
	aRepo "github.com/ChoTracker-C23-PS308/ChoTracker-CC/internal/repository/article/pg"
	hRepo "github.com/ChoTracker-C23-PS308/ChoTracker-CC/internal/repository/history/pg"
	uRepo "github.com/ChoTracker-C23-PS308/ChoTracker-CC/internal/repository/user/pg"
	"github.com/gin-gonic/gin"
	"log"

	cfg "github.com/ChoTracker-C23-PS308/ChoTracker-CC/common/config"
	firebaseCommon "github.com/ChoTracker-C23-PS308/ChoTracker-CC/common/firebase/admin"
	firebaseAuthCommon "github.com/ChoTracker-C23-PS308/ChoTracker-CC/common/firebase/auth"
	httpCommon "github.com/ChoTracker-C23-PS308/ChoTracker-CC/common/http"
	pgCommon "github.com/ChoTracker-C23-PS308/ChoTracker-CC/common/pg"
)

func main() {

	config := cfg.LoadConfig()
	store := pgCommon.New(config.DatabaseURL)
	//store := pgCommon.New("postgres://dev_user:dev_password@db:5432/dev_database?sslmode=disable")
	defer store.Db.Close()

	app, err := firebaseCommon.NewFirebaseAdmin(config.CredentialType, config.CredentialValue)
	if err != nil {
		panic(err)
	}
	//app, err := firebaseCommon.NewFirebaseAdmin("json", "configs/var/chotracker-c23-ps308-firebase-adminsdk-akrc3-a53cb2d6ef.json")
	//if err != nil {
	//	panic(err)
	//}
	fAuth, err := firebaseAuthCommon.NewFirebaseAuth(app)
	if err != nil {
		panic(err)
	}
	//_, err = firebaseStgCommon.NewFirebaseStorage(app, config.BucketName)
	//if err != nil {
	//	panic(err)
	//}

	h := httpCommon.NewHTTPServer()
	api := h.Router.Group("/api/v1", gin.Logger(), httpCommon.CORS())

	//aur := auRepo.NewFirebaseAuthRepository(fAuth)

	ur := uRepo.NewPGUserRepository(store.Querier)
	uDelivery.NewHTTPUserDelivery(api, ur, fAuth)

	ar := aRepo.NewPGArticleRepository(store.Querier)
	aDelivery.NewHTTPArticleDelivery(api, ar, fAuth)

	hr := hRepo.NewPGHistoryRepository(store.Querier)
	hDelivery.NewHTTPHistoryDelivery(api, hr, fAuth)

	log.Fatal(h.Router.Run(fmt.Sprintf(":%d", config.Port)))
}
