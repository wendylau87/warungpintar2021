package inbound

import "github.com/wendylau87/warungpintar2021/transactionsvc/entities"

func(d *domain) createInbound(v entities.Inbound)(entities.Inbound, error){
	// NOTE: this is a transaction example.
	tx, err := d.SQLHandler.Begin()
	if err != nil {
		return v, err
	}

	result, err := tx.Exec(CreateInbound, v.PONumber, v.CreatedAt)
	if err != nil {
		_ = tx.Rollback()
		return v, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return v, nil
	}
	v.ID = int(id)

	for key, detail := range v.Details{
		res, err := tx.Exec(CreateInboundDetail, v.ID, detail.ItemID, detail.Quantity)
		if err != nil {
			_ = tx.Rollback()
			return v, err
		}
		detailID, err := res.LastInsertId()
		if err != nil {
			return v, nil
		}
		v.Details[key].ID = int(detailID)
		v.Details[key].InboundID = v.ID
	}

	if err = tx.Commit(); err != nil {
		return v, err
	}

	return v, nil
}
