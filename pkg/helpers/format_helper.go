package helpers

import (
	"fmt"
	"ragamaya-api/pkg/exceptions"
	"time"

	"github.com/midtrans/midtrans-go"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func FormatFileSize(size int64) string {
	if size < 1024 {
		return fmt.Sprintf("%d B", size)
	} else if size < 1024*1024 {
		return fmt.Sprintf("%d KB", size/1024)
	}
	return fmt.Sprintf("%d MB", size/(1024*1024))
}

func FormatIndonesianTime(time time.Time) string {
	return time.Format("January 2, 2006 • 15:04 PM")
}

func FormatIndonesianLocaleString(value uint) string {
	p := message.NewPrinter(language.Indonesian)
	return p.Sprintf("%d", value)
}

func FormatMidtransErrorToException(err *midtrans.Error) *exceptions.Exception {
	return exceptions.NewException(err.StatusCode, err.Message)
}
