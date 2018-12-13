Just Log it, with lit
====
[![GoDoc](https://godoc.org/github.com/bwmarrin/lit?status.svg)](https://godoc.org/github.com/bwmarrin/lit) [![Go report](http://goreportcard.com/badge/bwmarrin/lit)](http://goreportcard.com/report/bwmarrin/lit) [![Build Status](https://travis-ci.org/bwmarrin/lit.svg?branch=master)](https://travis-ci.org/bwmarrin/lit) [![Discord Gophers](https://img.shields.io/badge/Discord%20Gophers-%23info-blue.svg)](https://discord.gg/0f1SbxBZjYq9jLBk)

lit is a [Go](https://golang.org/) package that provides
* A very simple and opinionated message logger
* Logging levels that allow you to set the verbosity of what is logged.
* Added detail, such as file, line, and function name to each logged message.

**For help with this package or general Go discussion, please join the [Discord 
Gophers](https://discord.gg/0f1SbxBZjYq9jLBk) chat server.**

## Status @ 2018-08-20
This package is is the stage of figuring out what it wants to be. It may change
wildly or stay mostly the same.  I am definitely open to input on any changes 
you feel would be a good fit for this package.

## Design Goals
I find somethings, like logging, a bit tedious.  So I wanted a way to have a
very accessible logger that I could use anywhere without needing much setup.  

So, lit has minimal configuration options, it doesn't require to be instantiated
as a variable you pass around, or a global one you setup somewhere.  You can just
call the package functions from anywhere and there are handy methods for each of
the four log levels it supports.

## Usage

Add the package to your project.

Look around your code, and find a place that needs something logged.

If it's an error, just add a line like

```go
lit.Error("message here, %s", err)
```

Now that error message will be logged.  If it's something kind of spammy and not
even an error at all - just detail you need when debugging your application.

There's also lit.Warning() and lit.Informational() methods that can be used 
similarly.

```go
lit.Debug("message here")
```

Now that will be logged anytime your lit.LogLevel is set to LogDebug.

## Options

### lit.LogLevel
Can be set to LogError, LogWarning, LogInformational, and LogDebug.  The default 
is LogError.

### lit.Prefix
Can be set to any string you want to prefix all logged messages.

### lit.Writer
This can be set to any io.Writer and that's where your logged messages will go.

