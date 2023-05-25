console.log("selamat pagi!");

let hamburgerIsOpen = false;

function openHamburger() {
  let hamburgerNavContainer = document.getElementById(
    "hamburger-nav-container"
  );

  // !hamburgerIsOpen > hamburgerIsOpen == false
  // hamburgerIsOpen > hamburgerIsOpen == true
  if (!hamburgerIsOpen) {
    console.log(hamburgerIsOpen);
    hamburgerNavContainer.style.display = "block";
    hamburgerNavContainer.style.backgroundColor = "red";
    hamburgerIsOpen = true;
  } else {
    console.log(hamburgerIsOpen);
    hamburgerNavContainer.style.display = "none";
    hamburgerIsOpen = false;
  }
}
