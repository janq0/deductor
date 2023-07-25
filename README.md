# deductor

Takes in a deductive argument and determines if it's valid. See
<https://en.wikipedia.org/wiki/Deductive_reasoning> for more info.

## Usage

Specify any number of premises separated by a newline. To input a conclusion,
start the line with `therefore <conclusion>`. The validity is evaluated after
the conclusion line.

### Features

- [x] Logical connectives
    - [x] `<x> and <y>`
    - [x] `<x> or <y>`
    - [x] `<x> if <y>`
    - [x] `<x> only if <y>`
    - [x] `not <x>`
- [x] Nesting using parenthesis
- [ ] Detecting a negative statement without `not`

### Examples

```
it is raining only if there are clouds in the sky
not there are clouds in the sky
therefore not it is raining
true
```