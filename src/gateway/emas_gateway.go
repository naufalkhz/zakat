package gateway

import (
	"context"
	"crypto/tls"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/naufalkhz/zakat/src/models"
	"github.com/tidwall/gjson"
)

const apiURL = "https://logam-mulia-api.vercel.app/prices/anekalogam"

type EmasGateway interface {
	GetHargaEmas(ctx context.Context) (*models.Emas, error)
}

type emasGateway struct {
	client *resty.Client
}

func NewEmasGateway() EmasGateway {
	client := resty.New()
	client.SetDebug(true)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	client.SetBaseURL(apiURL)

	return &emasGateway{
		client: client,
	}
}

func (r *emasGateway) GetHargaEmas(ctx context.Context) (*models.Emas, error) {
	resp, err := r.client.R().SetContext(ctx).Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("error get harga emas: %s", err.Error())
	}

	if resp.Error() != nil {
		return nil, fmt.Errorf("error get harga emas: %s", resp.String())
	}

	return &models.Emas{
		Harga:  gjson.GetBytes(resp.Body(), "data.0.buy").Int(),
		Tipe:   gjson.GetBytes(resp.Body(), "data.0.type").String(),
		Sumber: gjson.GetBytes(resp.Body(), "meta.url").String(),
	}, nil
}
