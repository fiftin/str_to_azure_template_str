## What is it?

Simple tool for converting strings to valid Azure Resource Manager template strings.

It useful for passing a lot of data to VM by custom data in ARM templates.

For example you need pass some data to VM. In ARM template it can looks like:

```
{
...
"customData": "[base64(concat('#!/bin/sh\n\ncat > /home/ec2-user/config.json << CONFIG\n{\n   \"domain\": \"',reference(parameters('frontIpName')).dnsSettings.fqdn,'\",\n    \"sslKey\": \"',parameters('domainSslKey'),'\"\n}\nCONFIG\n\nsh /home/ec2-user/init.sh\n'))]",
...
}
```

It is hell. And it is not a lot of data.

This tool converts next readable text to string above:

```
#!/bin/sh

cat > /home/ec2-user/config.json << CONFIG
{
	"domain": "${reference(parameters('frontIpName')).dnsSettings.fqdn}",
	"sslKey": "${parameters('domainSslKey')}"
}
CONFIG

sh /home/ec2-user/init.sh


```

## How to install?

```
go get github.com/fiftin/str_to_azure_template_str
go install github.com/fiftin/str_to_azure_template_str
```

## How to use?
```
str_to_azure_template_str < test.txt
```