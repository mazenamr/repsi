# repsi
a go tool to generate an NFA or DFA from a regular expression

## usage

```
repsi convert <regex>
```

this outputs the following files:
    + nfa.json
    + nfa.svg
    + dfa.json
    + dfa.svg
    + dfa-minimized.json
    + dfa-minimized.svg