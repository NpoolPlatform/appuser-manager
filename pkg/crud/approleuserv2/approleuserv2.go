package approleuserv2

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/approleuser"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/appusermgrv2/approleuser"
	"github.com/google/uuid"
)

func Create(ctx context.Context, in *npool.AppRoleUser) (*ent.AppRoleUser, error) {
	var info *ent.AppRoleUser
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		c := cli.AppRoleUser.Create()
		if in.AppID != nil {
			c.SetAppID(uuid.MustParse(in.GetAppID()))
		}
		if in.RoleID != nil {
			c.SetRoleID(uuid.MustParse(in.GetRoleID()))
		}
		if in.UserID != nil {
			c.SetUserID(uuid.MustParse(in.GetUserID()))
		}
		if in.ID != nil {
			c.SetID(uuid.MustParse(in.GetID()))
		}
		info, err = c.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func CreateBulk(ctx context.Context, in []*npool.AppRoleUser) ([]*ent.AppRoleUser, error) {
	rows := []*ent.AppRoleUser{}
	var err error
	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		bulk := make([]*ent.AppRoleUserCreate, len(in))
		for i, info := range in {
			bulk[i] = tx.AppRoleUser.Create()
			if info.AppID != nil {
				bulk[i].SetAppID(uuid.MustParse(info.GetAppID()))
			}
			if info.RoleID != nil {
				bulk[i].SetRoleID(uuid.MustParse(info.GetRoleID()))
			}
			if info.UserID != nil {
				bulk[i].SetUserID(uuid.MustParse(info.GetUserID()))
			}
			if info.ID != nil {
				bulk[i].SetID(uuid.MustParse(info.GetID()))
			}
		}
		rows, err = tx.AppRoleUser.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func Update(ctx context.Context, in *npool.AppRoleUser) (*ent.AppRoleUser, error) {
	var info *ent.AppRoleUser
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		u := cli.AppRoleUser.UpdateOneID(uuid.MustParse(in.GetID()))
		if in.AppID != nil {
			u.SetAppID(uuid.MustParse(in.GetAppID()))
		}
		if in.RoleID != nil {
			u.SetRoleID(uuid.MustParse(in.GetRoleID()))
		}
		if in.UserID != nil {
			u.SetUserID(uuid.MustParse(in.GetUserID()))
		}
		info, err = u.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func Row(ctx context.Context, id uuid.UUID) (*ent.AppRoleUser, error) {
	var info *ent.AppRoleUser
	var err error
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.AppRoleUser.Query().Where(approleuser.ID(id)).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

//nolint
func setQueryConds(conds *npool.Conds, cli *ent.Client) (*ent.AppRoleUserQuery, error) {
	stm := cli.AppRoleUser.Query()
	if conds.ID != nil {

		id := uuid.MustParse(conds.GetID().GetValue())
		switch conds.GetID().GetOp() {
		case cruder.EQ:
			stm.Where(approleuser.ID(id))

		case cruder.IN:
			stm.Where(approleuser.ID(id))

		default:
			return nil, fmt.Errorf("invalid approleuser field")
		}

	}
	if conds.AppID != nil {

		appID := uuid.MustParse(conds.GetAppID().GetValue())
		switch conds.GetAppID().GetOp() {
		case cruder.EQ:
			stm.Where(approleuser.AppID(appID))

		case cruder.IN:
			stm.Where(approleuser.AppID(appID))

		default:
			return nil, fmt.Errorf("invalid approleuser field")
		}

	}
	if conds.RoleID != nil {

		roleID := uuid.MustParse(conds.GetRoleID().GetValue())
		switch conds.GetRoleID().GetOp() {
		case cruder.EQ:
			stm.Where(approleuser.RoleID(roleID))

		case cruder.IN:
			stm.Where(approleuser.RoleID(roleID))

		default:
			return nil, fmt.Errorf("invalid approleuser field")
		}

	}
	if conds.UserID != nil {

		userID := uuid.MustParse(conds.GetUserID().GetValue())
		switch conds.GetUserID().GetOp() {
		case cruder.EQ:
			stm.Where(approleuser.UserID(userID))

		case cruder.IN:
			stm.Where(approleuser.UserID(userID))

		default:
			return nil, fmt.Errorf("invalid approleuser field")
		}

	}
	return stm, nil
}

func Rows(ctx context.Context, conds *npool.Conds, offset, limit int) ([]*ent.AppRoleUser, int, error) {
	rows := []*ent.AppRoleUser{}
	var total int
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := setQueryConds(conds, cli)
		if err != nil {
			return err
		}
		total, err = stm.Count(_ctx)
		if err != nil {
			return err
		}

		rows, err = stm.
			Offset(offset).
			Order(ent.Desc(approleuser.FieldUpdatedAt)).
			Limit(limit).
			All(_ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func RowOnly(ctx context.Context, conds *npool.Conds) (*ent.AppRoleUser, error) {
	var info *ent.AppRoleUser

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := setQueryConds(conds, cli)
		if err != nil {
			return err
		}

		info, err = stm.Only(_ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func Count(ctx context.Context, conds *npool.Conds) (uint32, error) {
	var err error
	var total int

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := setQueryConds(conds, cli)
		if err != nil {
			return err
		}

		total, err = stm.Count(_ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return 0, err
	}

	return uint32(total), nil
}

func Exist(ctx context.Context, id uuid.UUID) (bool, error) {
	var err error
	exist := false

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.AppRoleUser.Query().Where(approleuser.ID(id)).Exist(_ctx)
		return err
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}

func ExistConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	var err error
	exist := false

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := setQueryConds(conds, cli)
		if err != nil {
			return err
		}

		exist, err = stm.Exist(_ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}

func Delete(ctx context.Context, id uuid.UUID) (*ent.AppRoleUser, error) {
	var info *ent.AppRoleUser
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.AppRoleUser.UpdateOneID(id).
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}