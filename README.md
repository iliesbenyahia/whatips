# whatips
![alt text](https://i.ibb.co/23Z9KS0z/kawai-cloud.png)  
**whatips** is a small Go application that retrieves your current public IP address and compares it to a DNS record.  
Currently, it only supports [Gandi](https://www.gandi.net/) as the DNS provider because that is where my personal domain name is registered.  

That said, it is not impossible that the project could be extended in the future to support other DNS registrars and become more generic.

---

## 🌐 Features

- Fetches your public IP address using `https://api.ipify.org`
- Compares it against a DNS record from Gandi (only supported provider for now)
- Designed for automation (e.g., cron jobs, CI, etc.)

---

## 🐳 Docker Usage

An official image is available on Docker Hub:  
👉 [`yayadu08/whatips`](https://hub.docker.com/r/yayadu08/whatips)

### 🔧 Environment Variables

You must provide the following environment variables for the app to work:

| Variable      | Description                                | Example                           |
|---------------|--------------------------------------------|-----------------------------------|
| `GANDI_PAT`   | Gandi Personal Access Token                 | `abc123xyz...`                    |
| `FQDN`        | Fully Qualified Domain Name to check       | `home.example.com`                |
| `RECORD_TYPE` | DNS record type                            | `A`                               |
| `RECORD_NAME` | Subdomain part of the FQDN                 | `home` (for `home.example.com`)   |

---

## 🧪 Example (Docker CLI)

```bash
docker run --rm \
  -e GANDI_PAT=your-gandi-token \
  -e FQDN=example.com \
  -e RECORD_TYPE=A \
  -e RECORD_NAME=home \
  yayadu08/whatips
