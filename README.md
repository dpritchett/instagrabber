# instagrabber

## What is this?

This is a simple program that grabs a recent thumbnail from my instagram feed and serves it up to end users.

## Why?

1. To experiment with simple serverless programming
2. To enable some future chatbot hacks I have in mind

## How?

[Click here to see a random photo](https://clkpnz3810.execute-api.us-east-2.amazonaws.com/development/) served up by this app on a lambda.

### Deploying your own

You'll need to have golang installed with a properly configured `$GOPATH`.

You'll want a copy of `up`.

`npm i -g up`

You'll need to download this:

`go get -u github.com/dpritchett/instagrabber`

You'll need to be logged into an AWS account.

Then you use `up` to deploy your own lambda with APIGateway frontend:
```sh
bash-3.2$ up && up url

     build: 8 files, 6.3 MB (506ms)
     deploy: complete (14.881s)

https://w120cddvm8.execute-api.us-east-2.amazonaws.com/development/
```

#### Huh?
There's a lot of prereqs here. The magic is that `up` and Lambda save you from needing to deal with your own server. Maybe you'd rather write Node or Python instead of Go.  That's cool! If you have *any* questions just Tweet me or email me and I'd love to help.

## Who?

[Daniel Pritchett](http://blog.dpritchett.net)
