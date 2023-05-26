// custom hof that return another function
function multiplyBy(n) {
  return function (x) {
    return x * n;
  };
}

// fill the parameter n
const double = multiplyBy(2);
// fill the parameter x
console.log(double(5));

// custom hof that implement callbacks
function repeat(n, action) {
  for (let i = 0; i < n; i++) {
    action(i);
  }
}

function logNumber(x) {
  console.log(`The number is ${x}`);
}

repeat(3, logNumber);
