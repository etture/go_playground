package main

import "fmt"

// LivestreamFieldMask is the field mask of the livestream.
type LivestreamFieldMask uint8

const (
	// LiveStreamFieldMaskProject is the field mask for project of the livestream.
	LiveStreamFieldMaskProject LivestreamFieldMask = 1 << iota

	// LiveStreamFieldMaskStreamerID is the field mask for StreamerID of the livestream.
	LiveStreamFieldMaskStreamerID

	// LiveStreamFieldMaskStreamerType is the field mask for StreamerType of the livestream.
	LiveStreamFieldMaskStreamerType

	// LiveStreamFieldMaskStatus is the field mask for Status of the livestream.
	LiveStreamFieldMaskStatus

	// LiveStreamFieldMaskTitle is the field mask for Title of the livestream.
	LiveStreamFieldMaskTitle

	// LiveStreamFieldMaskPublishTime is the field mask for PublishTime of the livestream.
	LiveStreamFieldMaskPublishTime

	// LiveStreamFieldMaskFinishTime is the field mask for FinishTime of the livestream.
	LiveStreamFieldMaskFinishTime
)

func main() {
	fmt.Println("LiveStreamFieldMaskProject:", LiveStreamFieldMaskProject)
	fmt.Println("LiveStreamFieldMaskStreamerID:", LiveStreamFieldMaskStreamerID)
	fmt.Println("LiveStreamFieldMaskStreamerType:", LiveStreamFieldMaskStreamerType)
	fmt.Println("LiveStreamFieldMaskStatus:", LiveStreamFieldMaskStatus)
	fmt.Println("LiveStreamFieldMaskTitle:", LiveStreamFieldMaskTitle)
	fmt.Println("LiveStreamFieldMaskPublishTime:", LiveStreamFieldMaskPublishTime)
	fmt.Println("LiveStreamFieldMaskFinishTime:", LiveStreamFieldMaskFinishTime)
}
