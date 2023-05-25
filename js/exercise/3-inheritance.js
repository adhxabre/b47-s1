class Car {
  constructor(brand, model) {
    this.brand = brand;
    this.model = model;
  }

  getInfo() {
    return `The car is a ${this.brand} ${this.model}`;
  }
}

class ElectricCar extends Car {
  constructor(brand, model, batteryCapacity) {
    super(brand, model);
    this.batteryCapacity = batteryCapacity;
  }

  getInfo() {
    return `${super.getInfo()} and has a battery capacity of ${
      this.batteryCapacity
    }`;
  }
}

let myCar = new ElectricCar("Tesla", "Model S", 100);
console.log(myCar.getInfo());

// output: The car is a Tesla Model S and has a battery capacity of 100
