package helper

import (
	"net/url"
	"runtime"
)

func JoinURL(baseUrlStr, pathUrlStr string) (joinedUrlStr string, err error) {
	baseUrl, err := url.Parse(baseUrlStr)
	if err != nil {
		return
	}

	pathUrl, err := url.Parse(pathUrlStr)
	if err != nil {
		return
	}

	joinedUrlStr = baseUrl.ResolveReference(pathUrl).String()

	return
}

func GetFrame(skipFrames int) runtime.Frame {
	// We need the frame at index skipFrames+2, since we never want runtime.Callers and GetFrame
	targetFrameIndex := skipFrames + 2

	// Set size to targetFrameIndex+2 to ensure we have room for one more caller than we need
	programCounters := make([]uintptr, targetFrameIndex+2)
	n := runtime.Callers(0, programCounters)

	frame := runtime.Frame{Function: "unknown"}
	if n > 0 {
		frames := runtime.CallersFrames(programCounters[:n])
		frame = getTargetFrame(frames, targetFrameIndex)
	}

	return frame
}

func getTargetFrame(frames *runtime.Frames, targetFrameIndex int) runtime.Frame {
	frame := runtime.Frame{Function: "unknown"}

	for more, frameIndex := true, 0; more && frameIndex <= targetFrameIndex; frameIndex++ {
		var frameCandidate runtime.Frame
		frameCandidate, more = frames.Next()

		if frameIndex == targetFrameIndex {
			frame = frameCandidate
		}
	}

	return frame
}
