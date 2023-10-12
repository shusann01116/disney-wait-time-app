package tdr

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"github.com/disney-wait-time-app/app/collector_go/core"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
)

type TdrClient struct {
	Ctx    context.Context
	Client *http.Client
	rand   *rand.Rand
}

func (t *TdrClient) getRandInt31() int {
	if t.rand == nil {
		t.rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	}

	return int(t.rand.Int31())
}

func (t *TdrClient) getRequest(url string) (*http.Request, error) {
	req, err := http.NewRequestWithContext(
		t.Ctx,
		http.MethodGet,
		fmt.Sprintf("%s?%d", url, time.Now().UnixMilli()),
		nil,
	)
	if err != nil {
		return nil, err
	}
	req.Header.Add(
		"User-Agent",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) ",
	)
	req.Header.Add(
		"Accept",
		"application/json, text/javascript, */*; q=0.01",
	)
	req.Header.Add(
		"Accept-Encoding",
		"gzip, deflate, br",
	)
	req.Header.Add(
		"Referer",
		"https://www.tokyodisneyresort.jp/tdl/attraction.html",
	)
	req.Header.Add(
		"Cookie",
		fmt.Sprintf("%d", t.getRandInt31()),
	)
	return req, nil
}

func (t *TdrClient) FetchStandbys(urls []string) ([]core.Standby, error) {
	g := new(errgroup.Group)

	var results []core.Standby
	for _, url := range urls {
		url := url
		g.Go(func() error {
			req, err := t.getRequest(url)
			if err != nil {
				log.Err(err).Msgf("failed to create request: %v", err)
				return err
			}
			log.Debug().Msgf("request: %+v", req)

			resp, err := t.Client.Do(req)
			if err != nil {
				log.Err(err).Msgf("failed to get %s: %v", url, err)
				return err
			}
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Err(err).Msgf("failed to read body: %v", err)
				return err
			}
			_ = resp.Body.Close()
			log.Debug().Msgf("body: %s", body)

			standby, err := ParseStandby(body)
			if err != nil {
				log.Err(err).Msgf("failed to parse standby: %v", err)
				return err
			}

			results = append(results, standby...)
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		log.Err(err).Msg("")
	}

	return results, nil
}

func ParseStandby(body []byte) ([]core.Standby, error) {
	var o interface{}
	err := json.Unmarshal(body, &o)
	if err != nil {
		return []core.Standby{}, err
	}

	log.Debug().Msgf("parsed: %T", o)
	switch o := o.(type) {
	case []interface{}:
		standbys := make([]core.Standby, len(o))
		for i, f := range o {
			f := f.(map[string]interface{})
			updateTimeStr := fmt.Sprintf("%sT%s:00+09:00", time.Now().Format(time.DateOnly), f["UpdateTime"].(string))
			updateTime, err := time.Parse(time.RFC3339, updateTimeStr)
			if err != nil {
				log.Err(err).Msgf("failed to parse time: %v", err)
				continue
			}
			facilityId, err := strconv.ParseUint(f["FacilityID"].(string), 10, 64)
			if err != nil {
				log.Err(err).Msgf("failed to parse facility id: %v", err)
				continue
			}
			if f["StandbyTime"] == nil {
				f["StandbyTime"] = "0"
			}
			if reflect.TypeOf(f["StandbyTime"]).Kind() == reflect.Bool {
				f["StandbyTime"] = "0"
			}
			duration, err := strconv.ParseInt(f["StandbyTime"].(string), 10, 64)
			if err != nil {
				log.Err(err).Msgf("failed to parse standby time: %v", err)
				continue
			}
			if f["OperatingStatus"] == nil {
				f["OperatingStatus"] = ""
			}
			standbys[i] = core.NewStandby(
				updateTime,
				core.FacilityId(facilityId),
				core.Status{Name: f["OperatingStatus"].(string)},
				time.Minute*time.Duration(duration),
			)
		}
		return standbys, nil
	default:
		return []core.Standby{}, errors.New("unknown type")
	}
}
