# Future Log

A series of items that I want to implement.

## Better Errors

I would like to return more advanced errors or use the existing errors to show more detailed error output. When tokens are not expected or rules fail it would be good to show the standard `Unexpected token X, expected Y` or such when possible. Maybe even with a code snippet showing the line and column where the invalid token is found.

## Metasyntax

It goes without saying that once its possible to parse a language it becomes possible to support a `meta syntax`. I think the most suitable based on functionality would have to be `EBNF` although it will need to be extended further to allow token definitions.

I would also like to have a crack at my own although only as an exercise. I think the support for various existing syntax's will be more beneficial if possible.

## Complexity analysis.

Analysing the tokens and rules (per defined tree and globally) to see what affect it has.
Tokens will be ranked based on complexity and type, for example literals will have a smaller cost to regex.
Rules also have complexity depending on type, choice rules and optional have a high chance of wasting time.

This data can be used in future to suggest the following optimisations:
* Re-arrangement of rules?
* Rule compresssion?
* Breaking regex in to rules and tokens?

## Token occurance analysis.

Based on complexity (or purely on type) and the length of the input tokens could be counted before tokenisation happens.
This would prevent tokens that can be costly from running when there is no chance it will match.
For example counting all the times a whitespace token matches against the input could provide a significant efficiency in some cases.
