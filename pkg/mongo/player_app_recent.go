package mongo

import (
	"strconv"

	"github.com/gamedb/gamedb/pkg/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PlayerRecentApp struct {
	PlayerID        int64  `bson:"player_id"`
	AppID           int    `bson:"app_id"`
	AppName         string `bson:"name"`
	PlayTime2Weeks  int    `bson:"playtime_2_weeks"` // Minutes
	PlayTimeForever int    `bson:"playtime_forever"` // Minutes
	Icon            string `bson:"icon"`
	Logo            string `bson:"logo"`
}

func (g PlayerRecentApp) BSON() bson.D {

	return bson.D{
		{"_id", g.getKey()},
		{"player_id", g.PlayerID},
		{"app_id", g.AppID},
		{"name", g.AppName},
		{"playtime_2_weeks", g.PlayTime2Weeks},
		{"playtime_forever", g.PlayTimeForever},
		{"icon", g.Icon},
		{"logo", g.Logo},
	}
}

func (g PlayerRecentApp) getKey() (ret interface{}) {

	return strconv.FormatInt(g.PlayerID, 10) + "-" + strconv.Itoa(g.AppID)
}

func DeleteRecentApps(playerID int64, apps []int) (err error) {

	if len(apps) < 1 {
		return nil
	}

	client, ctx, err := getMongo()
	if err != nil {
		return err
	}

	keys := bson.A{}
	for _, appID := range apps {

		player := PlayerRecentApp{}
		player.PlayerID = playerID
		player.AppID = appID

		keys = append(keys, player.getKey())
	}

	collection := client.Database(MongoDatabase).Collection(CollectionPlayerAppsRecent.String())
	_, err = collection.DeleteMany(ctx, bson.M{"_id": bson.M{"$in": keys}})
	return err
}

func UpdateRecentApps(apps []PlayerRecentApp) (err error) {

	if len(apps) < 1 {
		return nil
	}

	client, ctx, err := getMongo()
	if err != nil {
		return err
	}

	var writes []mongo.WriteModel
	for _, app := range apps {

		write := mongo.NewReplaceOneModel()
		write.SetFilter(bson.M{"_id": app.getKey()})
		write.SetReplacement(app.BSON())
		write.SetUpsert(true)

		writes = append(writes, write)
	}

	collection := client.Database(MongoDatabase).Collection(CollectionPlayerAppsRecent.String())
	_, err = collection.BulkWrite(ctx, writes, options.BulkWrite())
	return err
}

func GetRecentApps(playerID int64, offset int64, limit int64, sort bson.D) (apps []PlayerRecentApp, err error) {

	filter := bson.D{{"player_id", playerID}}

	cur, ctx, err := Find(CollectionPlayerAppsRecent, offset, limit, sort, filter, nil, nil)
	if err != nil {
		return apps, err
	}

	defer func(cur *mongo.Cursor) {
		err = cur.Close(ctx)
		log.Err(err)
	}(cur)

	for cur.Next(ctx) {

		var app PlayerRecentApp
		err := cur.Decode(&app)
		if err != nil {
			log.Err(err, app.getKey())
		}
		apps = append(apps, app)
	}

	return apps, cur.Err()
}
