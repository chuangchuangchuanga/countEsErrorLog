package main

import (
	"context"
	"encoding/json"
	"fmt"
	e "github.com/olivere/elastic/v7"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"strings"
	"time"
)



func getESData() []string {
	client , err := e.NewClient(e.SetSniff(false), e.SetURL("http://192.168.2.22:1234"))
	if err != nil {
	// Handle error
		fmt.Println(err)
	}


	timeQuery := e.NewRangeQuery("@timestamp").Gte("now-1m").Lte("now")
	matchQuery := e.NewMatchQuery("message","error")

	generalQuery := e.NewBoolQuery()
	generalQuery = generalQuery.Must(timeQuery).Must(matchQuery)

	resule,err := client.Search().Index("logstash-platform").Query(generalQuery).Size(1000).Do(context.Background())
	if err != nil {
		fmt.Println("faild")
		fmt.Println(err)
	}else {
		fmt.Println(resule.TotalHits())
	}

	hits := resule.Hits.Hits
	var pathList []string
	for _,b := range hits {
		//fmt.Println(b)
		jsonStr, err := b.Source.MarshalJSON()
		if err != nil{
			fmt.Println(err)
		}
		x := map[string]string{}

		json.Unmarshal(jsonStr, &x)

		pathList = append(pathList,x["path"] )
	}

	return pathList
}

func getLog()  {

	for _, i := range getESData(){
		//fmt.Println(i)
		a := strings.Split(i, "/")
		b := a[len(a)-1:]
		c := strings.Split(b[0], ".")


	switch c[0] {
	case "rental-admin":
		rentaladmin.Add(1)
	case "rental-asset":
		rentalasset.Add(1)
	case "rental-auth":
		rentalauth.Add(1)
	case "rental-bizsupport":
		rentalbizsupport.Add(1)
	case "rental-config":
		rentalconfig.Add(1)
	case "rental-device":
		rentaldevice.Add(1)
	case "rental-discovery":
		rentaldiscovery.Add(1)
	case "rental-esign":
		rentalesign.Add(1)
	case "rental-file":
		rentalfile.Add(1)
	case "rental-gateway":
		rentalgateway.Add(1)
	case "rental-gps":
		rentalgps.Add(1)
	case "rental-leasing":
		rentalleasing.Add(1)
	case "rental-logistics":
		rentallogistics.Add(1)
	case "rental-maintain":
		rentalmaintain.Add(1)
	case "rental-mall":
		rentalmall.Add(1)
	case "rental-marketing":
		rentalmarketing.Add(1)
	case "rental-metadata":
		rentalmetadata.Add(1)
	case "rental-mini":
		rentalmini.Add(1)
	case "rental-notification":
		rentalnotification.Add(1)
	case "rental-operation":
		rentaloperation.Add(1)
	case "rental-paas":
		rentalpaas.Add(1)
	case "rental-payment":
		rentalpayment.Add(1)
	case "rental-settle":
		rentalsettle.Add(1)
	case "rental-usercenter":
		rentalusercenter.Add(1)
	case "rental-workflow":
		rentalworkflow.Add(1)
	default:
		nullll.Add(1)
	}

}
}


func recordMetrics()  {
	for true {
		go getLog()
		time.Sleep(10 * time.Second)
	}




}

var (
	rentaladmin = promauto.NewCounter(prometheus.CounterOpts{
		Name: "rentaladmin",
		Help: "The total number of processed events",
	})

	rentalasset = promauto.NewCounter(prometheus.CounterOpts{
		Name: "rentalasset",
		Help: "The total number of processed events",
	})

	rentalauth = promauto.NewCounter(prometheus.CounterOpts{
		Name: "rentalauth",
		Help: "The total number of processed events",
	})

	rentalbizsupport = promauto.NewCounter(prometheus.CounterOpts{
		Name: "rentalbizsupport",
		Help: "The total number of processed events",
	})

	rentalconfig = promauto.NewCounter(prometheus.CounterOpts{
		Name: "rentalconfig",
		Help: "The total number of processed events",
	})

	rentaldevice = promauto.NewCounter(prometheus.CounterOpts{
		Name: "rentaldevice",
		Help: "The total number of processed events",
	})

	rentaldiscovery = promauto.NewCounter(prometheus.CounterOpts{
		Name: "rentaldiscovery",
		Help: "The total number of processed events",
	})

	rentalesign = promauto.NewCounter(prometheus.CounterOpts{
		Name: "rentalesign",
		Help: "The total number of processed events",
	})

	rentalfile = promauto.NewCounter(prometheus.CounterOpts{
		Name: "rentalfile",
		Help: "The total number of processed events",
	})

	rentalgateway = promauto.NewCounter(prometheus.CounterOpts{
		Name: "rentalgateway",
		Help: "The total number of processed events",
	})

	rentalgps = promauto.NewCounter(prometheus.CounterOpts{
		Name: "rentalgps",
		Help: "The total number of processed events",
	})

	rentalleasing = promauto.NewCounter(prometheus.CounterOpts{
		Name: "rentalleasing",
		Help: "The total number of processed events",
	})

	rentallogistics = promauto.NewCounter(prometheus.CounterOpts{
		Name: "rentallogistics",
		Help: "The total number of processed events",
	})

	rentalmaintain = promauto.NewCounter(prometheus.CounterOpts{
		Name: "rentalmaintain",
		Help: "The total number of processed events",
	})

	rentalmall = promauto.NewCounter(prometheus.CounterOpts{
		Name: "rentalmall",
		Help: "The total number of processed events",
	})

	rentalmarketing = promauto.NewCounter(prometheus.CounterOpts{
		Name: "rentalmarketing",
		Help: "The total number of processed events",
	})

	rentalmetadata = promauto.NewCounter(prometheus.CounterOpts{
		Name: "rentalmetadata",
		Help: "The total number of processed events",
	})

	rentalmini = promauto.NewCounter(prometheus.CounterOpts{
		Name: "rentalmini",
		Help: "The total number of processed events",
	})

	rentalnotification = promauto.NewCounter(prometheus.CounterOpts{
		Name: "rentalnotification",
		Help: "The total number of processed events",
	})

	rentaloperation = promauto.NewCounter(prometheus.CounterOpts{
		Name: "rentaloperation",
		Help: "The total number of processed events",
	})

	rentalpaas = promauto.NewCounter(prometheus.CounterOpts{
		Name: "rentalpaas",
		Help: "The total number of processed events",
	})

	rentalpayment = promauto.NewCounter(prometheus.CounterOpts{
		Name: "rentalpayment",
		Help: "The total number of processed events",
	})

	rentalsettle = promauto.NewCounter(prometheus.CounterOpts{
		Name: "rentalsettle",
		Help: "The total number of processed events",
	})

	rentalusercenter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "rentalusercenter",
		Help: "The total number of processed events",
	})


	rentalworkflow = promauto.NewCounter(prometheus.CounterOpts{
		Name: "rentalworkflow",
		Help: "The total number of processed events",
	})
	nullll =  promauto.NewCounter(prometheus.CounterOpts{
		Name: "nullll",
		Help: "The total number of processed events",
	})

)

func main() {
	go recordMetrics()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}