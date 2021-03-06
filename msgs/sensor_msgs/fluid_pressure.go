// Autogenerated with msg-import, do not edit.
package sensor_msgs

import (
	"github.com/aler9/goroslib/msgs"
	"github.com/aler9/goroslib/msgs/std_msgs"
)

type FluidPressure struct {
	msgs.Package  `ros:"sensor_msgs"`
	Header        std_msgs.Header
	FluidPressure float64
	Variance      float64
}
