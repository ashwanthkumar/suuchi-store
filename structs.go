package main

// type Entry struct {
// 	Hash  uint64
// 	Key   []byte
// 	Value []byte
// }

// // Block is a set of Key-Values inside a List. We model the data as blocks so we read
// // and write data in chunks than individual records. We would do sequential reads of
// // each key during Lookup within a block to find the right entry.
// // On Disk, this would of the format:
// // [fixedHashSize][keySize][valueSize][Hash][Key][Value]
// type Block struct {
// 	Entries []Entry
// }

// // List represents an SST file inside a region. SST is a sorted collection of Blocks
// // written where blocks have non-overlapping keys sorted on their hash value.
// // Each List is effectively immutable once it's written to Disk. So we can do lot of
// // interesting things like storing the BloomFilter, hash to block offsets index, etc.
// // at the front of the file and use that for reading rest of the data.
// type List struct {
// 	Id       int    // greater id, means more recent file
// 	Filename string // filename on disk
// 	Start    uint64
// 	End      uint64
// 	Blocks   []Block
// 	// - Add BloomFilter per list to improve read performance
// 	// - Add hash to block offsets for better identifying the right block faster
// }

// // Region represents a hash key range within the Ring. Ring is a logical collection
// // of Lists which might have overlapping keys among them. Every time we flush the
// // memstore, we create a new List and write it to Disk. These files would be consolidated
// // as part of the consolidation process that we run in the background (async).
// // As part of Read, we would query each List in the reverse order of List.Id to see if
// // the given key exist among them.
// type Region struct {
// 	Id    int
// 	Start uint64 // inclusive
// 	End   uint64 // exclusive
// 	Lists []List
// 	// TODO: Add the options for tuning Regions in here
// }

// // Ring contains the whole key space of the Hash Function. Logically this represents
// // the entire DB. Ring can have multiple regions, who have non-overlapping keys spaces
// // that they're responsible for. You can't change the Hasher implementation once the
// // DB is created if you do, you'll loose the ability to query the older data.
// type Ring struct {
// 	RWLock  sync.RWMutex
// 	Hasher  Hasher
// 	Regions []Region
// 	// TODO: Add options for when the Regions would be split
// }

type FileMetadata struct {
	StartKey []byte
	LastKey  []byte
}
