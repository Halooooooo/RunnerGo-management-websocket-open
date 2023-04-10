package folder

import (
	"context"
	"fmt"

	"RunnerGo-management/internal/pkg/biz/consts"
	"RunnerGo-management/internal/pkg/biz/record"
	"RunnerGo-management/internal/pkg/dal"
	"RunnerGo-management/internal/pkg/dal/query"
	"RunnerGo-management/internal/pkg/dal/rao"
	"RunnerGo-management/internal/pkg/packer"
)

func Save(ctx context.Context, userID string, req *rao.SaveFolderReq) error {
	target := packer.TransSaveFolderReqToTargetModel(req, userID)
	return query.Use(dal.DB()).Transaction(func(tx *query.Query) error {
		// 接口名排重
		_, err := tx.Target.WithContext(ctx).Where(tx.Target.TeamID.Eq(req.TeamID), tx.Target.Name.Eq(req.Name),
			tx.Target.TargetType.Eq(consts.TargetTypeFolder), tx.Target.TargetID.Neq(req.TargetID),
			tx.Target.Status.Eq(consts.TargetStatusNormal)).First()
		if err == nil {
			return fmt.Errorf("名称已存在")
		}

		// 查询当前接口是否存在
		_, err = tx.Target.WithContext(ctx).Where(tx.Target.TargetID.Eq(req.TargetID)).First()
		if err != nil { // 需新增
			if err := tx.Target.WithContext(ctx).Create(target); err != nil {
				return err
			}
			return record.InsertCreate(ctx, target.TeamID, userID, record.OperationOperateCreateFolder, target.Name)
		}

		if _, err := tx.Target.WithContext(ctx).Where(tx.Target.TargetID.Eq(req.TargetID)).Updates(target); err != nil {
			return err
		}

		return record.InsertUpdate(ctx, target.TeamID, userID, record.OperationOperateUpdateFolder, target.Name)
	})
}

func GetByTargetID(ctx context.Context, teamID string, targetID string) (*rao.Folder, error) {
	tx := query.Use(dal.DB()).Target
	t, err := tx.WithContext(ctx).Where(
		tx.TargetID.Eq(targetID),
		tx.TeamID.Eq(teamID),
		tx.TargetType.Eq(consts.TargetTypeFolder),
		tx.Status.Eq(consts.TargetStatusNormal),
		tx.Source.Eq(consts.TargetSourceNormal),
	).First()

	if err != nil {
		return nil, err
	}

	return packer.TransTargetToRaoFolder(t), nil
}
