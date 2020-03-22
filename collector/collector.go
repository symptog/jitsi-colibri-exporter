package collector

import (
	"encoding/json"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

type ColibriMetrics struct {
	Threads                               float64  `json:"threads"`
	UsedMemory                            float64  `json:"used_memory"`
	TotalMemory                           float64  `json:"total_memory"`
	CPUUsage                              float64  `json:"cpu_usage"`
	BitRateDownload                       float64  `json:"bit_rate_download"`
	BitRateUpload                         float64  `json:"bit_rate_upload"`
	PacketRateDownload                    float64  `json:"packet_rate_download"`
	PacketRateUpload                      float64  `json:"packet_rate_upload"`
	LossRateDownload                      float64  `json:"loss_rate_download"`
	LossRateUpload                        float64  `json:"loss_rate_upload"`
	RTPLoss                               float64  `json:"rtp_loss"`
	JitterAggregate                       float64  `json:"jitter_aggregate"`
	RTTAggregate                          float64  `json:"rtt_aggregate"`
	LargestConference                     float64  `json:"largest_conference"`
	ConferenceSizes                       []uint64 `json:"conference_sizes"`
	Audiochannels                         float64  `json:"audiochannels"`
	Videochannels                         float64  `json:"videochannels"`
	Conferences                           float64  `json:"conferences"`
	Participants                          float64  `json:"participants"`
	Videostreams                          float64  `json:"videostreams"`
	TotalLossControlledParticipantSeconds float64  `json:"total_loss_controlled_participant_seconds"`
	TotalLossLimitedParticipantSeconds    float64  `json:"total_loss_limited_participant_seconds"`
	TotalLossDegradedParticipantSeconds   float64  `json:"total_loss_degraded_participant_seconds"`
	TotalConferenceSeconds                float64  `json:"total_conference_seconds"`
	TotalConferencesCreated               float64  `json:"total_conferences_created"`
	TotalFailedConferences                float64  `json:"total_failed_conferences"`
	TotalPartiallyFailedConferences       float64  `json:"total_partially_failed_conferences"`
	TotalDataChannelMessagesReceived      float64  `json:"total_data_channel_messages_received"`
	TotalDataChannelMessagesSent          float64  `json:"total_data_channel_messages_sent"`
	TotalColibriWebsocketMessagesReceived float64  `json:"total_colibri_web_socket_messages_received"`
	TotalColibriWebsocketMessagesSent     float64  `json:"total_colibri_web_socket_messages_sent"`
}

type Collector struct {
	client *http.Client
	target string
	log    *logrus.Logger

	up                                         *prometheus.Desc
	threads                                    *prometheus.Desc
	used_memory                                *prometheus.Desc
	total_memory                               *prometheus.Desc
	cpu_usage                                  *prometheus.Desc
	bit_rate_download                          *prometheus.Desc
	bit_rate_upload                            *prometheus.Desc
	packet_rate_download                       *prometheus.Desc
	packet_rate_upload                         *prometheus.Desc
	loss_rate_download                         *prometheus.Desc
	loss_rate_upload                           *prometheus.Desc
	rtp_loss                                   *prometheus.Desc
	jitter_aggregate                           *prometheus.Desc
	rtt_aggregate                              *prometheus.Desc
	largest_conference                         *prometheus.Desc
	conference_sizes                           *prometheus.Desc
	audiochannels                              *prometheus.Desc
	videochannels                              *prometheus.Desc
	conferences                                *prometheus.Desc
	participants                               *prometheus.Desc
	videostreams                               *prometheus.Desc
	total_loss_controlled_participant_seconds  *prometheus.Desc
	total_loss_limited_participant_seconds     *prometheus.Desc
	total_loss_degraded_participant_seconds    *prometheus.Desc
	total_conference_seconds                   *prometheus.Desc
	total_conferences_created                  *prometheus.Desc
	total_failed_conferences                   *prometheus.Desc
	total_partially_failed_conferences         *prometheus.Desc
	total_data_channel_messages_received       *prometheus.Desc
	total_data_channel_messages_sent           *prometheus.Desc
	total_colibri_web_socket_messages_received *prometheus.Desc
	total_colibri_web_socket_messages_sent     *prometheus.Desc
}

func New(client *http.Client, target string) *Collector {
	return &Collector{
		client: client,
		target: target,

		up:                   prometheus.NewDesc("jitsi_colibri_up", "Whether the Azure ServiceBus scrape was successful", nil, nil),
		threads:              prometheus.NewDesc("jitsi_colibri_threads", "threads", nil, nil),
		used_memory:          prometheus.NewDesc("jitsi_colibri_used_memory", "used_memory", nil, nil),
		total_memory:         prometheus.NewDesc("jitsi_colibri_total_memory", "total_memory", nil, nil),
		cpu_usage:            prometheus.NewDesc("jitsi_colibri_cpu_usage", "cpu_usage", nil, nil),
		bit_rate_download:    prometheus.NewDesc("jitsi_colibri_bit_rate_download", "bit_rate_download", nil, nil),
		bit_rate_upload:      prometheus.NewDesc("jitsi_colibri_bit_rate_upload", "bit_rate_upload", nil, nil),
		packet_rate_download: prometheus.NewDesc("jitsi_colibri_packet_rate_download", "packet_rate_download", nil, nil),
		packet_rate_upload:   prometheus.NewDesc("jitsi_colibri_packet_rate_upload", "packet_rate_upload", nil, nil),
		loss_rate_download:   prometheus.NewDesc("jitsi_colibri_loss_rate_download", "loss_rate_download", nil, nil),
		loss_rate_upload:     prometheus.NewDesc("jitsi_colibri_loss_rate_upload", "loss_rate_upload", nil, nil),
		rtp_loss:             prometheus.NewDesc("jitsi_colibri_rtp_loss", "rtp_loss", nil, nil),
		jitter_aggregate:     prometheus.NewDesc("jitsi_colibri_jitter_aggregate", "jitter_aggregate", nil, nil),
		rtt_aggregate:        prometheus.NewDesc("jitsi_colibri_rtt_aggregate", "rtt_aggregate", nil, nil),
		largest_conference:   prometheus.NewDesc("jitsi_colibri_largest_conference", "largest_conference", nil, nil),
		audiochannels:        prometheus.NewDesc("jitsi_colibri_audiochannels", "audiochannels", nil, nil),
		videochannels:        prometheus.NewDesc("jitsi_colibri_videochannels", "videochannels", nil, nil),
		conferences:          prometheus.NewDesc("jitsi_colibri_conferences", "conferences", nil, nil),
		participants:         prometheus.NewDesc("jitsi_colibri_participants", "participants", nil, nil),
		videostreams:         prometheus.NewDesc("jitsi_colibri_videostreams", "videostreams", nil, nil),
		total_loss_controlled_participant_seconds:  prometheus.NewDesc("jitsi_colibri_total_loss_controlled_participant_seconds", "total_loss_controlled_participant_seconds", nil, nil),
		total_loss_limited_participant_seconds:     prometheus.NewDesc("jitsi_colibri_total_loss_limited_participant_seconds", "total_loss_limited_participant_seconds", nil, nil),
		total_loss_degraded_participant_seconds:    prometheus.NewDesc("jitsi_colibri_total_loss_degraded_participant_seconds", "total_loss_degraded_participant_seconds", nil, nil),
		total_conference_seconds:                   prometheus.NewDesc("jitsi_colibri_total_conference_seconds", "total_conference_seconds", nil, nil),
		total_conferences_created:                  prometheus.NewDesc("jitsi_colibri_total_conferences_created", "total_conferences_created", nil, nil),
		total_failed_conferences:                   prometheus.NewDesc("jitsi_colibri_total_failed_conferences", "total_failed_conferences", nil, nil),
		total_partially_failed_conferences:         prometheus.NewDesc("jitsi_colibri_total_partially_failed_conferences", "total_partially_failed_conferences", nil, nil),
		total_data_channel_messages_received:       prometheus.NewDesc("jitsi_colibri_total_data_channel_messages_received", "total_data_channel_messages_received", nil, nil),
		total_data_channel_messages_sent:           prometheus.NewDesc("jitsi_colibri_total_data_channel_messages_sent", "total_data_channel_messages_sent", nil, nil),
		total_colibri_web_socket_messages_received: prometheus.NewDesc("jitsi_colibri_total_colibri_web_socket_messages_received", "total_colibri_web_socket_messages_received", nil, nil),
		total_colibri_web_socket_messages_sent:     prometheus.NewDesc("jitsi_colibri_total_colibri_web_socket_messages_sent", "total_colibri_web_socket_messages_sent", nil, nil),
		conference_sizes:                           prometheus.NewDesc("jitsi_colibri_conference_sizes", "conference_sizes", nil, nil),
	}
}

func (c *Collector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.up
	ch <- c.threads
	ch <- c.used_memory
	ch <- c.total_memory
	ch <- c.cpu_usage
	ch <- c.bit_rate_download
	ch <- c.bit_rate_upload
	ch <- c.packet_rate_download
	ch <- c.packet_rate_upload
	ch <- c.loss_rate_download
	ch <- c.loss_rate_upload
	ch <- c.rtp_loss
	ch <- c.jitter_aggregate
	ch <- c.rtt_aggregate
	ch <- c.largest_conference
	ch <- c.conference_sizes
	ch <- c.audiochannels
	ch <- c.videochannels
	ch <- c.conferences
	ch <- c.participants
	ch <- c.videostreams
	ch <- c.total_loss_controlled_participant_seconds
	ch <- c.total_loss_limited_participant_seconds
	ch <- c.total_loss_degraded_participant_seconds
	ch <- c.total_conference_seconds
	ch <- c.total_conferences_created
	ch <- c.total_failed_conferences
	ch <- c.total_partially_failed_conferences
	ch <- c.total_data_channel_messages_received
	ch <- c.total_data_channel_messages_sent
	ch <- c.total_colibri_web_socket_messages_received
	ch <- c.total_colibri_web_socket_messages_sent
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

func conferenceSizesHelper(conferenceSizes []uint64) (conferenceSizesHistogram map[float64]uint64, sum uint64) {
	var sizes = make(map[float64]uint64)
	var values []uint64
	for _, v := range conferenceSizes {
		values = append(values, v)
	}

	//calculate sum (makes this metric independent from conferences metric)
	sum = 0
	for _, v := range values {
		sum += v
	}

	//for the histgram buckets we need to omit the last field b/c the +inf bucket is added automatically
	values = values[:len(values)-1]

	//the bucket values have to be cumulative
	var i int
	for i = len(values) - 1; i >= 0; i-- {
		var cumulative uint64
		var j int
		for j = i; j >= 0; j-- {
			cumulative += values[j]
		}
		values[i] = cumulative
	}

	for i, v := range values {
		sizes[float64(i)] = v
	}

	return sizes, sum
}

func (c *Collector) Collect(ch chan<- prometheus.Metric) {
	jsonData, err := probeColibri(c.client, c.target)

	if err != nil {
		log.Error(err)
		// client call failed, set the up metric value to 0
		ch <- prometheus.MustNewConstMetric(c.up, prometheus.GaugeValue, 0)

	} else {
		// client call succeeded, set the up metric value to 1
		ch <- prometheus.MustNewConstMetric(c.up, prometheus.GaugeValue, 1)
		ch <- prometheus.MustNewConstMetric(c.threads, prometheus.GaugeValue, jsonData.Threads)
		ch <- prometheus.MustNewConstMetric(c.used_memory, prometheus.GaugeValue, jsonData.UsedMemory)
		ch <- prometheus.MustNewConstMetric(c.total_memory, prometheus.GaugeValue, jsonData.TotalMemory)
		ch <- prometheus.MustNewConstMetric(c.cpu_usage, prometheus.GaugeValue, jsonData.CPUUsage)
		ch <- prometheus.MustNewConstMetric(c.bit_rate_download, prometheus.GaugeValue, jsonData.BitRateDownload)
		ch <- prometheus.MustNewConstMetric(c.bit_rate_upload, prometheus.GaugeValue, jsonData.BitRateUpload)
		ch <- prometheus.MustNewConstMetric(c.packet_rate_download, prometheus.GaugeValue, jsonData.PacketRateDownload)
		ch <- prometheus.MustNewConstMetric(c.packet_rate_upload, prometheus.GaugeValue, jsonData.PacketRateUpload)
		ch <- prometheus.MustNewConstMetric(c.loss_rate_download, prometheus.GaugeValue, jsonData.LossRateUpload)
		ch <- prometheus.MustNewConstMetric(c.loss_rate_upload, prometheus.GaugeValue, jsonData.LossRateDownload)
		ch <- prometheus.MustNewConstMetric(c.rtp_loss, prometheus.GaugeValue, jsonData.RTPLoss)
		ch <- prometheus.MustNewConstMetric(c.jitter_aggregate, prometheus.GaugeValue, jsonData.JitterAggregate)
		ch <- prometheus.MustNewConstMetric(c.rtt_aggregate, prometheus.GaugeValue, jsonData.RTTAggregate)
		ch <- prometheus.MustNewConstMetric(c.largest_conference, prometheus.GaugeValue, jsonData.LargestConference)
		ch <- prometheus.MustNewConstMetric(c.audiochannels, prometheus.GaugeValue, jsonData.Audiochannels)
		ch <- prometheus.MustNewConstMetric(c.videochannels, prometheus.GaugeValue, jsonData.Videochannels)
		ch <- prometheus.MustNewConstMetric(c.conferences, prometheus.GaugeValue, jsonData.Conferences)
		ch <- prometheus.MustNewConstMetric(c.participants, prometheus.GaugeValue, jsonData.Participants)
		ch <- prometheus.MustNewConstMetric(c.videostreams, prometheus.GaugeValue, jsonData.Videochannels)
		ch <- prometheus.MustNewConstMetric(c.total_loss_controlled_participant_seconds, prometheus.CounterValue, jsonData.TotalLossControlledParticipantSeconds)
		ch <- prometheus.MustNewConstMetric(c.total_loss_limited_participant_seconds, prometheus.GaugeValue, jsonData.TotalLossLimitedParticipantSeconds)
		ch <- prometheus.MustNewConstMetric(c.total_loss_degraded_participant_seconds, prometheus.GaugeValue, jsonData.TotalLossDegradedParticipantSeconds)
		ch <- prometheus.MustNewConstMetric(c.total_conference_seconds, prometheus.GaugeValue, jsonData.TotalConferenceSeconds)
		ch <- prometheus.MustNewConstMetric(c.total_conferences_created, prometheus.GaugeValue, jsonData.TotalConferencesCreated)
		ch <- prometheus.MustNewConstMetric(c.total_failed_conferences, prometheus.GaugeValue, jsonData.TotalFailedConferences)
		ch <- prometheus.MustNewConstMetric(c.total_partially_failed_conferences, prometheus.GaugeValue, jsonData.TotalPartiallyFailedConferences)
		ch <- prometheus.MustNewConstMetric(c.total_data_channel_messages_received, prometheus.GaugeValue, jsonData.TotalDataChannelMessagesReceived)
		ch <- prometheus.MustNewConstMetric(c.total_data_channel_messages_sent, prometheus.GaugeValue, jsonData.TotalDataChannelMessagesSent)
		ch <- prometheus.MustNewConstMetric(c.total_colibri_web_socket_messages_received, prometheus.GaugeValue, jsonData.TotalColibriWebsocketMessagesReceived)
		ch <- prometheus.MustNewConstMetric(c.total_colibri_web_socket_messages_sent, prometheus.GaugeValue, jsonData.TotalColibriWebsocketMessagesSent)

		var combinedConferenceSizes = make(map[float64]uint64)
		var combinedSum uint64
		conSizes, sum := conferenceSizesHelper(jsonData.ConferenceSizes)
		for bucket, numConferences := range conSizes {
			combinedConferenceSizes[bucket] += numConferences
		}
		combinedSum += sum

		ch <- prometheus.MustNewConstHistogram(c.conference_sizes, combinedSum, float64(combinedSum), combinedConferenceSizes)

	}
}
