package influx

import (
	"encoding/json"
	"reflect"
	"sort"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/gamedb/gamedb/pkg/log"
	influx "github.com/influxdata/influxdb1-client"
	influxModels "github.com/influxdata/influxdb1-client/models"
)

const (
	InfluxGameDB     = "GameDB"
	InfluxTelegrafDB = "Telegraf-Web"

	InfluxRetentionPolicyAllTime InfluxRetentionPolicy = "alltime"
	InfluxRetentionPolicy7Day    InfluxRetentionPolicy = "7d"
	InfluxRetentionPolicy14Day   InfluxRetentionPolicy = "14d"

	InfluxMeasurementAPICalls    InfluxMeasurement = "api_calls"
	InfluxMeasurementApps        InfluxMeasurement = "apps"
	InfluxMeasurementChanges     InfluxMeasurement = "changes"
	InfluxMeasurementChatBot     InfluxMeasurement = "chat_bot"
	InfluxMeasurementGroups      InfluxMeasurement = "groups"
	InfluxMeasurementPackages    InfluxMeasurement = "packages"
	InfluxMeasurementPlayers     InfluxMeasurement = "players"
	InfluxMeasurementRabbitQueue InfluxMeasurement = "rabbitmq_queue"
	InfluxMeasurementStats       InfluxMeasurement = "stats"
	InfluxMeasurementSignups     InfluxMeasurement = "signups"
	InfluxMeasurementTags        InfluxMeasurement = "tags"
)

type InfluxRetentionPolicy string

func (irp InfluxRetentionPolicy) String() string {
	return string(irp)
}

type InfluxMeasurement string

func (im InfluxMeasurement) String() string {
	return string(im)
}

func InfluxWrite(retention InfluxRetentionPolicy, point influx.Point) (resp *influx.Response, err error) {

	return InfluxWriteMany(retention, influx.BatchPoints{
		Points: []influx.Point{point},
	})
}

func InfluxWriteMany(retention InfluxRetentionPolicy, batch influx.BatchPoints) (resp *influx.Response, err error) {

	if len(batch.Points) == 0 {
		return &influx.Response{}, nil
	}

	batch.Database = InfluxGameDB
	batch.RetentionPolicy = string(retention)
	batch.Precision = batch.Points[0].Precision // Must be in batch and point

	if batch.Time.IsZero() || batch.Time.Unix() == 0 {
		batch.Time = time.Now()
	}

	client, err := getInfluxClient()
	if err != nil {
		return resp, err
	}

	policy := backoff.NewExponentialBackOff()
	policy.InitialInterval = time.Second

	operation := func() (err error) {
		resp, err = client.Write(batch)
		return err
	}

	err = backoff.RetryNotify(operation, backoff.WithMaxRetries(policy, 5), func(err error, t time.Duration) { log.Info(err) })
	return resp, err
}

func InfluxQuery(query string) (resp *influx.Response, err error) {

	client, err := getInfluxClient()
	if err != nil {
		return resp, err
	}

	resp, err = client.Query(influx.Query{
		Command:         query,
		Database:        InfluxGameDB,
		RetentionPolicy: string(InfluxRetentionPolicyAllTime),
	})

	return resp, err
}

type (
	HighChartsJSON      map[string][][]interface{}
	HighChartsJSONMulti struct {
		Key   string         `json:"key"`
		Value HighChartsJSON `json:"value"`
	}
)

func InfluxResponseToHighCharts(series influxModels.Row) HighChartsJSON {

	resp := HighChartsJSON{}

	for k, v := range series.Columns {
		if k > 0 {

			// var hasValue bool

			for _, vv := range series.Values {

				// if !hasValue && trimLeft {
				// 	func() {
				// 		for k, vvv := range vv {
				// 			if k > 0 {
				// 				if val, ok := vvv.(json.Number); ok {
				// 					i, err := val.Float64()
				// 					if err == nil && math.Abs(i) > 0 {
				// 						hasValue = true
				// 						return
				// 					}
				// 				}
				// 			}
				// 		}
				// 	}()
				// }

				// if hasValue || !trimLeft {

				t, err := time.Parse(time.RFC3339, vv[0].(string))
				if err != nil {
					log.Err(err)
					continue
				}

				resp[v] = append(resp[v], []interface{}{t.Unix() * 1000, vv[k]})
				// }
			}
		}
	}

	for k := range resp {

		sort.Slice(resp[k], func(i, j int) bool {
			return resp[k][i][0].(int64) < resp[k][j][0].(int64)
		})

	}

	return resp
}

func InfluxResponseToImageChartData(series influxModels.Row) (x []time.Time, y []float64) {

	for k := range series.Columns {
		if k > 0 {
			for _, vv := range series.Values {

				t, err := time.Parse(time.RFC3339, vv[0].(string))
				if err != nil {
					log.Err(err)
					continue
				}

				// Hide some incorrect data in Influx
				if t.Year() < 2000 {
					continue
				}

				val, ok := vv[k].(json.Number)
				if ok {
					i, err := val.Float64()
					if err == nil {
						x = append(x, t)
						y = append(y, i)
					}
				}
			}
		}
	}

	return x, y
}

func GetFirstInfluxInt(resp *influx.Response) int64 {

	if resp != nil &&
		len(resp.Results) > 0 &&
		len(resp.Results[0].Series) > 0 &&
		len(resp.Results[0].Series[0].Values) > 0 &&
		len(resp.Results[0].Series[0].Values[0]) > 1 {

		switch v := resp.Results[0].Series[0].Values[0][1].(type) {
		case int:
			return int64(v)
		case int64:
			return v
		case json.Number:
			i, err := v.Int64()
			log.Err(err)
			return i
		default:
			log.Warning("Unknown type from Influx DB: " + reflect.TypeOf(v).String())
		}
	}

	return 0
}

func GetFirstInfluxFloat(resp *influx.Response) float64 {

	if resp != nil &&
		len(resp.Results) > 0 &&
		len(resp.Results[0].Series) > 0 &&
		len(resp.Results[0].Series[0].Values) > 0 &&
		len(resp.Results[0].Series[0].Values[0]) > 1 {

		switch v := resp.Results[0].Series[0].Values[0][1].(type) {
		case float64:
			return v
		case json.Number:
			i, err := v.Float64()
			log.Err(err)
			return i
		default:
			log.Warning("Unknown type from Influx DB: " + reflect.TypeOf(v).String())
		}
	}

	return 0
}