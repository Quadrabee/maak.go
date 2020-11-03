FROM alpine

WORKDIR /

COPY src/ /
COPY assets/ /

RUN echo `date +'%s'` > greeting.txt

CMD ["sh", "/script.sh"]
