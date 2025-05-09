package assets

import (
	"github.com/Dobefu/csb/cmd/cs_sdk/structs"
	db_structs "github.com/Dobefu/csb/cmd/database/structs"
)

func SetAsset(
	asset structs.Asset,
) error {
	err := queryUpsert("assets", getAssetValues(asset))

	if err != nil {
		return err
	}

	return nil
}

func getAssetValues(asset structs.Asset) []db_structs.QueryValue {
	return []db_structs.QueryValue{
		{
			Name:  "id",
			Value: utilsGenerateId(asset.Uid, asset.Locale),
		},
		{
			Name:  "uid",
			Value: asset.Uid,
		},
		{
			Name:  "title",
			Value: asset.Title,
		},
		{
			Name:  "content_type",
			Value: asset.ContentType,
		},
		{
			Name:  "locale",
			Value: asset.Locale,
		},
		{
			Name:  "url",
			Value: asset.Url,
		},
		{
			Name:  "parent",
			Value: asset.Parent,
		},
		{
			Name:  "version",
			Value: asset.Version,
		},
		{
			Name:  "filesize",
			Value: asset.Filesize,
		},
		{
			Name:  "height",
			Value: asset.Height,
		},
		{
			Name:  "width",
			Value: asset.Width,
		},
		{
			Name:  "updated_at",
			Value: asset.UpdatedAt,
		},
		{
			Name:  "published",
			Value: asset.Published,
		},
	}
}
