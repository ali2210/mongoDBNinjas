package dataModel



import (
	 "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	ctx "context"
	"fmt"
)

type Collection struct {
	Name string `json:"name, omitempty" bson:"name, omitemty"`
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Contents string `json:"content,omitempty" bson:"content,omitempty"`
}

// MONGODB FUNC
type Archieve interface {
	AddContentInCollection(*mongo.Collection, *Collection)(*mongo.InsertOneResult, error)
	// DeleteContentFromCollection()
	// DisplayContentFromCollection()
	// FindCollectionFromCollection()
}

type CollectionArcheive struct {}

func NewClient() Archieve{
	return &CollectionArcheive{}
}

func (*CollectionArcheive) AddContentInCollection(recordDB *mongo.Collection, record *Collection)(*mongo.InsertOneResult, error){
	docs , err := recordDB.InsertOne(ctx.TODO(), record)
	if err != nil {
		fmt.Println("insertion fail:", err)
		return docs, err
	}
	return docs, nil
}
// func (*CollectionArcheive) DeleteContentFromCollection(){}
// func (*CollectionArcheive) FindCollectionFromCollection(){}
// func (*CollectionArcheive) DisplayContentFromCollection(){}

