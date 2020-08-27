## Custom MUX

In this code, we are creating a custom struct called **CustomServeMux** , which is going to take care of our routing. We implemented a function called **ServeHTTP** in order to capture the request and write a response back to it. The **fmt** package is usually used to create strings.
Fprinf composes the string out of supplied parameters.
In the main function, we are creating an instance of our CustomServeMux and passing it to
the **ListenAndServe** function on http . **math/rand** is the library that takes care of
generating random numbers. This basic foundation is going to be helpful for us when we
discuss adding authentication to our API server.