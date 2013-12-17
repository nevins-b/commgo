package commgo

import (
	"labix.org/v2/mgo/bson"
	"time"
)

// for output of dbstats command

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
		Batches TimedMetric `bson:"batches"`
		Ops     int64       `bson:"ops"`
	} `bson:"apply"`
	Buffer struct {
		Count        int64 `bson:"count"`
		MaxSizeBytes int   `bson:"maxSizeBytes"`
		SizeBytes    int64 `bson:"sizeBytes"`
	} `bson:"buffer"`
	Network struct {
		Bytes          int64       `bson:"bytes"`
		Getmores       TimedMetric `bson:"getmores"`
		Ops            int64       `bson:"ops"`
		ReadersCreated int64       `bson:"readersCreated"`
	} `bson:"network"`
	Oplog struct {
		Insert      TimedMetric `bson:"insert"`
		InsertBytes int64       `bson:"insertBytes"`
	} `bson:"oplog"`
	Preload struct {
		Docs    TimedMetric `bson:"docs"`
		Indexes TimedMetric `bson:"indexes"`
	} `bson:"preload"`
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
	Repl ReplMetrics `bson:"repl"`
	TTL  TTLMetrics  `bson:"ttl"`
	OK   int         `bson:"ok"`
}

type NetworkStats struct {
	BytesIn     int `bson:"bytesIn"`
	BytesOut    int `bson:"bytesOut"`
	NumRequests int `bson:"numRequests"`
}

type OpcounterStats struct {
	Command int `bson:"command"`
	Delete  int `bson:"delete"`
	Getmore int `bson:"getmore"`
	Insert  int `bson:"insert"`
	Query   int `bson:"query"`
	Update  int `bson:"update"`
}

type DBRecordStats struct {
	AccessesNotInMem    int `bson:"accessesNotInMemory"`
	PageFaultExceptions int `bson:"pageFaultExceptionsThrown"`
}

type RecordStatInfo struct {
	AccessesNotInMem    int           `bson:"accessesNotInMemory"`
	PageFaultExceptions int           `bson:"pageFaultExceptionsThrown"`
	Admin               DBRecordStats `bson:"admin"`
	Local               DBRecordStats `bson:"local"`
	DBs                 bson.M        `bson:",inline"`
}

type ReplStats struct {
	SetName     string   `bson:"setName"`
	SetVersion  int      `bson:"setVersion"`
	IsMaster    bool     `bson:"ismaster"`
	IsSecondary bool     `bson:"secondary"`
	Hosts       []string `bson:"hosts"`
	Primary     string   `bson:"primary"`
	Me          string   `bson:"me"`
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
	OpcountersRepl     OpcounterStats       `bson:"opcountersRepl"`
	PID                int                  `bson:"pid"`
	Process            string               `bson:"process"`
	RecordStats        RecordStatInfo       `bson:"recordStats"`
	Repl               ReplStats            `bson:"repl,omitempty"`
	Uptime             int                  `bson:"uptime"`
	UptimeEstimate     int                  `bson:"uptimeEstimate"`
	UptimeMillis       int                  `bson:"uptimeMillis"`
	Version            string               `bson:"version"`
	WriteBacksQueued   bool                 `bson:"writeBacksQueued"`
	Misc               bson.M               `bson:",inline"`
}

/*func check( err error ) {
	if err != nil {
		panic(err)
	}
}

// test these on a running mongod.
func main() {
	session, err := mgo.Dial("localhost:27017?connect=direct")
	check(err)

	ss := ServerStatus{}
	err = session.Run("serverStatus", &ss)
	check(err)

	fmt.Println("Running " + ss.Process + " version " + ss.Version)
	fmt.Println(ss.LocalTime)
	fmt.Printf("%v currently requires %v MB of RAM.\n", ss.Process, ss.Mem.Resident)
	fmt.Printf("Full memory stats:\n\t%+v\n\n\n", ss.Mem)

	// test some sub-documents, make sure they print
	fmt.Printf("Number of documents returned: %v\n\n", ss.Metrics.Document.Returned)
	fmt.Printf("Number of page faults thrown from admin db: %v\n\n", ss.RecordStats.Admin.PageFaultExceptions)
	fmt.Printf("Time spent in read lock in local db: %v\n\n", ss.Locks.Local.TimeLockedMicros.R)

	// if a replica set, test some stuff
	fmt.Printf("Replica set name: %v\n\n", ss.Repl.SetName)
	fmt.Printf("Replica set members: ")
	for i, host := range ss.Repl.Hosts {
		if i != 0 {
			fmt.Printf(", ")
		}
		fmt.Printf("%v", host)
	}
	fmt.Printf("\n\n")

	fmt.Println(ss)
}*/
