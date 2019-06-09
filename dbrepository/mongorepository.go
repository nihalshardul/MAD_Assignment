package dbrepository

import (
	//	"github.com/priteshgudge/mongorestaurantsample/domain"
	"mongorestaurantsample-master/domain"

	mgo "gopkg.in/mgo.v2"
	bson "gopkg.in/mgo.v2/bson"
)

//MongoRepository mongodb repo
type MongoRepository struct {
	mongoSession *mgo.Session
	db           string
}

var collectionName = "restaurant"

//NewMongoRepository create new repository
func NewMongoRepository(mongoSession *mgo.Session, db string) *MongoRepository {
	return &MongoRepository{
		mongoSession: mongoSession,
		db:           db,
	}
}

//Find a Restaurant
func (r *MongoRepository) Get(id domain.ID) (*domain.Restaurant, error) {
	result := domain.Restaurant{}
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	err := coll.Find(bson.M{"_id": id}).One(&result)
	switch err {
	case nil:
		return &result, nil
	case mgo.ErrNotFound:
		return nil, domain.ErrNotFound
	default:
		return nil, err
	}
}

// getAll
func (r *MongoRepository) GetAll() ([]*domain.Restaurant, error) {
	result := []*domain.Restaurant{}
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	err := coll.Find(bson.M{}).All(&result)
	switch err {
	case nil:
		return result, nil
	case mgo.ErrNotFound:
		return nil, domain.ErrNotFound
	default:
		return nil, err
	}
}

//FindByName
func (r *MongoRepository) FindByName(name string) ([]*domain.Restaurant, error) {
	result := []*domain.Restaurant{}
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	err := coll.Find(bson.M{"name": bson.RegEx{name, "i"}}).All(&result)
	switch err {
	case nil:
		return result, nil
	case mgo.ErrNotFound:
		return nil, domain.ErrNotFound
	default:
		return nil, err
	}
}

//Delete by id
func (r *MongoRepository) Delete(id domain.ID) error {
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	err := coll.Remove(bson.M{"_id": id})
	switch err {
	case nil:
		return nil
	case mgo.ErrNotFound:
		return domain.ErrNotFound
	default:
		return err
	}
}

//Search fields by substring
func (r *MongoRepository) Search(query string) ([]*domain.Restaurant, error) {
	result := []*domain.Restaurant{}
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	err := coll.Find(bson.M{"$or": []bson.M{bson.M{"name": bson.RegEx{query, "i"}}, bson.M{"address": bson.RegEx{query, "i"}}, bson.M{"addressline2": bson.RegEx{query, "i"}}, bson.M{"url": bson.RegEx{query, "i"}}, bson.M{"outcode": bson.RegEx{query, "i"}}, bson.M{"postcode": bson.RegEx{query, "i"}}, bson.M{"type_of_food": bson.RegEx{query, "i"}}}}).All(&result)
	switch err {
	case nil:
		return result, nil
	case mgo.ErrNotFound:
		return nil, domain.ErrNotFound
	default:
		return nil, err
	}
}

//Find a Restaurant by type_of_food
func (r *MongoRepository) FindByTypeOfFood(type_of_food string) ([]*domain.Restaurant, error) {
	result := []*domain.Restaurant{}
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	err := coll.Find(bson.M{"type_of_food": type_of_food}).All(&result)
	switch err {
	case nil:
		return result, nil
	case mgo.ErrNotFound:
		return nil, domain.ErrNotFound
	default:
		return nil, err
	}
}

//Find a Restaurant by postcode
func (r *MongoRepository) FindByTypeOfPostCode(postcode string) ([]*domain.Restaurant, error) {
	result := []*domain.Restaurant{}
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	err := coll.Find(bson.M{"postcode": postcode}).All(&result)
	switch err {
	case nil:
		return result, nil
	case mgo.ErrNotFound:
		return nil, domain.ErrNotFound
	default:
		return nil, err
	}
}

//Find a Restaurant count
func (r *MongoRepository) CountRestaurantByFood(criteria string) (int, error) {
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	count, err := coll.Find(bson.M{"type_of_food": criteria}).Count()
	switch err {
	case nil:
		return count, nil
	case mgo.ErrNotFound:
		return 0, domain.ErrNotFound
	default:
		return 0, err
	}
}

//update the field
func (r *MongoRepository) Update(b *domain.Restaurant, ID string) error {
	newID := domain.StringToID(ID)
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	err := coll.Update(bson.M{"_id": newID}, bson.M{"$set": bson.M{"name": b.Name, "address": b.Address, "addressline2": b.AddressLine2, "url": b.URL, "opcode": b.Outcode, "postcode": b.Postcode, "rating": b.Rating, "type_of_food": b.TypeOfFood}})
	return err
}

//Store a Restaurantrecord
func (r *MongoRepository) Store(b *domain.Restaurant) (domain.ID, error) {
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	b.DBID = domain.NewID()
	if domain.ID(0) == b.DBID {
		b.DBID = domain.NewID()
	}

	_, err := coll.UpsertId(b.DBID, b)

	if err != nil {
		return domain.ID(0), err
	}
	return b.DBID, nil
}
