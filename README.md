# r53dyndns

This is a small utility to fetch the external IP of a client and update an A record in Route53 with the IP value.

I wrote this as a hack for basic dynamic DNS without signing up for anything other than AWS.

### Usage

```
$ go get github.com/cgansen/r53dyndns

$ R53_DOMAIN_NAME=example.com R53_HOSTED_ZONE_ID=Z1234EXAMPLE R53_TTL=600 $GOBIN/r53dyndns
2015/10/13 23:17:17 external ip is: 1.2.3.4
2015/10/13 23:17:17 route53 update was a success.
```

Note: the AWS Go SDK assumes that you have your AWS credentials available in `$HOME/.aws/credentials`. More complex authentication, including using alternate IAM credentials, is left as an exercise for the reader.

### Inspiration

These fine folks made similar tools that do roughly the same job. Sadly, none of them are written in Go.

- http://holgr.com/blog/2014/08/using-amazon-route-53-as-dynamic-dns-service/
- https://github.com/holgr/php-ddns53
