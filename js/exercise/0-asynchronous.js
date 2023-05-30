// synchronous concept:
// console.log("Hello!");
// console.log("Javascript");
// console.log("Coder");

// asynchronous concept:
setTimeout(function () {
  console.log("Hello!");
}, 3000);
setTimeout(() => {
  console.log("Javascript");
}, 5000);
console.log("Coder");

// => : = >
