# Billbee empty house number field assistant

Introducing the Billbee Empty House Number Field Assistant, an innovative open-source solution crafted in Go (Golang) to
streamline and enhance the efficiency of order processing. This specialized tool is designed to address a common pain
point encountered in Billbee orders by automatically checking and rectifying empty house number fields within address
details. By leveraging seamless API access, the application empowers users to significantly reduce the need for
time-consuming manual post-processing, ensuring that order data is accurate and complete right from the start.

The Billbee Empty House Number Field Assistant operates in a continuous loop, tirelessly monitoring and optimizing
address information until the user decides to halt the process. This ensures a hands-free and proactive approach to
order management, enabling businesses to focus their resources on more strategic and value-added tasks. With its
user-friendly design and robust functionality, this open-source project serves as a valuable asset for anyone seeking to
optimize their Billbee order workflow, ultimately saving time and minimizing errors in the address details of their
transactions. Embrace automation and elevate your order processing experience with the Billbee Empty House Number Field
Assistant in Go – a reliable ally for efficient and error-free business operations.

## Configuration

The following environment variables are used for configuration

| Name             | Description                                                   |
|------------------|---------------------------------------------------------------|
| RUN_INTERVAL     | **Optional:** How often to scan for new orders (default: 300) |
| BILLBEE_USER     | Billbee user name                                             |
| BILLBEE_PASSWORD | The API password of the Billbee user                          |
| BILLBEE_API_KEY  | The Billbee API Key                                           |

### Billbee API Key

In [the documentation](https://app.billbee.io//swagger/ui/index) you can read how to get an API key:
> In addition you need a Billbee API Key identifying the application you develop. To get an API key, send a mail to
> support@billbee.io and send us a short note about what you are building.

## Installation

You can easily run this as a docker compose project:

```yaml
version: "3"

services:
  worker:
    image: lippertsweb/billbee-house-number-assistant:latest
    restart: unless-stopped
    environment:
      # Billbee user name
      BILLBEE_USER: "..."

      # The API password of the Billbee user
      BILLBEE_PASSWORD: "..."

      # The Billbee API Key
      BILLBEE_API_KEY: "..."

      # Optional: Run all five minutes (5 minutes * 60 seconds = 300 seconds)
      # RUN_INTERVAL: 300
```
