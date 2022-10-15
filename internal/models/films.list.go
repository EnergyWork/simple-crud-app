package models

import errs "simple-crud-app/internal/lib/errors"

type FilmList struct {
	UserID uint64
	Offset uint64
	Limit  uint64
}

func (fl *FilmList) GetList(db DB) ([]Film, uint64, *errs.Error) {
	var total uint64
	if err := db.QueryRow(`SELECT count(*) FROM films WHERE user_id=$1`).Scan(&total); err != nil {
		return nil, 0, errs.New().SetCode(errs.ErrorInternal).SetMsg("unable to count total: %s", err)
	}
	sqlStr := `SELECT * FROM films WHERE user_id=$1 OFFSET $2 LIMIT $3`
	rows, err := db.Query(sqlStr, fl.UserID, fl.Offset, fl.Limit)
	if err != nil {
		return nil, 0, errs.New().SetCode(errs.ErrorInternal).SetMsg("%s", err)
	}
	defer rows.Close()

	list := []Film{}
	for rows.Next() {
		tmp := Film{}
		if err := rows.Scan(&tmp.ID, &tmp.UserID, &tmp.Name, &tmp.ReleaseDate, &tmp.Duration, &tmp.Score); err != nil {
			return nil, 0, errs.New().SetCode(errs.ErrorInternal).SetMsg("%s", err)
		}
		list = append(list, tmp)
	}
	if err = rows.Err(); err != nil {
		return nil, 0, errs.New().SetCode(errs.ErrorInternal).SetMsg("%s", err)
	}

	return list, total, nil
}
