# Go Levy

Get your shared household expenses in order! Levy because `go-taxes` seemed too damn boring...

### Intention

This project brings an easy to use "tax/levy" calculation for shared expenses. This way each member of a living arrengement can pay up a share based on their income.

It will only work for situations in which all members are OK to provide different amounts for day-to-day expenses.

### Functionality

It expects three cli arguments:
- `TOTAL_TARGET`: The amount that you want to get into the shared account.
- `INCOME_1`: The income of the first contributor.
- `INCOME_2`: The income of the second contributor.

It has three **constant values** (that can be tweaked):
- `FRACTION_SIZE` (default: 100): The chunks in which the provided amount will be divided in.
- `BASE_PERCENTAGE` (default: 30%): The percentage calculation in which we start.
- `PERCENTAGE_INCREMENT` (default: 1%): How much the percentage gets incremented per fraction.

The algorithm divides the `inputAmounts` in `n` chunks of size `FRACTION_SIZE`. It then loops over those chunks, for each iteration adds `BASE_PERCENTAGE + PERCENTAGE_INCREMENT` to a calcuation of how much a member is due. That calculation gets normalized based on the calculation of the other member and then the contribution of each one is based on that percentage applied to the provided total amount.

