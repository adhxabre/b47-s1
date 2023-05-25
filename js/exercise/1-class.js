// Class

class Car {
  constructor(brand, model) {
    this.brand = brand;
    this.model = model;
  }

  getInfo() {
    return `The car is a ${this.brand} ${this.model}`;
  }
}
