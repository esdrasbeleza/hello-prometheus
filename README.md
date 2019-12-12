# Prometheus & Grafana

A small example of Prometheus and Grafana monitoring stats from a Go project.

To run it:

1. Install Docker

2. Run

    ```shell
    $ docker-compose up
    ```

3. Go to:

    - [http://localhost:8080/metrics](http://localhost:8080/metrics) to see the metrics used by Prometheus
    - [http://localhost:9090](http://localhost:9090) to see the Prometheus console
    - [http://localhost:3000](http://localhost:3000) to access Grafana. User and password are `admin`. There's a friendly [dashboard](http://localhost:3000/d/KjLA0RfWz/sample-dashboard?orgId=1) there!