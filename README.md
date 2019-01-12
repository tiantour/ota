# ota
a golang sdk for ota such as mafengwo, ctrip, fliggy

## how to use

mafengwo

```
package main

import (
	"fmt"
	"log"

	"github.com/tiantour/ota/mafengwo"
	"github.com/tiantour/ota/mafengwo/order"
	"github.com/tiantour/ota/mafengwo/product"
)

func init() {
	// set conf
	mafengwo.ClientID = 0
	mafengwo.ClientSecret = "input your client secret"
	mafengwo.PartnerID = 0
	mafengwo.AseKey = "input your ase key"

	// get token
	auth, err := mafengwo.NewOauth2().Token()
	if err != nil {
		log.Fatal(err)
	}

	// set token
	mafengwo.AccessToken = auth.AccessToken
}

func main() {

	x, err := order.NewOrder().List(1, 10)
	fmt.Println(x, err)

	y, err := product.NewProduct().List(1, 10)
	fmt.Println(y, err)
}

```

