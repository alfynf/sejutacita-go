package databases

import (
	"context"
	"log"
	"sejutacita/config"
	"sejutacita/lib/plugins"
	"sejutacita/middlewares"
	"sejutacita/models"

	"go.mongodb.org/mongo-driver/bson"
)

// fungsi untuk mendaftarkan user baru
func CreateUser(user models.User) (interface{}, error) {
	var collection = config.DB.Database("sejutacita").Collection("users")
	insertResult, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		return nil, err
	}
	return insertResult.InsertedID, nil
}

// fungsi untuk melihat daftar semua user yang sudah pernah terdaftar
func GetUser(filter bson.M) ([]*models.GetUser, error) {
	var users []*models.GetUser
	var collection = config.DB.Database("sejutacita").Collection("users")
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	for cur.Next(context.TODO()) {
		var user models.GetUser
		err := cur.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

// fungsi untuk melihat informasi user berdasarkan email
func GetUserByEmail(filter bson.M) (*models.User, error) {
	var result *models.User
	var collection = config.DB.Database("sejutacita").Collection("users")
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// fungsi untuk melihat informasi user berdasarkan username
func GetUserByUsername(filter bson.M) (*models.GetUser, error) {
	var result *models.GetUser
	var collection = config.DB.Database("sejutacita").Collection("users")
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// fungsi untuk menghapus user
func DeleteUser(filter bson.M) (int64, error) {
	var collection = config.DB.Database("sejutacita").Collection("users")
	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return -1, err
	}
	return deleteResult.DeletedCount, nil

}

// fungsi untuk melakukan update user
func UpdateUser(updateData models.User, filter bson.M) (int64, error) {
	collection := config.DB.Database("sejutacita").Collection("users")
	atualizacao := bson.D{{Key: "$set", Value: updateData}}
	updatedResult, err := collection.UpdateOne(context.TODO(), filter, atualizacao)
	if err != nil {
		return -1, err
	}
	return updatedResult.ModifiedCount, nil
}

// fungsi untuk melakukan login user
func Login(userLogin models.User) (interface{}, error) {
	var result *models.User
	user, err := GetUserByEmail(bson.M{"email": userLogin.Email})
	if err != nil {
		return nil, err
	}

	check := plugins.Decrypt(user.Password, userLogin.Password)
	if !check {
		return 0, nil
	}

	user.Token, err = middlewares.CreateToken(user.Username, user.Role)
	if err != nil {
		return nil, err
	}
	updateData, _ := ToDoc(user)
	collection := config.DB.Database("sejutacita").Collection("users")
	atualizacao := bson.D{{Key: "$set", Value: updateData}}
	_, err = collection.UpdateOne(context.TODO(), bson.M{"username": user.Username}, atualizacao)
	if err != nil {
		return -1, err
	}
	result, _ = GetUserByEmail(bson.M{"email": userLogin.Email})
	return result, nil
}

// fungsi untuk mengubah tipe data struct menjadi bson
func ToDoc(v interface{}) (doc *bson.M, err error) {
	data, err := bson.Marshal(v)
	if err != nil {
		return
	}

	err = bson.Unmarshal(data, &doc)
	return
}

// fungsi untuk membuat data admin awal
func MockDataAdmin() {
	var collection = config.DB.Database("sejutacita").Collection("users")
	_, err := collection.DeleteOne(context.TODO(), bson.M{"username": "john"})
	if err != nil {
		log.Fatal(err)
	}
	user := models.User{
		Name:     "John Doe",
		Username: "john",
		Email:    "john@mail.com",
		Password: "12345678",
		Role:     "admin",
	}
	user.Password, _ = plugins.Encrypt(user.Password)
	_, err = collection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}
}
