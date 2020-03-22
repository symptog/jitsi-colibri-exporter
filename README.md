# Jitsi Prometheus Exporter
Exporter that grabs various metrics from [Jitsi](https://jitsi.org), especially form the video bridges, and publishes them as [Prometheus](https://prometheus.io) metrics.

There is a [documentation](https://github.com/jitsi/jitsi-videobridge/blob/master/doc/statistics.md) of the published statistics by the video bridges.

[Jitsi Stats HTTP-API documentation](https://github.com/jitsi/jitsi-videobridge/blob/master/doc/rest.md)

# Jitsi Settings

In `/etc/jitsi/videobridge/sip-communicator.properties`:

```
org.jitsi.videobridge.ENABLE_STATISTICS=true
org.jitsi.videobridge.STATISTICS_TRANSPORT=colibri
org.jitsi.videobridge.STATISTICS_INTERVAL=1000
```

In `/etc/jitsi/videobridge/config`:

* Add `--apis=rest` to `JVB_OPTS`
* Add `--add-opens jdk.management/com.sun.management.internal=ALL-UNNAMED` to `JAVA_SYS_PROPS`


# Install

```
go get github.com/symptog/jitsi-colibri-exporter
```

# Run

```bash
./jitsi-colibri-exporter -h                                                                                                                                                                       [14:56:02]
Usage of ./jitsi-colibri-exporter:
  -colibri.url string
    	Colibiri URL (default "http://127.0.0.1:8080/colibri/stats")
  -loglevel string
    	Log level (default "info")
  -metrics.addr string
    	Metrics address (default ":9210")
  -metrics.path string
    	Metrics path (default "/metrics")
```

# Result

```text
# HELP jitsi_colibri_audiochannels audiochannels
# TYPE jitsi_colibri_audiochannels gauge
jitsi_colibri_audiochannels 0
# HELP jitsi_colibri_bit_rate_download bit_rate_download
# TYPE jitsi_colibri_bit_rate_download gauge
jitsi_colibri_bit_rate_download 0
# HELP jitsi_colibri_bit_rate_upload bit_rate_upload
# TYPE jitsi_colibri_bit_rate_upload gauge
jitsi_colibri_bit_rate_upload 0
# HELP jitsi_colibri_conference_sizes conference_sizes
# TYPE jitsi_colibri_conference_sizes histogram
jitsi_colibri_conference_sizes_bucket{le="0"} 0
jitsi_colibri_conference_sizes_bucket{le="1"} 0
jitsi_colibri_conference_sizes_bucket{le="2"} 0
jitsi_colibri_conference_sizes_bucket{le="3"} 0
jitsi_colibri_conference_sizes_bucket{le="4"} 0
jitsi_colibri_conference_sizes_bucket{le="5"} 0
jitsi_colibri_conference_sizes_bucket{le="6"} 0
jitsi_colibri_conference_sizes_bucket{le="7"} 0
jitsi_colibri_conference_sizes_bucket{le="8"} 0
jitsi_colibri_conference_sizes_bucket{le="9"} 0
jitsi_colibri_conference_sizes_bucket{le="10"} 0
jitsi_colibri_conference_sizes_bucket{le="11"} 0
jitsi_colibri_conference_sizes_bucket{le="12"} 0
jitsi_colibri_conference_sizes_bucket{le="13"} 0
jitsi_colibri_conference_sizes_bucket{le="14"} 0
jitsi_colibri_conference_sizes_bucket{le="15"} 0
jitsi_colibri_conference_sizes_bucket{le="16"} 0
jitsi_colibri_conference_sizes_bucket{le="17"} 0
jitsi_colibri_conference_sizes_bucket{le="18"} 0
jitsi_colibri_conference_sizes_bucket{le="19"} 0
jitsi_colibri_conference_sizes_bucket{le="20"} 0
jitsi_colibri_conference_sizes_bucket{le="+Inf"} 0
jitsi_colibri_conference_sizes_sum 0
jitsi_colibri_conference_sizes_count 0
# HELP jitsi_colibri_conferences conferences
# TYPE jitsi_colibri_conferences gauge
jitsi_colibri_conferences 0
# HELP jitsi_colibri_cpu_usage cpu_usage
# TYPE jitsi_colibri_cpu_usage gauge
jitsi_colibri_cpu_usage 0
# HELP jitsi_colibri_jitter_aggregate jitter_aggregate
# TYPE jitsi_colibri_jitter_aggregate gauge
jitsi_colibri_jitter_aggregate 0
# HELP jitsi_colibri_largest_conference largest_conference
# TYPE jitsi_colibri_largest_conference gauge
jitsi_colibri_largest_conference 0
# HELP jitsi_colibri_loss_rate_download loss_rate_download
# TYPE jitsi_colibri_loss_rate_download gauge
jitsi_colibri_loss_rate_download 0
# HELP jitsi_colibri_loss_rate_upload loss_rate_upload
# TYPE jitsi_colibri_loss_rate_upload gauge
jitsi_colibri_loss_rate_upload 0
# HELP jitsi_colibri_packet_rate_download packet_rate_download
# TYPE jitsi_colibri_packet_rate_download gauge
jitsi_colibri_packet_rate_download 0
# HELP jitsi_colibri_packet_rate_upload packet_rate_upload
# TYPE jitsi_colibri_packet_rate_upload gauge
jitsi_colibri_packet_rate_upload 0
# HELP jitsi_colibri_participants participants
# TYPE jitsi_colibri_participants gauge
jitsi_colibri_participants 0
# HELP jitsi_colibri_rtp_loss rtp_loss
# TYPE jitsi_colibri_rtp_loss gauge
jitsi_colibri_rtp_loss 0
# HELP jitsi_colibri_rtt_aggregate rtt_aggregate
# TYPE jitsi_colibri_rtt_aggregate gauge
jitsi_colibri_rtt_aggregate 0
# HELP jitsi_colibri_threads threads
# TYPE jitsi_colibri_threads gauge
jitsi_colibri_threads 81
# HELP jitsi_colibri_total_colibri_web_socket_messages_received total_colibri_web_socket_messages_received
# TYPE jitsi_colibri_total_colibri_web_socket_messages_received counter
jitsi_colibri_total_colibri_web_socket_messages_received 0
# HELP jitsi_colibri_total_colibri_web_socket_messages_sent total_colibri_web_socket_messages_sent
# TYPE jitsi_colibri_total_colibri_web_socket_messages_sent counter
jitsi_colibri_total_colibri_web_socket_messages_sent 0
# HELP jitsi_colibri_total_conference_seconds total_conference_seconds
# TYPE jitsi_colibri_total_conference_seconds counter
jitsi_colibri_total_conference_seconds 12718
# HELP jitsi_colibri_total_conferences_created total_conferences_created
# TYPE jitsi_colibri_total_conferences_created counter
jitsi_colibri_total_conferences_created 4
# HELP jitsi_colibri_total_data_channel_messages_received total_data_channel_messages_received
# TYPE jitsi_colibri_total_data_channel_messages_received counter
jitsi_colibri_total_data_channel_messages_received 16617
# HELP jitsi_colibri_total_data_channel_messages_sent total_data_channel_messages_sent
# TYPE jitsi_colibri_total_data_channel_messages_sent counter
jitsi_colibri_total_data_channel_messages_sent 16556
# HELP jitsi_colibri_total_failed_conferences total_failed_conferences
# TYPE jitsi_colibri_total_failed_conferences counter
jitsi_colibri_total_failed_conferences 0
# HELP jitsi_colibri_total_loss_controlled_participant_seconds total_loss_controlled_participant_seconds
# TYPE jitsi_colibri_total_loss_controlled_participant_seconds counter
jitsi_colibri_total_loss_controlled_participant_seconds 10230
# HELP jitsi_colibri_total_loss_degraded_participant_seconds total_loss_degraded_participant_seconds
# TYPE jitsi_colibri_total_loss_degraded_participant_seconds counter
jitsi_colibri_total_loss_degraded_participant_seconds 0
# HELP jitsi_colibri_total_loss_limited_participant_seconds total_loss_limited_participant_seconds
# TYPE jitsi_colibri_total_loss_limited_participant_seconds counter
jitsi_colibri_total_loss_limited_participant_seconds 0
# HELP jitsi_colibri_total_memory total_memory
# TYPE jitsi_colibri_total_memory gauge
jitsi_colibri_total_memory 16820
# HELP jitsi_colibri_total_partially_failed_conferences total_partially_failed_conferences
# TYPE jitsi_colibri_total_partially_failed_conferences counter
jitsi_colibri_total_partially_failed_conferences 0
# HELP jitsi_colibri_up Whether the Azure ServiceBus scrape was successful
# TYPE jitsi_colibri_up gauge
jitsi_colibri_up 1
# HELP jitsi_colibri_used_memory used_memory
# TYPE jitsi_colibri_used_memory gauge
jitsi_colibri_used_memory 1832
# HELP jitsi_colibri_videochannels videochannels
# TYPE jitsi_colibri_videochannels gauge
jitsi_colibri_videochannels 0
# HELP jitsi_colibri_videostreams videostreams
# TYPE jitsi_colibri_videostreams gauge
jitsi_colibri_videostreams 0
```