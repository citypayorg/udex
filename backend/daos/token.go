package daos

import (
	"time"

	"github.com/citypayorg/udex/tree/udex/backend/app"
	"github.com/citypayorg/udex/tree/udex/backend/types"
	mgo "github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// TokenDao contains:
// collectionName: MongoDB collection name
// dbName: name of mongodb to interact with
type TokenDao struct {
	collectionName string
	dbName         string
}

// NewTokenDao returns a new instance of TokenDao.
func NewTokenDao() *TokenDao {
	dbName := app.Config.DBName
	collection := "tokens"
	index := mgo.Index{
		Key:    []string{"asset"},
		Unique: true,
	}

	err := db.Session.DB(dbName).C(collection).EnsureIndex(index)
	if err != nil {
		panic(err)
	}
	return &TokenDao{collection, dbName}
}

// Create function performs the DB insertion task for token collection
func (dao *TokenDao) Create(token *types.Token) error {
	if err := token.Validate(); err != nil {
		logger.Error(err)
		return err
	}

	token.ID = bson.NewObjectId()
	token.CreatedAt = time.Now()
	token.UpdatedAt = time.Now()

	err := db.Create(dao.dbName, dao.collectionName, token)
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

// GetAll function fetches all the tokens in the token collection of mongodb.
func (dao *TokenDao) GetAll() ([]types.Token, error) {
	var res []types.Token

	sort := []string{"-rank"}

	err := db.GetAndSort(dao.dbName, dao.collectionName, bson.M{}, sort, 0, 0, &res)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return res, nil
}

func (dao *TokenDao) GetListedTokens() ([]types.Token, error) {
	var res []types.Token

	sort := []string{"-rank"}
	q := bson.M{"listed": true}

	err := db.GetAndSort(dao.dbName, dao.collectionName, q, sort, 0, 0, &res)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return res, nil
}

// GetQuote function fetches all the quote tokens in the token collection of mongodb.
func (dao *TokenDao) GetQuoteTokens() ([]types.Token, error) {
	var res []types.Token

	sort := []string{"-rank"}
	q := bson.M{"quote": true}

	err := db.GetAndSort(dao.dbName, dao.collectionName, q, sort, 0, 0, &res)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return res, nil
}

// GetBase function fetches all the base tokens in the token collection of mongodb.
func (dao *TokenDao) GetBaseTokens() ([]types.Token, error) {
	var res []types.Token

	sort := []string{"-rank"}
	q := bson.M{"quote": false}
	err := db.GetAndSort(dao.dbName, dao.collectionName, q, sort, 0, 0, &res)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	if res == nil {
		res = []types.Token{}
	}

	return res, nil
}

func (dao *TokenDao) GetListedBaseTokens() ([]types.Token, error) {
	var res []types.Token

	sort := []string{"-rank"}
	q := bson.M{"quote": false, "listed": true}

	err := db.GetAndSort(dao.dbName, dao.collectionName, q, sort, 0, 0, &res)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	if res == nil {
		res = []types.Token{}
	}

	return res, nil
}

func (dao *TokenDao) GetUnlistedTokens() ([]types.Token, error) {
	var res []types.Token

	sort := []string{"-rank"}
	q := bson.M{"quote": false, "listed": false}

	err := db.GetAndSort(dao.dbName, dao.collectionName, q, sort, 0, 0, &res)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	if res == nil {
		res = []types.Token{}
	}

	return res, nil
}

// GetByID function fetches details of a token based on its mongo id
func (dao *TokenDao) GetByID(id bson.ObjectId) (*types.Token, error) {
	var res *types.Token
	err := db.GetByID(dao.dbName, dao.collectionName, id, &res)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return res, nil
}

// GetByAsset function fetches details of a token based on its asset ID
func (dao *TokenDao) GetByAsset(asset string) (*types.Token, error) {
	q := bson.M{"asset": asset}
	var resp []types.Token

	err := db.Get(dao.dbName, dao.collectionName, q, 0, 1, &resp)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	if len(resp) == 0 {
		return nil, nil
	}

	return &resp[0], nil
}

// GetBySymbol function fetches details of a token based on its asset ID
func (dao *TokenDao) GetBySymbol(symbol string) (*types.Token, error) {
	q := bson.M{"symbol": symbol}
	var resp []types.Token

	err := db.Get(dao.dbName, dao.collectionName, q, 0, 1, &resp)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	if len(resp) == 0 {
		return nil, nil
	}

	return &resp[0], nil
}

func (dao *TokenDao) GetByAssetOrSymbol(assetOrsymbol string) (*types.Token, error) {
	t, err := dao.GetByAsset(assetOrsymbol)
	if t != nil || err != nil {
		return t, err
	}
	return dao.GetBySymbol(assetOrsymbol)
}

// Drop drops all the order documents in the current database
func (dao *TokenDao) Drop() error {
	err := db.DropCollection(dao.dbName, dao.collectionName)
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}
