package main

import (
	"crypto/tls"
	"encoding/json"
	"flag"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

var (
	LogLevel = flag.String("loglevel", "info", "Log level")

	MetricsAddr = flag.String("metrics.addr", ":9210", "Metrics address")
	MetricsPath = flag.String("metrics.path", "/metrics", "Metrics path")

	ColibriUrl = flag.String("colibri.url", "http://127.0.0.1:8080/colibri/stats", "Colibiri URL")
	ColibriUpdateInterval = flag.Int("colibri.update.interval", 30, "Colibiri update interval in seconds")

	threads = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "jitsi_colibri_threads",
			Help: "threads",
		},
	)
	used_memory = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "jitsi_colibri_used_memory",
			Help: "used_memory",
		},
	)
	total_memory = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "jitsi_colibri_total_memory",
			Help: "total_memory",
		},
	)
	cpu_usage = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "jitsi_colibri_cpu_usage",
			Help: "cpu_usage",
		},
	)
	bit_rate_download = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "jitsi_colibri_bit_rate_download",
			Help: "bit_rate_download",
		},
	)
	bit_rate_upload = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "jitsi_colibri_bit_rate_upload",
			Help: "bit_rate_upload",
		},
	)
	packet_rate_download = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "jitsi_colibri_packet_rate_download",
			Help: "packet_rate_download",
		},
	)
	packet_rate_upload = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "jitsi_colibri_packet_rate_upload",
			Help: "packet_rate_upload",
		},
	)
	loss_rate_download = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "jitsi_colibri_loss_rate_download",
			Help: "loss_rate_download",
		},
	)
	loss_rate_upload = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "jitsi_colibri_loss_rate_upload",
			Help: "loss_rate_upload",
		},
	)
	rtp_loss = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "jitsi_colibri_rtp_loss",
			Help: "rtp_loss",
		},
	)
	jitter_aggregate = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "jitsi_colibri_jitter_aggregate",
			Help: "jitter_aggregate",
		},
	)
	rtt_aggregate = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "jitsi_colibri_rtt_aggregate",
			Help: "rtt_aggregate",
		},
	)
	largest_conference = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "jitsi_colibri_largest_conference",
			Help: "largest_conference",
		},

	)
	conference_sizes = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "jitsi_colibri_conference_sizes",
			Help:    "conference_sizes",
			Buckets: []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		},
	)
	audiochannels = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "jitsi_colibri_audiochannels",
			Help: "audiochannels",
		},
	)
	videochannels = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "jitsi_colibri_videochannels",
			Help: "videochannels",
		},
	)
	conferences = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "jitsi_colibri_conferences",
			Help: "conferences",
		},
	)
	participants = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "jitsi_colibri_participants",
			Help: "participants",
		},
	)
	videostreams = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "jitsi_colibri_videostreams",
			Help: "videostreams",
		},
	)
	total_loss_controlled_participant_seconds = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "jitsi_colibri_total_loss_controlled_participant_seconds",
			Help: "total_loss_controlled_participant_seconds",
		},
	)
	total_loss_limited_participant_seconds = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "jitsi_colibri_total_loss_limited_participant_seconds",
			Help: "total_loss_limited_participant_seconds",
		},
	)
	total_loss_degraded_participant_seconds = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "jitsi_colibri_total_loss_degraded_participant_seconds",
			Help: "total_loss_degraded_participant_seconds",
		},
	)
	total_conference_seconds = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "jitsi_colibri_total_conference_seconds",
			Help: "total_conference_seconds",
		},
	)
	total_conferences_created = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "jitsi_colibri_total_conferences_created",
			Help: "total_conferences_created",
		},
	)
	total_failed_conferences = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "jitsi_colibri_total_failed_conferences",
			Help: "total_failed_conferences",
		},
	)
	total_partially_failed_conferences = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "jitsi_colibri_total_partially_failed_conferences",
			Help: "total_partially_failed_conferences",
		},
	)
	total_data_channel_messages_received = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "jitsi_colibri_total_data_channel_messages_received",
			Help: "total_data_channel_messages_received",
		},
	)
	total_data_channel_messages_sent = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "jitsi_colibri_total_data_channel_messages_sent",
			Help: "total_data_channel_messages_sent",
		},
	)
	total_colibri_web_socket_messages_received = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "jitsi_colibri_total_colibri_web_socket_messages_received",
			Help: "total_colibri_web_socket_messages_received",
		},
	)
	total_colibri_web_socket_messages_sent = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "jitsi_colibri_total_colibri_web_socket_messages_sent",
			Help: "total_colibri_web_socket_messages_sent",
		},
	)

	httpClient *http.Client
)

func init() {
	prometheus.MustRegister(threads)
	prometheus.MustRegister(used_memory)
	prometheus.MustRegister(total_memory)
	prometheus.MustRegister(cpu_usage)
	prometheus.MustRegister(bit_rate_download)
	prometheus.MustRegister(bit_rate_upload)
	prometheus.MustRegister(packet_rate_download)
	prometheus.MustRegister(packet_rate_upload)
	prometheus.MustRegister(loss_rate_download)
	prometheus.MustRegister(loss_rate_upload)
	prometheus.MustRegister(rtp_loss)
	prometheus.MustRegister(jitter_aggregate)
	prometheus.MustRegister(rtt_aggregate)
	prometheus.MustRegister(largest_conference)
	prometheus.MustRegister(conference_sizes)
	prometheus.MustRegister(audiochannels)
	prometheus.MustRegister(videochannels)
	prometheus.MustRegister(conferences)
	prometheus.MustRegister(participants)
	prometheus.MustRegister(videostreams)
	prometheus.MustRegister(total_loss_controlled_participant_seconds)
	prometheus.MustRegister(total_loss_limited_participant_seconds)
	prometheus.MustRegister(total_loss_degraded_participant_seconds)
	prometheus.MustRegister(total_conference_seconds)
	prometheus.MustRegister(total_conferences_created)
	prometheus.MustRegister(total_failed_conferences)
	prometheus.MustRegister(total_partially_failed_conferences)
	prometheus.MustRegister(total_data_channel_messages_received)
	prometheus.MustRegister(total_data_channel_messages_sent)
	prometheus.MustRegister(total_colibri_web_socket_messages_received)
	prometheus.MustRegister(total_colibri_web_socket_messages_sent)

	httpClient = &http.Client{
		Transport: &http.Transport{
			MaxIdleConns: 100,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

}

func httpServer() {
	http.Handle(*MetricsPath, promhttp.Handler())
	log.Fatal(http.ListenAndServe(*MetricsAddr, nil))
}

type ColibriMetrics struct {
	Threads                               float64 `json:"threads"`
	UsedMemory                            float64 `json:"used_memory"`
	TotalMemory                           float64 `json:"total_memory"`
	CpuUsage                              float64 `json:"cpu_usage"`
	BitRateDownload                       float64 `json:"bit_rate_download"`
	BitRateUpload                         float64 `json:"bit_rate_upload"`
	PacketRateDownload                    float64 `json:"packet_rate_download"`
	PacketRateUpload                      float64 `json:"packet_rate_upload"`
	LossRateDownload                      float64 `json:"loss_rate_download"`
	LossRateUpload                        float64 `json:"loss_rate_upload"`
	RTPLoss                               float64 `json:"rtp_loss"`
	JitterAggregate                       float64 `json:"jitter_aggregate"`
	RTTAggregate                          float64 `json:"rtt_aggregate"`
	LargestConference                     float64 `json:"largest_conference"`
	ConferenceSizes                       []int   `json:"conference_sizes"`
	Audiochannels                         float64 `json:"audiochannels"`
	Videochannels                         float64 `json:"videochannels"`
	Conferences                           float64 `json:"conferences"`
	Participants                          float64 `json:"participants"`
	Videostreams                          float64 `json:"videostreams"`
	TotalLossControlledParticipantSeconds float64 `json:"total_loss_controlled_participant_seconds"`
	TotalLossLimitedParticipantSeconds    float64 `json:"total_loss_limited_participant_seconds"`
	TotalLossDegradedParticipantSeconds   float64 `json:"total_loss_degraded_participant_seconds"`
	TotalConferenceSeconds                float64 `json:"total_conference_seconds"`
	TotalConferencesCreated               float64 `json:"total_conferences_created"`
	TotalFailedConferences                float64 `json:"total_failed_conferences"`
	TotalPartiallyFailedConferences       float64 `json:"total_partially_failed_conferences"`
	TotalDataChannelMessagesReceived      float64 `json:"total_data_channel_messages_received"`
	TotalDataChannelMessagesSent          float64 `json:"total_data_channel_messages_sent"`
	TotalColibriWebsocketMessagesReceived float64 `json:"total_colibri_web_socket_messages_received"`
	TotalColibriWebsocketMessagesSent     float64 `json:"total_colibri_web_socket_messages_sent"`
}

func probeColibri(client *http.Client, target string) (ColibriMetrics, error) {
	var jsonData ColibriMetrics

	resp, err := client.Get(target)
	if err != nil {
		return jsonData, err
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return jsonData, err
	}

	err = json.Unmarshal([]byte(bytes), &jsonData)
	if err != nil {
		return jsonData, err
	}

	return jsonData, nil
}

func collectJitsiColibriStats(target string) {
	jsonData, err := probeColibri(httpClient, target)

	if err != nil {
		log.Error(err)
	}

	threads.Set(jsonData.Threads)
	used_memory.Set(jsonData.UsedMemory)
	total_memory.Set(jsonData.TotalMemory)
	cpu_usage.Set(jsonData.CpuUsage)
	bit_rate_download.Set(jsonData.BitRateDownload)
	bit_rate_upload.Set(jsonData.BitRateUpload)
	packet_rate_download.Set(jsonData.PacketRateDownload)
	packet_rate_upload.Set(jsonData.PacketRateUpload)
	loss_rate_download.Set(jsonData.LossRateDownload)
	loss_rate_upload.Set(jsonData.LossRateUpload)
	rtp_loss.Set(jsonData.RTPLoss)
	jitter_aggregate.Set(jsonData.JitterAggregate)
	rtt_aggregate.Set(jsonData.RTTAggregate)
	largest_conference.Set(jsonData.LargestConference)
	for index, value := range jsonData.ConferenceSizes {
		if value > 0 {
			for i := 0; i < value; i++ {
				conference_sizes.Observe(float64(index))
			}
		}
	}
	audiochannels.Set(jsonData.Audiochannels)
	videochannels.Set(jsonData.Videochannels)
	conferences.Set(jsonData.Conferences)
	participants.Set(jsonData.Participants)
	videostreams.Set(jsonData.Videostreams)
	total_loss_controlled_participant_seconds.Set(jsonData.TotalLossControlledParticipantSeconds)
	total_loss_limited_participant_seconds.Set(jsonData.TotalLossLimitedParticipantSeconds)
	total_loss_degraded_participant_seconds.Set(jsonData.TotalLossDegradedParticipantSeconds)
	total_conference_seconds.Set(jsonData.TotalConferenceSeconds)
	total_conferences_created.Set(jsonData.TotalConferencesCreated)
	total_failed_conferences.Set(jsonData.TotalFailedConferences)
	total_partially_failed_conferences.Set(jsonData.TotalPartiallyFailedConferences)
	total_data_channel_messages_received.Set(jsonData.TotalDataChannelMessagesReceived)
	total_data_channel_messages_sent.Set(jsonData.TotalDataChannelMessagesSent)
	total_colibri_web_socket_messages_received.Set(jsonData.TotalColibriWebsocketMessagesReceived)
	total_colibri_web_socket_messages_sent.Set(jsonData.TotalColibriWebsocketMessagesSent)

}

func main() {
	flag.Parse()

	lvl, _ := log.ParseLevel(*LogLevel)
	log.SetLevel(lvl)

	log.Info("Starting Jitsi Colibri Exporter")
	go httpServer()

	for true {
		collectJitsiColibriStats(*ColibriUrl)
		var interval = time.Duration(*ColibriUpdateInterval)
		time.Sleep(interval * time.Second)
	}

}
