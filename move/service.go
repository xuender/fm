package move

import (
	"os"
	"path/filepath"

	"github.com/samber/lo"
	"github.com/xuender/fm/meta"
	"github.com/xuender/fm/pb"
	"github.com/xuender/kit/logs"
	"github.com/xuender/kit/oss"
	"github.com/youthlin/t"
)

type Service struct {
	cfg *pb.Config
	ms  *meta.Service
}

func NewService(
	cfg *pb.Config,
	metaService *meta.Service,
) *Service {
	if len(cfg.Dirs) == 0 {
		panic(t.T("missing .fm.toml file, run fm init"))
	}

	return &Service{
		cfg: cfg,
		ms:  metaService,
	}
}

func (p *Service) Move(paths []string) {
	for _, path := range paths {
		p.move(lo.Must1(oss.Abs(path)))
	}
}

func (p *Service) move(path string) {
	if lo.Contains(p.cfg.Ignore, path) {
		return
	}

	info := p.ms.Info(path)

	if dir, has := p.cfg.Dirs[info.Meta.String()]; has {
		if Move(path, info.Target(dir)) == nil {
			logs.I.Println("mv", path, info.Target(dir))
		}
	}
}

func (p *Service) Scan(paths ...string) {
	for _, path := range paths {
		path = lo.Must1(oss.Abs(path))

		for _, entry := range lo.Must1(os.ReadDir(path)) {
			p.move(filepath.Join(path, entry.Name()))
		}
	}
}
