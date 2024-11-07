FROM scratch

COPY ./arkeod /usr/bin/arkeod

WORKDIR /root/.arkeo

EXPOSE 1317

EXPOSE 26656

EXPOSE 26657

ENTRYPOINT [ "/usr/bin/arkeod" ]

CMD [ "help" ]