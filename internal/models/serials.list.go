package models

import errs "simple-crud-app/internal/lib/errors"

type SerialList struct {
	UserID uint64
	Offset uint64
	Limit  uint64
}

type SerialFull struct {
	SerialData  Serial
	SeasonsData []Season
}

func (fl *SerialList) GetList(db DB) ([]SerialFull, uint64, *errs.Error) {
	// calc total
	var total uint64
	if err := db.QueryRow(`SELECT count(*) FROM serials WHERE user_id=$1`).Scan(&total); err != nil {
		return nil, 0, errs.New().SetCode(errs.ERROR_INTERNAL).SetMsg("unable to count total: %s", err)
	}
	// do sql request for serials
	sqlStr := `SELECT * FROM serials WHERE user_id=$1 OFFSET $2 LIMIT $3`
	rows, err := db.Query(sqlStr, fl.UserID, fl.Offset, fl.Limit)
	if err != nil {
		return nil, 0, errs.New().SetCode(errs.ERROR_INTERNAL).SetMsg("%s", err)
	}
	defer func() {
		_ = rows.Close()
	}()
	// serials handler
	var list []SerialFull
	for rows.Next() {
		tmpFull := SerialFull{} // variable for full one serial data store
		tmpSerial := Serial{}
		if err = rows.Scan(&tmpSerial.ID, &tmpSerial.UserID, &tmpSerial.Name, &tmpSerial.ReleaseDate, &tmpSerial.Score); err != nil {
			return nil, 0, errs.New().SetCode(errs.ERROR_INTERNAL).SetMsg("%s", err)
		}
		tmpFull.SerialData = tmpSerial
		seasons, errDb := fl.getSeasons(db) // get seasons for this serial
		if errDb != nil {
			return nil, 0, errDb
		}
		tmpFull.SeasonsData = seasons
		list = append(list, tmpFull)
	}
	if err = rows.Err(); err != nil {
		return nil, 0, errs.New().SetCode(errs.ERROR_INTERNAL).SetMsg("%s", err)
	}
	return list, total, nil
}

func (fl *SerialList) getSeasons(db DB) ([]Season, *errs.Error) {
	sqlStr := `SELECT * FROM seasons WHERE serial_id = $1`
	rows, err := db.Query(sqlStr, fl.UserID)
	if err != nil {
		return nil, errs.New().SetCode(errs.ERROR_INTERNAL).SetMsg("%s", err)
	}
	defer func() {
		_ = rows.Close()
	}()
	var seasons []Season
	for rows.Next() {
		tmpSeason := Season{}
		if err = rows.Scan(); err != nil {
			return nil, errs.New().SetCode(errs.ERROR_INTERNAL).SetMsg("%s", err)
		}
		seasons = append(seasons, tmpSeason)
	}
	return seasons, nil
}
