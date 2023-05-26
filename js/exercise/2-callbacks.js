//callbacks | bek kiri bek kanan
function hello() {
  console.log("Hello World!");
}

function goodbye() {
  console.log("Goodbye World!");
}

function print(callback) {
  callback();
}

print(hello);
print(goodbye);
