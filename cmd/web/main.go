package main

import (
	"fmt"
	"log"

	uDelivery "github.com/ChoTracker-C23-PS308/ChoTracker-CC/internal/delivery/user/http"
	uRepo "github.com/ChoTracker-C23-PS308/ChoTracker-CC/internal/repository/user/pg"
	"github.com/gin-gonic/gin"

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

	//uuid := uuid.NewRandom()

	h := httpCommon.NewHTTPServer()
	api := h.Router.Group("/api/v1", gin.Logger(), httpCommon.CORS())

	//aur := auRepo.NewFirebaseAuthRepository(fAuth)

	ur := uRepo.NewPGUserRepository(store.Querier)
	uDelivery.NewHTTPUserDelivery(api, ur, fAuth)

	log.Fatal(h.Router.Run(fmt.Sprintf(":%d", 4001)))

}
