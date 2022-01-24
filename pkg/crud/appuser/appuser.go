package appuser

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/appusermgr"

	constant "github.com/NpoolPlatform/appuser-manager/pkg/const"
	db "github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appuser"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

func validateAppUser(info *npool.AppUser) error {
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		return xerrors.Errorf("invalid app id: %v", err)
	}
	return nil
}

func dbRowToAppUser(row *ent.AppUser) *npool.AppUser {
	return &npool.AppUser{
		ID:            row.ID.String(),
		AppID:         row.AppID.String(),
		EmailAddress:  row.EmailAddress,
		PhoneNO:       row.PhoneNo,
		ImportFromApp: row.ImportFromApp.String(),
	}
}

func Create(ctx context.Context, in *npool.CreateAppUserRequest) (*npool.CreateAppUserResponse, error) {
	if err := validateAppUser(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	importApp, err := uuid.Parse(in.GetInfo().GetImportFromApp())
	if err != nil {
		importApp = uuid.UUID{}
	}

	ctx, cancel := context.WithTimeout(ctx, constant.DBTimeout)
	defer cancel()

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		AppUser.
		Create().
		SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
		SetEmailAddress(in.GetInfo().GetEmailAddress()).
		SetPhoneNo(in.GetInfo().GetPhoneNO()).
		SetImportFromApp(importApp).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create app user: %v", err)
	}

	return &npool.CreateAppUserResponse{
		Info: dbRowToAppUser(info),
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateAppUserRequest) (*npool.UpdateAppUserResponse, error) {
	id, err := uuid.Parse(in.GetInfo().GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app user id: %v", err)
	}

	if err := validateAppUser(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, constant.DBTimeout)
	defer cancel()

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		AppUser.
		UpdateOneID(id).
		SetEmailAddress(in.GetInfo().GetEmailAddress()).
		SetPhoneNo(in.GetInfo().GetPhoneNO()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail update app user: %v", err)
	}

	return &npool.UpdateAppUserResponse{
		Info: dbRowToAppUser(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetAppUserRequest) (*npool.GetAppUserResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app user id: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, constant.DBTimeout)
	defer cancel()

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		AppUser.
		Query().
		Where(
			appuser.And(
				appuser.ID(id),
				appuser.DeleteAt(0),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query app user: %v", err)
	}

	var myAppUser *npool.AppUser
	for _, info := range infos {
		myAppUser = dbRowToAppUser(info)
		break
	}

	return &npool.GetAppUserResponse{
		Info: myAppUser,
	}, nil
}

func GetByApp(ctx context.Context, in *npool.GetAppUsersByAppRequest) (*npool.GetAppUsersByAppResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, constant.DBTimeout)
	defer cancel()

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		AppUser.
		Query().
		Where(
			appuser.AppID(appID),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query app user: %v", err)
	}

	appUsers := []*npool.AppUser{}
	for _, info := range infos {
		appUsers = append(appUsers, dbRowToAppUser(info))
	}

	return &npool.GetAppUsersByAppResponse{
		Infos: appUsers,
	}, nil
}