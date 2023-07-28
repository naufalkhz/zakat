package src

import (
	"os"

	"github.com/naufalkhz/zakat/src/gateway"
	"github.com/naufalkhz/zakat/src/services"
	"github.com/robfig/cron/v3"
)

func CronTask(gatewayEmas gateway.EmasGateway, serviceEmas services.EmasService) {
	c := cron.New()

	c.AddFunc(os.Getenv("SCHEDULE_GET_HARGA_EMAS"), serviceEmas.InquryHargaEmas)

	c.Start()
}
