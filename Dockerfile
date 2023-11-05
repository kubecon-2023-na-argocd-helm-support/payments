FROM scratch
ADD dist/payments /bin/payments
ENTRYPOINT ["/bin/payments"]