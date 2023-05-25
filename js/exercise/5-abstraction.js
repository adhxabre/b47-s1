// Abstraction

class Car {
  #brand = "";
  #model = "";
  _house = "";

  constructor(brand, model) {
    this._brand = brand;
    this.#model = model;
  }

  get brand() {
    return this.#brand;
  }

  get model() {
    return this.#model;
  }
}

let myCar = new Car("Toyota", "Vios");
console.log(myCar.brand); // Toyota
console.log(myCar.model); // Vios
