package commgo

import (
	"time"
)

// for ReplSet Config files

type Host struct {
	ID          int       `bson:"_id"`
	Host        string    `bson:"host"`
}

type RsConf struct {
	ID          int       `bson:"_id"`
	Version     int       `bson:"version"`
        Members     []Host    `bson:"members"`
}

// rs.status()

type RsMemberStats struct {
        ID         int        `bson:"_id"`
	Name       string     `bson:"name"`
	Health     int        `bson:"health"`
	State      int        `bson:"state"`
	StateStr   string     `bson:"stateStr"`
	Uptime     int        `bson:"uptime"`
	Optime     int        `bson:"optime"`
	OptimeDate time.Time  `bson:"optimeDate"`
	Self          *bool    `bson:"self" json:",omitempty"`
	LastHeartbeat *time.Time `bson:"lastHeartbeat" json:",omitempty"`
	LastHeartbeatRecv *time.Time `bson:"lastHeatbeatRecv" json:",omitempty"`
	PingMS     *int         `bson:"pingMs" json:",omitempty"`
	SyncingTo *string       `bson:"syncingTo" json:",omitempty"`
}

type RsStatus struct {
	Set         string    `bson:"set"`
	Date        time.Time `bson:"date"`
	MyState     int       `bson:"myState"`
	Members     []*RsMemberStats `bson:"members"`
	Fake1       int       `bson:"fake1" json:",omitempty"`
	Fake2       *int      `bson:"fake2" json:",omitempty"`
	Ok          int       `bson:"ok"`
}


