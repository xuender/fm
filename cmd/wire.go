//go:build wireinject
// +build wireinject

package cmd

import (
	"github.com/google/wire"
	"github.com/spf13/cobra"
	"github.com/xuender/fm/meta"
	"github.com/xuender/fm/move"
	"github.com/xuender/fm/pb"
	"github.com/xuender/fm/ui"
)

func InitMeta(cmd *cobra.Command) *meta.Service {
	wire.Build(
		meta.NewService,
	)

	return &meta.Service{}
}

func InitUI(cmd *cobra.Command) *ui.Service {
	wire.Build(
		ui.NewService,
		pb.NewConfig,
	)

	return &ui.Service{}
}

func InitMove(cmd *cobra.Command) *move.Service {
	wire.Build(
		meta.NewService,
		pb.NewConfig,
		move.NewService,
	)

	return &move.Service{}
}
