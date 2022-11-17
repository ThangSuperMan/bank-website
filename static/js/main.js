// Query selectors
const headerLogin = document.querySelector(".header-login")
const headerLoginSubmenu = document.querySelector(".header-login-submenu")


// Add event listnners
headerLogin.addEventListener('mouseenter', () => {
  headerLoginSubmenu.classList.add("active")
})

headerLogin.addEventListener('mouseleave', () => {
  headerLoginSubmenu.classList.remove("active")
})

