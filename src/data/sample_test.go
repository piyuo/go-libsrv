package data

import (
	"context"
)

type SampleDB interface {
	DB
	SampleTable() *Table
	Counters() *SampleCounters
	Serials() *SampleSerials
	Coders() *SampleCoders
}

// global connection
//
type SampleGlobalDB struct {
	BaseDB
}

func NewSampleGlobalDB(ctx context.Context) (*SampleGlobalDB, error) {
	conn, err := FirestoreGlobalConnection(ctx)
	if err != nil {
		return nil, err
	}
	db := &SampleGlobalDB{
		BaseDB: BaseDB{Connection: conn},
	}
	return db, nil
}

func (db *SampleGlobalDB) SampleTable() *Table {
	return &Table{
		Connection: db.Connection,
		TableName:  "Sample",
		Factory: func() Object {
			return &Sample{}
		},
	}
}

func (db *SampleGlobalDB) Counters() *SampleCounters {
	return &SampleCounters{
		Counters: Counters{
			Connection: db.Connection,
			TableName:  "Count",
		},
	}
}

func (db *SampleGlobalDB) Serials() *SampleSerials {
	return &SampleSerials{
		Serials: Serials{
			Connection: db.Connection,
			TableName:  "Serial",
		},
	}
}

func (db *SampleGlobalDB) Coders() *SampleCoders {
	return &SampleCoders{
		Coders: Coders{
			Connection: db.Connection,
			TableName:  "Code",
		},
	}
}

// regional connection
//
type SampleRegionalDB struct {
	BaseDB
}

func NewSampleRegionalDB(ctx context.Context) (*SampleRegionalDB, error) {
	conn, err := FirestoreRegionalConnection(ctx)
	if err != nil {
		return nil, err
	}
	db := &SampleRegionalDB{
		BaseDB: BaseDB{Connection: conn},
	}
	return db, nil
}

func (db *SampleRegionalDB) SampleTable() *Table {
	return &Table{
		Connection: db.Connection,
		TableName:  "Sample",
		Factory: func() Object {
			return &Sample{}
		},
	}
}

func (db *SampleRegionalDB) Counters() *SampleCounters {
	return &SampleCounters{
		Counters: Counters{
			Connection: db.Connection,
			TableName:  "Count",
		},
	}
}

func (db *SampleRegionalDB) Serials() *SampleSerials {
	return &SampleSerials{
		Serials: Serials{
			Connection: db.Connection,
			TableName:  "Serial",
		},
	}
}

func (db *SampleRegionalDB) Coders() *SampleCoders {
	return &SampleCoders{
		Coders: Coders{
			Connection: db.Connection,
			TableName:  "Code",
		},
	}
}

type PlainObject struct {
	ID   string
	Name string
}

// Sample
//
type Sample struct {
	DomainObject
	Name    string
	Value   int
	Map     map[string]string
	Array   []string
	Numbers []int
	Obj     *PlainObject
}

// SampleCoders  represent collection of code
//
type SampleCoders struct {
	Coders `firestore:"-"`
}

// SampleCoder return sample code
//
func (ss *SampleCoders) SampleCoder() Coder {
	return ss.Coder("SampleCode", 10)
}

// SampleCoder100 return sample code with 100 shards
//
func (ss *SampleCoders) SampleCoder1000() Coder {
	return ss.Coder("SampleCode", 1000)
}

// SampleSerials  represent collection of serial
//
type SampleSerials struct {
	Serials `firestore:"-"`
}

func (ss *SampleSerials) SampleSerial() Serial {
	return ss.Serial("SampleSerial")
}

// SampleCounters represent collection of counter
//
type SampleCounters struct {
	Counters `firestore:"-"`
}

// SampleCounter return sample counter
//
func (scs *SampleCounters) SampleCounter() Counter {
	return scs.Counter("SampleCount", 3, DateHierarchyNone)
}

// SampleCounter100 return sample counter with 100 shards
//
func (scs *SampleCounters) SampleCounter1000() Counter {
	return scs.Counter("SampleCount", 1000, DateHierarchyNone)
}