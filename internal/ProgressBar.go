package internal

import (
	"github.com/cheggaaa/pb/v3"
	"time"
)

func Progress(count int) {
	bar := pb.StartNew(count).SetRefreshRate(time.Millisecond * 10)

	for i := 0; i < count; i++ {
		bar.Increment()
		time.Sleep(time.Millisecond)
	}

	bar.Finish()
}
