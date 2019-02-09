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

//Find a Restaurant by type_of_food
func (r *MongoRepository) FindByTypeOfFood(type_of_food string) (*[]domain.Restaurant, error) {
        result := []domain.Restaurant{}
        session := r.mongoSession.Clone()
        defer session.Close()
        coll := session.DB(r.db).C(collectionName)
        err := coll.Find(bson.M{"type_of_food": type_of_food}).All(&result)
        switch err {
        case nil:
                return &result, nil
        case mgo.ErrNotFound:
                return nil, domain.ErrNotFound
        default:
                return nil, err
        }
}

//Find a Restaurant by postcode
func (r *MongoRepository) FindByTypeOfPostCode(postcode string) (*[]domain.Restaurant, error) {
        result := []domain.Restaurant{}
        session := r.mongoSession.Clone()
        defer session.Close()
        coll := session.DB(r.db).C(collectionName)
        err := coll.Find(bson.M{"postcode": postcode}).All(&result)
        switch err {
        case nil:
                return &result, nil
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
        count,err := coll.Find(bson.M{"type_of_food": criteria}).Count()
        switch err {
        case nil:
                return count, nil
        case mgo.ErrNotFound:
                return 0, domain.ErrNotFound
        default:
                return 0, err
        }
}



//Store a Restaurantrecord
func (r *MongoRepository) Store(b *domain.Restaurant) (domain.ID, error) {
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	if domain.ID(0) == b.DBID {
		b.DBID = domain.NewID()
	}

	_, err := coll.UpsertId(b.DBID, b)

	if err != nil {
		return domain.ID(0), err
	}
	return b.DBID, nil
}
