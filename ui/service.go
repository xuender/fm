package ui

import (
	"fmt"
	"os"
	"sort"

	"github.com/manifoldco/promptui"
	"github.com/pelletier/go-toml/v2"
	"github.com/samber/lo"
	"github.com/xuender/fm/pb"
	"github.com/xuender/kit/base"
	"github.com/xuender/kit/logs"
	"github.com/youthlin/t"
)

type Service struct {
	cfg *pb.Config
}

func NewService(cfg *pb.Config) *Service {
	logs.D.Println("new", cfg)

	return &Service{cfg: cfg}
}

func (p *Service) Init() {
	for {
		if len(p.cfg.Dirs) == len(pb.Meta_name) {
			logs.W.Println(t.T("No configurable meta."))

			return
		}

		meta := p.SelectMeta()
		p.cfg.Dirs[meta.String()] = p.Prompt(t.T("Input %v to dir", meta), fmt.Sprintf("~/%v/$yyyy/$mm/$dd", meta))

		if len(p.cfg.Dirs) == len(pb.Meta_name) || !p.Confirm(t.T("Continue input")) {
			break
		}
	}

	logs.I.Println(p.cfg)

	p.Save()
}

func (p *Service) Save() {
	file := lo.Must1(os.Create(pb.ConfigFile))
	defer file.Close()

	encoder := toml.NewEncoder(file)
	_ = encoder.Encode(p.cfg)

	logs.I.Println(t.T("save: %s", pb.ConfigFile))
}

func (p *Service) SelectMeta() pb.Meta {
	keys := make([]string, 0, len(pb.Meta_value))
	for k := range pb.Meta_value {
		keys = append(keys, k)
	}

	keys2 := make([]string, 0, len(p.cfg.Dirs))
	for k := range p.cfg.Dirs {
		keys2 = append(keys2, k)
	}

	items := base.Result1(lo.Difference(keys, keys2))
	sort.Strings(items)

	logs.D.Println(p.cfg, keys2, items)

	prompt := promptui.Select{
		Label: t.T("Select file type"),
		Items: items,
	}

	_, key, err := prompt.Run()
	if err != nil {
		panic(err)
	}

	return pb.Meta(pb.Meta_value[key])
}

func (p *Service) Prompt(label, def string) string {
	prompt := promptui.Prompt{
		Label:   label,
		Default: def,
		Validate: func(s string) error {
			if s == "" {
				return ErrDirEmpty
			}

			return nil
		},
	}

	return lo.Must1(prompt.Run())
}

func (p *Service) Confirm(label string) bool {
	prompt := promptui.Prompt{
		Label:     label,
		IsConfirm: true,
	}
	_, err := prompt.Run()

	return err == nil
}
