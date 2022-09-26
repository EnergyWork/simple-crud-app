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
	// do sql request
	sqlStr := `SELECT * FROM serials WHERE user_id=$1 OFFSET $2 LIMIT $3`
	rows, err := db.Query(sqlStr, fl.UserID, fl.Offset, fl.Limit)
	if err != nil {
		return nil, 0, errs.New().SetCode(errs.ERROR_INTERNAL).SetMsg("%s", err)
	}
	defer rows.Close()
	// result handler
	list := []SerialFull{}
	for rows.Next() {
		tmp := SerialFull{}
		if err := rows.Scan(&tmp.ID, &tmp.UserID, &tmp.Name, &tmp.ReleaseDate, &tmp.Score); err != nil {
			return nil, 0, errs.New().SetCode(errs.ERROR_INTERNAL).SetMsg("%s", err)
		}
		list = append(list, tmp)
	}
	if err = rows.Err(); err != nil {
		return nil, 0, errs.New().SetCode(errs.ERROR_INTERNAL).SetMsg("%s", err)
	}
	return list, total, nil
}
