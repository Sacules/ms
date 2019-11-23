package ms

import (
	"path/filepath"

	"github.com/adrg/xdg"
	"gitlab.com/Sacules/jsonfile"
)

const (
	app      = "ms"
	savefile = "current.json"
)

var (
	// DataDir is the directoy in which to save the queue
	DataDir = filepath.Join(xdg.DataHome, app)

	// DataPath is the full path for the queue
	DataPath = filepath.Join(DataDir, savefile)
)

// Queue is the overall scheduler, consisting of 4 blocks.
type Queue [4]*Block

// Add a block to the queue and get rid of any old ones.
func (q *Queue) Add(b *Block) {
	for i := len(q) - 1; i > 0; i-- {
		q[i] = q[i-1]
	}

	q[0] = b
}

// Load queue from disk.
func (q *Queue) Load() error {
	err := createDirNotExist(DataDir)
	if err != nil {
		return err
	}

	return jsonfile.Load(q, DataPath)
}

// Replace goes through the queue and replaces an album for a new one.
func (q *Queue) Replace(old, actual Album) {
	for _, block := range q {
		block.Replace(old, actual)
	}
}

// Save queue to disk.
func (q *Queue) Save() error {
	err := createDirNotExist(DataDir)
	if err != nil {
		return err
	}

	return jsonfile.Save(q, DataPath)
}
