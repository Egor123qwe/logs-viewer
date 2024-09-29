package log

import (
	"context"

	proto "github.com/Egor123qwe/logs-storage/pkg/proto"
	"github.com/Egor123qwe/logs-viewer/api"
	"github.com/Egor123qwe/logs-viewer/internal/model"
	"github.com/Egor123qwe/logs-viewer/internal/util"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Service interface {
	GetLogs(ctx context.Context, req model.LogFilter) (model.LogResp, error)

	GetModules(ctx context.Context, req model.ModuleReq) ([]string, error)
	InitModule(ctx context.Context, module string) (int64, error)
}

type service struct {
	api api.Service
}

func New(api api.Service) Service {
	return &service{
		api: api,
	}
}

func (s service) GetLogs(ctx context.Context, req model.LogFilter) (model.LogResp, error) {
	apiReq := &proto.LogFilter{
		TraceID:     req.TraceID,
		ModuleID:    req.ModuleID,
		Message:     req.Message,
		CountOnPage: req.CountOnPage,
		Page:        req.Page,
	}

	if req.Level != nil {
		apiReq.Level = util.Ptr((*req.Level).String())
	}

	if req.StartTime != nil {
		apiReq.StartTime = timestamppb.New(*req.StartTime)
	}

	if req.EndTime != nil {
		apiReq.EndTime = timestamppb.New(*req.EndTime)
	}

	resp, err := s.api.Log().GetLogs(ctx, apiReq)
	if err != nil {
		return model.LogResp{}, err
	}

	result := model.LogResp{
		Total: resp.PagesCount,
	}

	for _, l := range resp.Logs {
		log := model.Log{
			ID:      l.Id,
			TraceID: l.TraceID,
			Module:  l.Module,
			Level:   l.Level,
			Time:    l.Time.AsTime(),
			Message: l.Message,
		}

		result.Logs = append(result.Logs, log)
	}

	return result, nil
}
