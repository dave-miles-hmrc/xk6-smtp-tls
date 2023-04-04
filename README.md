**xk6-smtp-tls** is an [xk6](https://github.com/grafana/xk6) extension to send emails via the [gomail](https://github.com/go-gomail/gomail) package.

## Build
```shell
xk6 build --with github.com/dave-miles-hmrc/xk6-smtp-tls@latest
```

## Example
An example of using the extension:

```javascript
import smtp_tls from 'k6/x/smtp_tls';


const server = JSON.parse(open('./ses-account.json'));

export default function() {

  const server = {
    "host": "email-smtp.eu-west-2.amazonaws.com",
    "port": 465,
    "account": "AKIAAKIAAKIAAKIAAKIA",
    "password": "PasswordPasswordPasswordPasswordPassword",
    "secure_mode": "ssl",
    "helo_string": "",
    "disable_cert_validation": false
  }

  const mail = {
    from: "some.one@example.com",
    to: ["success@simulator.amazonses.com"],
    subject: "Success subject",
    body: "Hello World!\nWhat a world it still is.\n"
  }

  smtp_tls.sendMail(server, mail);
  sleep(2);

}
```
