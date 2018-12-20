# Goraes

[![standard-readme compliant](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)

A simple cli utility to encrypt/decrypt plaintext files with AES. It uses a defaul configuration file, `conf.json` to load the files' path or they can be given as cli arguments.

It can be used to encrypt or decrypt any text file but I've specifically written it to handle login credentials for my online and offline accounts. It thus uses a json file similar to:

```json
	"AccountName": {
		"username": "myusername",
		"password": "mypassword",
		"email": "emailtologin"
	}
```

and so on, one entry for each account. Multiple accounts in the same site are grouped under a further json object level

```json
    "MainSite": {
        "url": "websiteurl",
        "account1": {
            "username": "myusername",
            "password": "mypassword",
			"email": "emailtologin"
        },
        "account2": {
			...
```

for as many accounts you have.

## How to use it

```go
	Arguments:
		-s|-searckey <word>
			Search for matching account names

		-i|-inputfile <file>
			The input file

		-o|-outputfile <file>
			The output file

		-d|-decrypt
			Set program to decrypt mode

		-e|-encrypt
			Set program to encrypt mode

		-p|-password
			Give encryption/decryption password directly on the command line
```

If no inputfile or outputfile are given, defaults to those set up in the config file. Encryption or decryption mode must be set. If no `-p` password given, a prompt will ask for it.

## TODO

+ search on plaintext file for an account and print the credentials in a readable format (easy to copy/paste)

## Contribute

PRs are welcome. Especially testing ones.

## License

MIT © Gianluca Fiore

[![ko-fi](https://www.ko-fi.com/img/donate_sm.png)](https://ko-fi.com/W7W7KA0Z)

