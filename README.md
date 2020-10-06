# Grafana-shell

Call the shell instruction of the specified machine through grafana's alert channel webhook

## Usage

1. Use 'go build' to compile after downloading.
2. Copy `config.simple.xml` to `config.xml`, Modify the corresponding configuration.
3. Add webhook to grafana alert channel, address: `http://{IP or domain}:8850/webhook/`
