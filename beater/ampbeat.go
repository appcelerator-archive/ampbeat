package beater

import (
	"fmt"
	"log"
	"time"

	"github.com/appcelerator/ampbeat/config"
	"github.com/docker/docker/client"
	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/libbeat/publisher"
)

// Ampbeat the amp beat struct
type Ampbeat struct {
	done               chan struct{}
	config             config.Config
	client             publisher.Client
	dockerClient       *client.Client
	eventStreamReading bool
	containers         map[string]*ContainerData
	lastUpdate         time.Time
}

// New Creates beater
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	config := config.DefaultConfig
	if err := cfg.Unpack(&config); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}
	log.Printf("Config: %+v\n", config)
	bt := &Ampbeat{
		done:   make(chan struct{}),
		config: config,
	}
	return bt, nil
}

// Run ampbeat main loop
func (bt *Ampbeat) Run(b *beat.Beat) error {
	logp.Info("starting ampbeat")
	fmt.Printf("config: %v\n", bt.config)

	bt.client = b.Publisher.Connect()
	ticker := time.NewTicker(bt.config.Period)
	err := bt.start(&bt.config)
	if err != nil {
		log.Fatal(err)
	}
	logp.Info("ampbeat is running! Hit CTRL-C to stop it.")

	for {
		select {
		case <-bt.done:
			return nil
		case <-ticker.C:
		}

		bt.tick()
	}
}

// Stop Ampbeat stop
func (bt *Ampbeat) Stop() {
	bt.client.Close()
	bt.Close()
	close(bt.done)
}
