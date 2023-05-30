// What is async await? The point is that async-await is simply other to call promise but it looks like in synchronous way.
let condition = true;

let promise = new Promise((resolve, reject) => {
  if (condition) {
    resolve("Promise is resolved");
  } else {
    reject("Janji kok gak ditepatin");
  }
});

// if we previously call the promise like this:
// promise
//   .then((value) => {
//     console.log(value);
//   })
//   .catch((reason) => {
//     console.log(reason);
//   });

// now, with async-await, we can call it like this:
async function asyncFunction() {
  const response = await promise;
  console.log(response);
}

asyncFunction(); // we call the function here
