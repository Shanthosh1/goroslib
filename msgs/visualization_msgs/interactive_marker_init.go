// Autogenerated with msg-import, do not edit.
package visualization_msgs

import (
	"github.com/aler9/goroslib/msgs"
)

type InteractiveMarkerInit struct {
	msgs.Package `ros:"visualization_msgs"`
	ServerId     string
	SeqNum       uint64
	Markers      []InteractiveMarker
}
