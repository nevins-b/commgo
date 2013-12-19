package commgo

import (
	"labix.org/v2/mgo/bson"
	"time"
)

// db.stats()

type DBStats struct {
	Db              string  `bson:"db"`
	Collections     int     `bson:"collections"`
	Objects         int     `bson:"objects"`
	AvgObjSize      float64 `bson:"avgObjSize"`
	DataSize        int     `bson:"dataSize"`
	StorageSize     int     `bson:"storageSize"`
	NumExtents      int     `bson:"numExtents"`
	Indexes         int     `bson:"indexes"`
	IndexSize       int     `bson:"indexSize"`
	FileSize        int     `bson:"fileSize"`
	NsSizeMB        int     `bson:"nsSizeMB"`
	DataFileVersion struct {
		Major int `bson:"major"`
		Minor int `bson:"minor"`
	} `bson:"dataFileVersion"`
	Ok int `bson:"ok"`
}

// db.runCommand({ buildInfo : 1 })

type BuildInfo struct {
	Version            string `bson:"version"` 
	GitVersion         string `bson:"gitVersion"`
	SysInfo            string `bson:"sysInfo"`
	LoaderFlags        string `bson:"loaderFlags"`
	CompilerFlags      string `bson:"compilerFlags"`
	Allocator          string `bson:"allocator"`
	VersionArray       []int  `bson:"versionArray"`
	JavascriptEngine   string `bson:"javascriptEngine"`
	Bits               int    `bson:"bits"`
	Debug              bool   `bson:"debug"`
	MaxBsonObjectSize  int    `bson:"maxBsonObjectSize"`
	Ok                 int    `bson:"ok"`
}

// db.adminCommand({ getLog : 1 })

type GetLog struct {
	LinesWritten int      `bson:"totalLinesWritten"`
	Logs         []string `bson:"log"`
	Ok           int      `bson:"ok"`
}

/* -----------------------------------
   These are all for db.serverStatus()
   ----------------------------------- */

type AssertStats struct {
	Msg       int `bson:"msg"`
	Regular   int `bson:"regular"`
	Rollovers int `bson:"rollovers"`
	User      int `bson:"user"`
	Warning   int `bson:"warning"`
}

type BackgroundFlushStats struct {
	Flushes      int       `bson:"flushes"`
	TotalMS      int       `bson:"total_ms"`
	AverageMS    float64   `bson:"average_ms"`
	LastMS       int       `bson:"last_ms"`
	LastFinished time.Time `bson:"last_finishes"`
}

type ConnectionStats struct {
	Available    int `bson:"available"`
	Current      int `bson:"current"`
	TotalCreated int `bson:"totalCreated"`
}

type CursorStats struct {
	TotalOpen        int `bson:"totalOpen"`
	ClientCursorSize int `bson:"clientCursors_size"`
	TimedOut         int `bson:"timedOut"`
}

type DurTimeStats struct {
	DT               int `bson:"dt"`
	PrepLogBuffer    int `bson:"prepLogBuffer"`
	WriteToJournal   int `bson:"writeToJournal"`
	WriteToDataFiles int `bson:"writeToDataFiles"`
	RemapPrivateView int `bson:"remapPrivateView"`
}

type DurStats struct {
	Commits            int          `bson:"commits"`
	JournaledMB        int          `bson:"journaledMB"`
	WriteToDataFilesMB int          `bson:"writeToDataFilesMB"`
	Compression        int          `bson:"compression"`
	CommitsInWriteLock int          `bson:"commitsInWriteLock"`
	EarlyCommits       int          `bson:"earlyCommits"`
	TimeMS             DurTimeStats `bson:"timeMS"`
}

// adapt this to be more flexible?
type ExtraInfo struct {
	Note       string `bson:"note"`
	PageFaults int    `bson:"page_faults"`
	Misc       bson.M `bson:",inline"`
}

type LockQueueInfo struct {
	Total   int `bson:"total"`
	Readers int `bson:"readers"`
	Writers int `bson:"writers"`
}

type GlobalLockStats struct {
	TotalTime     int           `bson:"totalTime"`
	LockTime      int           `bson:"lockTime"`
	CurrentQueue  LockQueueInfo `bson:"currentQueue"`
	ActiveClients LockQueueInfo `bson:"activeClients"`
}

type IndexStats struct {
	Accesses  int `bson:"accesses"`
	Hits      int `bson:"hits"`
	Misses    int `bson:"misses"`
	Resets    int `bson:"resets"`
	MissRatio int `bson:"missRatio"`
}

type LockTimeStats struct {
	R int `bson:"r"`
	W int `bson:"w"`
}

type LockDetails struct {
	TimeLockedMicros    LockTimeStats `bson:"timeLockedMicros"`
	TimeAcquiringMicros LockTimeStats `bson:"timeAcquiringMicros"`
}

type LockInfo struct {
	Dot       LockDetails `bson:"."`
	Admin     LockDetails `bson:"admin"`
	Local     LockDetails `bson:"local"`
	Databases bson.M      `bson:",inline"`
}

type MemInfo struct {
	Bits              int  `bson:"bits"`
	Resident          int  `bson:"resident"`
	Supported         bool `bson:"supported"`
	Virtual           int  `bson:"virtual"`
	Mapped            int  `bson:"mapped"`
	MappedWithJournal int  `bson:"mappedWithJournal"`
}

type DocMetrics struct {
	Deleted  int64 `bson:"deleted"`
	Inserted int64 `bson:"inserted"`
	Returned int64 `bson:"returned"`
	Updated  int64 `bson:"updated"`
}

type OpMetrics struct {
	Fastmod      int64 `bson:"fastmod"`
	IDHack       int64 `bson:"idhack"`
	ScanAndOrder int64 `bson:"scanAndOrder"`
}

type TimedMetric struct {
	Num     int `bson:"num"`
	TotalMs int `bson:"totalMillis"`
}

type ReplMetrics struct {
	Apply struct {
		Batches TimedMetric `bson:"batches" json:",omitempty"`
		Ops     int64       `bson:"ops" json:",omitempty"`
	} `bson:"apply" json:",omitempty"`
	Buffer struct {
		Count        int64 `bson:"count" json:",omitempty"`
		MaxSizeBytes int   `bson:"maxSizeBytes" json:",omitempty"`
		SizeBytes    int64 `bson:"sizeBytes" json:",omitempty"`
	} `bson:"buffer" json:",omitempty"`
	Network struct {
		Bytes          int64       `bson:"bytes" json:",omitempty"`
		Getmores       TimedMetric `bson:"getmores" json:",omitempty"`
		Ops            int64       `bson:"ops" json:",omitempty"`
		ReadersCreated int64       `bson:"readersCreated" json:",omitempty"`
	} `bson:"network" json:",omitempty"`
	Oplog struct {
		Insert      TimedMetric `bson:"insert" json:",omitempty"`
		InsertBytes int64       `bson:"insertBytes" json:",omitempty"`
	} `bson:"oplog" json:",omitempty"`
	Preload struct {
		Docs    TimedMetric `bson:"docs" json:",omitempty"`
		Indexes TimedMetric `bson:"indexes" json:",omitempty"`
	} `bson:"preload" json:",omitempty"`
}

type TTLMetrics struct {
	DeletedDocs int64 `bson:"deletedDocuments"`
	Passes      int64 `bson:"passes"`
}

type MetricsInfo struct {
	Document     DocMetrics `bson:"document"`
	Operation    OpMetrics  `bson:"operation"`
	GetLastError struct {
		WTimeouts int64 `bson:"wtimeouts"`
		WTime     struct {
			Num         int `bson:"num"`
			TotalMillis int `bson:"totalMillis"`
		} `bson:"wtime"`
	} `bson:"getLastError"`
	QueryExecutor struct {
		Scanned int64 `bson:"scanned"`
	} `bson:"queryExecutor"`
	Record struct {
		Moves int64 `bson:"moves"`
	} `bson:"record"`
	Repl ReplMetrics `bson:"repl" json:",omitempty"`
	TTL  TTLMetrics  `bson:"ttl"`
	OK   int         `bson:"ok"`
}

type NetworkStats struct {
	BytesIn     int `bson:"bytesIn"`
	BytesOut    int `bson:"bytesOut"`
	NumRequests int `bson:"numRequests"`
}

type OpcounterStats struct {
	Command int `bson:"command" json:",omitempty"`
	Delete  int `bson:"delete" json:",omitempty"`
	Getmore int `bson:"getmore" json:",omitempty"`
	Insert  int `bson:"insert" json:",omitempty"`
	Query   int `bson:"query" json:",omitempty"`
	Update  int `bson:"update" json:",omitempty"`
}

type DBRecordStats struct {
	AccessesNotInMem    int `bson:"accessesNotInMemory"`
	PageFaultExceptions int `bson:"pageFaultExceptionsThrown"`
}

type RecordStatInfo struct {
	AccessesNotInMem    int                      `bson:"accessesNotInMemory"`
	PageFaultExceptions int                      `bson:"pageFaultExceptionsThrown"`
	Admin               DBRecordStats            `bson:"admin"`
	Local               DBRecordStats            `bson:"local"`
	DBs                 map[string]DBRecordStats `bson:",inline" json:",omitempty"`
}

type ReplStats struct {
	SetName     string   `bson:"setName" json:",omitempty"`
	SetVersion  int      `bson:"setVersion" json:",omitempty"`
	IsMaster    bool     `bson:"ismaster" json:",omitempty"`
	IsSecondary bool     `bson:"secondary" json:",omitempty"`
	Hosts       []string `bson:"hosts" json:",omitempty"`
	Primary     string   `bson:"primary" json:",omitempty"`
	Me          string   `bson:"me" json:",omitempty"`
}

type ServerStatus struct {
	Asserts            AssertStats          `bson:"asserts"`
	BackgroundFlushing BackgroundFlushStats `bson:"backgroundFlushing"`
	Connections        ConnectionStats      `bson:"connections"`
	Cursors            CursorStats          `bson:"cursors"`
	Dur                DurStats             `bson:"dur"`
	ExtraInfo          ExtraInfo            `bson:"extra_info"`
	GlobalLock         GlobalLockStats      `bson:"globalLock"`
	Host               string               `bson:"host"`
	IndexCounters      IndexStats           `bson:"indexCounters"`
	LocalTime          time.Time            `bson:"localTime"`
	Locks              LockInfo             `bson:"locks"`
	Mem                MemInfo              `bson:"mem"`
	Metrics            MetricsInfo          `bson:"metrics"`
	Network            NetworkStats         `bson:"network"`
	Ok                 int                  `bson:"ok"`
	Opcounters         OpcounterStats       `bson:"opcounters"`
	OpcountersRepl     OpcounterStats       `bson:"opcountersRepl" json:",omitempty"`
	PID                int                  `bson:"pid"`
	Process            string               `bson:"process"`
	RecordStats        RecordStatInfo       `bson:"recordStats"`
	Repl               ReplStats            `bson:"repl,omitempty" json:",omitempty"`
	Uptime             int                  `bson:"uptime"`
	UptimeEstimate     int                  `bson:"uptimeEstimate"`
	UptimeMillis       int                  `bson:"uptimeMillis"`
	Version            string               `bson:"version"`
	WriteBacksQueued   bool                 `bson:"writeBacksQueued"`
}
