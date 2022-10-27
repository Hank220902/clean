package todo

import (
	model "clean/app/domain/model/todo"
	"clean/app/infra/presistence/mongo"
	"context"
	"fmt"
	"log"
	"time"


	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *repo) GetAll(ctx context.Context, email string) ([]*model.Get, error) {
	var results []*model.Get
	updateStatus(ctx,email)
	filter := bson.D{{Key: "email", Value: email}}
	
	cur, err := mongo.TodoCollection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(ctx) {

		var result model.Get
		err := cur.Decode(&result)

		if err != nil {
			log.Fatal(err)
		}
		zone := time.FixedZone("", +8*60*60)
		result.CreateTime = result.CreateTime.In(zone)
		result.UpdateTime = result.UpdateTime.In(zone)
		

		results = append(results, &result)

	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	return results, nil

}
func updateStatus(ctx context.Context,email string) {
	var results []*model.UpdateStatus

	filter := bson.D{{Key: "email", Value: email}}

	cur, err := mongo.TodoCollection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(ctx) {

		var result model.UpdateStatus
		err := cur.Decode(&result)

		if err != nil {
			log.Fatal(err)
		}
		zone := time.FixedZone("", +8*60*60)
		result.CreateTime = result.CreateTime.In(zone)
		id, _ := primitive.ObjectIDFromHex(result.ID)
		if result.EndTime.After(time.Now()) {
			filter := bson.D{{Key: "_id", Value: id}, {Key: "email", Value: result.Email}}
			opts := options.Update().SetUpsert(true)
			update := bson.D{
				{Key: "$set", Value: bson.D{
					{Key: "status", Value: "未逾期"},
				},
				}}
				mongo.TodoCollection.UpdateOne(ctx, filter, update, opts)

		} else {
			filter := bson.D{{Key: "_id", Value: id}, {Key: "email", Value: result.Email}}
			opts := options.Update().SetUpsert(false)
			update := bson.D{
				{Key: "$set", Value: bson.D{
					{Key: "status", Value: "逾期"},
				},
				}}
				mongo.TodoCollection.UpdateOne(ctx, filter, update, opts)

		}

		results = append(results, &result)

	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(ctx)
	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)

	// spew.Dump(results)
	// return results, "success"

}
