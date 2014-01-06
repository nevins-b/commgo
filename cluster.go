package commgo

import (
       "labix.org/v2/mgo/bson"
)

// sh.status()

type ShardVersion struct {
	ID                     int           `bson:"_id"`
	Version                int           `bson:"version"`
	MinCompatibleVersion   int           `bson:"minCompatibleVersion"`
	CurrentVersion         int           `bson:"currentVersion"`
	ClusterID              bson.ObjectId `bson:"clusterId"`
}

type Shard struct {
	ID     int    `bson:"_id"`
	Host   string `bson:"host"`
}

type Chunk struct {
	Min string
	Max string
	On  string
}

type ShardedColl struct {
	NS         string 
	ShardKey   bson.M 
	ChunkCount map[string]int
	Chunks     []Chunk
}

type ShardedDB struct {
	ID           string          `bson:"_id"`
	Partitioned  bool            `bson:"partitioned"`
	Primary      string          `bson:"primary"`
	Collections  []*ShardedColl
}

type ShStatus struct {
	ShardingVersion  ShardVersion
	Shards           []Shard
	Databases        map[string]*ShardedDB
}
