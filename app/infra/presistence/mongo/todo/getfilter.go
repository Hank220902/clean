package todo

import (
	model "clean/app/domain/model/todo"
	"clean/app/infra/presistence/mongo"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *repo) GetFilter(ctx context.Context, data *model.GetFilter) ([]*model.Get, error) {
	var results []*model.Get
	updateStatus(ctx, data.Email)
	filter := bson.D{{Key: "email", Value: data.Email}}
	if data.Status != "" && data.FinishedCondition != ""{
		filter =bson.D{{Key: "email", Value: data.Email},{Key: "status", Value: data.Status},{Key: "finishedCondition", Value: data.FinishedCondition}}
	}else if data.Status != ""{
		filter =bson.D{{Key: "email", Value: data.Email},{Key: "status", Value: data.Status}}
	}else if data.FinishedCondition != ""{
		filter =bson.D{{Key: "email", Value: data.Email},{Key: "finishedCondition", Value: data.FinishedCondition}}
	}

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
		result.EndTime = result.EndTime.In(zone)

		results = append(results, &result)

	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	return results, nil

}


