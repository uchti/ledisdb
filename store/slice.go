package store

import (
	"github.com/uchti/ledisdb/store/driver"
)

type Slice interface {
	driver.ISlice
}
