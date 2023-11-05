FROM scratch
COPY ./dist/payments /bin/payments
ENTRYPOINT ["/bin/payments"]