package midtrans

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"fp_pinjaman_online/model/dto"
	"log"
	"os"

	"github.com/go-resty/resty/v2"
)

type MidtransService interface {
	Pay(payload dto.MidtransSnapRequest) (dto.MidtransSnapResponse, error)
}

type midtransService struct {
	client *resty.Client
	url    string
}

func (m *midtransService) Pay(payload dto.MidtransSnapRequest) (dto.MidtransSnapResponse, error) {

	serverKey := os.Getenv("MIDTRANS_SERVER_KEY")
	encodedKey := base64.StdEncoding.EncodeToString([]byte(serverKey))

	resp, err := m.client.R().
		SetHeader("Authorization", "Basic "+encodedKey).
		SetBody(payload).Post(m.url)

	fmt.Println("resp: ", resp)

	if err != nil {
		log.Println("Error bayar: ", err.Error())
		return dto.MidtransSnapResponse{}, err
	}

	var snapResponse dto.MidtransSnapResponse

	if err := json.Unmarshal(resp.Body(), &snapResponse); err != nil {
		log.Println("Error: ", err.Error())
		return dto.MidtransSnapResponse{}, err
	}

	redirectUrl := fmt.Sprintf("https://app.sandbox.midtrans.com/snap/v2/vtweb/%s", snapResponse.Token)
	snapResponse.RedirectUrl = redirectUrl
	return snapResponse, nil
}

func NewMidtransService(client *resty.Client) MidtransService {
	return &midtransService{
		client: client,
		url:    "https://app.sandbox.midtrans.com/snap/v1/transactions",
	}
}
