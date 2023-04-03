package meta_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/fm/meta"
	"github.com/xuender/kit/base"
)

func TestInfo_Target(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	info := &meta.Info{Date: base.Result1(time.Parse("20060102150405", "20060102150405"))}

	ass.Equal("aa", info.Target("aa"))
	ass.Equal("[2006]", info.Target("[$yyyy]"))
	ass.Equal("[06]", info.Target("[$Yy]"))
	ass.Equal("[01]", info.Target("[$mm]"))
	ass.Equal("[2]", info.Target("[$d]"))
	ass.Equal("06[2]", info.Target("06[$d]"))
}
