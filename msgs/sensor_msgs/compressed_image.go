// Autogenerated with msg-import, do not edit.
package sensor_msgs

import (
	"github.com/aler9/goroslib/msg"
	"github.com/aler9/goroslib/msgs/std_msgs"
)

type CompressedImage struct {
	Header std_msgs.Header
	Format msg.String
	Data   []msg.Uint8
}