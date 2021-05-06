package item

import "github.com/wendylau87/warungpintar2021/mastersvc/entities"

func(d *domain) createItem(v entities.Item)(entities.Item, error){
	// NOTE: this is a transaction example.
	tx, err := d.SQLHandler.Begin()
	if err != nil {
		return v, err
	}

	result, err := tx.Exec(CreateItem, v.Name)
	if err != nil {
		_ = tx.Rollback()
		return v, err
	}

	if err = tx.Commit(); err != nil {
		return v, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return v, nil
	}
	v.ID = int(id)

	return v, nil
}

func(d *domain) findAllItem()([]entities.Item, error){
	items := []entities.Item{}
	rows, err := d.SQLHandler.Query(ReadItem)

	defer rows.Close()

	if err != nil {
		return items, err
	}

	for rows.Next() {
		var result entities.Item
		if err = rows.Scan(&result.ID, &result.Name, &result.Total); err != nil {
			return items, err
		}

		items = append(items, result)
	}

	//if err = rows.Err(); err != nil {
	//	return
	//}

	return items, nil
}