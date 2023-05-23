// console.log("Hello student!");
// alert("Good morning madafaka");
// document.write("Batch 47");

// Variable

// VAR
// LET
// CONST

// bisa di deklarasikan ulang dan bisa diubah value/data nya
// var gelas = "Air Putih";
// var gelas = "Kopi";
// console.log(gelas);

// tidak bisa dideklarasikan ulang, namun value/datanya bisa kita ubah
// let pemerintah = "PDIP";
// pemerintah = "Nasdem";
// console.log(pemerintah);

// tidak bisa dideklarasikan ulang dan tidak bisa diubah value/data nya
// const rog = "asus";
// rog = "lenovo";
// console.log(rog);

// gender = "attack helicopter";
// console.log(gender);

// data type
// let nama = "Abel Dustin";
// let umur = 15;
// let isMarried = false;

// nama saya Abel Dustin umur saya 15 tahun
// console.log("nama saya Abel Dustin umur saya 15 tahun");
// console.log(`nama saya ${nama} umur saya ${umur} tahun`);
// console.log("nama saya", nama, "umur saya", umur, "tahun");
// console.log("nama saya " + nama + " umur saya " + umur + " tahun");

// operator
// let x = 48;
// let y = "4";

// let result = x + y;

// console.log(result);

// condition
// jika nilai sama dengan atau lebih dari 75, maka lulus

// let nilai = "75";

// if (nilai === 75) {
//   console.log("Kamu lulus!");
// } else if (nilai === "75") {
//   console.log("Salah input!");
// } else {
//   console.log("Kamu tidak lulus!");
// }

// function
// function Hello() {
//   let x = 5;
//   let y = 10;

//   let result = x * y;
//   console.log(result);
// }

// Hello();

// function Hello2(x, y) {
//   console.log(x);
//   console.log(y);

//   let result = x * y;
//   console.log(result);
// }

// Hello2(5, 10);

// camelcase = namaSayaAdalah

// function namaSaya(name) {
//   console.log(name);
// }

// namaSaya("Abel");

let name = "abel";

function submitData() {
  let name = document.getElementById("input-name").value;
  let email = document.getElementById("input-email").value;
  let phone = document.getElementById("input-phone").value;
  let subject = document.getElementById("input-subject").value;
  let message = document.getElementById("input-message").value;

  if (name == "") {
    return alert("Nama harus diisi!");
  } else if (email == "") {
    return alert("Email harus diisi!");
  } else if (phone == "") {
    return alert("Phone harus diisi!");
  } else if (subject == "") {
    return alert("Subject harus dipilih!");
  } else if (message == "") {
    return alert("Message harus diisi!");
  }

  let emailReceiver = "abeldustin06@gmail.com";

  //   <a href="mailto:abeldustin06@gmail.com?subject=frontend&body=Halo, nama saya Abel, P mabar. Silakan kontak saya di nomor 08971406428, terima kasih."></a>

  let a = document.createElement("a");
  a.href = `mailto:${emailReceiver}?subject=${subject}&body=Halo, nama saya ${name}, ${message}. Silakan kontak saya di nomor ${phone}, terima kasih.`;
  a.click();

  console.log(name);
  console.log(email);
  console.log(phone);
  console.log(subject);
  console.log(message);

  let emailer = {
    name,
    email,
    phone,
    subject,
    message,
  };

  console.log(emailer);
}

if (cowo === "ganteng") {
  dapetCewe();
} else if (cowo === "kaya") {
  dapetCewe();
} else {
  cumaTemenan();
}

if (jokowi === "3 periode") {
  indonesiaMaju();
} else {
  indonesiaMundur();
}
