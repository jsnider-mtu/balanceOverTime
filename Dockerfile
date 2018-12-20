FROM ubuntu:18.04

COPY balanceOverTime.sh .

RUN chmod +x balanceOverTime.sh

RUN echo "Acquire::http::Pipeline-Depth \"0\";" > /etc/apt/apt.conf.d/71pipelinedepth && apt-get update && apt-get install -y bc

ENTRYPOINT ["./balanceOverTime.sh"]