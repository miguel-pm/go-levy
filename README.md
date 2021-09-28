# Go Levy

Get your shared household expenses in order!

### Intention

This project brings an easy to use "tax/levy" calculation for shared expenses. This way each member of a living arrengement can pay up a share based on their income.

It will only works for situations in which all members are OK to provide different amounts for day-to-day expenses.

### Functionality

It expects an `inputAmount` (income of a member) as a cli argument and has three **constant values** (that can be tweaked):

- `FRACTION_SIZE` (default: 100): The chunks in which the provided amount will be divided in.
- `BASE_PERCENTAGE` (default: 30%): The percentage calculation in which we start.
- `PERCENTAGE_INCREMENT` (default: 1%): How much the percentage gets incremented per fraction.

The algorithm divides the `inputAmount` in `n` chunks of size `FRACTION_SIZE`. It then loops over those chunks, for each iteration adds `BASE_PERCENTAGE + PERCENTAGE_INCREMENT * n` to the result. The output is the amount to be given to the shared account. This way the amount "taxed" increments the more income you have.

