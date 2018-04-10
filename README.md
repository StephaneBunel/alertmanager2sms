# am2sms

am2sms (alertmanager2sms) is an SMS alerting tool designed to listen alerts from [Prometheus Alertmanager](https://prometheus.io/docs/alerting/alertmanager/)
and forward them as [SMS](https://en.wikipedia.org/wiki/SMS).
am2sms uses [sendsms](https://github.com/StephaneBunel/sendsms) as a library to handle the SMS part.

## Usage

```bash
am2sms
```

## Configuration

/etc/am2sms/config.yml:

```yaml
---

am2sms:
  logLevel:                  "INFO"
  eventBufferSize:           128

  http:
    address:                 "127.0.0.1:7171"
    endpoint:                "/am2sms"

  sms:
    worker:                  2
    profile:                 "default"
    template: |
      {{ $numAlerts := len .Alerts }}Hello {{ .Receiver }},
      Alertmanger raised {{ $numAlerts }} {{ .Status | ToUpper }} alert(s):
      {{ range $index, $alert := .Alerts }}
      ({{ add $index 1 }}/{{ $numAlerts }}) {{ $alert.Status | ToUpper }} {{ $alert.Labels.severity | ToUpper }} {{ $alert.Labels.instance }}/{{ $alert.Labels.job }}
      Begin: {{ $alert.StartsAt }}
      {{ $alert.Annotations.description }}
      {{ end }}
      Thank you.

  stat:
    handler: prom
    handlerConfig:
      address: 127.0.0.1:1471
      endpoint: /metrics

  recipients:
    enabled:
      - from_csv_file

    config:
      from_csv_file:
        engine: "csv"
        filename: "recipients.csv"

sendsms:
  profiles:
    default:
      provider:              "ovh"
      providerConfig:
        api:
          location:          "ovh-eu"
          appKey:            "<app key>"
          appSecret:         "<app secret>"
          consumerKey:       "<consumer key>"
          servicename:       "<service name>"
        smsOptions:
          sender:            "<sender>"
        smsOptionsCaps:
          nostopclause:      "noStopClause"
          servicename:       "serviceName"
          senderforresponse: "senderForResponse"
```

Alertmanager configuration snippet:

```yaml
receivers:
  - name: 'sysadmin'
    webhook_configs:
      - url: 'http://127.0.0.1:7171/am2sms'
        send_resolved: true
```
