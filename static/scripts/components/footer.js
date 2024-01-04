function updateFooterTheme(theme) {
    // Select the LinkedIn and GitHub image elements
    var linkedinImage = document.querySelector('.linkedin-icon img');
    var githubImage = document.querySelector('.github-icon img');

    // Update the src attribute based on the theme
    if (theme != "dark") {
        // If the theme is not dark, set to light theme images
        linkedinImage.src = "/static/assets/linkedin-light.svg";
        githubImage.src = "/static/assets/github-light.svg";
    } else {
        linkedinImage.src = "/static/assets/linkedin-dark.svg";
        githubImage.src = "/static/assets/github-dark.svg";
    }
}

if (window.changeThemeCallbacks !== undefined) {
    window.changeThemeCallbacks.push(updateFooterTheme)
} else {
    window.changeThemeCallbacks = [updateFooterTheme];
}
