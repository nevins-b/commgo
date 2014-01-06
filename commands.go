package commgo

import (
	"errors"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

// sh.status()
//
// GetShStatus acts similarly to the shell command sh.status().  It
// gathers information about the current cluster configuration, if run
// on a cluster, and returns that information to the caller wrapped in
// a ShStatus object.  Errors will be thrown if GetShStatus is run
// on a non-cluster.
//
func GetShStatus( session *mgo.Session ) (ShStatus, error) {
	config := session.DB("config")
	status := ShStatus{}

	version := ShardVersion{}
	err := config.C("version").Find(nil).One(&version)
	if err != nil {
		return status, errors.New("Sharding is not enabled.  Be sure to connect to a mongos and not to a mongod.")
	}

	// get shards
	iter := config.C("shards").Find(nil).Iter()
	err = iter.All(&status.Shards)
	if err != nil { return status, err }
	err = iter.Close()
	if err != nil { return status, err }

	// get database information
	iter = config.C("databases").Find(nil).Sort("name").Iter()
	hostdoc := bson.M{}
	status.Databases = make(map[string]*ShardedDB)
	for iter.Next(&hostdoc) {
		var name string
		name = hostdoc["_id"].(string)

		db := &ShardedDB{}
		db.ID = name
		db.Partitioned = hostdoc["partitioned"].(bool)
		db.Primary = hostdoc["primary"].(string)

		if hostdoc["partitioned"].(bool) {

			// add the collection information
			db.Collections = []*ShardedColl{}
			reg := bson.RegEx{ "^" + hostdoc["_id"].(string) + "\\.", ""}
			iter2 := config.C("collections").Find(bson.M{ "_id" : reg }).Sort("_id").Iter()
			colldoc := bson.M{}

			for iter2.Next(&colldoc) {
				if colldoc["dropped"] == false {
					coll := &ShardedColl{}
					coll.NS = colldoc["_id"].(string)
					coll.ShardKey = colldoc["key"].(bson.M)
					coll.ChunkCount = make(map[string]int)

					// get chunks
					chunkdoc := bson.M{}
					iter3 := config.C("chunks").Find(bson.M{ "ns" : colldoc["_id"] }).Sort("min").Iter()
					for iter3.Next(&chunkdoc) {
						chunk := Chunk{}

						// if "min" is a document, it's Minkey
						switch chunkdoc["min"].(bson.M)["n"].(type) {
						default:
							chunk.Min = "$minKey"
						case string:
							chunk.Min = chunkdoc["min"].(bson.M)["n"].(string)
						}

						// if "max" is a document, it's Maxkey
						switch chunkdoc["max"].(bson.M)["n"].(type) {
						default:
							chunk.Max = "$maxKey"
						case string:
							chunk.Max = chunkdoc["max"].(bson.M)["n"].(string)
						}

						chunk.On = chunkdoc["shard"].(string)
						coll.ChunkCount[chunk.On]++
						coll.Chunks = append(coll.Chunks, chunk)
					}
					db.Collections = append(db.Collections, coll)
				}
			}
		}
		status.Databases[name] = db
	}

	status.ShardingVersion = version
	return status, nil
}
