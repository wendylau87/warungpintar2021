package inventory

import (
	"errors"
	"github.com/wendylau87/warungpintar2021/inventorysvc/entities"
)

func(d *domain) createInventory(v entities.Inventory)(entities.Inventory, error){
	totalQty, err := d.sumQuantityByItem(v.ItemID)

	d.logger.LogAccess("Total Current Qty Item[%d] : %d", v.ItemID, totalQty)
	if totalQty + v.Quantity > 1000{
		return entities.Inventory{}, errors.New("Maximum quantity must lower than 1000.")
	}

	// NOTE: this is a transaction example.
	tx, err := d.SQLHandler.Begin()
	if err != nil {
		return v, err
	}

	result, err := tx.Exec(CreateInventory, v.InboundDetailID, v.ItemID, v.Quantity, v.CreatedAt)
	if err != nil {
		_ = tx.Rollback()
		return v, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return v, nil
	}
	v.ID = int(id)

	if err = tx.Commit(); err != nil {
		return v, err
	}

	return v, nil
}

func(d *domain) sumQuantityByItem(id int)(int, error){
	var result int
	rows, err := d.SQLHandler.Query(SumInventoryQuantityByItem, id)

	defer rows.Close()

	if err != nil {
		return result, err
	}

	for rows.Next() {
		if err = rows.Scan(&result); err != nil {
			return result, err
		}
	}
	if err != nil {
		return result, err
	}

	return result, nil
}

func(d *domain) getInventoryByInboundDetail(id int)(entities.Inventory, error){
	var result entities.Inventory
	rows, err := d.SQLHandler.Query(GetInventoryByInboundDetail, id)

	defer rows.Close()

	if err != nil {
		return result, err
	}

	for rows.Next() {
		if err = rows.Scan(&result.ID, &result.InboundDetailID, &result.ItemID, &result.Quantity, &result.CreatedAt); err != nil {
			return result, err
		}
	}
	if err != nil {
		return result, err
	}

	return result, nil
}

