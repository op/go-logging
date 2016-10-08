package logging

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/olivere/elastic"
)

type elasticSearchBackend struct {
	Client   *elastic.Client
	appName  string
	nodeName string
}

func NewElasticSearchBackendFrom(ec *elastic.Client, appName, nodeName string) (*elasticSearchBackend, error) {
	return &elasticSearchBackend{
		appName:  appName,
		Client:   ec,
		nodeName: nodeName,
	}, nil
}

func NewElasticSearchBackendTo(ip, port, appName string) (*elasticSearchBackend, error) {
	c, err := elastic.NewClient(
		elastic.SetURL("http://"+ip+":"+port),
		elastic.SetMaxRetries(10))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect elastic, %v", err)
	}
	return &elasticSearchBackend{
		appName: appName,
		Client:  c,
	}, nil
}

func (be *elasticSearchBackend) Log(level Level, calldepth int, rec *Record) error {
	msgMap := map[string]string{
		"level":      level.String(),
		"message":    rec.Message(),
		"@timestamp": rec.Time.Format(time.RFC3339Nano),
		"node":       be.nodeName,
	}
	msg, err := json.Marshal(msgMap)
	if err != nil {
		return fmt.Errorf("Failed to marshal message '%#v' to JSON, %v", msgMap, err)
	}
	msg = append(msg, '\n')
	currDay := time.Now().Format(`-2006-01-02`)
	_, err = be.Client.Index().Index(be.appName + currDay).Type("logs").BodyString(string(msg)).Do()
	return err
}
