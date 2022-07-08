# mark-offline

> Simple wrapper around the web based app

### Why?

I need an offline version of the app in cases where I don't really have web and
would like to extend the app to the functionality of being able to store and
read notes offline.

### Concept

The basic concept for now is to use go lang to build a single binary that
handles serving the assets for the web app and since it's a single page app it
works well.

Though for future, I'd like the go side to be able to pass through data to the
templates and thing won't be possible using the SPA arch and I might re-write
the app as a go web app

### Usage

Download a binary from the [releases](/releases) page for your specific system.

Once you have the binary in your system's `PATH`, you can run it from anywhere
like so

```sh
mark-offline
```

If you wish to run it on a different **port** than 3000, use the `PORT`
environment variable

```sh
PORT=8080 mark-offline
```
