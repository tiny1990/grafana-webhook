FROM alpine:3.6

COPY dp-grafana-webhook /


ENTRYPOINT ["/dp-grafana-webhook"]