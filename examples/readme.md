
# Config Files

The `.ini` files in this directory were used for the development of this parser.

## Peek at `example01.ini`

```ini
[Test INI Config] ; This is an inline-comment
key001 = value001 # This also is an inline-comment

key002 = value002:semi-colon: ; * `:semi-colon:` will be replaced with an actual semi-colon.
; This is a comment

key003 = value003:hashtag: ; * Same goes for the the `:hashtag:`

[Another Section] # I don't know the use case of these sections, apart from being useful for visual navigation for the user.
# A comment
_this-IS-_key097124_ = "Some string value, even though every thing is a string here"

some_key = ';#more"hello" string == = = == - 102neioapsas][fu\"n'

; A Last Comment from Yours Truly.

a;asd
b#;asdnasodi;asd

You could write comments like this as well, since the parser will ignore all lines that don't have an assignment operator, start and end with square-brackets [], like a section,
or if it starts with a comment, be it either with `#` or `;`.

```

## Peek at `example02.ini`

```ini
; A comment
# Another Comment

[A Section]
akey01 = avalue01 ; Comment

akey02 = avalue02#Comment

asdio a$@AS iu # OISD ASD
[Another Section]
asddasd
akey03 = avalue03    ; MOOOOORE COMMENT

```
