package commgo

import (
	"time"
)

// for ReplSet Config files

type RsConf struct {
	ID       string                 `bson:"_id"`
	Version  int64                  `bson:"version"`
	Members  []*Host                `bson:"members"`
	Settings map[string]interface{} `bson:"settings"`
}

type Host struct {
	ID           int64             `bson:"_id"`
	Host         string            `bson:"host"`
	ArbiterOnly  bool              `bson:"arbiterOnly"`
	BuildIndexes bool              `bson:"buildIndexes"`
	Hidden       bool              `bson:"hidden"`
	Priority     int64             `bson:"priority"`
	Tags         map[string]string `bson:"tags"`
	SlaveDelay   int64             `bson:"slaveDelay"`
	Votes        int64             `bson:"votes"`
}

// rs.status()

type RsMemberStats struct {
	ID                int        `bson:"_id"`
	Name              string     `bson:"name"`
	Health            int        `bson:"health"`
	State             int        `bson:"state"`
	StateStr          string     `bson:"stateStr"`
	Uptime            int        `bson:"uptime"`
	Optime            int        `bson:"optime"`
	OptimeDate        time.Time  `bson:"optimeDate"`
	Self              *bool      `bson:"self" json:",omitempty"`
	LastHeartbeat     *time.Time `bson:"lastHeartbeat" json:",omitempty"`
	LastHeartbeatRecv *time.Time `bson:"lastHeatbeatRecv" json:",omitempty"`
	PingMS            *int       `bson:"pingMs" json:",omitempty"`
	SyncingTo         *string    `bson:"syncingTo" json:",omitempty"`
}

type RsStatus struct {
	Set     string           `bson:"set"`
	Date    time.Time        `bson:"date"`
	MyState int              `bson:"myState"`
	Members []*RsMemberStats `bson:"members"`
	Fake1   int              `bson:"fake1" json:",omitempty"`
	Fake2   *int             `bson:"fake2" json:",omitempty"`
	Ok      int              `bson:"ok"`
}
