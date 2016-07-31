# FROST

FROST is for people who write desktop applications that read text based input
file decks. The apps often end up having to read and parse a number of files in
a whole variety of formats. Maybe a mix of XML, JSON, CSV, text delimitted etc.
The file formats are often not standardised and can be frankly eccentric.

Consequently, a great deal of time is expended in writing file parsers and
interpreters in multiple languages and for multiple platforms.

Frost seeks to replace this with a single web service API that sucks in all the
files in an input directory, makes a stab at deducing what the formats mean,
and then returning the whole thing as one big structured JSON representation.

It tries to recognize lists and tables etc rather like a human being would. It
doesn't expect to get this right. But it does expect to make it repeatable,
which removes much of the heavy lifting from the original coder.

It serves interactive web pages too as well as the API. It's written in Go and
the current work in progress is hosted on Google App Engine here

http://frost-1001.appspot.com/quickstart

You can read more in the docs directory.


