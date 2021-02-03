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
- Client
  - makes requests.
- Server
  - handles request to:
    - increments counter
    - keep the count to its own memory and flush it to db at proper rate
- DB
  - holds the current count, so that external system can retrieve it.

## Define how it will be done
- have db running.
- create one process for requester and handler respectively and measure the amount.
- request should be following layout in bits
  - 0-7: unsigned integer ... indicates which button was pressed

## Implementation
- [x] have containers ready
  - workspace
  - redis

- [ ] implement initial try
  - the simplemst way
  - client
    - just make request within loop
  - server
    - just handle requests within loop
    - write down the count on the closest last milli second

- [ ] implement requester
  - [ ] make sure to understand some of the go features
    - worker pool with go routine and channel
    - channel buffer
    - measure the time it takes to do context switch
  - [ ] try make it increment like
    - increase number of workers
      - try to find the most optimized number of workers to handle requests
    - try utilize multiple cores
    - measure how many jobs(increment) it can process
- [ ] implement handler

## Todo
- [ ] learn about bufio
- [ ] learn about atomic function, should be best choice to increment concurrently

## Reference
- https://ops.tips/blog/udp-client-and-server-in-go/
- https://medium.com/a-journey-with-go/go-string-conversion-optimization-767b019b75ef
- https://medium.com/a-journey-with-go/go-introduction-to-the-escape-analysis-f7610174e890
