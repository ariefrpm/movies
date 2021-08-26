package log

import (
	"errors"
	"testing"

	"github.com/tokopedia/logistic/svc/mp-logistic/infra/config"
	e "github.com/tokopedia/logistic/svc/mp-logistic/infra/errors"
)

func TestLog(t *testing.T) {
	cfg := config.LoggerCfg{
		JsonFormat: false,
		Level:      5,
	}
	InitLogger(cfg)
	cfg.JsonFormat = true
	InitLogger(cfg)
	InitMockLogger()

	Info("i", WithMessage("m"), WithField("t", "tes"))
	Infof("i")
	Debug("d", WithFieldf("t", "%s", "f"), WithDeliveryID(1))
	Debugf("d")

	err := errors.New("err")
	Errorw(err)
	Errorw(e.Wrap(err), WithOrderID(1), WithHistoryID(2))
	Errorw(e.Wrap(err), WithMessagef("%s,%d", "a", 1))

	Errors("error", WithOrderID(2))
}
