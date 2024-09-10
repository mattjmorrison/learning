# Selenium


I really want to get this to work reliabily and repeatabily completely in docker


https://medium.com/@sharmila.may5/steps-to-run-selenium-tests-in-docker-7610281a5581

Ok, I got this working pretty reliabily. 

To start the selenium hub and firefox node run

```
    docker compose --profile selenium up
```

once that is started up you can open a browser to http://localhost:4444/ to see what's going on

To run your tests you can run

```
    docker compose run --rm test
```

I ran into some issues with selenium locking up if I would try to run the tests multiple times.

I tried multiple variations of using `quit` and `close` in either `tearDown` or `tearDownClass`.

What seems to work pretty consistently is having `quit` in `tearDownClass`.