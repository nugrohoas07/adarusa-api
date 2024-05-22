package debiturRepository

import (
	"database/sql"
	"errors"
	"fp_pinjaman_online/model/dto"
	"fp_pinjaman_online/model/dto/debiturDto"
	"fp_pinjaman_online/model/dto/json"
	"fp_pinjaman_online/src/debitur"
	"fp_pinjaman_online/src/midtrans"

	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
)

type debiturRepository struct {
	db *sql.DB
}

func NewDebiturRepository(db *sql.DB) debitur.DebiturRepository {
	return &debiturRepository{db}
}

func (d *debiturRepository) AddPengajuanPinjaman(id int, jumlahPinjaman float64, tenor int, description string) error {
	var sukuBunga float64
	if tenor < 3 {
		sukuBunga = 0.05
	} else if tenor >= 3 && tenor < 6 {
		sukuBunga = 0.07
	} else if tenor >= 6 && tenor < 9 {
		sukuBunga = 0.09
	} else if tenor >= 9 && tenor <= 12 {
		sukuBunga = 0.11
	}

	//kurang user_id
	_, err := d.db.Exec("INSERT INTO pinjaman (user_id, jumlah_pinjaman, tenor, description, bunga_per_bulan, status_pengajuan) VALUES ($1, $2, $3, $4, $5, 'pending')", id, jumlahPinjaman, tenor, description, sukuBunga)
	if err != nil {
		return err
	}
	return nil
}

func (d *debiturRepository) GetPengajuanPinjaman(id int) ([]debiturDto.GetPengajuanResponse, error) {
	var data []debiturDto.GetPengajuanResponse
	rows, err := d.db.Query("SELECT id, jumlah_pinjaman, tenor, description, bunga_per_bulan, status_pengajuan, created_at, updated_at FROM pinjaman WHERE user_id = $1 AND deleted_at IS NULL ORDER BY created_at DESC", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var pengajuan debiturDto.GetPengajuanResponse
		if err := rows.Scan(&pengajuan.Id, &pengajuan.JumlahPinjaman, &pengajuan.Tenor, &pengajuan.Description, &pengajuan.BungaPerBulan, &pengajuan.StatusPengajuan, &pengajuan.CreatedAt, &pengajuan.UpdatedAt); err != nil {
			return nil, err
		}
		data = append(data, pengajuan)
	}
	return data, nil
}

func (d *debiturRepository) GetCicilan(page, limit, offset int, id string, status string) ([]debiturDto.GetCicilanResponse, json.Paging, error) {
	var paging json.Paging
	var totalData int
	err := d.db.QueryRow("SELECT COUNT(*) FROM cicilan WHERE pinjaman_id = $1", id).Scan(&totalData)
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, json.Paging{}, err
	}

	paging.Page = page
	paging.TotalData = totalData

	var rows *sql.Rows
	if status == "" {
		rows, err = d.db.Query("SELECT id, pinjaman_id, tanggal_jatuh_tempo, tanggal_selesai_bayar, jumlah_bayar, status FROM cicilan WHERE pinjaman_id = $1", id)
	} else {
		rows, err = d.db.Query("SELECT id, pinjaman_id, tanggal_jatuh_tempo, tanggal_selesai_bayar, jumlah_bayar, status FROM cicilan WHERE pinjaman_id = $1 AND status = $2", id, status)
	}
	var data []debiturDto.GetCicilanResponse
	if err != nil {
		return nil, json.Paging{}, err
	}
	defer rows.Close()
	for rows.Next() {
		var cicilan debiturDto.GetCicilanResponse
		if err := rows.Scan(&cicilan.Id, &cicilan.PinjamanId, &cicilan.TanggalJatuhTempo, &cicilan.TanggalBayarSelesai, &cicilan.JumlahBayar, &cicilan.Status); err != nil {
			return nil, json.Paging{}, err
		}
		data = append(data, cicilan)
	}

	return data, paging, nil
}

func (d *debiturRepository) CicilanPayment(pinjamanId int, totalBayar float64) error {

	var cicilanId int
	var jumlahBayar float64
	err := d.db.QueryRow("SELECT id, jumlah_bayar FROM cicilan WHERE pinjaman_id = $1 AND status = 'unpaid'", pinjamanId).Scan(&cicilanId, &jumlahBayar)
	if err != nil {
		return errors.New("anda tidak mempunyai cicilan")
	}

	if totalBayar != jumlahBayar {
		return errors.New("total bayar tidak sesuai")
	}

	client := resty.New()
	midtransService := midtrans.NewMidtransService(client)

	payload := dto.MidtransSnapRequest{
		TransactionDetails: struct {
			OrderID  int     `json:"order_id"`
			GrossAmt float64 `json:"gross_amount"`
		}{OrderID: cicilanId, GrossAmt: totalBayar},
		PaymentType: "gopay",
		Customer:    "user_id->name", //belum setup
	}

	_, err = midtransService.Pay(payload)
	if err != nil {
		log.Error().Msg(err.Error())
		return err
	}

	_, err = d.db.Exec("UPDATE cicilan SET status = 'paid' WHERE id = $1", cicilanId)
	if err != nil {
		return err
	}

	return err
}
