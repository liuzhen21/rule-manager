package service

import (
	"context"
	"strings"

	"git.internal.yunify.com/manage/common/log"
	"github.com/tkeel-io/core-broker/pkg/auth"
	"github.com/tkeel-io/core-broker/pkg/core"
	"github.com/tkeel-io/core-broker/pkg/pagination"
	tkeelLog "github.com/tkeel-io/kit/log"
	pb "github.com/tkeel-io/rule-manager/api/rule/v1"
	"github.com/tkeel-io/rule-manager/constant"
	"github.com/tkeel-io/rule-manager/internal/dao"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// Log prefix
const (
	CreatePrefixTag = "[RuleCreate]"
	UpdatePrefixTag = "[RuleUpdate]"
	DeletePrefixTag = "[RuleDelete]"
	QueryPrefixTag  = "[RuleQuery]"
)

type RulesService struct {
	pb.UnimplementedRulesServer
	Core *core.Client
}

func NewRulesService() *RulesService {
	if dao.CoreClient == nil {
		if err := dao.SetCoreClientUp(); err != nil {
			tkeelLog.Fatal("setup core client failed", err)
		}
	}
	return &RulesService{
		Core: dao.CoreClient,
	}
}

func (s *RulesService) RuleCreate(ctx context.Context, req *pb.RuleCreateReq) (res *pb.RuleCreateResp, err error) {
	printInputDebug(CreatePrefixTag, req)
	user, err := auth.GetUser(ctx)
	if err != nil {
		return nil, pb.ErrUnauthorized()
	}
	rule := dao.Rule{
		UserID: user.ID,
		Name:   req.Name,
		Status: constant.RuleStatusStop,
		Desc:   req.Desc,
	}

	result := dao.DB().Model(&rule).Create(&rule)
	if result.Error != nil {
		log.ErrorWithFields(CreatePrefixTag, log.Fields{
			"error": result.Error,
		})
		return nil, pb.ErrInternalError()
	}
	return &pb.RuleCreateResp{
		Id:        uint64(rule.ID),
		Name:      rule.Name,
		Desc:      rule.Desc,
		Status:    uint32(rule.Status),
		Type:      uint32(rule.Type),
		CreatedAt: rule.CreatedAt.Unix(),
		UpdatedAt: rule.UpdatedAt.Unix(),
	}, nil
}

func (s *RulesService) RuleUpdate(ctx context.Context, req *pb.RuleUpdateReq) (*pb.RuleUpdateResp, error) {
	printInputDebug(UpdatePrefixTag, req)
	user, err := auth.GetUser(ctx)
	if err != nil {
		return nil, pb.ErrUnauthorized()
	}
	rule := &dao.Rule{
		Model:  gorm.Model{ID: uint(req.Id)},
		UserID: user.ID,
	}

	var c int
	result := dao.DB().Model(&rule).Select("1").
		Where(&rule).
		First(&c)
	if errors.Is(
		result.Error,
		gorm.ErrRecordNotFound,
	) || result.RowsAffected == 0 {
		return nil, pb.ErrForbidden()
	}

	result = dao.DB().Model(&rule).First(&rule)
	if result.Error != nil {
		tkeelLog.Error(UpdatePrefixTag, result.Error)
		return nil, pb.ErrInternalError()
	}

	rule.Name = req.Name
	rule.Desc = req.Desc

	result = dao.DB().Save(&rule)
	if result.Error != nil {
		return nil, pb.ErrInternalError()
	}

	return &pb.RuleUpdateResp{
		Id:        uint64(rule.ID),
		Name:      rule.Name,
		Desc:      rule.Desc,
		Status:    uint32(rule.Status),
		Type:      uint32(rule.Type),
		CreatedAt: rule.CreatedAt.Unix(),
		UpdatedAt: rule.UpdatedAt.Unix(),
	}, nil
}

func (s *RulesService) RuleDelete(ctx context.Context, req *pb.RuleDeleteReq) (*pb.RuleDeleteResp, error) {
	//print request [debug]
	printInputDebug(DeletePrefixTag, req)
	user, err := auth.GetUser(ctx)
	if err != nil {
		return nil, pb.ErrUnauthorized()
	}
	rule := &dao.Rule{
		Model:  gorm.Model{ID: uint(req.Id)},
		UserID: user.ID,
	}
	result := dao.DB().Model(&rule).Where(&rule).First(&rule)
	if result.Error != nil {
		tkeelLog.Error(DeletePrefixTag, result.Error)
		return nil, pb.ErrForbidden()
	}

	result = dao.DB().Delete(&rule)
	if result.Error != nil {
		tkeelLog.Error(DeletePrefixTag, result.Error)
		return nil, pb.ErrInternalError()
	}

	return &pb.RuleDeleteResp{}, nil
}

func (s *RulesService) RuleGet(ctx context.Context, req *pb.RuleGetReq) (*pb.Rule, error) {
	user, err := auth.GetUser(ctx)
	if err != nil {
		return nil, pb.ErrUnauthorized()
	}
	rule := &dao.Rule{
		Model:  gorm.Model{ID: uint(req.Id)},
		UserID: user.ID,
	}
	if result := rule.Select(); result.Error != nil {
		tkeelLog.Error(QueryPrefixTag, result.Error)
		return nil, pb.ErrInternalError()
	}
	return &pb.Rule{
		Id:        uint64(rule.ID),
		Name:      rule.Name,
		Desc:      rule.Desc,
		Status:    uint32(rule.Status),
		Type:      uint32(rule.Type),
		CreatedAt: rule.CreatedAt.Unix(),
		UpdatedAt: rule.UpdatedAt.Unix(),
	}, nil
}

func (s *RulesService) RuleQuery(ctx context.Context, req *pb.RuleQueryReq) (*pb.RuleQueryResp, error) {
	//print request [debug]
	printInputDebug(QueryPrefixTag, req)
	user, err := auth.GetUser(ctx)
	if err != nil {
		return nil, pb.ErrUnauthorized()
	}

	page, err := pagination.Parse(req)
	if err != nil {
		tkeelLog.Error(QueryPrefixTag, err)
		return nil, pb.ErrInternalError()
	}

	rule := &dao.Rule{UserID: user.ID}
	tx := dao.DB().Model(&rule).Where(&rule)

	fillPagination(tx, page)

	if req.Id != nil && req.Ids != nil && len(req.Ids) > 0 {
		return nil, pb.ErrInvalidArgument()
	}

	if req.Id != nil {
		tx.Where("id = ?", req.Id.Value)
	}

	if req.Ids != nil && len(req.Ids) > 0 {
		tx.Where("id in (?)", req.Ids)
	}

	if req.Name != nil {
		tx.Where("name = ?", req.Name.Value)
	}

	if req.Type != nil {
		tx.Where("type = ?", req.Type.Value)
	}

	if req.Status != nil {
		tx.Where("status = ?", req.Status.Value)
	}

	rules := make([]*dao.Rule, 0)
	result := tx.Find(&rules)
	if result.Error != nil {
		tkeelLog.Error(QueryPrefixTag, result.Error)
		return nil, pb.ErrInternalError()
	}
	var count int64
	result = tx.Count(&count)
	if result.Error != nil {
		tkeelLog.Error(QueryPrefixTag, result.Error)
		return nil, pb.ErrInternalError()
	}

	resp := &pb.RuleQueryResp{}

	page.SetTotal(uint(count))
	if err = page.FillResponse(resp); err != nil {
		tkeelLog.Error(QueryPrefixTag, err)
		return nil, err
	}
	resp.Data = make([]*pb.Rule, len(rules))
	for i, r := range rules {
		resp.Data[i].Id = uint64(r.ID)
		resp.Data[i].Name = r.Name
		resp.Data[i].Desc = r.Desc
		resp.Data[i].Status = uint32(r.Status)
		resp.Data[i].Type = uint32(r.Type)
		resp.Data[i].CreatedAt = r.CreatedAt.Unix()
		resp.Data[i].UpdatedAt = r.UpdatedAt.Unix()
	}
	return resp, nil
}

func (s *RulesService) RuleStatusSwitch(ctx context.Context, req *pb.RuleStatusSwitchReq) (*pb.RuleStatusSwitchResp, error) {
	return nil, errors.New("impl me")
}

func fillPagination(tx *gorm.DB, p pagination.Page) {
	if p.Required() {
		tx.Limit(int(p.Limit())).Offset(int(p.Offset()))
	}
	if p.IsDescending {
		if p.SearchKey != "" && !strings.Contains(p.SearchKey, ",") {
			tx.Order(p.SearchKey + " desc")
		}
	}
}
