FROM alpine

WORKDIR /

COPY . /

RUN echo `date +'%s'` > buildtime.txt

CMD ["sh", "script.sh"]
