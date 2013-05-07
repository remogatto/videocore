# What's that?

<tt>videocore</tt> is a [Go](http://golang.org) package for accessing
the Broadcom VideoCore GPU on the [Raspberry
PI](http://www.raspberrypi.org).

<tt>videocore</tt> implements bindings to the following libraries:

* EGL
* OpenGL ES
* OpenVG
* OpenMAX

It's still in a non-working state.

# Install

~~~bash
$ go get github.com/remogatto/videocore/egl
$ go get github.com/remogatto/videocore/opengles
$ go get github.com/remogatto/videocore/openvg
$ go get github.com/remogatto/videocore/openmax
~~~

# Thanks

* Anthony Starks for his [openvg](https://github.com/ajstarks/openvg) library.
* Roger Roach for his [egl/opengles](https://github.com/mortdeus/egles) libraries.

# ToDo

* OpenMAX bindings
