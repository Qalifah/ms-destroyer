package handler

import (
	"context"
	"destroyer"
	"destroyer/database"
	"destroyer/lib/pubsub"
	"destroyer/proto"
	"encoding/json"
	"github.com/brianvoe/gofakeit/v6"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const eventName = "targets.acquired"

type handler struct {
	pubSubService pubsub.Service
	store         database.Repository
	proto.UnimplementedDestroyerServer
}

func (h *handler) AcquireTargets(ctx context.Context, _ *proto.TargetParams) (*proto.Response, error) {
	var event destroyer.Event
	if err := gofakeit.Struct(&event); err != nil {
		return nil, err
	}
	event.Name = eventName
	data, err := json.Marshal(event)
	if err != nil {
		return nil, err
	}
	if err := h.pubSubService.Publish(ctx, data); err != nil {
		return nil, err
	}
	return nil, nil
}

func (h *handler) ListTargets(ctx context.Context, _ *proto.TargetParams) (*proto.Response, error) {
	targets, err := h.store.ListTarget(ctx)
	if err != nil {
		return nil, err
	}
	mshTargetColl := marshalTargetColl(targets)
	return &proto.Response{Targets: mshTargetColl}, nil
}

func (h *handler) GetTarget(ctx context.Context, req *proto.TargetParams) (*proto.Response, error) {
	target, err := h.store.GetTarget(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &proto.Response{Target: marshalTarget(target)}, nil
}

func New(ps pubsub.Service, store database.Repository) *handler {
	return &handler{
		pubSubService: ps,
		store:         store,
	}
}

func marshalTargetColl(targets []destroyer.Target) []*proto.Target {
	var result []*proto.Target
	for _, target := range targets {
		result = append(result, marshalTarget(&target))
	}
	return result
}

func marshalTarget(target *destroyer.Target) *proto.Target {
	val := &proto.Target{
		Id:        target.Id,
		Message:   target.Message,
		CreatedOn: timestamppb.New(target.CreatedOn),
	}
	if target.UpdatedOn != nil {
		val.UpdatedOn = timestamppb.New(*target.UpdatedOn)
	}
	return val
}
