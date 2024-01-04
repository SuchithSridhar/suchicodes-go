function updateThemeOnServer(theme) {
    const daysUntilExpire = 365;

    let date = new Date();
    date.setTime(date.getTime() + (daysUntilExpire * 24 * 60 * 60 * 1000));

    let expires = "expires=" + date.toUTCString();
    document.cookie = "theme=" + theme + ";" + expires + ";path=/";
}

function callThemeCallbacks(oldColor, counter) {
    let newColor = style.getPropertyValue('--color-foreground-1000').trim();

    // Color has not changed yet - wait for 1ms
    if (newColor == oldColor && counter < 100) {
        setTimeout(() => {callThemeCallbacks(oldColor, counter+1)}, 1);
        return;
    }

    let themeLink = document.getElementById('theme-css');
    let newTheme = themeLink.getAttribute('data-theme');
    for (let callback of window.changeThemeCallbacks) {
        if (typeof callback === 'function') {
            callback(newTheme);
        }
    }

}

function changeTheme() {
    let themeLink = document.getElementById('theme-css');
    let themeIcons = document.querySelectorAll('#theme-changer img');
    let togglerIcons = document.querySelectorAll('#navbar-toggler img');
    let currentTheme = themeLink.getAttribute('data-theme');
    let oldColor = style.getPropertyValue('--color-foreground-1000').trim();
    let newTheme;

    if (currentTheme === 'light') {
        newTheme = 'dark'
        themeLink.setAttribute('href', '/static/styles/colors-dark.css');
        themeLink.setAttribute('data-theme', 'dark');
        themeIcons.forEach(icon => {
            icon.setAttribute('src', '/static/assets/moon.svg')
        })
        togglerIcons.forEach(icon => {
            icon.setAttribute('src', '/static/assets/menu-dark.svg')
        })
        updateThemeOnServer(newTheme)
    } else {
        newTheme = 'light'
        themeLink.setAttribute('href', '/static/styles/colors-light.css');
        themeLink.setAttribute('data-theme', 'light');
        themeIcons.forEach(icon => {
            icon.setAttribute('src', '/static/assets/sun.svg')
        })
        togglerIcons.forEach(icon => {
            icon.setAttribute('src', '/static/assets/menu-light.svg')
        })
        updateThemeOnServer(newTheme)
    }

    setTimeout(() => {callThemeCallbacks(oldColor, 0)}, 1);
}

let themeChangerButtons = document.querySelectorAll('#theme-changer');
themeChangerButtons.forEach(btn => {
    btn.addEventListener('click', function() {
        changeTheme();
    });
});

const tabletToggler = document.querySelector('.nav-tablet #navbar-toggler');
const tabletNavbar = document.querySelector('.nav-tablet .navbar-nav');

tabletToggler.addEventListener('click', function() {
    tabletNavbar.classList.toggle('show');
});
