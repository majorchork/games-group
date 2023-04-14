package mongodb

import (
	"context"
	"fmt"
	"github.com/majorchork/tech-crib-africa/config"
	"github.com/majorchork/tech-crib-africa/internal/database"
	"github.com/majorchork/tech-crib-africa/internal/models"
	"github.com/majorchork/tech-crib-africa/internal/services/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"math/rand"
	"time"
)

const (
	timeout         = 10
	databaseName    = "mongodb"
	userCollection  = "users"
	guestCollection = "guests"
)

type Adapter struct {
	userCol  *mongo.Collection
	guestCol *mongo.Collection
}

func NewMongoDatabaseAdapter(config config.Config) (*Adapter, error) {

	db, err := database.NewDriver(database.Config{
		URI:     config.MongoURI,
		Timeout: timeout,
	})
	if err != nil {
		return nil, err
	}

	return &Adapter{
		userCol:  db.Database(databaseName).Collection(userCollection),
		guestCol: db.Database(databaseName).Collection(guestCollection),
	}, nil

}

func (a *Adapter) CreateUser(ctx context.Context, userRequest models.UserRequest) error {
	salt := generateSalt()

	objId, err := primitive.ObjectIDFromHex(utils.ComputeHash(userRequest.Email, ""))
	if err != nil {
		return err
	}

	user := models.Admin{
		ID:           objId,
		FullName:     userRequest.FullName,
		PhoneNumber:  userRequest.PhoneNumber,
		Email:        userRequest.Email,
		PasswordHash: utils.ComputeHash(userRequest.Password, salt),
		Salt:         salt,
		CreatedAt:    primitive.NewDateTimeFromTime(time.Now().UTC()),
	}

	_, err = a.userCol.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (a *Adapter) CreateGuests(ctx context.Context, guests []models.PeopleRequest) error {
	var people []interface{}
	for _, gst := range guests {
		objId, err := primitive.ObjectIDFromHex(utils.ComputeHash(gst.Email, ""))
		if err != nil {
			log.Println(err, "error creating object id")
			return err
		}
		log.Println(objId, gst.FullName, "object id")
		guest := models.Guest{
			ID:          objId,
			FullName:    gst.FullName,
			PhoneNumber: gst.PhoneNumber,
			Email:       gst.Email,
			Gender:      gst.Gender,
			Group:       gst.Group,
			CreatedAt:   primitive.NewDateTimeFromTime(time.Now().UTC()),
		}

		people = append(people, guest)
	}

	_, err := a.guestCol.InsertMany(ctx, people)
	if err != nil {
		log.Println(err, "error inserting guest")
		return err
	}
	return nil
}

func (a *Adapter) GetUserByEmail(ctx context.Context, email string) (*models.Admin, error) {
	objId, err := primitive.ObjectIDFromHex(utils.ComputeHash(email, ""))
	if err != nil {
		return nil, err
	}

	var user models.Admin
	err = a.userCol.FindOne(ctx, bson.M{"_id": objId}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (a *Adapter) GetGuestByEmail(ctx context.Context, email string) (*models.Guest, error) {
	objId, err := primitive.ObjectIDFromHex(utils.ComputeHash(email, ""))
	if err != nil {
		return nil, err
	}

	var guest models.Guest
	err = a.guestCol.FindOne(ctx, bson.M{"_id": objId}).Decode(&guest)
	if err != nil {
		return nil, err
	}

	return &guest, nil
}

func (a *Adapter) GetGuestsByGroup(ctx context.Context, group int) (*[]models.Guest, error) {
	var guests []models.Guest
	filter := bson.M{"group": group}
	opts := options.Find().SetSort(bson.D{{"name", 1}})
	cursor, err := a.guestCol.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.TODO(), &guests); err != nil {
		return nil, err
	}
	return &guests, nil
}

func (a *Adapter) GetGuests(ctx context.Context) (*[]models.Guest, error) {
	var guests []models.Guest
	cursor, err := a.guestCol.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.TODO(), &guests); err != nil {
		return nil, err
	}
	return &guests, nil
}

func generateSalt() string {
	rand.Seed(time.Now().Unix())
	result := ""

	for i := 0; i <= 8; i++ {
		result += fmt.Sprint('0' + rand.Intn(41))
	}
	return result
}
