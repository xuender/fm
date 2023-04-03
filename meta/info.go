package meta

import (
	"fmt"
	"io"
	"regexp"
	"time"

	"github.com/xuender/fm/pb"
	"github.com/youthlin/t"
)

// nolint: gochecknoglobals
var _formats = [...][2]string{
	{"\\$[yY]{4}", "2006"},
	{"\\$[yY]{2}", "06"},
	{"\\$[mM]{2}", "01"},
	{"\\$[mM]{1}", "1"},
	{"\\$[dD]{2}", "02"},
	{"\\$[dD]{1}", "2"},
}

// Info 文件信息.
type Info struct {
	Path  string
	Meta  pb.Meta
	Date  time.Time
	Error error
}

func NewInfoError(path string, err error) *Info {
	return &Info{
		Path:  path,
		Error: err,
	}
}

// Target return dir.
func (p Info) Target(format string) string {
	for _, str := range _formats {
		reg := regexp.MustCompile(str[0])

		if reg.MatchString(format) {
			format = reg.ReplaceAllString(format, p.Date.Format(str[1]))
		}
	}

	return format
}

func (p Info) String() string {
	return fmt.Sprintf("%s: %v %s %v", p.Path, p.Meta, p.Date.Format("2006-01-02 15:04:05"), p.Error)
}

// Output write file meta.
func (p Info) Output(writer io.Writer) {
	fmt.Fprintln(writer, p.Path)

	if p.Error == nil {
		fmt.Fprintln(writer, t.T("Meta: %v", p.Meta))
		fmt.Fprintln(writer, t.T("Time: %s", p.Date.Format("2006-01-02 15:04:05")))

		return
	}

	fmt.Fprintf(writer, "Error: %v\n", p.Error)
}
