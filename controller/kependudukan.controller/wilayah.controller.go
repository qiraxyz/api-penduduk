package kependudukan_controller

import (
	"context"
	"datawarehouse/config/database"
	"datawarehouse/helper"
	kependudukan_response "datawarehouse/model/response/kependudukan.response"
	"github.com/gofiber/fiber/v2"
	"github.com/stoewer/go-strcase"
	"strings"
)

func Csql(conditionType, fieldName, value string) string {
	if value == "" {
		return ""
	}
	condition := ""
	if conditionType == "WHERE" {
		condition = " WHERE " + fieldName + " = '" + value + "'"
	} else {
		condition = " AND " + fieldName + " = '" + value + "'"
	}
	return condition
}

func Wilayah(c *fiber.Ctx) error {
	var (
		wilayah     []kependudukan_response.Wilayah
		rowWilayah  kependudukan_response.Wilayah
		regionLevel = strings.ToLower(c.Query("RegionLevel"))
		grouping    = ""
		regionCode  = c.Query("RegionCode")
		cond        = ""
		count       = 0
	)
	db, _ := database.DBConnection()
	defer db.Close()
	ctx := context.Background()

	if regionLevel == "kota" || regionLevel == "kabupaten" {
		cond = Csql("WHERE", "kd_propinsi", regionCode)
	} else if regionLevel == "kecamatan" {
		cond = Csql("WHERE", "kd_kabupaten", regionCode)
	} else if regionLevel == "kelurahan" {
		cond = Csql("WHERE", "kd_kecamatan", regionCode)
	}

	qry, err := db.QueryContext(ctx, `
    SELECT kd_propinsi, kd_kabupaten, kd_kecamatan, kd_kelurahan, propinsi, kabupaten, kecamatan, kelurahan
    FROM mapwil
    `+cond+grouping+`;`)
	if err != nil {
		res := helper.GetResponse(500, nil, err)
		return c.Status(res.Status).JSON(res)
	}
	defer qry.Close()
	for qry.Next() {
		count++
		err := qry.Scan(
			&rowWilayah.KodeProvinsi,
			&rowWilayah.KodeKabupaten,
			&rowWilayah.KodeKecamatan,
			&rowWilayah.KodeKelurahan,
			&rowWilayah.Provinsi,
			&rowWilayah.Kabupaten,
			&rowWilayah.Kecamatan,
			&rowWilayah.Kelurahan)
		if err != nil {
			res := helper.GetResponse(500, nil, err)
			return c.Status(res.Status).JSON(res)
		}

		wilayah = append(wilayah, rowWilayah)
	}

	camelCaseRegionLevel := strcase.UpperCamelCase(regionLevel)
	res := helper.GetResponse(200, fiber.Map{
		"Jumlah" + camelCaseRegionLevel: count,
		"Wilayah":                       wilayah,
	}, nil)
	return c.JSON(res)
}
