package main


import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	ctx "context"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"github.com/gorilla/mux"
	."github.com/ali2210/mongoDBNinjas/dataModel"
)


const(
	connection_mongodb string = "mongodb://localhost:27017/"
)

var archieve CollectionArcheive = CollectionArcheive{}
var collect *mongo.Collection

func main() {
	clientops := options.Client().ApplyURI(connection_mongodb)
	client , err = mongo.Connect(ctx.TODO(), clientops)
	if err != nil {
		fmt.Println("Conenction reject ", err)
		return 
	}
	docs := client.Database("Books").Collection("MyDairy")
	fmt.Println("Database instance starting...", docs)

	WebHandler := mux.NewRouter()
	WebHandler.HandleFunc("/addpage", NewPage).Method("GET")
	WebHandler.HandleFunc("/displayPage", Display).Method("POST")
	err = http.ListenAndServe(":9000", WebHandler)
	if err != nil {
		return
	}
}

func NewPage(w http.ResponseWriter, r * http.Request){
	
	dairy := Collection{
		Name : "Dairy",
		Contents : "Hello Nigna",
	}
	page, err := archieve.AddContentInCollection(collect, dairy)
	if err != nil {
		return 
	}
	fmt.Println("Dairy contains:", page)
}
func Display(w http.ResponseWriter, r * http.Request){

}