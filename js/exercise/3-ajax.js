// What is AJAX?
// basically ajax is simply an technique in Javascript to make us able to interact with asynchronous request.
// Example, if you get data from internet.
const xhr = new XMLHttpRequest();

// Basic syntax for we learn looks like this
xhr.open("GET", "https://your-url", true);
// param 1 : is the method
// param 2 : is the url
// param 3 : true or false, if true means that we using asynchronous WebAssembly, if false means that we using synchronous way. It is recommended to make it true (async).

xhr.onload = function () {}; // to load and check the status of the request
xhr.onerror = function () {}; // loaded when it is error when we requesting the data
xhr.send(); // send the request to the server

// AJAX is a technique for making asynchronous HTTP reqest from a web page to a server, using the XMLHttpRequest object in the browser. Since the XMLHttpRequest object is specific to the browser environment, it cannot be used directly in Node.JS. So we gonna make it directly in our project.

// And it is gonna using connection too, because we are gonna using AJAX to request data from internet.
