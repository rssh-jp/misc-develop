FROM docker.elastic.co/elasticsearch/elasticsearch:8.15.2

RUN bin/elasticsearch-plugin install https://github.com/WorksApplications/elasticsearch-sudachi/releases/download/v3.3.0/elasticsearch-8.15.2-analysis-sudachi-3.3.0.zip
RUN curl -Lo sudachi-dictionary-20250129-core.zip https://github.com/WorksApplications/SudachiDict/releases/download/v20250129/sudachi-dictionary-20250129-core.zip && \
    unzip sudachi-dictionary-20250129-core.zip && \
    mkdir -p /usr/share/elasticsearch/config/sudachi/ && \
    mv sudachi-dictionary-20250129/system_core.dic /usr/share/elasticsearch/config/sudachi/ && \
    rm -rf sudachi-dictionary-20250129-core.zip sudachi-dictionary-20250129/

COPY ./resources/elasticsearch/sudachi.json  /usr/share/elasticsearch/config/sudachi/

USER elasticsearch
