# TypeScript

Initial impressions, this is a hot mess... I installed typescript with npm and tsc just transpiles to javascript, which is
fine I suppose, but I would like to just be able to run my code, not have an extra step.

I found the ts-node module, which claim to be able to just run your typescript code directly, except that I get an error
saying that ".ts" is an unknown file extension... why? who knows, I read through a huge github issue about it and I'm still
not sure why. 

I would move forward with just using tsc to build and node to run and build that into my docker-compose file, but then I would need to do that for each different file maybe, I'm not sure, seems like a mess. Someone said that you don't need to use ts-node you can use Deno or Bun - I don't know what those are or why I would choose to use them.

I will revisit this later when things seem a little less chaotic.