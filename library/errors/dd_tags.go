package errors

import "fmt"

type DDTags struct {
	funcName string
	pic      string
	repo     string
}

//for anonymous func use this to manually set func name
func (d DDTags) SetFuncName(funcName string) DDTags {
	d.funcName = funcName
	return d
}

func (d DDTags) GetString() string {
	return fmt.Sprintf("func:%s|pic:%s|repo:%s", d.funcName, d.pic, d.repo)
}

func ErrDBDeliveries() DDTags {
	return DDTags{
		repo: "db",
		pic:  "mp-logistic",
	}
}

func ErrDBOrder() DDTags {
	return DDTags{
		repo: "db",
		pic:  "order",
	}
}

func ErrDBCore() DDTags {
	return DDTags{
		repo: "db",
		pic:  "core",
	}
}

func ErrDBLogistic() DDTags {
	return DDTags{
		repo: "db",
		pic:  "logistic",
	}
}

func ErrAPIKero() DDTags {
	return DDTags{
		repo: "api",
		pic:  "kero",
	}
}

func ErrAPIKrab() DDTags {
	return DDTags{
		repo: "api",
		pic:  "mrkrab",
	}
}

func ErrAPISonic() DDTags {
	return DDTags{
		repo: "api",
		pic:  "sonic",
	}
}

func ErrAPIOrder() DDTags {
	return DDTags{
		repo: "api",
		pic:  "order",
	}
}

func ErrAPIElastic() DDTags {
	return DDTags{
		repo: "api",
		pic:  "elastic",
	}
}

func ErrAPITxV2() DDTags {
	return DDTags{
		repo: "api",
		pic:  "txv2",
	}
}

func ErrAPINSQ() DDTags {
	return DDTags{
		repo: "api",
		pic:  "nsq",
	}
}

func ErrRedis(redis string) DDTags {
	return DDTags{
		repo: "redis",
		pic:  redis,
	}
}

func ErrAPITome() DDTags {
	return DDTags{
		repo: "api",
		pic:  "tome",
	}
}

func ErrAPIPayment() DDTags {
	return DDTags{
		repo: "api",
		pic:  "payment",
	}
}
