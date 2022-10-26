package todo

// import (
// 	model "clean/app/domain/model/todo"
// 	"clean/app/infra/presistence/mongo"
// 	"context"
// 	"log"
// 	"time"

// 	"go.mongodb.org/mongo-driver/bson"
// )

// func (r *repo) GetAll(ctx context.Context, data *model.Todo) (*model.Todo, error) {
// 	var results []*model.HaveIDTodo
// 	filter := bson.D{{Key: "email", Value: data.Email}}

// 	var blog *model.Todo
// 	cur, err := mongo.TodoCollection.Find(ctx, filter)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	for cur.Next(ctx) {

// 		var result model.HaveIDTodo
// 		err := cur.Decode(&result)

// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		zone := time.FixedZone("", +8*60*60)
// 		result.CreateTime = result.CreateTime.In(zone)

// 		results = append(results, &result)

// 	}

// 	if err := cur.Err(); err != nil {
// 		log.Fatal(err)
// 	}
// 	cur.Close(ctx)

// 	return blog, nil
// 	// fmt.Println(r)
// 	// _, err := mongo.TodoCollection.InsertOne(ctx, data)
// 	// fmt.Println("qweqeqe")
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// return nil
// }
