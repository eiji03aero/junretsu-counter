# junretsu-counter
- An experiment to emulate vote-counting backend server used for junretsu's performance on 2020 kohaku.

## Background
On kohaku utagassen 2020, junretsu had a performance with a playful interactive voting system, which was
like if total number of voting achieve certain amount, the back-screen gets additional effect.
The overall voting count got pretty huge (30 million per minute), and got me thinking
how further performance-heavy the backend system has to be to deal with it.
So here I am, trying to emulate, see if it is feasible, and use it as an opportunity to learn.


# Plan
## Define what it does
### Requirements
- Create program to make requests which operates as fast as possible.
- Create program to make handle requests and increment counter.
- Prepare DB to write the result.
  - given data is not really a one that has to be accurate but fast enough to support throughput, redis is fine.
- Counting rate should achieve 500,000/s
  - The real number was around 90,000,000/3-minutes.
  - For brevity, just measure the amount processed within 1 second.
  - The real number looks pretty much impossible to achieve with local machine, so take it more like
    just a target performance and try to optimize codes as much as possible.

### Components
- Requester
  - makes request concurrently.
  - assign 1 cpu core.
- Handler
  - handles request to:
    - increments counter
    - keep the count to its own memory and flush it to db once in a while?
- DB
  - holds the current count, so that external system can retrieve it.

## Define how it will be done
- have db running.
- create one process for requester and handler respectively and measure the amount.
- request should be following layout in bits
  - 0-7: represents an unsigned integer indicating which button was pressed

## Implementation
- [ ] have containers ready
  - requester
  - handler
  - redis
- [ ] implement requester
- [ ] implement handler
