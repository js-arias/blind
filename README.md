# Blind

Package `blind` implement color gradients
that are friendly for color blind people.

The color schemes are taken from [Paul Tol](https://personal.sron.nl/~pault/) page.

Here is a color gradient build by the [gradient example](blind_gradient_test.go)
based on [rainbow color scheme (purple to red)](https://personal.sron.nl/~pault/#fig:scheme_rainbow_smooth).

![Gradient example](gradient.png)

## Implemented color schemes

Here are another gradients:

[Iridescent color scheme](https://personal.sron.nl/~pault/#fig:scheme_iridescent)

![Iridescent](iridescent.png)

Full [rainbow color scheme](https://personal.sron.nl/~pault/#fig:scheme_rainbow_smooth).
This scheme is not intended to be used in full,
rather use an slice frame
(such as the purple to red used by default in `Gradient` function).

![Full rainbow](full-rainbow.png)

[Incandescent color scheme](https://personal.sron.nl/~pault/#fig:scheme_incandescent).
This scheme is not printer friendly.

![Incandescent](incandescent.png)

## Authorship and license

Color schemes copyright © 2022 Paul Tol <https://personal.sron.nl/~pault/>
Code copyright © 2023 J. Salvador Arias <jsalarias@gmail.com>.
All rights reserved.
Distributed under BSD3 licenses that can be found in the LICENSE file.
