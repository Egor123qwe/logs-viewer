package log

import (
	"context"

	proto "github.com/Egor123qwe/logs-storage/pkg/proto"
	//proto "github.com/Egor123qwe/logs-storage/pkg/proto"
	"github.com/Egor123qwe/logs-viewer/internal/model"
)

func (s service) InitModule(ctx context.Context, module string) (int64, error) {
	resp, err := s.api.Log().InitModule(ctx, &proto.InitModuleReq{Module: module})
	if err != nil {
		return 0, err
	}

	return resp.ModuleId, nil
}

func (s service) GetModules(ctx context.Context, req model.ModuleReq) ([]string, error) {
	resp, err := s.api.Log().GetModules(ctx, &proto.ModuleReq{NameFilter: req.NameFilter})
	if err != nil {
		return nil, err
	}

	return resp.Modules, nil
}
