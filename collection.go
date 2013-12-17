package commgo

// for collection-level statistics

type IndexInfo struct {
	V    int            `bson:"v"`
	Key  map[string]int `bson:"key"`
	Name string         `bson:"name"`
	NS   string         `bson:"ns"`
}

type CollStats struct {
	NS             string         `bson:"ns"`
	Count          int            `bson:"count"`
	Size           int            `bson:"size"`
	AvgObjSize     int            `bson:"avgObjSize"`
	StorageSize    int            `bson:"storageSize"`
	NumExtents     int            `bson:"numExtents"`
	NIndexes       int            `bson:"nindexes"`
	LastExtentSize int            `bson:"lastExtentSize"`
	PaddingFactor  int            `bson:"paddingFactor"`
	SystemFlags    int            `bson:"systemFlags"`
	UserFlags      int            `bson:"userFlags"`
	TotalIndexSize int            `bson:"totalIndexSize"`
	IndexSizes     map[string]int `bson:"indexSizes"`
	Ok             int            `bson:"ok"`
}
