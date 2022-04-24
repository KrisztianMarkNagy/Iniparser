
# Iniparser

This is my implementation of an ini config file parser.

It is very simple, yet it allows the following:

- multiple config file parsing, one after another
- allows the usage of sections, example: `[This is a Section]` (While the sections are useless, it is good to have for a better visual navigation for the user)
- comments, both inline and not

This parser immediately sets the content of the ini file to the process' environment.
